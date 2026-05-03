---
name: VulnCheck Stage 1 -- MVP
overview: "~2-week MVP. Ships two parallel tracks: (A) Node.js scope-aware SCA scanner CLI + minimal API -- reads package-lock.json, queries OSV.dev + npm audit, computes noise-reduced effective severity; (B) LLM/AI Security module -- Python AI ecosystem SCA, prompt injection SAST rules, OWASP LLM Top 10 test runner, AI model SBOM, guardrails config auditor, RAG pipeline scanner, CLI commands, dashboard. Goal: demonstrate both traditional SCA and AI-specific security in a single demoable product."
todos:
  - id: s1-foundation
    content: "Foundation: go.mod, config, domain stubs, middleware, main.go, Makefile, migration 001, docker-compose"
    status: pending
  - id: s1-auth-min
    content: "Minimal auth: register + login + JWT only. No invites, no orgs as first-class, no RBAC. Just a User -> JWT flow"
    status: pending
  - id: s1-projects-min
    content: "Projects minimal: POST /projects (name, repo URL optional, ecosystem='npm'), GET /projects, GET /projects/:id"
    status: pending
  - id: s1-scan-orchestrator
    content: "Scan orchestrator + River job: POST /scans (upload lockfile OR path reference), scan_job.go queues and executes"
    status: pending
  - id: s1-npm-parser
    content: "npm lockfile parser v1/v2/v3 + dependency graph builder with scope (dependencies/devDependencies/optionalDependencies/peerDependencies)"
    status: pending
  - id: s1-vuln-sources
    content: "Vulnerability sources minimal: OSV.dev client (batched) + npm audit bulk endpoint client. Cache results in DB 24h."
    status: pending
  - id: s1-dedup-severity
    content: "Deduplicate findings across sources (package+version+CVE key). Compute effective severity Stage 1 only (base CVSS adjusted by scope: dev -> -2 levels, optional -> -1, prod -> 0)"
    status: pending
  - id: s1-findings-api
    content: "Finding storage + query API: GET /scans/:id, GET /scans/:id/findings with filters (severity, scope, package)"
    status: pending
  - id: s1-cli-test
    content: "CLI command `vulncheck test --path .`: detects package-lock.json, uploads or parses locally, calls scan API, prints pretty output + optional --json"
    status: pending
  - id: s1-e2e-smoke
    content: "End-to-end smoke test: docker-compose up -> seed a sample buggy Node.js repo -> `vulncheck test` -> verify output matches expected findings"
    status: pending
  - id: s1-readme
    content: "README with install / run / first scan instructions"
    status: pending
  - id: s1-llm-1
    content: "LLM SCA: Python AI ecosystem plugin (requirements.txt + pyproject.toml parser, curated AI package list, IsAIPackage flag). Extends pkg/ecosystem/ shared interface."
    status: completed
  - id: s1-llm-2
    content: "LLM SAST: Semgrep rules for prompt injection (llm-prompt-injection, llm-missing-guardrail, llm-rag-unsanitized). Go embed loader + RunLLMRulesOnFile()."
    status: completed
  - id: s1-llm-3
    content: "OWASP LLM Top 10 runner: OWASPRunner targeting OpenAI-compatible endpoint, tests for LLM01/02/06, httptest-verified, DB migration 007_llm_findings.sql."
    status: completed
  - id: s1-llm-4
    content: "AI Model SBOM: DetectModels() scans Python source for model ID references (OpenAI, Anthropic, HuggingFace), flags deprecated/known-vulnerable models."
    status: pending
  - id: s1-llm-5
    content: "Guardrails config auditor: ScanNeMoConfig() checks NeMo Guardrails YAML for missing PII output filter, jailbreak input check, NSFW filter."
    status: pending
  - id: s1-llm-6
    content: "RAG security scanner: ScanPythonRAG() detects unauthenticated ChromaDB access and unsanitized retrieval injection into prompts."
    status: pending
  - id: s1-llm-7
    content: "CLI: vulncheck llm test / guardrails / rag subcommands via Cobra. Wires all llmsec packages into CLI."
    status: pending
  - id: s1-llm-8
    content: "Web dashboard: /llm-security React page with OWASP scorecard, model inventory, summary cards. TanStack Query + Tailwind."
    status: pending
isProject: false
---

# Stage 1 -- MVP (5-7 days)

Parent roadmap: [vulncheck_roadmap.plan.md](vulncheck_roadmap.plan.md). Master plan: [vulncheck_snyk-like_platform_41691b25.plan.md](vulncheck_snyk-like_platform_41691b25.plan.md).

## Goal

Ship a working CLI + minimal API that, given a Node.js `package-lock.json`, produces a findings report where **dev-only** and **optional** dependency vulnerabilities are automatically downgraded in severity. A repo that returns 150 "high" CVEs from `npm audit` should return ~15 "high" from `vulncheck test` -- that's the demo.

**Success = on a real public Node.js repo (e.g. a moderate-size Express app), after Stage 1 ships:**

- `vulncheck test --path .` completes in under 30 seconds for 500+ transitive deps
- Output clearly distinguishes production-affecting CVEs from dev-only
- JSON output is valid and stable (schema frozen at v1)
- Exit code is 0 if no production findings at severity >= HIGH, else non-zero
- `docker-compose up` brings up Postgres + server, and the CLI talks to it via HTTP

## Scope IN

Only these components from the master plan, and only minimal slices:

- **Foundation** (master Phase A): [Phase A detailed](vulncheck_snyk-like_platform_41691b25.plan.md)
  - `go.mod` with minimal deps: gin, gorm, postgres driver, river, zerolog, viper, cobra, bcrypt, golang-jwt
  - `internal/config/` -- viper-based env + file
  - `internal/domain/` -- only models needed: User, Project, Scan, Finding
  - `internal/middleware/` -- JWT auth, request logging, recover
  - `cmd/server/main.go` -- gin server wiring
  - `Makefile` -- build, run, test, docker-up
  - `migrations/001_initial.sql`
  - `docker-compose.yml` with Postgres 16
- **Auth** (minimal slice of master Phase B)
  - `POST /auth/register`, `POST /auth/login`
  - bcrypt hashing, JWT issuance (24h)
  - No invite flow. No orgs as first-class table; keep `user_id` as the only scoping key for Stage 1.
- **Projects** (minimal slice of master Phase C)
  - `internal/project/handler.go` + `service.go` + `repo.go`
  - CRUD: create, list, get (no update/delete for MVP)
- **Scan orchestration** (minimal slice of master Phase D)
  - `internal/scan/handler.go` + `service.go` + `repo.go`
  - `POST /scans` -- accepts `{"project_id": "...", "lockfile_content": "<base64>"}` OR `{"project_id": "...", "lockfile_path": "/tmp/..."}`
  - River job `internal/jobs/scan_job.go` calls the SCA scanner
- **SCA scanner** (minimal slice of master Phase E)
  - `internal/sca/scanner.go` -- orchestrates parser -> sources -> dedup -> severity
  - `internal/sca/parser.go` -- delegates to ecosystem-specific parsers
  - `internal/sca/graph.go` -- builds dep graph with scope tags
  - `internal/sca/context.go` -- minimal: attaches scope metadata to each finding
  - `internal/sca/effective_severity.go` -- **Stage 1 formula only** (see below)
  - `internal/sca/npm.go` -- npm audit bulk API client
  - `internal/sca/osv.go` -- OSV.dev batch API client
  - `internal/sca/dedup.go` -- merge findings on `(package, version, vuln_id)` tuple, prefer most complete record
  - `pkg/ecosystem/npm/parser.go` -- v1/v2/v3 lockfile parsing
  - `pkg/ecosystem/ecosystem.go` + `pkg/ecosystem/registry.go` -- interface + npm registration
- **CLI** (minimal slice of master Phase N)
  - `cmd/cli/main.go` with cobra
  - `vulncheck auth login`, `vulncheck auth logout`
  - `vulncheck test --path . [--json] [--fail-on=high]`
  - Config at `~/.vulncheck/config.yaml` with server URL + token

## Scope OUT (explicitly deferred to later stages)

| Feature | Stage |
| --- | --- |
| Git integration (GitHub/GitLab/Bitbucket) | 2 |
| AI Impact Analyzer | 2 |
| Remediation engine (fix plans, upgrade paths) | 2 |
| Dual policies (DB + `.vulncheck` file) | 2 |
| SAST for general app code (Semgrep non-LLM rules) | 2 (minimal) / 3 (full) |
| Supply chain detection (hooks / typosquat / new package) | 2 |
| NVD, GHSA, VulnCheck, Trivy DB sources | 2 (NVD + GHSA) / 3 (rest) |
| License, SBOM, Secrets (non-AI) | 3 |
| Fix PRs, Monitor, Jira, Notifications, PR Checks, Reports | 3 |
| Behavioral package analysis, Call-graph reachability | 3 |
| Rego, VEX consumer, SLSA/Sigstore | 4 |
| Container / OCI scanning | 4 |
| Multi-ecosystem for SCA (Python/.NET/Go as general targets) | 4 |
| Full web dashboard (non-LLM pages) | 3 minimal / 4 full |
| IDE plugin | 4 |
| ASPM, AI feedback loop, runtime eBPF agent | 5 |

**Moved into Stage 1 scope (LLM SecOps track):**
- Python AI ecosystem SCA (`pkg/ecosystem/python/`) — originally Stage 4, pulled in for LLM supply chain
- Prompt injection SAST rules (`internal/sast/rules/llm/`) — originally Stage 2, pulled in for OWASP LLM01
- LLM Security web dashboard page (`/llm-security`) — originally Stage 3+, pulled in for AI security posture
- AI Model SBOM (LLM-specific, not general SBOM) — originally Stage 3, pulled in for model inventory

## LLM SecOps Module (merged into Stage 1)

Derived from OWASP LLM Top 10 and LLM SecOps Specialist job requirements. Eight tasks shipped alongside the npm SCA track, sharing the same `pkg/ecosystem/` plugin interface and `cmd/cli/` cobra structure.

| Task | Feature | Status | Files |
| --- | --- | --- | --- |
| s1-llm-1 | Python AI Ecosystem SCA | ✅ done | `pkg/ecosystem/python/` |
| s1-llm-2 | Prompt Injection SAST Rules | ✅ done | `internal/sast/rules/llm/` |
| s1-llm-3 | OWASP LLM Top 10 Runner | ✅ done | `internal/llmsec/`, `migrations/007_llm_findings.sql` |
| s1-llm-4 | AI Model SBOM | pending | `internal/llmsec/model_sbom.go` |
| s1-llm-5 | Guardrails Config Auditor | pending | `internal/llmsec/guardrails/` |
| s1-llm-6 | RAG Pipeline Security Scanner | pending | `internal/llmsec/rag/` |
| s1-llm-7 | CLI: `vulncheck llm` commands | pending | `cmd/cli/llm.go` |
| s1-llm-8 | LLM Security Dashboard | pending | `web/src/pages/llm-security/` |

**Shared infrastructure (no duplication):**
- `pkg/ecosystem/ecosystem.go` + `registry.go` — created by s1-llm-1, will also be used by npm parser (s1-npm-parser)
- `pkg/ecosystem/registry.go` — single plugin registry for both npm and python-ai
- `cmd/cli/` — shared cobra root, `vulncheck test` (SCA) and `vulncheck llm` (LLM) are sibling subcommands

## Effective Severity -- Stage 1 Formula

We only implement Stage 1 of the formula described in master plan's Phase E Context Analyzer section:

```
base_severity = max(osv.cvss_base, npm_audit.severity_as_cvss)
scope_adjustment:
  prod         -> 0
  optional     -> -1 CVSS level (critical->high, high->medium, etc.)
  peer         -> 0  (peer deps are caller-resolved; treat as prod for safety)
  dev          -> -2 CVSS levels (critical->medium, high->low)
  test-only    -> -2 CVSS levels  (detected via heuristic: package name contains "test"/"jest"/"mocha" OR path under dev+no prod import)

effective_severity = clamp(low, base_severity + scope_adjustment)
```

Reachability heuristic, call-graph, AI -- all deferred. A package is "prod" if it appears in the prod dep graph AT ALL, transitively; no reachability gate.

## Milestones (day-by-day for a solo dev)

- **Day 1** -- Foundation: Go project scaffold, docker-compose, first migration, basic gin server returning `/healthz`. Tests green.
- **Day 2** -- Auth: register/login working end-to-end with a curl script. Projects CRUD (create/list/get) with JWT.
- **Day 3** -- Scan orchestration: `POST /scans` accepting lockfile content, River worker picks it up, persists a `Scan` row with status transitions.
- **Day 4** -- npm parser + dep graph + scope. Unit tests with fixtures for lockfile v1, v2, v3. Graph has `prod`/`dev`/`optional`/`peer` tags.
- **Day 5** -- OSV.dev + npm audit clients with response caching. Scanner service wires parser -> sources -> dedup -> severity. Findings persist.
- **Day 6** -- CLI: `vulncheck auth login`, `vulncheck test --path .`, pretty output (color by severity, grouped by scope) + `--json`. Exit code wiring.
- **Day 7** -- End-to-end smoke test on a real Node.js repo. README. Small fixes and polish.

## Detailed Checklist

### Foundation

- [ ] `go mod init github.com/artfend/vulncheck`
- [ ] Add deps: `gin`, `gorm.io/gorm`, `gorm.io/driver/postgres`, `github.com/riverqueue/river`, `github.com/riverqueue/river/rivertype`, `github.com/rs/zerolog`, `github.com/spf13/viper`, `github.com/spf13/cobra`, `golang.org/x/crypto/bcrypt`, `github.com/golang-jwt/jwt/v5`
- [ ] `internal/config/config.go` reads env + `config.yaml`
- [ ] `internal/domain/*.go` with GORM models: User, Project, Scan, Finding, DependencyGraph
- [ ] `internal/middleware/auth.go` (JWT), `logging.go`, `recover.go`
- [ ] `cmd/server/main.go` wires gin, middleware, handlers
- [ ] `migrations/001_initial.sql`
- [ ] `docker-compose.yml` with Postgres 16
- [ ] `Makefile` targets: `build`, `run`, `test`, `up`, `down`, `migrate`
- [ ] `/healthz` returns 200

### Auth minimal

- [ ] `POST /auth/register` (email, password) -> 201 with user id
- [ ] `POST /auth/login` -> 200 with JWT
- [ ] JWT middleware on `/api/*` routes
- [ ] Passwords hashed with bcrypt cost 12
- [ ] JWT secret from env, expiry 24h
- [ ] Unit tests for auth service

### Projects minimal

- [ ] `POST /api/projects` (name, ecosystem='npm', repo_url optional)
- [ ] `GET /api/projects` -- list user's projects
- [ ] `GET /api/projects/:id` -- get by id with ownership check
- [ ] Returns 403 on cross-user access
- [ ] Integration test via httptest

### Scan orchestrator

- [ ] `POST /api/scans` with body `{project_id, lockfile_content (base64) OR lockfile_path}`
- [ ] Returns 202 with `scan_id` and `status=queued`
- [ ] River job `scan_job.go` picks up queued scans, transitions `queued -> running -> completed/failed`
- [ ] `GET /api/scans/:id` returns scan status + counts
- [ ] Scan timeout 5 minutes -> marks failed

### npm parser + graph

- [ ] `pkg/ecosystem/npm/parser.go` with `ParseLockfile(content []byte) (*DepGraph, error)`
- [ ] Handles v1 (legacy `dependencies` tree), v2 (packages map + dependencies tree), v3 (packages map only)
- [ ] Each node has: name, version, resolved, integrity, scope{prod,dev,optional,peer} bitmask, parents, children
- [ ] Scope bitmask union'd across all paths (a dep can be both prod and dev)
- [ ] Fixtures in `testdata/npm/lockfile-v1.json`, `v2.json`, `v3.json` with known outputs
- [ ] Unit tests cover cycles, workspaces (flatten), dedup paths

### Vulnerability sources

- [ ] `internal/sca/osv.go`: `Query(packages []PkgRef) []Finding` -- uses `https://api.osv.dev/v1/querybatch`, chunk at 1000/request
- [ ] `internal/sca/npm.go`: `Query(packages []PkgRef) []Finding` -- uses `https://registry.npmjs.org/-/npm/v1/security/advisories/bulk`
- [ ] Both: retry with exponential backoff, timeout 30s per batch
- [ ] Cache layer keyed on `(package, version)` with 24h TTL in a `vuln_cache` table
- [ ] Findings have consistent internal shape: `package, version, vuln_id, cvss_base, severity, title, description, fixed_in[], source`

### Dedup + effective severity

- [ ] `internal/sca/dedup.go`: group by `(package, version, normalized_vuln_id)`, normalize aliases (CVE, GHSA, OSV IDs). Merge fields, prefer highest CVSS, union fixed_in.
- [ ] `internal/sca/effective_severity.go`: implements Stage 1 formula (scope adjustment only). Unit tests with matrix of scope x severity.
- [ ] Scanner pipeline: parse -> context (attach scope) -> query sources -> dedup -> severity -> persist

### Findings API

- [ ] `GET /api/scans/:id/findings` with query params: `severity`, `scope`, `package`, `pagination`
- [ ] Response includes: raw severity, effective severity, scope, paths (top 3), fixed_in suggestions
- [ ] Integration test with seeded scan + 20 findings

### CLI

- [ ] `cmd/cli/main.go` with cobra subcommands
- [ ] `vulncheck auth login --email ... --password ...` stores token in `~/.vulncheck/config.yaml`
- [ ] `vulncheck auth logout` clears token
- [ ] `vulncheck test --path .` detects `package-lock.json`, submits scan, polls until done, prints:
  - Header: total findings, by effective severity
  - Groups: PROD / OPTIONAL / PEER / DEV
  - Within group: table of finding (package, version, vuln, effective sev, raw sev, fixed_in)
- [ ] `--json` emits a stable JSON schema version 1
- [ ] `--fail-on=<level>` sets exit code 1 if any prod finding >= level
- [ ] Colors from `fatih/color`; NO color if not tty

### End-to-end smoke

- [ ] Pick a public repo (e.g. a pinned commit of an express-based boilerplate) with known CVEs
- [ ] `docker-compose up` -> `vulncheck auth register` -> `test --path ./fixtures/real-repo` succeeds
- [ ] Output manually verified: known dev-only CVEs are at lower severity vs raw npm audit

### README

- [ ] `README.md` with: what it is, 1-minute pitch, quickstart (docker-compose + curl + CLI), architecture diagram (from master plan), FAQ
- [ ] Link to master plan for roadmap

## Definition of Done

All boxes above checked. All tests green. `make test` clean. `docker-compose up` + smoke test passes.

## Key Risks

- **npm audit API undocumented changes** -- mitigate by testing against fixtures captured from real responses; degrade gracefully if unavailable (just use OSV).
- **OSV.dev rate limits** -- mitigate via 24h caching + request batching.
- **Lockfile v3 edge cases** (workspaces, git deps, tarball deps) -- mitigate by fixture-driven tests; tolerate unparseable entries by logging and continuing.
- **Scope bitmask correctness** -- a package reached via both prod and dev paths must be treated as prod. Unit test this explicitly.
- **JSON schema drift** -- freeze v1 schema in `docs/cli-json-schema-v1.json`. Any change in Stage 2+ requires a new version field.
- **Time pressure** -- if Day 5 slips, cut the npm audit source (keep only OSV). If Day 6 slips, cut `--fail-on` and `--json`.
