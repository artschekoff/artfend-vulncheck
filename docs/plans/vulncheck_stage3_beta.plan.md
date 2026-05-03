---
name: VulnCheck Stage 3 -- Beta
overview: ~1 month extension (4-5 weeks). Turns the Node.js product into a fully customer-ready offering. Adds Behavioral Package Analysis (sandboxed tarball execution + capability fingerprinting), Hallucination / Slopsquat detection, Function-Level Call-Graph Reachability (Node.js + Go self-hosting), License scanner, SBOM generation (CycloneDX + SPDX with VEX export), Secrets detection, Automated Fix PR creation, Jira integration, Monitor mode, Notifications (Slack + email + webhook), PR checks + PR Baseline Diff, Reports (PDF + CSV), remaining vuln sources (VulnCheck KEV + Trivy DB), full Node.js SAST rule set (20+ rules) + AI/LLM security rules (15+ rules), multi-file repo-level AI context, Anthropic + self-hosted LLM providers, GitLab + Bitbucket providers, full CLI set, and a minimal read-only web dashboard. By end of Stage 3, the product is sellable to Node.js-first shops as a Snyk replacement with demonstrable differentiators.
todos:
  - id: s3-behavioral-analysis
    content: "Behavioral Package Analysis: tarball fetch + sandboxed exec + AST capability extractor (network/fs/shell/env/eval/native/obfuscation/install_hook_shell) + version diff. River job behavior_profile_job.go."
    status: pending
  - id: s3-hallucination-detection
    content: "Hallucinated / Slopsquat detection: cross-reference imports in Git repo against npm registry existence + maintainer reputation + creation date + known-bad feed."
    status: pending
  - id: s3-callgraph
    content: "Function-level call-graph reachability (Node.js via ts-morph, Go via golang.org/x/tools/go/callgraph). Entry-point discovery + path tracing to vulnerable functions. finding.reachable=true/false/unknown."
    status: pending
  - id: s3-vuln-sources-complete
    content: "Add VulnCheck KEV feed + Trivy DB. Update dedup to 6-source merge. Severity boost if CVE in KEV."
    status: pending
  - id: s3-sast-full
    content: "Full Node.js SAST (20+ rules) + AI/LLM security rules (15+ rules). Note: 3 LLM rules already shipped in Stage 1 (s1-llm-2): llm-prompt-injection, llm-missing-guardrail, llm-rag-unsanitized. Stage 3 adds the remaining ~12 AI/LLM patterns (insecure direct LLM call, unvalidated LLM output, secret leakage via LLM, insecure MCP configs, vector DB injection, eval of LLM output, SSRF via LLM URL, PII leak, rate-limit missing on LLM endpoints, LLM call without timeout/cost-cap, secrets in trace)."
    status: pending
  - id: s3-ai-multifile
    content: "Multi-file repo-level AI context: cross-file data flow, framework context (Express/Nest/Koa), config awareness (.env, nest-cli.json), previous-verdict feedback."
    status: pending
  - id: s3-ai-providers
    content: "Anthropic provider + self-hosted LLM provider (OpenAI-compatible endpoint). Provider failover."
    status: pending
  - id: s3-gitlab-bitbucket
    content: "GitLab + Bitbucket providers. Provider abstraction tightened."
    status: pending
  - id: s3-license-sbom-secrets
    content: "License scanner (SPDX ID + license compatibility matrix), SBOM exporter (CycloneDX 1.5 + SPDX 2.3 + embedded VEX statements), secrets scanner (gitleaks-style regex + entropy + provider patterns)."
    status: pending
  - id: s3-fix-prs
    content: "Automated Fix PR: creates branch, applies lockfile upgrade, opens PR with diff + policy summary + AI confidence. Per provider (GitHub/GitLab/Bitbucket)."
    status: pending
  - id: s3-jira
    content: "Jira integration: OAuth + project mapping + auto-create issue on blocking findings + sync status changes."
    status: pending
  - id: s3-monitor
    content: "Monitor mode: continuous scan of a project daily + on new advisory; notifications on change."
    status: pending
  - id: s3-notifications
    content: "Notifications: Slack (bot + channel), email (SMTP + templated), webhook (signed HMAC). Fan-out queue."
    status: pending
  - id: s3-pr-checks-baseline
    content: "PR checks: GitHub App / GitLab webhook; annotate PR with findings. PR Baseline Diff: compare PR findings vs base branch, only fail on NEW findings."
    status: pending
  - id: s3-reports
    content: "Reports: PDF (findings summary + executive summary + compliance mapping) + CSV export. Scheduled weekly optional."
    status: pending
  - id: s3-cli-full
    content: "Full CLI set: `test`, `monitor`, `fix`, `fix --apply`, `ignore`, `policy test`, `sbom`, `license`, `secrets`, `behavior`, `graph`, `report`, `pr-diff`, `auth`, `projects`."
    status: pending
  - id: s3-dashboard-minimal
    content: "Minimal read-only dashboard (React + Vite): login, projects list, scan list, findings table with AI verdict cards, SBOM download. No write operations yet."
    status: pending
  - id: s3-dogfood
    content: "Dogfood: run VulnCheck on VulnCheck's own codebase in CI. Fix findings or ignore with rationale."
    status: pending
isProject: false
---

# Stage 3 -- Beta (~1 month / 4-5 weeks)

Parent roadmap: [vulncheck_roadmap.plan.md](vulncheck_roadmap.plan.md). Master plan: [vulncheck_snyk-like_platform_41691b25.plan.md](vulncheck_snyk-like_platform_41691b25.plan.md). Requires Stages 1-2 complete.

## Goal

Ship the full Node.js AppSec story. By end of Stage 3:

- A Node.js shop can replace Snyk / GitHub Dependabot / Socket for the full workflow: scan, prioritize, fix PR, Jira ticket, Slack notify, PR check, weekly PDF report.
- The product demonstrates 3 clear wins over incumbents:
  1. **Socket-grade behavioral analysis** for malicious packages
  2. **Function-level reachability** for noise reduction (Snyk only does library-level)
  3. **AI Impact Analyzer with multi-file context** answering "does this matter in MY code"

## Scope IN

### Behavioral Package Analysis (master Phase E-bis)

- `internal/behavior/fetcher.go` -- pulls tarball from npm registry, verifies integrity
- `internal/behavior/sandbox.go` -- runs install + `node -r ./package/index.js` in gvisor / firecracker (preferred) or chroot+seccomp fallback with strict egress deny
- `internal/behavior/ast.go` -- AST walks to extract capabilities: `network`, `filesystem`, `shell`, `env`, `unsafe_eval`, `native_binding`, `obfuscated`, `install_hook_shell`
- `internal/behavior/diff.go` -- compares capabilities between versions; flags suspicious drift (e.g. added `network` + `env` between 1.2.3 and 1.2.4)
- `internal/behavior/profile.go` -- stores as `BehavioralProfile` model
- Cache: one profile per `(package, version)`; rebuild on registry republish
- River job `behavior_profile_job.go`

### Hallucinated / Slopsquat Detection (master Phase E-bis)

- `internal/supply/hallucination.go`:
  - During AI indexer pass, record every imported package name
  - For each imported name NOT in `package.json`: check npm registry existence
  - For each imported name IN `package.json`: check maintainer tenure, creation date, weekly downloads, known-bad feed
- Flags: `ghost_import` (imported but not in package.json), `slopsquat_candidate` (new + few downloads + name similar to popular), `abandoned_maintainer`
- Curated feed `data/known-bad-packages.json` (sourced manually + community)

### Call-Graph Reachability (master Phase E2-bis)

- `internal/callgraph/engine.go` -- dispatches per ecosystem
- `internal/callgraph/nodejs.go` -- uses ts-morph via child process OR a Go-native TS/JS parser (preferred: fork / wrap); extracts call graph
- `internal/callgraph/golang.go` -- uses `golang.org/x/tools/go/callgraph/cha` or `vta` for self-hosting scans (target ecosystems Python / .NET land in Stage 4)
- `internal/callgraph/entrypoints.go` -- identifies entry points per framework: Express routes, Nest controllers, Koa handlers, Next.js pages/API, Cobra CLI commands (for Go self-host)
- `internal/callgraph/reachability.go` -- path search from entry points to each vulnerable function
- Vulnerable function data: OSV `affected_functions` field when available; otherwise fallback to reachability at package level (current behavior)
- `finding.reachable` enum: `reachable` / `unreachable` / `unknown`
- Effective severity Stage 3: adds `reachability_adj` (-2 for `unreachable`, 0 for `reachable` or `unknown`)

### Vuln Sources Complete (master Phase E)

- `internal/sca/vulncheck.go` -- VulnCheck KEV feed sync (authenticated API), adds `in_kev=true` to findings; KEV forces severity >= high regardless of scope/AI
- `internal/sca/trivydb.go` -- Trivy DB git repo sync, covers OS + language advisories (used later for containers, introduced here)
- Dedup: 6-source merge with priority GHSA > NVD > VulnCheck > Trivy > OSV > npm audit

### Full Node.js SAST + AI/LLM Rules (master Phase E)

> **Stage 1 overlap note:** 3 LLM SAST rules are already implemented (`internal/sast/rules/llm/`): `llm-prompt-injection`, `llm-missing-guardrail`, `llm-rag-unsanitized` (s1-llm-2 ✅). Stage 3 extends `rules/semgrep/ai-llm.yaml` with the remaining ~12 patterns listed below — do not re-implement the 3 done rules.

- `rules/semgrep/node-express-nest.yaml` -- 20+ rules covering: auth bypasses, CORS misconfig, CSRF, cookie insecure flags, JWT weak algs, rate-limit missing, open redirect, XXE, insecure file upload, MIME-type spoof, HTTP response splitting, mass assignment, etc.
- `rules/semgrep/ai-llm.yaml` -- 15+ new rules (adds ~12 to the 3 from Stage 1):
  - Unsanitized user input concatenated into LLM prompt (prompt injection vector)
  - LLM output rendered without validation (XSS / code-exec via LLM)
  - API key in LLM client construction (prefer env)
  - Insecure MCP server configuration (world-readable socket, no auth)
  - Vector DB query with user-tainted text (injection)
  - Unbounded LLM context growth (cost + data leak)
  - Missing rate-limit on user-facing LLM endpoints
  - PII leak into LLM request (regex-based heuristic)
  - `eval` of LLM output
  - `exec` of LLM-generated shell commands
  - SSRF via LLM-suggested URL
  - Insecure deserialization of LLM structured output
  - LLM call without timeout
  - LLM call without cost cap
  - Secrets logged in LLM request/response trace
- All rules ship with positive + negative fixtures; CI runs `semgrep --test`

### Multi-file Repo-level AI Context (master Phase E2)

- `internal/ai/repo_context.go`:
  - Framework detection (Express, Nest, Koa, Fastify, Next.js) from package.json + file patterns
  - Config awareness: reads `.env.example`, `nest-cli.json`, `next.config.js`, `tsconfig.json`
  - Cross-file dataflow: which files call the vulnerable function (shallow 1-2 hops via imports)
  - Previous-verdict feedback: if a finding was previously dismissed with reason, inject reason into prompt
- Context bundle injected into prompt; context token budget capped (default 8K) with truncation heuristic

### AI Providers Expansion (master Phase E2)

- `internal/ai/anthropic.go` -- Anthropic API client, Claude models
- `internal/ai/selfhosted.go` -- OpenAI-compatible endpoint (for Ollama / vLLM / Azure OpenAI)
- Provider selection per org; failover chain configurable

### GitLab + Bitbucket (master Phase C2)

- `internal/git/gitlab.go`, `internal/git/bitbucket.go`
- PAT + OAuth for GitLab; app password + OAuth for Bitbucket
- Full parity with GitHub: clone, PR status, PR comment, webhook

### License / SBOM / Secrets (master Phase F2)

- `internal/license/scanner.go`: maps each dep to SPDX ID; compatibility matrix flags mixing (e.g. GPL in MIT project)
- `internal/sbom/cyclonedx.go`, `internal/sbom/spdx.go`: exporters; embed VEX not_affected statements for findings marked unreachable or dismissed
- `internal/secrets/scanner.go`: gitleaks-style regex + entropy + provider patterns (AWS, GCP, GitHub, Stripe, etc.); scans files in the clone

### Fix PR Automation (master Phase F3)

- `internal/fixpr/planner.go`: computes lockfile delta for a fix
- `internal/fixpr/github.go` / `gitlab.go` / `bitbucket.go`: branch + commit + PR with:
  - Title referencing CVE / GHSA
  - Body: finding summary, AI verdict, fix risk tier, upgrade chain, breaking notes, link to platform
  - Label (`vulncheck:safe` / `:moderate` / `:breaking`)
- Throttling: max N open fix PRs per repo (configurable, default 5)
- Re-open on new advisory if closed without merge

### Jira Integration (master Phase H)

- OAuth 2.0 + 3LO
- Project mapping table: `vulncheck_project_id -> jira_project_key`
- Auto-create issue on `decision=block` finding; sync status (finding resolved -> transition issue to "Done")
- Manual `vulncheck jira link <finding>` for ad-hoc

### Monitor Mode (master Phase I)

- Daily re-scan per project on schedule
- On new advisory affecting a monitored project: immediate rescan + notify
- `monitor` state on project: active/paused

### Notifications (master Phase J)

- Slack: bot token per org, channel routing by project, rich blocks
- Email: SMTP config per org, templated (MJML -> HTML)
- Webhook: HMAC-SHA256 signed, retries with exponential backoff

### PR Checks + PR Baseline Diff (master Phase K)

- GitHub App (install via marketplace stub; for beta: Personal App + manual install)
- GitLab / Bitbucket: webhook-based
- On PR open / push: scan PR branch, compare to base, compute diff
- PR check status `pass/fail` based on policy + diff (NEW findings only)
- Inline PR comment with critical NEW findings
- `PR Baseline Diff` snapshot persisted

### Reports (master Phase L)

- PDF via `gofpdf` or Chromium headless: cover, executive summary, findings by severity, SBOM summary, compliance mapping (OWASP ASVS, NIST SSDF, CIS)
- CSV export: one row per finding
- Weekly schedule per project, delivered by email or webhook

### Full CLI Set (master Phase N)

All subcommands wired:
- `vulncheck test` (scan)
- `vulncheck monitor` (enable / disable monitor)
- `vulncheck fix` (`--dry-run` / `--apply [--tier=safe|moderate|breaking]`)
- `vulncheck ignore <finding>`
- `vulncheck policy test --path .` (evaluates `.vulncheck` without scan)
- `vulncheck sbom --format=cyclonedx|spdx --output=file`
- `vulncheck license`
- `vulncheck secrets --path .`
- `vulncheck behavior <package>@<version>`
- `vulncheck graph --finding <id>` (show reachability path)
- `vulncheck report --format=pdf|csv`
- `vulncheck pr-diff --base <sha> --head <sha>`
- `vulncheck auth` / `projects` / `scans`

### Minimal Read-only Dashboard (master Phase O slice)

- React 18 + Vite + TailwindCSS + React Query
- Pages: Login, Projects list, Project detail (scans), Scan detail (findings table with AI verdict cards), SBOM download
- **No write ops in Beta.** Add in Stage 4.

## Scope OUT (deferred)

| Feature | Stage |
| --- | --- |
| Python / .NET / Go (as scanned target) | 4 |
| Rego policies | 4 |
| VEX consumer (upstream statements) | 4 |
| SLSA / Sigstore verification | 4 |
| DAST + IaC scanners (scaffold only) | 4 |
| Container / OCI image scanning | 4 |
| Full dashboard with writes (policies, fix PR triggers, user management) | 4 |
| IDE plugin | 4 |
| ASPM correlation | 5 |
| AI verdict feedback loop + prompt quality metrics | 5 |
| Runtime eBPF agent | 5 |

## Milestones

- **Week 1**: Call-graph engine + reachability (biggest risk item first).
- **Week 2**: Behavioral analysis + sandbox infrastructure + hallucination detection.
- **Week 3**: Fix PRs + Jira + Monitor + Notifications + PR checks + baseline diff. GitLab + Bitbucket providers.
- **Week 4**: Full SAST rule set + AI/LLM rules + multi-file AI context + remaining AI providers. License + SBOM + Secrets. Reports.
- **Week 5**: Full CLI. Minimal read-only dashboard. Dogfood + fixes + docs.

## Detailed Checklist

Each section above maps to a todos entry; treat the "Scope IN" bullets as the acceptance criteria for that todo. Cross-cutting must-haves:

- [ ] All new GORM models migrated: `BehavioralProfile`, `CallGraph`, `ReachabilityPath`, `FixPR`, `JiraLink`, `MonitorState`, `Notification`, `PRCheck`, `PRBaselineDiff`, `Report`, `License`, `SBOMExport`, `SecretFinding`, `HallucinatedImport`
- [ ] Every new River job: `behavior_profile_job`, `hallucination_feed_job`, `callgraph_job`, `fix_pr_job`, `jira_sync_job`, `monitor_rescan_job`, `notification_job`, `pr_check_job`, `baseline_diff_job`, `report_generation_job`
- [ ] Every external dependency pinned; Dockerfile multi-stage to keep image lean
- [ ] Every new CLI subcommand has `--help` and at least one integration test
- [ ] Dashboard: E2E test via Playwright for login -> projects -> scan -> findings
- [ ] Load test: 50 concurrent scans of 500-dep repos complete in <15 min total
- [ ] Self-scan (dogfood) in CI: every PR runs `vulncheck test` against VulnCheck's own repo; red if blocking finding introduced

## Definition of Done

- All Stage 1-2 DoD still passes
- All todos above marked complete
- Dogfood gate: platform's own code passes VulnCheck with no unresolved blocking findings
- README upgraded to full feature list with screenshots
- User documentation: policy file reference, CLI reference, webhook HMAC verification guide
- API reference regenerated (OpenAPI 3.1)
- Demo video: 3-minute "scan a real Node.js repo, see behavioral capture of a malicious package, merge a safe fix PR"

## Key Risks

- **Sandbox escape in behavioral analysis** -- treat as highest-severity risk. Run sandboxes in dedicated isolated hosts; no shared networking; egress deny by default; allow-list for registry-only access during install.
- **Tarball storage cost** -- cache only profile artifacts, not raw tarballs; TTL 30 days.
- **Call-graph performance on large repos** -- cap analysis at 100K LoC per scan in Stage 3; fallback to package-level reachability on timeout.
- **Call-graph accuracy for dynamic JS** -- well-known limitation (dynamic `require`, dependency injection). Label `reachable=unknown` on ambiguity; never use unknown to downgrade.
- **Fix PR churn / noise** -- rate limit per repo, coalesce multiple findings into one PR when upgrade chains overlap.
- **LLM provider costs at scale** -- aggregate reporting per org; hard-cap per scan; allow "AI off" mode for cost-sensitive customers (product degrades to Stage 1+2 non-AI features).
- **Jira OAuth flakiness** -- thorough retry + expiry refresh; status page on integrations dashboard.
- **Notification spam** -- per-org quiet hours + digest mode; dedup within a 15-minute window.
- **Dashboard scope creep** -- strictly read-only in Stage 3; writes land in Stage 4 to avoid multiplying auth / audit logic now.
- **GitLab / Bitbucket API drift** -- isolate provider specifics behind interface; add provider-specific integration tests gated by env.
- **SBOM VEX embedding correctness** -- use official CycloneDX / SPDX validators in CI.
- **Semgrep rule FP rate** -- track per-rule FP from day one; any rule > 20% FP gets disabled until reworked.
- **Dogfood breakage** -- can block CI unexpectedly; have a break-glass flag to let urgent PRs through with an audit log entry.
