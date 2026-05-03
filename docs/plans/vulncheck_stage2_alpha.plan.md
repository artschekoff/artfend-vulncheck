---
name: VulnCheck Stage 2 -- Alpha
overview: 2-week extension of Stage 1 MVP. Adds the killer product differentiator -- AI Impact Analyzer -- plus Git integration (GitHub PAT), Remediation Engine (safe fix suggestions), Dual Policies (DB + .vulncheck file), supply chain detection (lifecycle hooks, typosquatting, new-package baseline), two more vuln sources (NVD + GHSA), and a minimal SAST layer for Node.js (top 10 Semgrep rules). At the end of Stage 2, the product answers the key question "is this CVE actually dangerous in MY code?" for a Node.js repository scanned from GitHub.
todos:
  - id: s2-github-integration
    content: "Git integration: GitHub PAT auth + repo clone + on-demand scan trigger. /api/projects can now accept a GitHub URL."
    status: pending
  - id: s2-vuln-sources-expand
    content: "Add NVD + GHSA Advisory DB sources. Update dedup to handle 4-source merge. Daily sync job."
    status: pending
  - id: s2-supply-chain
    content: "Supply chain detectors: install-hook scanner (postinstall/preinstall/install), typosquat detector (Levenshtein vs top-5000 packages), new-package detector (registry time < 30 days)."
    status: pending
  - id: s2-ai-analyzer
    content: "AI Impact Analyzer (Node.js): code indexer, snippet extractor, prompt templates, OpenAI provider, verdict parser, persist AIVerdict. Run per finding async."
    status: pending
  - id: s2-effective-severity-v2
    content: "Effective severity Stage 2: combine scope + AI verdict. Formula extended to incorporate verdict=no_impact/low/medium/high."
    status: pending
  - id: s2-sast-minimal
    content: "Semgrep integration + top 10 Node.js/Express/NestJS rules (XSS, SSRF, SQLi in raw query, path traversal, prototype pollution, eval, child_process injection, insecure randomness, insecure deserialization, SSTI). Run as part of scan."
    status: pending
  - id: s2-remediation
    content: "Remediation engine: computes fix plans (patch upgrade / minor upgrade / major upgrade with breaking-change risk tier). Returns fix-risk tier per finding."
    status: pending
  - id: s2-dual-policies
    content: "Dual policies: DB-backed org policies + .vulncheck policy file in repo. Policy engine evaluates findings -> decisions (block / warn / allow). CLI respects both."
    status: pending
  - id: s2-cli-extend
    content: "CLI extensions: `vulncheck fix --dry-run`, `vulncheck ignore <finding-id> --reason ...`, `vulncheck monitor` stub, output shows AI verdict + fix risk tier"
    status: pending
  - id: s2-tests-docs
    content: "Integration test: scan a seeded GitHub repo end-to-end with AI disabled+enabled. Update README, CLI help, API docs."
    status: pending
isProject: false
---

# Stage 2 -- Alpha (~2 weeks)

Parent roadmap: [vulncheck_roadmap.plan.md](vulncheck_roadmap.plan.md). Master plan: [vulncheck_snyk-like_platform_41691b25.plan.md](vulncheck_snyk-like_platform_41691b25.plan.md). Requires Stage 1 complete.

## Goal

By the end of Stage 2, a user can:

1. `vulncheck auth login` and link a GitHub PAT
2. Create a project from a GitHub URL (clones on scan)
3. Run `vulncheck test` (remote repo) or `vulncheck test --path .` (local)
4. See findings annotated with:
   - Scope (prod/dev/optional/peer) **[from Stage 1]**
   - AI verdict (`no_impact` / `low` / `medium` / `high`) **[NEW]**
   - Fix risk tier (`safe` / `moderate` / `breaking`) **[NEW]**
   - Policy decision (`block` / `warn` / `allow`) **[NEW]**
   - Supply chain flags (has install hook, typosquat candidate, new package) **[NEW]**
5. Run `vulncheck fix --dry-run` to see proposed fix plan
6. Drop a `.vulncheck` file in repo root to override org policy

**The demo pitch becomes: "Point us at your repo. For each CVE, we tell you with AI whether it actually matters to YOUR code, and a safe fix suggestion. We catch malicious packages that Snyk misses."**

## Scope IN

- **Git Integration** (master Phase C2)
  - GitHub PAT-only for Stage 2. GitLab + Bitbucket deferred to Stage 3.
  - `internal/git/provider.go` interface + `internal/git/github.go` implementation
  - Clone to temp dir; scan; discard
  - Project links (`project.repo_url`, `project.default_branch`, `project.provider='github'`)
- **Vulnerability Sources: NVD + GHSA** (master Phase E)
  - `internal/sca/nvd.go` -- NVD JSON feed sync (daily), local DB cache table `nvd_advisories`
  - `internal/sca/ghsa.go` -- GitHub GraphQL API for GHSA advisories
  - Dedup updated to 4-source merge
  - Daily `vuln_sync_job.go` refreshes local caches
- **Supply Chain Detection** (master Phase E -- supply chain subsection)
  - `internal/supply/hooks.go` -- parses package.json for `scripts.{pre,post}install`, flags shell / network / filesystem operations
  - `internal/supply/typosquat.go` -- Levenshtein distance 1-2 to a curated top-5000 package names list
  - `internal/supply/new_package.go` -- query npm registry metadata, flag if `time.created < 30d`
  - Results stored as `finding.type='supply_chain'` with subtype
- **AI Impact Analyzer** (master Phase E2 -- Node.js only, OpenAI only)
  - `internal/ai/indexer.go` -- indexes repo source files (AST or regex-based import extraction)
  - `internal/ai/snippet.go` -- extracts relevant code snippets referencing the vulnerable package
  - `internal/ai/prompts.go` -- prompt templates for: impact analysis, fix safety assessment
  - `internal/ai/provider.go` + `internal/ai/openai.go` -- OpenAI client wrapper (GPT-4o or later, configurable)
  - `internal/ai/verdict.go` -- parses structured JSON verdict from LLM; persists AIVerdict model
  - Queued after SCA; one job per high/critical finding; timeout 60s per finding
- **Effective Severity Stage 2** (master Phase E)
  - `internal/sca/effective_severity.go` extended: combines scope + AI verdict
  - Formula: `eff = base + scope_adj + ai_adj`, where `ai_adj` is -3/-2/-1/0 for `no_impact/low/medium/high`
- **SAST Minimal** (master Phase E -- SAST subsection)
  - `internal/sast/semgrep.go` -- wraps Semgrep CLI, parses JSON output
  - `rules/semgrep/node-express-nest-top10.yaml` -- curated 10 rules, each with positive + negative fixtures
  - Run as part of scan pipeline for Node.js projects
  - SAST findings use same `Finding` model with `type='sast'`
- **Remediation Engine** (master Phase F)
  - `internal/remediation/planner.go` -- for each finding, computes fix plan:
    - `safe` = patch version bump within semver
    - `moderate` = minor version bump
    - `breaking` = major version bump OR no fixed version yet
  - `internal/remediation/upgrades.go` -- queries npm registry for versions
  - Fix plans stored on finding (`finding.fix_plan` JSONB)
- **Dual Policies** (master Phase G)
  - DB-backed: `policy` table with rules (`finding_type`, `scope`, `severity_threshold`, `decision`)
  - File-backed: `.vulncheck` YAML in repo root, parsed and merged (file overrides DB on conflicts)
  - `internal/policy/engine.go` -- evaluates finding against merged policy, returns decision
  - Policy decisions added to findings output
- **CLI Extensions**
  - `vulncheck fix --dry-run` -- shows computed fix plan per finding, no writes
  - `vulncheck ignore <finding-id> --reason "..."` -- creates an ignore rule (stored DB + suggests `.vulncheck` snippet)
  - `vulncheck monitor` -- stub command, prints "Monitor mode lands in Beta"
  - Output includes new columns: AI Verdict, Fix Risk, Policy Decision, Supply Chain Flags

## Scope OUT (explicitly deferred)

| Feature | Stage |
| --- | --- |
| GitLab + Bitbucket providers | 3 |
| Anthropic provider, self-hosted LLM | 3 |
| Multi-file AI context | 3 |
| AI feedback loop | 5 |
| VulnCheck + Trivy DB sources | 3 |
| Full Node.js SAST rule set + AI/LLM rules | 3 |
| Behavioral package analysis | 3 |
| Hallucination detection | 3 |
| Call-graph reachability | 3 |
| License, SBOM, Secrets | 3 |
| Fix PR automation (actual PR creation) | 3 |
| Jira, Notifications, PR Checks, Reports, Monitor mode | 3 |
| Web dashboard | 3 minimal |
| Multi-ecosystem | 4 |
| Container scanning | 4 |
| Rego, VEX, SLSA | 4 |
| IDE plugin | 4 |
| ASPM, runtime agent | 5 |

## Effective Severity -- Stage 2 Formula

```
base_severity = max over sources
scope_adjustment: as Stage 1
ai_adjustment:
  no_impact    -> -3 levels (often pushes to none / dismiss)
  low          -> -2 levels
  medium       -> -1 level
  high         -> 0
  (no verdict) -> 0
effective_severity = clamp(none, base + scope_adj + ai_adj)
```

If an ignore rule matches, decision is `ignored` regardless of severity.

## Milestones

- **Week 1**
  - Days 1-2: GitHub integration + clone + repo-based scan flow
  - Day 3: NVD + GHSA sources + dedup expansion
  - Day 4: Supply chain detectors (hooks, typosquat, new package)
  - Day 5: SAST integration + top 10 Node.js rules + fixtures
- **Week 2**
  - Days 1-3: AI Impact Analyzer end-to-end (Node.js, OpenAI)
  - Day 4: Remediation engine + fix plans
  - Day 5: Dual policies (DB + file)
  - Days 6-7: CLI extensions, integration test, README + docs, polish

## Detailed Checklist

### GitHub integration

- [ ] `internal/git/provider.go` interface: `Clone`, `FetchDefaultBranch`, `ListRepos`
- [ ] `internal/git/github.go` implementation with go-github
- [ ] PAT stored encrypted (AES-256-GCM) in `user_credential` table
- [ ] `POST /api/git/github/connect` with PAT, validates via `/user`, stores
- [ ] `POST /api/projects` now accepts `repo_url` + `provider='github'`
- [ ] Scan job clones repo to `/tmp/vulncheck/<scan_id>`, cleans up after

### Vuln sources expansion

- [ ] `internal/sca/nvd.go` -- syncs NVD JSON 2.0 feed, persists to `nvd_advisories`
- [ ] `internal/sca/ghsa.go` -- GitHub GraphQL `securityAdvisories` pagination
- [ ] `internal/jobs/vuln_sync_job.go` -- daily at 03:00 UTC
- [ ] Dedup updated with priority: GHSA > NVD > OSV > npm audit for conflicts
- [ ] CVE/GHSA/OSV alias normalization

### Supply chain

- [ ] `internal/supply/hooks.go`: detects `preinstall`, `postinstall`, `install` scripts; flags suspicious patterns (curl, wget, eval, base64, | sh, network calls, writes outside node_modules)
- [ ] `internal/supply/typosquat.go`: Levenshtein 1-2 to curated list `data/npm-top-5000.txt`; only flags if different package (not aliases)
- [ ] `internal/supply/new_package.go`: queries `registry.npmjs.org/<pkg>`, checks `time[version].created`; flag if < 30 days AND also not from known publishers list
- [ ] Supply chain finding type + subtype + severity mapping
- [ ] Unit tests with fixture packages

### AI Analyzer

- [ ] `internal/ai/indexer.go`: walks repo, indexes `.js/.ts/.mjs/.cjs/.jsx/.tsx` files; for each file records imports (ES + CommonJS)
- [ ] `internal/ai/snippet.go`: given `(package_name)`, returns list of `(file, line_start, line_end, snippet)` up to 10 snippets, max 200 lines total
- [ ] `internal/ai/prompts.go`: IMPACT_ANALYSIS prompt with structured JSON output schema (`verdict`, `reasoning`, `confidence`, `key_evidence`)
- [ ] `internal/ai/provider.go` interface + `openai.go` implementation (model from env, default `gpt-4o-mini`)
- [ ] `internal/ai/verdict.go`: parse JSON, validate, persist to `ai_verdicts` table linked to `finding_id`
- [ ] River job `ai_analysis_job.go`: one job per finding where `severity >= medium AND scope != dev-only`
- [ ] Rate limiting per user: max 100 AI calls per scan (configurable); excess findings skipped with a warning
- [ ] Cost logging per call (tokens in / out)

### Effective severity v2

- [ ] `effective_severity.go` updated with AI adjustment
- [ ] Unit tests for all 4 AI verdicts x 4 scopes x 5 base severities matrix
- [ ] Regression test: Stage 1 behavior preserved when AI is disabled

### SAST minimal

- [ ] Semgrep installed in Docker image
- [ ] `internal/sast/semgrep.go`: runs `semgrep --config <rules> --json .`, parses output
- [ ] `rules/semgrep/node-express-nest-top10.yaml`: 10 rules:
  1. Express `res.send(req.query.X)` -> XSS
  2. `child_process.exec(req.X)` -> Command injection
  3. `eval(req.X)` / `new Function(req.X)` -> Code injection
  4. Insecure randomness (`Math.random()` in security contexts)
  5. Path traversal (`fs.readFile(req.params.path)`)
  6. Prototype pollution (`Object.assign(target, req.body)`)
  7. SSRF (`axios/fetch(req.query.url)`)
  8. Raw SQL injection (`knex.raw(\`${req.X}\`)`)
  9. Insecure deserialization (`serialize-javascript unsafe: true`)
  10. SSTI (`ejs.render(template + req.X)`)
- [ ] Each rule has `positive.test.js` + `negative.test.js` fixtures
- [ ] CI runs `semgrep --test rules/` on every PR

### Remediation

- [ ] `internal/remediation/planner.go`: for each finding with fixed_in, compute best fix version
- [ ] Fix risk tiers: `safe` (patch within same minor), `moderate` (minor bump within same major), `breaking` (major bump or no fix)
- [ ] Stored as `finding.fix_plan` JSONB `{recommended_version, tier, upgrade_chain, breaking_notes}`
- [ ] Unit tests with semver fixtures

### Dual policies

- [ ] `policy` table with fields: id, org_id, rules JSONB, is_default, priority
- [ ] `.vulncheck` YAML parser: `policies:`, `ignores:`, `fix_risk_threshold:`, `fail_on:`
- [ ] `internal/policy/merge.go`: merges DB policy + file policy (file overrides DB)
- [ ] `internal/policy/engine.go`: evaluates finding -> decision
- [ ] CLI respects merged policy for exit code
- [ ] Policy decision stored per finding

### CLI extensions

- [ ] `vulncheck fix --dry-run`: prints proposed lockfile changes + risk tiers, no write
- [ ] `vulncheck ignore <finding-id> --reason "..."`: creates DB ignore AND prints `.vulncheck` YAML snippet to paste into repo
- [ ] `vulncheck monitor`: stub that prints "Monitor mode lands in Beta, see stage 3"
- [ ] Output table now includes: Scope / Raw Sev / Eff Sev / AI Verdict / Fix Risk / Policy / Supply Chain

### Tests + Docs

- [ ] Integration test: seed a test GitHub repo + fixtures, run full pipeline, assert counts
- [ ] README updated with AI setup instructions (OpenAI key)
- [ ] API reference (OpenAPI) regenerated
- [ ] `docs/policy-file.md` describing `.vulncheck` schema

## Definition of Done

- All Stage 1 Definition of Done still passes
- All boxes above checked
- End-to-end demo: connect GitHub -> scan a real repo -> AI verdicts populated for high/crit findings -> fix plan suggested -> `.vulncheck` ignore roundtrips
- `vulncheck test --json` schema bumped to v2, with v1 still accepted (opt-in flag)

## Key Risks

- **OpenAI cost runaway** -- hard-cap AI calls per scan (default 100), log tokens, allow disabling via env. Document cost model.
- **LLM hallucination in verdicts** -- use structured JSON output with schema validation; retry on parse failure once; if still invalid, skip verdict (finding falls back to Stage 1 severity).
- **Snippet extraction missing key code** -- start simple (regex on imports), accept limitations; note that multi-file repo-level AI context is deferred to Stage 3 (significant accuracy boost).
- **Semgrep install size in Docker image** -- pin minimal install (no extra packs), target <200MB layer.
- **Policy file precedence confusion** -- document clearly. Default: file-level ignores additive; severity thresholds: lower of the two (more restrictive wins).
- **GitHub clone size limits** -- cap at 500MB, reject bigger repos with an actionable error.
- **NVD feed ingestion churn** -- retry + resume; persist `last_synced_at`; alert on sync failures.
- **AI verdict trust** -- clearly label AI verdicts in UI/CLI as "AI-assisted, review before relying on"; don't use AI to downgrade a finding to `none` without user opt-in (default cap: AI can only push down to `low`, never fully suppress).
