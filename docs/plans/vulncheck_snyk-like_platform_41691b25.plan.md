---
name: VulnCheck Snyk-Like Platform
overview: Build an AI-powered context-aware security platform that solves the core problems of npm audit (99% false positives, context blindness, dangerous auto-fix) and Snyk (alert fatigue, slow scans, high cost), while also closing the 2026 competitive gap vs Socket.dev (behavioral package analysis), Endor Labs (function-level call-graph reachability), Aikido (consolidated all-in-one), and ASPM platforms (attack story correlation). Multi-ecosystem support (progressive delivery -- Node.js first, then Python, then .NET + Go). Killer features -- (1) AI Impact Analyzer reads user's actual source code, finds import/call sites of vulnerable packages, and produces a verdict; (2) Behavioral Package Analysis detects malicious packages pre-CVE via capability diff (Socket-style); (3) Function-level Call-Graph Reachability gives deterministic 95%+ noise reduction without LLM cost (Endor-style). 6 vulnerability data sources (npm/OSV/NVD/GHSA/VulnCheck/Trivy) with cross-validation and exploit intelligence (EPSS, KEV). 100+ custom Semgrep SAST rules across 4 ecosystems: Node.js/Express/NestJS, Python/Django/Flask/FastAPI, .NET/ASP.NET Core, Go/Gin/GORM, PLUS dedicated AI/LLM security rule pack (prompt injection, unsafe LLM output handling, MCP server risks, insecure LangChain configs). 3-stage effective severity (scope + AI verdict + exploit intel), supply chain detection per ecosystem, safe remediation, incremental scanning. Full SCA with transitive dep graph, SAST concrete + DAST/IaC scaffolded. Container / OCI image scanning (OS packages, base image analysis, Dockerfile best practices, Sigstore/SLSA verification). License compliance (SPDX expression support, org-level allow/deny/warn policies). SBOM generation (SPDX 2.3 + CycloneDX 1.5 with VEX overlay) + VEX consumption (ingest upstream OpenVEX/CSAF statements to auto-dismiss not-affected CVEs). Automated Fix PRs (auto-create PRs for safe dependency upgrades with safety check). Secrets scanning (100+ patterns + entropy analysis + git history). Dual policy system + Rego/OPA policy-as-code for enterprise gates. PR baseline diff (show only NEW vulns in a PR, not entire inventory). ASPM attack-story correlation across SCA/SAST/IaC/secrets. AI verdict feedback loop (learn from user corrections). Hallucinated / slopsquatted package detection. IDE plugin (VSCode/JetBrains via LSP) for inline hints. Monitoring, Jira sync, notifications, PR checks, compliance reports. Go/Gin/GORM/PostgreSQL backend, React frontend, CLI, LSP server. LLM backend: OpenAI/Anthropic/self-hosted.
todos:
  - id: foundation
    content: "Phase A: Foundation -- go.mod, config, domain, middleware, main.go, Makefile, migration, docker-compose"
    status: pending
  - id: auth
    content: "Phase B: Auth -- org/user models, JWT, bcrypt, register/login/invite"
    status: pending
  - id: projects
    content: "Phase C: Projects -- CRUD scoped to org"
    status: pending
  - id: git-integration
    content: "Phase C2: Git Integration -- GitCredential model (encrypted), safe cloner (shallow/no-hooks/timeout), GitHub/GitLab/Bitbucket provider APIs (ownership verification, repo listing), file extractor, webhook receivers (HMAC verified), rate limiter, audit log, security checks"
    status: pending
  - id: scan-engine
    content: "Phase D: Scan Engine -- scan/finding models (effectiveSeverity, reachability, scope, fixRiskLevel), orchestrator, River async job, vuln cache with incremental support"
    status: pending
  - id: sca-depgraph
    content: "Phase E: Node.js Ecosystem + SAST Engine -- npm lockfile parser (v1/v2/v3) with scope extraction, dep graph builder, context analysis, effective severity (3-stage), supply chain (npm hooks/typosquat), 6 vuln source clients, dedup, SAST engine (Semgrep wrapper), Node.js/Express/NestJS SAST rules (20+)"
    status: pending
  - id: ai-impact
    content: "Phase E2: AI Impact Analyzer -- code indexer (find imports/calls of vuln packages), snippet extractor (data-flow context), LLM client (OpenAI/Anthropic/self-hosted), prompt templates per vuln type, verdict parser, analyzer orchestrator with caching+batching, async River job"
    status: pending
  - id: remediation
    content: "Phase F: Remediation Engine -- upgrade path analysis, fix risk level classification (patch/minor/major/no_fix), npm version lookup, safe-fix logic"
    status: pending
  - id: license-sbom-secrets
    content: "Phase F2: License Compliance + SBOM + Secrets -- LicensePolicy/LicenseFinding models, SPDX expression parser, license detection from lockfiles + registry APIs, license policy evaluation (allow/deny/warn), SBOM generation (SPDX 2.3 + CycloneDX 1.5), VEX overlay, Secrets scanner (100+ regex patterns + Shannon entropy), git history secrets scan, validator for false positive reduction"
    status: pending
  - id: fix-prs
    content: "Phase F3: Automated Fix PRs -- FixPR model, lockfile updater (npm/pip/go/.NET), safety check (re-scan after upgrade), branch creation + PR via GitHub/GitLab/Bitbucket APIs, async River job, merge status monitoring"
    status: pending
  - id: dual-policies
    content: "Phase G: Dual Policy System -- DB policies + .vulncheck file parser, merge logic, import/export"
    status: pending
  - id: jira-integration
    content: "Phase H: Jira Integration -- Jira REST client, auto-ticketing, two-way sync job, dedup"
    status: pending
  - id: monitor-mode
    content: "Phase I: Monitor Mode -- snapshot storage, periodic River job, diff logic"
    status: pending
  - id: python-ecosystem
    content: "Phase E3: Python Ecosystem -- pip/pipenv/poetry lockfile parsers, PyPI advisory client, Python/Django/Flask/FastAPI SAST rules (20+), PyPI supply chain detection (setup.py hooks, typosquat), AI code indexer for Python imports"
    status: pending
  - id: notifications
    content: "Phase J: Notifications -- email + Slack alerting for new vulns, monitor alerts, scan completions"
    status: pending
  - id: pr-checks
    content: "Phase K: GitHub PR Checks -- post scan status + finding comments on pull requests"
    status: pending
  - id: reporting
    content: "Phase L: Reporting -- PDF/CSV compliance report generation"
    status: pending
  - id: dotnet-go-ecosystem
    content: "Phase E4: .NET + Go Ecosystems -- NuGet parser (packages.lock.json/.csproj), go.mod/go.sum parser, ASP.NET Core SAST rules (20+), Go/Gin/GORM SAST rules (25+), NuGet/Go supply chain, AI code indexer for C#/Go imports"
    status: pending
  - id: scaffold-scanners
    content: "Phase M: Scaffold DAST (ZAP) + IaC (Checkov/Trivy misconfig) scanner interfaces"
    status: pending
  - id: cli
    content: "Phase N: CLI -- Cobra commands (test with scope/reachability output, fix with risk levels, supply-chain, monitor, ignore, auth, policy)"
    status: pending
  - id: web-dashboard
    content: "Phase O: Web Dashboard -- React frontend (actionable-first default view, scope/reachability filters, noise metrics, fix risk badges)"
    status: pending
  - id: behavioral-analysis
    content: "Phase E-bis: Behavioral Package Analysis -- tarball fetcher, AST capability extractor (network/filesystem/shell/env/unsafe-eval), capability diff between versions, malicious-pattern detector, hallucinated/slopsquat package detection, quarantine ignore-policy auto-creation"
    status: pending
  - id: callgraph-reachability
    content: "Phase E2-bis: Function-Level Call-Graph Reachability -- deterministic call-graph engine per ecosystem (ts-morph/TypeScript, jedi/Python, golang.org/x/tools, Roslyn/.NET), entry-point discovery, path tracing from entrypoint to vulnerable function, 95%+ deterministic noise reduction before LLM kicks in"
    status: pending
  - id: ai-llm-sast
    content: "Phase E extension: AI/LLM Security SAST Rules -- rules/semgrep/ai-llm.yaml (prompt injection, unsafe LLM output handling, LLM output to eval/SQL/HTML, hardcoded system prompts with secrets, MCP server unbounded shell exec, insecure LangChain/LlamaIndex agent configs, missing LLM rate limits, tokenizer DoS)"
    status: pending
  - id: container-scanning
    content: "Phase M2: Container / OCI Image Scanning -- image pull + layer extraction (go-containerregistry), OS package SCA (apt/apk/rpm vs Trivy DB), Dockerfile SAST rules (root user, ADD http, latest tag, missing USER), base image analysis + minimal-image recommendations (distroless/chainguard), Sigstore/cosign signature verification, SLSA provenance check"
    status: pending
  - id: vex-consumer
    content: "Phase Q: VEX Consumption -- ingest upstream OpenVEX JSON + CSAF from GitHub releases / vendor security pages, apply as implicit ignore-policy (source=upstream_vex), track provenance of each dismissal, refresh job"
    status: pending
  - id: rego-policy
    content: "Phase Q: Policy-as-Code (Rego/OPA) -- opt-in Rego engine via open-policy-agent/opa, composite gates (EPSS threshold + fix-risk + license + VEX status), org can write custom .rego files, policy evaluation hook in scan pipeline"
    status: pending
  - id: slsa-sigstore
    content: "Phase Q: SLSA Provenance + Sigstore Verification -- for each dep verify SLSA provenance attestation on registry, Sigstore Rekor lookup for signed artifacts, supply-chain signal types missing_provenance / provenance_mismatch / unsigned_artifact"
    status: pending
  - id: pr-baseline-diff
    content: "Phase K extension: PR Baseline Diff -- compare PR-head scan against last defaultBranch scan, post only NEW findings as blocking + RESOLVED findings as green, inline lockfile diff in PR comment"
    status: pending
  - id: ai-feedback
    content: "Phase O extension: AI Verdict Feedback Loop -- AIVerdictFeedback model, UI correction form (wrong impact / wrong explanation / missed usage), feedback aggregator, prompt-quality metrics dashboard, export anonymized eval set"
    status: pending
  - id: aspm-correlation
    content: "Phase P2: ASPM Attack-Story Correlation -- post-scan correlator groups findings across SCA/SAST/Secrets/IaC by shared entity (file, package, service, DB), builds attack-graph edges, surfaces prioritized attack stories in dashboard"
    status: pending
  - id: multifile-ai-context
    content: "Phase E2 extension: Multi-file Repository-level AI Context -- framework detector (Express/NestJS/Django/Flask/FastAPI/ASP.NET Core/Gin), entry-point + route extractor, DI-binding resolver, repo-level context builder injected into LLM prompts, token-budget manager with summarization"
    status: pending
  - id: ide-plugin
    content: "Phase P: IDE Plugin (LSP) -- Language Server Protocol server in Go, VSCode extension, JetBrains plugin, inline hover hints on package versions in package.json/requirements.txt/go.mod/.csproj, one-click fix PR trigger from IDE, live AI verdict preview"
    status: pending
  - id: runtime-reachability
    content: "Phase R (future, opt-in): Runtime Reachability Agent -- light-weight Go agent with eBPF (cilium/ebpf), pushes loaded module list per process + function-level syscall sampling, findings get runtimeObserved flag, drives prioritization in prod"
    status: pending
isProject: false
---

# Feature Plan: ArtFend VulnCheck (Snyk-Like)

## Competitive Differentiators (vs npm audit / Snyk)

Research into real developer pain points with npm audit and Snyk revealed systemic problems that ArtFend VulnCheck must solve from day one. These are not nice-to-haves -- they are the core reason to build a new tool.

### Problem 1: Context Blindness (npm audit treats all deps equally)
npm audit does not distinguish production, development, build-time, or optional dependencies. A ReDoS vulnerability in `browserslist` (a build-only CSS config parser) is reported with the same "High" severity as an RCE in a runtime HTTP library. Dan Abramov's "npm audit: Broken by Design" (overreacted.io) documented that 100% of npm audit findings in Create React App were false positives because all flagged packages were build-time-only tools that never reach production.

**Our answer: Dependency Scope Awareness.** Every dependency in the graph is tagged with a `scope` (production / development / optional / peer). Findings include an `effectiveSeverity` that is automatically downgraded when the vulnerable package is dev/build-only. By default, CLI and dashboard show only production-scope findings; dev-scope findings require an explicit `--include-dev` flag. This eliminates 80-90% of the noise developers currently experience.

### Problem 2: No Reachability Analysis (flagging unreachable code paths)
Both tools report a CVE if the vulnerable package version exists anywhere in the tree, even if the vulnerable function is never called. Example: a "Denial of Service via crafted CSS selector" in `css-what` is flagged as "High" even though the parser only runs at build time on developer-controlled SVG files.

**Our answer: Reachability Classification.** Each finding gets a `reachability` field: `confirmed` (exploitable code path exists in prod), `potentially_reachable` (package is in prod scope but call path not verified), `unreachable` (dev-only, or the vulnerable function is not in the call graph). Reachability is determined by scope analysis (Phase E) and optionally by static call-graph analysis (future SAST integration). Unreachable findings are shown as informational, not as blockers.

### Problem 3: Dangerous Auto-Fix (npm audit fix --force breaks things)
`npm audit fix --force` can downgrade dependencies to 3-year-old versions with real vulnerabilities, apply major version bumps that break APIs, and cause cascading dependency conflicts. Developers apply it out of desperation and break production.

**Our answer: Safe Remediation with Risk Assessment.** The remediation engine classifies each suggested fix as `patch` (safe, semver-compatible), `minor` (likely safe, minor version bump), `major` (breaking change risk -- requires review), or `no_fix` (no available fix). The CLI refuses to auto-apply `major` fixes without an explicit `--allow-major` flag. Each upgrade path includes the semver delta and a breaking-change risk score.

### Problem 4: Supply Chain Attacks Are Invisible
Neither npm audit nor Snyk detect typosquatting, malicious postinstall scripts, obfuscated code, or newly published packages with suspicious patterns. They only check known CVEs in advisory databases.

**Our answer: Supply Chain Signals (Phase E-bis).** The SCA scanner includes a lightweight supply chain analyzer that flags: packages with `postinstall`/`preinstall` hooks that execute shell commands, packages published less than 7 days ago (new package risk), packages where the maintainer changed recently, packages with names similar to popular packages (typosquat detection via Levenshtein distance), and packages with obfuscated or minified source in the tarball. These are surfaced as `supply_chain` type findings with their own severity.

### Problem 5: Alert Fatigue / Vulnerability Noise
Snyk users report hundreds of findings per scan, most of which are not actionable. Developers stop reading alerts entirely -- "the boy who cried wolf" effect. When a real critical vulnerability appears, it gets buried in noise.

**Our answer: Signal-over-Noise by Default.** Findings are sorted by `effectiveSeverity` (context-adjusted, not raw CVSS). The default view is "Actionable Only" -- production-scope, reachable, with available fixes. Noise metrics are tracked: if a project consistently generates >50 findings with <5% action rate, the system suggests tightening scope filters or adding policies. The CLI exit code is based on effective severity, not raw severity.

### Problem 6: Slow CI/CD Scans (Snyk on large repos)
Snyk scans can take minutes on medium-to-large repositories, blocking CI/CD pipelines.

**Our answer: Incremental Scanning + Vuln Cache.** The vuln cache (Phase D) stores known CVE data with TTLs, avoiding redundant API calls. Incremental scan mode compares the current lockfile hash against the previous scan -- if the dependency tree hasn't changed, only new CVEs published since the last scan are checked. This reduces scan time from minutes to seconds for unchanged lockfiles.

### Problem 7: No Understanding of Actual Impact on YOUR Code
Both npm audit and Snyk report "CVE exists in your dependency tree" but never answer the real question: **"Does this vulnerability actually affect MY application?"** A CVE might say "prototype pollution via `lodash.merge` when called with user-controlled input" -- but if the user's code never calls `lodash.merge`, or only calls it with hardcoded config objects, the vulnerability is irrelevant. Today, developers must manually trace every finding through their source code to determine if it matters. This is the single biggest waste of time in vulnerability management.

**Our answer: AI Impact Analyzer (Phase E2).** After the SCA scanner finds vulnerabilities, the AI Impact Analyzer:

1. **Indexes the user's source code** to find all import/require/usage sites of the vulnerable package.
2. **Extracts the relevant code snippets** -- the specific call sites where the vulnerable function is used, including the data flow (where does the input come from?).
3. **Sends the CVE description + the user's code snippets to an LLM** (OpenAI / Anthropic / self-hosted) with a structured prompt: "This is the vulnerability. This is how the user's code uses the library. Assess the real impact."
4. **Generates a human-readable AI verdict** per finding:
   - **Impact level**: `critical_impact` / `moderate_impact` / `low_impact` / `no_impact`
   - **Explanation**: "The vulnerable function `lodash.merge` is called at `src/api/handler.js:42` with `req.body` (user-controlled HTTP input), making prototype pollution directly exploitable in production."
   - Or: "The package `css-what` is only imported by the build tool `svgo` and is never called from your application code. This vulnerability has no impact on your application."
   - **Usage locations**: exact file:line references where the package is used
5. **Overrides the reachability classification**: AI verdict replaces the scope-based heuristic with actual code analysis. A finding with `no_impact` AI verdict is downgraded regardless of scope.

This is the **killer feature** -- the question every developer asks ("is this actually dangerous for ME?") is answered automatically by AI instead of requiring hours of manual code review.

### Problem 8: Malicious Packages Published Before Any CVE Exists

Socket.dev's 2024-2026 research shows that **99.8% of new malware originates from npm** and that 454k+ malicious packages were published in 2025 alone (e.g. Axios supply-chain compromise, March 2026 -- malicious dependency injected into a popular package, executed a RAT during `npm install` in a ~2-second breach window). Traditional SCA relies on CVE databases, which requires **post-hoc disclosure** -- by the time a CVE exists, thousands of developers have already been compromised. Snyk, npm audit, GHSA all miss this class entirely.

**Our answer: Behavioral Package Analysis (Phase E-bis).** Before trusting a package version, the scanner fetches the tarball, parses the AST, and builds a **capability profile**: does this package make network calls, write to the filesystem, spawn shell processes, read environment variables, use `eval`/`Function()`/`vm`, contain obfuscated strings, access native bindings? The profile is compared against the previous version's profile. A `capability_expansion` finding fires when a minor/patch bump adds capabilities that the library historically never had (e.g. `lodash` suddenly gains `network` + `filesystem` capabilities). Install-time scripts (`preinstall`/`postinstall`) are extracted and semantically analyzed, not just regex-matched. This catches malicious packages **before any CVE is assigned** -- the same class of defense that Socket.dev built its reputation on.

### Problem 9: Hallucinated / Slopsquatted Package Names

A new 2026 attack class: AI coding assistants (Copilot, Cursor, ChatGPT) routinely recommend **non-existent packages** in generated code. Attackers monitor popular hallucinations, register those names on npm/PyPI with malicious payloads, and wait for developers to `npm install` the suggestion. Traditional typosquatting detection (Levenshtein to popular packages) misses this entirely -- the hallucinated name may not resemble any real package.

**Our answer: Hallucinated Package Detector (Phase E-bis).** The scanner pulls a feed of known-hallucinated package names (maintained from `huntr.dev`, Snyk advisor, internal heuristics) and flags any dependency that (a) was published less than 30 days ago AND (b) appears in the hallucination-candidate list OR (c) has zero maintainer history + fewer than 5 weekly downloads + matches a "plausible but made-up" name pattern. Surfaced as `signal_type=hallucinated_pkg_candidate` with severity `critical` (block install).

### Problem 10: Function-Level Reachability Without LLM Cost

AI Impact Analysis is powerful but (a) expensive at scale, (b) non-deterministic, (c) can hallucinate. Endor Labs and Socket (via Coana acquisition) offer a cheaper alternative: **static function-level call-graph reachability** across 40+ languages, achieving 95-97% noise reduction **deterministically** -- no LLM required.

**Our answer: Call-Graph Reachability Engine (Phase E2-bis).** Before invoking the LLM, the scanner builds a static call graph of the user's source code + used library code, identifies entry points (HTTP routes, CLI main, message handlers, exported APIs), and traces whether the vulnerable function is reachable from any entry point. If unreachable -> `reachability=unreachable` without any LLM call. If reachable -> pass to AI Impact Analyzer for data-flow verdict. This combination (deterministic call-graph first, AI deep-dive second) is cheaper than pure-AI approaches, more accurate than scope-only heuristics, and gracefully degrades when AI is disabled.

### Problem 11: Container / OCI Image Blindness

The plan as originally scoped ignored container images -- but modern deployments ship 70%+ of vulnerabilities through OS packages in base images (glibc, openssl, libxml2), not through language-level dependencies. Snyk Container, Aqua, Trivy, Wiz all offer this; without it we cannot sell to production teams.

**Our answer: Container / OCI Scanning (Phase M2).** Full image scanning: pull the image, iterate over layers, identify installed OS packages (apt/apk/rpm), check against Trivy DB, detect base image + recommend a minimal alternative (distroless, chainguard), apply Dockerfile SAST rules (root user, `ADD` from http, `:latest` tag, missing `HEALTHCHECK`), verify Sigstore signatures and SLSA provenance attestations. Output a container-specific VEX-enriched SBOM.

### Problem 12: Developer Adoption Barrier (CLI + Web Is Not Enough)

Snyk, Socket, and Semgrep won adoption through **IDE integration**. A developer discovers a vulnerability when they hover over a version string in `package.json` -- not when CI fails hours later. A CLI-only tool, no matter how good, loses mindshare to a weaker tool that ships a VSCode plugin.

**Our answer: IDE Plugin via LSP (Phase P).** Ship a Go-based Language Server Protocol (LSP) server that the CLI exposes, plus thin VSCode and JetBrains extensions. Hover on a version in `package.json` / `requirements.txt` / `go.mod` / `.csproj` -> inline preview: CVE count, AI verdict one-liner, fix-PR button. Save-time scan (debounced) updates hints without running full CI.

### Problem 13: AI / LLM Code Risk (new attack class)

2026 AppSec platforms are racing to cover a new vulnerability class: **insecure LLM integration** in application code. Prompt injection (user input flowing into `openai.chat.completions`), unsafe LLM output handling (LLM output piped to `eval` / `dangerouslySetInnerHTML` / SQL), hardcoded system prompts with secrets, MCP server tools with unbounded shell exec, insecure LangChain/LlamaIndex agent configurations -- none of these appear in classical SAST rule packs.

**Our answer: AI/LLM Security Rule Pack (`rules/semgrep/ai-llm.yaml`, Phase E).** 15-20 dedicated Semgrep rules covering prompt injection, unsafe LLM output sinks, secret-bearing system prompts, MCP tool argument trust boundaries, LangChain agent insecure tool wiring, LlamaIndex missing output parsers, missing rate limits on LLM endpoints (cost-DoS), tokenizer DoS. Rules are ecosystem-agnostic and apply to Node.js, Python, and .NET codebases.

### Problem 14: Policy-as-Code for Enterprise Gates

Ignore-policies and allow/deny license rules are fine for small teams, but enterprises need **composite gates**: "fail CI if any prod finding has EPSS > 0.7 AND fix-risk is patch-level AND no VEX not_affected overlay AND no active Jira ticket". Expressing this in a YAML file quickly becomes unreadable. Snyk does not offer this; Endor Labs does.

**Our answer: Rego / OPA Policy-as-Code (Phase Q).** Opt-in Rego engine (`open-policy-agent/opa`). Org can commit `.vulncheck-policy/*.rego` files that express arbitrary composite gates against the full Finding + Dependency + License + VEX schema. Evaluation runs after the DB + `.vulncheck` YAML policies and can `block`, `warn`, or `notify`.

### Problem 15: Supply Chain Provenance (SLSA / Sigstore)

The 2026 supply-chain defense bar is **SLSA L3 + Sigstore**: every published artifact (npm tarball, PyPI wheel, container image) should have a verifiable provenance attestation proving it was built from the claimed git tag by the claimed CI system, and signed via Sigstore Rekor. None of the classical SCA tools verify this yet.

**Our answer: SLSA Provenance + Sigstore Verification (Phase Q).** For each dependency, query the registry for SLSA provenance attestations, verify Sigstore signatures via Rekor transparency log, and surface `signal_type=missing_provenance` / `provenance_mismatch` / `unsigned_artifact` as supply-chain findings. Registries that publish provenance (npm for `provenance: true` packages, GitHub Container Registry for ghcr images) are prioritized. Compliance tier for EU Cyber Resilience Act / US EO 14028.

### Problem 16: PR-Scoped Findings (only NEW, not everything)

When a developer opens a PR, they care about **what THEIR PR introduced**, not the full legacy finding inventory. Snyk and npm audit dump the entire scan result into PR comments, drowning out the signal. GitHub Dependabot got this right: show only what this PR adds.

**Our answer: PR Baseline Diff (Phase K extension).** On PR events, the scanner runs on the PR head commit and diffs against the last successful scan on `defaultBranch`. Only findings present on the PR head but absent on the baseline are posted as **blocking**. Findings fixed by the PR are posted as **resolved** in green. The PR comment also shows an inline lockfile diff (packages added/removed/bumped).

### Problem 17: Attack Story Correlation (ASPM)

Traditional AppSec platforms surface findings in silos: "you have an SCA finding", "you have a secret", "you have a SAST finding". An attacker chains these: the hardcoded DB credential (secret) + the SQL injection (SAST) + the prod-reachable `pg` package with known CVE (SCA) + public S3 bucket (IaC) = one exploit chain. ASPM platforms (Apiiro, OX, Cycode) **correlate** these.

**Our answer: ASPM Attack-Story Correlator (Phase P2).** Post-scan job groups findings across all scanners by shared entity (file path, package, service, DB connection string, HTTP route) and builds an attack-graph. Prioritized "Attack Stories" are surfaced in the dashboard: a single card that shows the chain of findings constituting a plausible attack path, with a combined severity and a single remediation plan. This is a dashboard-only feature (no new data required beyond joining existing findings).

### Problem 18: VEX Consumption (not just generation)

The original plan generates VEX statements as part of SBOM output, but does not **consume** upstream VEX statements. Many upstream maintainers (Kubernetes, Red Hat, Node.js TSC, Python Security WG) publish OpenVEX / CSAF documents explicitly stating "our version X is NOT affected by CVE-Y even though scanners flag it". Consuming these drastically reduces false positives without any heuristic.

**Our answer: VEX Consumer (Phase Q).** Periodic River job pulls OpenVEX JSON + CSAF documents from a configurable feed list (GitHub releases of major orgs, vendor security pages, OASIS CSAF repositories). Statements are parsed, validated (signature check where possible), and applied as implicit ignore-policies with `source=upstream_vex`. Provenance of each dismissal is tracked -- users see "dismissed per upstream VEX from kubernetes.io/security, published 2026-01-15".

### Problem 19: AI Verdict Feedback Loop

LLMs hallucinate. When a user disagrees with an AI verdict, the system must **learn**. The original plan offers a "re-analyze" button but does not capture the correction. Without a feedback loop, the LLM prompt never improves -- a compounding disadvantage vs competitors that tune on user signal.

**Our answer: AI Verdict Feedback + Prompt Quality Metrics (Phase O extension).** When a user corrects an AI verdict (marks a `no_impact` as actually exploitable, or vice versa), the correction is stored as `AIVerdictFeedback`. An aggregated dashboard shows per-model accuracy, per-CWE false-positive rate, per-language confidence. Anonymized evaluation sets are exported for prompt-tuning (org-level opt-in). Over time the prompt library self-improves.

---

## Business Requirements

- **SCA**: Parse `package-lock.json` to build a **full transitive dependency graph** with **dependency scope tracking** (production/development/optional/peer). Look up every dependency+version against npm audit API, OSV.dev, and NVD. Compute **effective severity** based on dependency scope and reachability. Report CVEs with **upgrade path remediation** ("bump express@4.17 to 4.18 to fix CVE in transitive dep qs@6.5") and **remediation risk level** (patch/minor/major/no_fix).
- **Supply Chain Analysis**: Detect suspicious dependency patterns: postinstall hooks with shell execution, newly published packages (<7 days), maintainer changes, typosquatting (Levenshtein distance to popular packages), obfuscated code. Surfaced as `supply_chain` findings.
- **AI Impact Analysis**: For each vulnerability finding, AI analyzes the user's actual source code to determine real impact. Indexes all import/call sites of the vulnerable package, extracts code snippets, sends CVE description + usage context to LLM, produces a structured verdict (critical_impact / moderate_impact / low_impact / no_impact) with human-readable explanation and exact file:line locations. AI verdict overrides scope-based reachability with actual code analysis. Supports OpenAI, Anthropic, and self-hosted LLM backends. Runs as an optional enrichment step (can be disabled for speed in CI).
- **License Compliance**: Scan all dependencies for open-source license types (MIT, Apache-2.0, GPL-2.0, GPL-3.0, AGPL-3.0, BSD, LGPL, MPL, ISC, Unlicense, proprietary, unknown). Org-level license policies: `allow` (whitelist), `deny` (blacklist, e.g. "block GPL in proprietary projects"), `warn` (flag for review). Handle dual-licensing (`OR` / `AND` SPDX expressions). Surface license violations as `LicenseFinding` -- separate from vulnerability findings but shown in the same dashboard. CLI: `vulncheck license --path .` for standalone license audit. Reports include license compliance section. Supports all ecosystems via the same parser (license info from lockfiles + registry API).
- **SBOM Generation**: Export Software Bill of Materials in **SPDX 2.3** (JSON/tag-value) and **CycloneDX 1.5** (JSON/XML) formats. Generated from the dependency graph (DependencySnapshot). Includes: all direct + transitive deps with versions, licenses, suppliers, dependency relationships, vulnerability status (optional). Required for US Executive Order 14028, EU Cyber Resilience Act compliance. Available via API endpoint, CLI command (`vulncheck sbom --format spdx-json`), and dashboard download.
- **Automated Fix PRs**: When a vulnerability has a safe fix (patch/minor risk level), the platform can **automatically create a pull request** on GitHub/GitLab/Bitbucket that upgrades the vulnerable dependency. Includes: updated `package.json` / `requirements.txt` / `go.mod` + lockfile regeneration. Safety check: only creates PR if the fix does not introduce new vulnerabilities of equal or higher severity. Configurable: auto (create PR on every scan), manual (button in dashboard), or disabled. Uses git credentials from Phase C2.
- **Secrets Scanning**: Dedicated secrets detection engine (beyond generic SAST rules). Scans source code for hardcoded credentials using 100+ regex patterns + entropy analysis: AWS keys (`AKIA...`), GCP service account JSON, Azure connection strings, GitHub/GitLab tokens (`ghp_`, `glpat-`), Slack tokens (`xoxb-`, `xoxp-`), database connection strings, JWT secrets, private keys (RSA/EC/Ed25519), generic high-entropy strings. Each detection includes: secret type, file:line, masked preview (first/last 4 chars), confidence score. Surfaced as findings with `scanType=secrets`. Integrates with git history scanning (optional: check if secret was ever committed, even if removed).
- **SAST**: Semgrep integration + custom rules (scaffolded, built out later).
- **DAST**: OWASP ZAP integration (scaffolded).
- **IaC**: Checkov/tfsec for Terraform, K8s, Docker, CloudFormation (scaffolded).
- **Dual Policy System**: `.vulncheck` policy file in the repo (version-controlled, Snyk `.snyk`-style) + org-level policies managed via API/dashboard. Both support required `reason` and `ignoreUntil` expiration.
- **Monitor Mode**: Snapshot a project's dependency tree. A periodic River job re-checks snapshots against vuln databases and creates new findings/alerts when new CVEs are published. **Incremental mode**: if lockfile hash unchanged, check only newly published CVEs.
- **Jira Integration**: Auto-create Jira tickets for findings above a configurable **effective severity** threshold. Two-way sync: auto-close tickets when vulns are fixed, read ticket status back to mark findings resolved. Deduplicate: one ticket per unique CVE+package per project.
- **Notifications**: Email + Slack alerts on new vulnerabilities (monitor mode), scan completions, and critical findings. Configurable per org. **Noise control**: rate-limit alerts to prevent alert fatigue.
- **GitHub PR Checks**: Post scan results as GitHub commit status checks (pass/fail based on **effective severity**, not raw). Comment on PRs with a findings summary table showing scope and reachability. Blocks merge if critical/high **production-scope** vulns found (configurable).
- **Compliance Reports**: Generate PDF and CSV reports for security audits. Include scan history, open findings, policy justifications, remediation status, **dependency scope breakdown**, **reachability summary**.
- **Private Registry Support**: Authenticate against private npm registries (GitHub Packages, Artifactory, Verdaccio) when resolving packages and looking up advisories.
- **Multi-Ecosystem Support (Progressive Delivery)**: Four ecosystems with full SCA + SAST coverage, delivered progressively:
  - **Phase 1 -- Node.js (npm)**: `package-lock.json` / `yarn.lock` / `pnpm-lock.yaml` parser, npm advisory client, Express/NestJS SAST rules. First and most complete implementation.
  - **Phase 2 -- Python (pip/pipenv/poetry)**: `requirements.txt` / `Pipfile.lock` / `poetry.lock` parser, PyPI advisory client, Django/Flask/FastAPI SAST rules, PyPI supply chain detection.
  - **Phase 3 -- .NET (NuGet) + Go (go.mod)**: `packages.lock.json` / `.csproj` / `packages.config` parser + `go.mod` / `go.sum` parser, NuGet/Go advisory clients, ASP.NET Core SAST rules, Go/Gin/GORM SAST rules, NuGet/Go supply chain detection.
  - **Future** -- Java (Maven/Gradle), Ruby (Bundler), PHP (Composer), Rust (Cargo) via the same `EcosystemParser` interface.
- **Web Dashboard**: React frontend for org management, project views, scan results (with scope/reachability filters), policy management, integration settings, report downloads. **Default view: actionable production findings only.**
- **Multi-tenant auth**: JWT, organizations, users (admin/member roles).
- **Delivery**: REST API + CLI + Web UI + **IDE Plugin (LSP server + VSCode/JetBrains extensions)**.
- **Behavioral Package Analysis**: For each dependency, fetch the published tarball, parse its AST, and build a `PackageBehaviorProfile` (capabilities: `network`, `filesystem`, `shell`, `env`, `unsafe_eval`, `native_binding`, `obfuscated`, `install_hook_shell`). Diff against the previous version's profile. Emit `capability_expansion` findings when a minor/patch bump gains capabilities (e.g. `lodash@4.17.21 -> 4.17.22` gained network + filesystem -> malicious). Semantically analyze install scripts (not regex): detect base64-decoded payloads, remote downloads, process spawns. Detect malicious packages **before any CVE is assigned**.
- **Hallucinated / Slopsquat Detection**: Maintain a feed of hallucinated package names (periodic River job pulls from public research sources + internal heuristics). Flag any dependency matching `(published < 30 days) AND (in hallucination list OR zero maintainer history + <5 weekly downloads + plausible-but-made-up name)`. Block install via auto-generated critical ignore-policy.
- **Function-Level Call-Graph Reachability**: Deterministic call-graph engine per ecosystem (`ts-morph` for JS/TS, `jedi`/`scalpel` for Python, `golang.org/x/tools/go/callgraph` for Go, Roslyn for C#). Identify entry points (HTTP routes, CLI main, message handlers, exported APIs, test files). Trace whether the vulnerable function is reachable from ANY entry point. If unreachable -> `reachability=unreachable` without any LLM call. Provides a cheap, deterministic 95%+ noise reduction layer BEFORE AI Impact Analyzer runs, with graceful fallback to scope-only mode when call-graph extraction fails.
- **Container / OCI Image Scanning**: Full image scanning via `go-containerregistry` + Trivy DB. Pull image -> iterate layers -> identify OS packages (apt/apk/rpm) per layer -> check against Trivy DB -> apply Dockerfile SAST rules (root user, `ADD http`, `:latest` tag, missing `USER`, missing `HEALTHCHECK`, sensitive files in image) -> base-image detection + recommend minimal alternative (distroless, chainguard, alpine-slim) with layer-size reduction estimate -> verify Sigstore signatures via Rekor + SLSA provenance attestations. Output: container-specific SBOM (CycloneDX) with layer-level breakdown + VEX overlay.
- **AI / LLM Security Rule Pack**: Dedicated Semgrep rule file (`rules/semgrep/ai-llm.yaml`, 15-20 patterns) covering the new attack class: prompt injection (user input -> LLM completion), unsafe LLM output sinks (LLM output -> `eval` / `dangerouslySetInnerHTML` / raw SQL / file write), hardcoded system prompts containing secrets, MCP tool unbounded shell exec, LangChain agent insecure tool wiring, LlamaIndex missing output parsers, missing rate limit on LLM endpoints (cost-DoS), tokenizer DoS via malformed input. Ecosystem-agnostic (Node.js, Python, .NET).
- **VEX Consumption (upstream)**: Periodic River job pulls OpenVEX JSON + OASIS CSAF documents from a configurable feed list (GitHub releases of kubernetes, nodejs, python-security-wg, rust-lang, etc., vendor security pages). Parses statements, validates signatures where available, applies `not_affected` / `fixed` / `under_investigation` statuses as implicit ignore-policies with `source=upstream_vex`. Every auto-dismissal retains provenance ("dismissed per Kubernetes VEX, published 2026-01-15, signed by sigstore identity X"). Combined with our own VEX generation gives bidirectional VEX workflow.
- **Policy-as-Code (Rego / OPA)**: Opt-in Rego policy engine via `open-policy-agent/opa`. Org commits `.vulncheck-policy/*.rego` files expressing composite gates against the full schema (Finding + Dependency + License + VEX + AIVerdict + ExploitIntel). Example: `block if any finding has epss_score > 0.7 and fix_risk == "patch" and scope == "production" and not has_vex_not_affected and not has_active_jira_ticket`. Evaluation runs after DB + `.vulncheck` YAML policies in the scan pipeline. Outcomes: `block` (fail CI), `warn` (comment only), `notify` (slack only).
- **SLSA Provenance + Sigstore Verification**: For each dependency, query the registry for SLSA provenance attestations (npm `provenance: true`, GitHub Artifact Attestations, SLSA.dev verifiable build metadata), verify Sigstore signatures via Rekor transparency log. Supply-chain signals: `missing_provenance` (medium), `provenance_mismatch` (high -- claimed git tag does not match published artifact), `unsigned_artifact` (info). Container images: same verification via cosign/Rekor. Compliance tier for EU CRA / US EO 14028.
- **PR Baseline Diff**: On PR webhook events, scanner runs on the PR head commit + diffs findings against the last successful scan on `defaultBranch`. Only NEW findings (present on PR head, absent on baseline) are posted as blocking. Findings RESOLVED by the PR are posted in green. Lockfile changes (packages added/removed/version-bumped) are rendered as an inline diff table in the PR comment. Eliminates "legacy finding dump" fatigue.
- **ASPM Attack-Story Correlation**: Post-scan correlator joins findings across SCA/SAST/Secrets/IaC/Container by shared entity (file path, package name, service, DB connection string, HTTP route, K8s resource). Builds an attack-graph where nodes are findings and edges are "reachable from" / "uses" / "exposes". Surfaces prioritized "Attack Stories": e.g. "Hardcoded DB credential at `config/db.js:5` (secret) + SQL injection in `db.query()` handler at `api/users.js:42` (SAST) + production-reachable `pg@8.5.0` with CVE-2024-XXX (SCA) + public-readable S3 bucket in terraform (IaC) = data exfiltration chain". Combined severity + single remediation plan.
- **AI Verdict Feedback Loop**: When a user corrects an AI verdict (marks `no_impact` as exploitable or vice versa), store as `AIVerdictFeedback` with the delta. Aggregate dashboard: per-model accuracy, per-CWE FP rate, per-language confidence, per-prompt-template quality. Org-level opt-in: anonymized eval sets exported for prompt-tuning. Over time the prompt library self-improves based on real signal.
- **Multi-file Repository-level AI Context**: Beyond per-finding snippets, the AI Impact Analyzer is fed a repo-level context envelope: detected framework (Express/NestJS/Django/Flask/FastAPI/ASP.NET Core/Gin), list of entry-points + routes, DI bindings, environment tags ("this is a Next.js API route", "this is a background worker"). Token-budget manager summarizes large contexts. Drastically improves accuracy on middleware chains, DI-based wiring, and framework-specific patterns.
- **IDE Plugin (LSP)**: Go-based Language Server Protocol server embedded in the CLI (`vulncheck lsp`). VSCode extension + JetBrains plugin communicate via LSP. Features: hover hints on package versions (CVE count, AI verdict one-liner, fix-risk badge), save-time debounced scan, quick-fix code actions (trigger Fix PR, apply safe upgrade locally, add `.vulncheck` ignore with template), inline capability badges for behavioral analysis, "Explain with AI" command.
- **Runtime Reachability Agent (future, opt-in)**: Light-weight Go agent deployed alongside prod workloads, using `cilium/ebpf` to observe which shared objects / JS modules / Python modules are actually loaded into memory per process, plus sampled syscall traces per vulnerable function. Findings gain `runtimeObserved: true/false` flag. Drives production prioritization: a CVE in an unloaded library is low priority regardless of everything else. Phase R, deferred until core platform is stable.

## Architecture

```mermaid
flowchart TB
  subgraph clients [Clients]
    WebUI["Web Dashboard (React)"]
    CLI["CLI (cobra)"]
    CICD["CI/CD Pipelines"]
  end

  subgraph api [Gin API Server]
    AuthMW["Auth Middleware"]
    AuthH["Auth Handler"]
    ProjH["Project Handler"]
    ScanH["Scan Handler"]
    PolicyH["Policy Handler"]
    LicenseH["License Handler"]
    SBOMH["SBOM Handler"]
    FixPRH["Fix PR Handler"]
    IntegH["Integration Handler"]
    MonitorH["Monitor Handler"]
  end

  subgraph services [Service Layer]
    AuthS["Auth Service"]
    ProjS["Project Service"]
    GitS["Git Integration Service"]
    ScanS["Scan Orchestrator"]
    PolicyS["Policy Service"]
    LicenseS["License Compliance Service"]
    SBOMS["SBOM Generator"]
    FixPRS["Fix PR Service"]
    TicketS["Ticket Service"]
    MonitorS["Monitor Service"]
    RemediationS["Remediation Engine"]
    NotifyS["Notification Service"]
    PRCheckS["PR Check Service"]
    ReportS["Report Service"]
  end

  subgraph scanners [Scanners]
    SCA["SCA Scanner"]
    BehavioralAnalyzer["Behavioral Package Analyzer"]
    HallucinationDetector["Hallucinated Package Detector"]
    CallGraphEngine["Call-Graph Reachability Engine"]
    ContextEngine["Context Analysis Engine"]
    SupplyChainDetector["Supply Chain Detector"]
    SLSAVerifier["SLSA / Sigstore Verifier"]
    AIAnalyzer["AI Impact Analyzer"]
    SecretsScanner["Secrets Scanner"]
    SAST["SAST Scanner"]
    ContainerScanner["Container / OCI Scanner"]
    DAST["DAST Scanner"]
    IAC["IaC Scanner"]
    VEXConsumer["Upstream VEX Consumer"]
    RegoEngine["Rego Policy Engine"]
    ASPMCorrelator["ASPM Attack-Story Correlator"]
  end

  subgraph external [External Services]
    NPM["npm Registry"]
    OSV["OSV.dev"]
    NVD["NVD"]
    GHSA["GitHub Advisory DB"]
    VulnCheckAPI["VulnCheck API"]
    TrivyDB["Trivy DB"]
    LLM["LLM API (OpenAI / Anthropic / self-hosted)"]
    Jira["Jira Cloud"]
    GH["GitHub API"]
    Lin["Linear"]
    Semgrep["Semgrep"]
    ZAP["OWASP ZAP"]
    Checkov["Checkov"]
    SMTP["SMTP / Email"]
    SlackAPI["Slack API"]
    PrivateReg["Private Registries"]
    OCIRegistry["OCI / Container Registries"]
    Sigstore["Sigstore Rekor"]
    OpenVEX["OpenVEX / CSAF Feeds"]
    OPA["Open Policy Agent"]
  end

  subgraph ide [IDE]
    VSCode["VSCode Extension"]
    JetBrains["JetBrains Plugin"]
    LSPServer["LSP Server (vulncheck lsp)"]
  end

  subgraph data [Data Layer]
    Repo["Repositories"]
    PG["PostgreSQL"]
    RiverQ["River Jobs"]
  end

  WebUI & CLI & CICD --> api
  VSCode & JetBrains --> LSPServer
  LSPServer --> api
  AuthMW --> AuthH & ProjH & ScanH & PolicyH & LicenseH & SBOMH & FixPRH & IntegH & MonitorH
  AuthH --> AuthS
  ProjH --> ProjS
  ScanH --> ScanS
  PolicyH --> PolicyS
  LicenseH --> LicenseS
  SBOMH --> SBOMS
  FixPRH --> FixPRS
  IntegH --> TicketS
  MonitorH --> MonitorS

  ScanS --> GitS
  GitS --> GH
  ScanS --> SCA & SAST & SecretsScanner & DAST & IAC & ContainerScanner
  SCA --> BehavioralAnalyzer
  SCA --> HallucinationDetector
  SCA --> SupplyChainDetector
  SCA --> SLSAVerifier
  SCA --> CallGraphEngine
  CallGraphEngine --> ContextEngine
  ContextEngine --> AIAnalyzer
  AIAnalyzer --> LLM
  ScanS --> RemediationS
  ScanS --> PolicyS
  ScanS --> LicenseS
  ScanS --> TicketS
  ScanS --> VEXConsumer
  ScanS --> RegoEngine
  ScanS --> ASPMCorrelator
  RegoEngine --> OPA
  VEXConsumer --> OpenVEX
  SLSAVerifier --> Sigstore
  ContainerScanner --> OCIRegistry & TrivyDB & Sigstore
  RemediationS --> FixPRS
  FixPRS --> GH
  LicenseS --> NPM & PrivateReg
  SBOMS --> SCA
  BehavioralAnalyzer --> NPM & PrivateReg

  SCA --> NPM & OSV & NVD & GHSA & VulnCheckAPI & TrivyDB & PrivateReg
  SAST --> Semgrep
  DAST --> ZAP
  IAC --> Checkov
  TicketS --> Jira & GH & Lin
  NotifyS --> SMTP & SlackAPI
  PRCheckS --> GH
  MonitorS --> RiverQ
  ScanS --> NotifyS & PRCheckS

  AuthS & ProjS & GitS & ScanS & PolicyS & TicketS & MonitorS --> Repo
  Repo --> PG
  ScanS --> RiverQ
```

## Project Structure

```
artfend-vulncheck/
  cmd/
    server/main.go                -- API entry point, DI wiring
    cli/main.go                   -- CLI entry point (cobra)
  internal/
    config/config.go              -- Viper config struct
    domain/
      errors.go                   -- Shared sentinel errors
      scanner.go                  -- Scanner interface, ScanType enum
      severity.go                 -- Severity enum + ordering
      types.go                    -- Shared types (pagination, etc.)
    middleware/
      auth.go                     -- JWT auth middleware
      logger.go                   -- Zap request logger
      recovery.go                 -- Panic recovery
      cors.go                     -- CORS middleware
    auth/
      model.go                    -- User, Organization models + DTOs
      repository.go
      service.go                  -- Register, login, JWT, invite
      handler.go
    project/
      model.go                    -- Project model + DTOs
      repository.go
      service.go
      handler.go
    git/
      model.go                    -- GitCredential model + DTOs
      repository.go               -- GitCredential CRUD
      service.go                  -- Git Integration Service: clone, verify ownership, extract files
      handler.go                  -- Git credential management + webhook endpoints
      cloner.go                   -- Safe repo cloner: shallow clone, timeout, size limits, no hooks
      provider_github.go          -- GitHub API: verify repo access, list repos, validate token scope
      provider_gitlab.go          -- GitLab API: verify repo access, list repos, validate token scope
      provider_bitbucket.go       -- Bitbucket API: verify repo access, list repos, validate token scope
      extractor.go                -- File extractor: pick lockfiles + source code from cloned repo, skip binaries
      webhook.go                  -- Webhook receiver: push events -> trigger scan
      security.go                 -- Ownership verification, rate limiting, audit logging
    scan/
      model.go                    -- Scan, Finding models + DTOs
      repository.go
      service.go                  -- Scan orchestrator
      handler.go
    policy/
      model.go                    -- IgnorePolicy model + DTOs
      repository.go
      service.go                  -- Evaluate policies (DB + file + Rego)
      handler.go
      file.go                     -- .vulncheck file parser/writer
      rego.go                     -- Rego / OPA policy engine: load .vulncheck-policy/*.rego, evaluate against full schema, emit block/warn/notify
      rego_schema.go              -- Rego input schema: Finding + Dependency + License + VEX + AIVerdict + ExploitIntel unified view
    sca/
      scanner.go                  -- SCA scanner (Scanner interface)
      parser.go                   -- package-lock.json parser, dep graph
      graph.go                    -- Dependency graph traversal + scope propagation
      context.go                  -- Context analysis: scope assignment + reachability classification
      effective_severity.go       -- Effective severity calculation engine
      remediation.go              -- Upgrade path analysis + fix risk level
      supply_chain.go             -- Supply chain signal detector (hooks, typosquat, new pkg, maintainer)
      npm.go                      -- npm bulk advisory client
      osv.go                      -- OSV.dev batch query client
      nvd.go                      -- NVD CVE API client + caching
      ghsa.go                     -- GitHub Advisory DB GraphQL client
      vulncheck.go                -- VulnCheck API client (exploit intel, EPSS, KEV)
      trivy.go                    -- Trivy DB integration (offline vuln DB + misconfig)
      dedup.go                    -- Cross-source finding dedup + multi-source merge
      incremental.go              -- Incremental scan: lockfile hash comparison + delta CVE check
    ai/
      analyzer.go                 -- AI Impact Analyzer orchestrator: per-finding analysis pipeline
      code_indexer.go             -- Source code indexer: find imports, require(), call sites of a given package
      snippet_extractor.go        -- Extract relevant code snippets around usage sites with data-flow context
      prompt.go                   -- Prompt templates for different vuln types (RCE, XSS, prototype pollution, ReDoS, etc.)
      llm.go                      -- LLM client abstraction (OpenAI, Anthropic, self-hosted) with retry/rate-limit
      verdict.go                  -- Structured AI verdict types + parsing LLM responses into typed verdicts
    license/
      model.go                    -- LicensePolicy, LicenseFinding models + DTOs
      repository.go               -- License policy + finding CRUD
      service.go                  -- License detection (lockfile + registry API) + policy evaluation
      handler.go                  -- License policy CRUD + license report endpoints
      spdx.go                     -- SPDX expression parser (handles OR/AND/WITH operators)
      detector.go                 -- Extract license from package metadata (npm, PyPI, NuGet, Go)
    sbom/
      model.go                    -- SBOMExport model
      service.go                  -- SBOM generation orchestrator
      handler.go                  -- SBOM export endpoints
      spdx_writer.go              -- Generate SPDX 2.3 JSON + tag-value format
      cyclonedx_writer.go         -- Generate CycloneDX 1.5 JSON + XML format
      vex.go                      -- VEX (Vulnerability Exploitability eXchange) overlay
    fixpr/
      model.go                    -- FixPR model + DTOs
      repository.go               -- FixPR CRUD
      service.go                  -- Fix PR orchestrator: compute changes, safety check, create PR
      handler.go                  -- Fix PR endpoints (trigger, list, status)
      lockfile_updater.go         -- Update lockfile content (npm/pip/dotnet/go ecosystem-specific)
      github_pr.go                -- GitHub PR creation API (create branch, commit, PR)
      gitlab_mr.go                -- GitLab MR creation API
      bitbucket_pr.go             -- Bitbucket PR creation API
    secrets/
      scanner.go                  -- Secrets scanner (Scanner interface implementation)
      patterns.go                 -- 100+ regex patterns for known secret types
      entropy.go                  -- Shannon entropy calculator for high-entropy string detection
      validator.go                -- False positive reduction: validate detected secrets (e.g. check if AWS key is real format)
      git_history.go              -- Optional: scan git log for previously committed secrets
    sast/
      scanner.go                  -- Semgrep wrapper (exec + parse SARIF/JSON output)
      rules.go                    -- Rule management: load, validate, merge custom + built-in rules
      mapper.go                   -- Map Semgrep SARIF findings to our Finding model
    behavior/
      model.go                    -- PackageBehaviorProfile model + Capability enum
      repository.go               -- CRUD for behavior profiles (keyed by package@version)
      service.go                  -- Behavioral analysis orchestrator: fetch tarball -> AST -> capabilities -> diff
      fetcher.go                  -- Package tarball fetcher (npm/PyPI/NuGet/Go module proxy) with sandboxed unpack
      ast_analyzer.go             -- AST-based capability extractor per ecosystem (uses tree-sitter for cross-lang consistency)
      capability_diff.go          -- Diff capabilities between two versions, detect capability_expansion
      install_script.go           -- Semantic analysis of install hooks (not regex): base64 payloads, remote downloads, process spawns
      hallucination.go            -- Hallucinated / slopsquat package detector with feed refresher
      feed.go                     -- Pull hallucination candidate feed from configured sources
    callgraph/
      engine.go                   -- Call-Graph Reachability Engine orchestrator: build graph -> find entrypoints -> trace paths
      model.go                    -- CallGraphNode, CallGraphEdge, ReachabilityPath models
      entrypoints.go              -- Entry-point discovery: HTTP routes, CLI main, message handlers, exported APIs, test files
      node_js.go                  -- Node.js/TS call-graph builder (ts-morph/TypeScript compiler API via external subprocess)
      python.go                   -- Python call-graph builder (jedi/scalpel via external subprocess)
      go.go                       -- Go call-graph builder (golang.org/x/tools/go/callgraph, native)
      dotnet.go                   -- .NET call-graph builder (Roslyn via external subprocess)
      tracer.go                   -- Path tracer: BFS from entry-points to target vulnerable function, returns reachable?+path
    container/
      scanner.go                  -- Container / OCI Scanner (Scanner interface, scanType=container)
      image.go                    -- Image pull + layer iteration via go-containerregistry
      os_packages.go              -- Extract installed apt/apk/rpm packages per layer
      dockerfile.go               -- Dockerfile parser + SAST rules (root user, ADD http, :latest, missing USER, missing HEALTHCHECK)
      base_image.go               -- Base image detection + minimal-alternative recommender (distroless, chainguard, alpine-slim)
      cosign.go                   -- Sigstore/cosign signature verification via Rekor transparency log
      slsa_attestation.go         -- SLSA provenance attestation verification for images
      sbom_container.go           -- Container-specific CycloneDX SBOM with layer breakdown
    slsa/
      verifier.go                 -- SLSA provenance verifier for language-level deps (npm provenance, GitHub Artifact Attestations)
      sigstore.go                 -- Sigstore Rekor transparency log client
      model.go                    -- SLSAAttestation model + ProvenanceStatus enum
    vex/
      consumer.go                 -- Upstream VEX consumer orchestrator
      openvex.go                  -- OpenVEX JSON parser
      csaf.go                     -- OASIS CSAF 2.0 parser
      feed.go                     -- Configurable feed list + periodic refresh
      applier.go                  -- Apply VEX statements as implicit ignore-policies with provenance tracking
      model.go                    -- VEXStatement, VEXSource models
    aspm/
      correlator.go               -- Post-scan ASPM correlator: join findings across scanners by shared entity
      attack_graph.go             -- Attack-graph builder (nodes = findings, edges = reachable-from/uses/exposes)
      story.go                    -- AttackStory aggregator + prioritized surfacing
      model.go                    -- AttackStory, AttackStoryEdge, AttackStoryEntity models
    feedback/
      model.go                    -- AIVerdictFeedback model
      repository.go
      service.go                  -- Capture corrections, aggregate quality metrics
      handler.go                  -- Feedback submission + metrics endpoints
      exporter.go                 -- Anonymized eval set exporter (org-level opt-in)
    dast/
      scanner.go                  -- ZAP wrapper [scaffold]
    iac/
      scanner.go                  -- Checkov wrapper [scaffold]
    integration/
      model.go                    -- IntegrationConfig, TicketMapping models
      repository.go
      service.go                  -- Ticket lifecycle (create/sync/close)
      handler.go                  -- Integration settings endpoints
      jira.go                     -- Jira Cloud REST API client
      github.go                   -- GitHub Issues client [scaffold]
      linear.go                   -- Linear client [scaffold]
    monitor/
      model.go                    -- Snapshot model
      repository.go
      service.go                  -- Snapshot + diff logic
      handler.go                  -- Monitor endpoints
    notify/
      model.go                    -- NotificationConfig model
      repository.go
      service.go                  -- Notification dispatch
      handler.go                  -- Notification settings endpoints
      email.go                    -- SMTP email sender
      slack.go                    -- Slack webhook/API sender
    prcheck/
      model.go                    -- PRCheckConfig model
      service.go                  -- Post status checks + PR comments
      github.go                   -- GitHub Checks API + PR comments client
    report/
      model.go                    -- ReportRequest model
      service.go                  -- Report generation orchestrator
      handler.go                  -- Report endpoints
      pdf.go                      -- PDF renderer
      csv.go                      -- CSV renderer
    jobs/
      scan_job.go                 -- River: async scan execution
      ai_analysis_job.go          -- River: async AI impact analysis for large scans
      monitor_job.go              -- River: periodic vuln re-check
      ticket_sync_job.go          -- River: two-way ticket sync
      notification_job.go         -- River: async notification dispatch
      report_job.go               -- River: async report generation
      sbom_job.go                 -- River: async SBOM generation (large dep trees)
      fix_pr_job.go               -- River: async fix PR creation (lockfile regen + safety check)
      fix_pr_monitor_job.go       -- River: periodic check if fix PRs were merged/closed
      git_credential_validation_job.go -- River: periodic credential validation
      behavior_profile_job.go     -- River: async behavioral profile build for new/updated packages
      hallucination_feed_job.go   -- River: periodic hallucination candidate feed refresh
      vex_feed_job.go             -- River: periodic upstream VEX/CSAF feed pull + apply
      callgraph_job.go            -- River: async call-graph build for large repos
      container_scan_job.go       -- River: async container image scan
      aspm_correlation_job.go     -- River: post-scan ASPM correlation + attack-story build
    vuln/
      model.go                    -- Cached vulnerability records
      repository.go               -- Vuln cache data access
      service.go                  -- Cache lookup/refresh
  pkg/
    httputil/response.go          -- Standard JSON responses
    ecosystem/
      ecosystem.go                -- EcosystemParser interface (multi-ecosystem design)
      registry.go                 -- Parser registry (auto-detect ecosystem from files)
      npm/parser.go               -- npm lockfile parser v1/v2/v3 (Phase E)
      python/parser.go            -- pip/pipenv/poetry lockfile parser (Phase E3)
      dotnet/parser.go            -- NuGet packages.lock.json + .csproj parser (Phase E4)
      gomod/parser.go             -- go.mod + go.sum parser (Phase E4)
  web/                            -- React frontend (separate build)
    src/
      ...
  ide/                            -- IDE plugins (separate build artifacts)
    lsp/                          -- Thin wrapper: cmd/cli exposes `vulncheck lsp` that serves Language Server Protocol
    vscode/                       -- VSCode extension (TypeScript, vsce-packaged .vsix)
    jetbrains/                    -- JetBrains plugin (Kotlin, Gradle-built .jar)
  rules/
    semgrep/
      node-express-nest.yaml      -- Node.js/Express/NestJS SAST rules (20+ patterns, Phase E)
      python-django-flask.yaml    -- Python/Django/Flask/FastAPI SAST rules (20+ patterns, Phase E3)
      dotnet-aspnet-core.yaml     -- ASP.NET Core SAST rules (20+ patterns, Phase E4)
      go-gin-gorm.yaml            -- Go/Gin/GORM SAST rules (25+ patterns, Phase E4)
      ai-llm.yaml                 -- AI/LLM security rules (15-20 patterns, Phase E): prompt injection, unsafe LLM output sinks, MCP tool trust boundaries, hardcoded system prompts, LangChain/LlamaIndex insecure configs, LLM rate-limit DoS, tokenizer DoS
      dockerfile.yaml             -- Dockerfile SAST rules (10+ patterns, Phase M2): root user, ADD http, :latest tag, missing USER, missing HEALTHCHECK, sensitive files
      generic.yaml                -- Language-agnostic secret/credential detection rules (Phase E)
  policies/
    rego-examples/                -- Example Rego policies shipped in docs
      no_high_epss_patch.rego     -- Block patch-fix findings with EPSS > 0.7
      kev_always_blocks.rego      -- KEV-listed findings always block regardless of other factors
      license_gate.rego           -- Composite license gate with exceptions
  migrations/
    ...                           -- Goose SQL migrations
  docs/                           -- swaggo generated
  go.mod
  Makefile
  Dockerfile
  docker-compose.yml
  .env.example
```

## Domain Models

### Core GORM Models

- **Organization**: `id`, `name`, `slug`, `createdAt`, `updatedAt`, `deletedAt`
- **User**: `id`, `orgID` (FK), `email`, `passwordHash`, `role` (admin/member), `createdAt`, `updatedAt`, `deletedAt`
- **Project**: `id`, `orgID` (FK), `name`, `description`, `repoURL`, `defaultBranch`, `gitCredentialID` (FK, nullable -- null for public repos), `lastClonedAt`, `repoSizeBytes` (cached), `createdAt`, `updatedAt`, `deletedAt`
- **GitCredential**: `id`, `orgID` (FK), `provider` (github/gitlab/bitbucket), `label` (user-friendly name, e.g. "My GitHub PAT"), `authMethod` (pat/oauth/app_installation), `credentials` (encrypted JSONB: token, or OAuth refresh token, or app installation ID + private key), `scope` (JSONB: list of permissions the token has, validated on creation), `ownerUsername` (the git provider username associated with this credential), `lastValidatedAt`, `valid` (bool -- set to false when token is revoked/expired), `createdAt`, `updatedAt`, `deletedAt`

### Scan + Findings

- **Scan**: `id`, `projectID`, `scanType`, `status` (pending/running/completed/failed), `startedAt`, `completedAt`, `summary` (JSONB: counts by severity + counts by effective severity), `lockfileHash` (SHA-256 of the parsed lockfile for incremental scan detection), `createdAt`
- **Finding**: `id`, `scanID`, `projectID`, `scanType`, `severity` (raw CVSS-based), `effectiveSeverity` (context-adjusted: accounts for scope + reachability + AI verdict + exploit intel), `reachability` (confirmed/potentially_reachable/unreachable), `dependencyScope` (production/development/optional/peer), `title`, `description`, `packageName`, `packageVersion`, `cveID`, `ghsaID` (nullable -- GitHub Advisory ID, e.g. GHSA-xxxx-xxxx-xxxx), `fixAvailable` (bool), `fixVersion`, `fixRiskLevel` (patch/minor/major/no_fix), `upgradePath` (JSONB: ordered list of direct dep upgrades needed), `source` (npm/osv/nvd/ghsa/vulncheck/trivy/supply_chain), `confirmedBySources` (JSONB array -- all sources that confirmed this CVE, e.g. ["npm","osv","ghsa"]), `exploitAvailable` (bool -- known exploit exists in the wild, from VulnCheck/NVD), `epssScore` (float, nullable -- EPSS probability of exploitation, 0.0-1.0, from VulnCheck), `kevListed` (bool -- in CISA Known Exploited Vulnerabilities catalog), `fingerprint` (hash of cveID+packageName+packageVersion for dedup), `raw` (JSONB), `status` (open/ignored/resolved), `aiVerdict` (JSONB, nullable -- see AI Impact Analysis below), `createdAt`

#### Effective Severity Calculation

The `effectiveSeverity` is derived from `severity` through a multi-stage pipeline:

**Stage 1 -- Scope + Reachability (always applied):**
- If `dependencyScope` is `development` or `optional`: downgrade by 2 levels (critical->low, high->info, medium->info, low->info)
- If `reachability` is `unreachable`: downgrade by 1 level (critical->high, high->medium, etc.)
- If `reachability` is `potentially_reachable`: no change
- If `reachability` is `confirmed`: no change (raw severity applies)
- Minimum effective severity is `info` (never hidden, but clearly deprioritized)

This means a "High" ReDoS in a dev-only `browserslist` becomes `info` -- visible but not blocking CI.

**Stage 2 -- AI Impact Analysis (when enabled):**
- If `aiVerdict.impactLevel` is `no_impact`: override to `info` regardless of other factors
- If `aiVerdict.impactLevel` is `low_impact`: downgrade by 1 level
- If `aiVerdict.impactLevel` is `moderate_impact`: no change
- If `aiVerdict.impactLevel` is `critical_impact`: upgrade by 1 level (AI found a confirmed dangerous usage pattern)
- AI verdict always takes precedence over scope-based heuristics when present.

**Stage 3 -- Exploit Intelligence (when VulnCheck/GHSA data available):**
- If `kevListed` is `true` (CISA KEV): upgrade by 1 level -- this CVE is actively exploited in the wild, requires urgent attention
- If `exploitAvailable` is `true` AND `epssScore` >= 0.5: upgrade by 1 level -- high probability of exploitation with known exploit code
- If `exploitAvailable` is `true` AND `epssScore` < 0.5: no change (exploit exists but low exploitation probability)
- If `confirmedBySources` has 3+ sources: boost confidence indicator (not severity, but shown in UI as "high confidence" badge)
- Exploit intelligence is additive with AI verdict: a finding can be upgraded by both AI (critical_impact) AND exploit intel (KEV-listed), making it maximally urgent

**Priority ordering for action:** KEV-listed + production-scope + AI critical_impact findings are always shown first, regardless of raw CVSS score.

#### AI Impact Analysis (`aiVerdict` JSONB structure)

```go
type AIVerdict struct {
    ImpactLevel   AIImpactLevel    `json:"impactLevel"`   // critical_impact / moderate_impact / low_impact / no_impact
    Explanation   string           `json:"explanation"`    // Human-readable: "lodash.merge is called at src/api/handler.js:42 with user-controlled input req.body..."
    UsageFound    bool             `json:"usageFound"`     // Whether the package is imported/required anywhere in user code
    UsageLocations []UsageLocation `json:"usageLocations"` // Exact file:line references
    VulnFunctionCalled bool        `json:"vulnFunctionCalled"` // Whether the specific vulnerable function/method is called
    DataFlowRisk  string           `json:"dataFlowRisk"`   // "user_controlled_input" / "hardcoded_input" / "internal_only" / "unknown"
    Confidence    float64          `json:"confidence"`     // 0.0-1.0 LLM confidence in this assessment
    ModelUsed     string           `json:"modelUsed"`      // e.g. "gpt-4o", "claude-sonnet-4-20250514"
    AnalyzedAt    time.Time        `json:"analyzedAt"`
}

type UsageLocation struct {
    FilePath    string `json:"filePath"`
    Line        int    `json:"line"`
    CodeSnippet string `json:"codeSnippet"` // 5-10 lines of context around the usage
    ImportType  string `json:"importType"`  // "direct_import" / "re_export" / "dynamic_require" / "type_only"
}

type AIImpactLevel string

const (
    AIImpactCritical AIImpactLevel = "critical_impact" // Vuln function called with user-controlled input in prod path
    AIImpactModerate AIImpactLevel = "moderate_impact"  // Vuln function called but input is not user-controlled
    AIImpactLow      AIImpactLevel = "low_impact"       // Package imported but vuln function not called directly
    AIImpactNone     AIImpactLevel = "no_impact"         // Package not imported or only used as type dependency
)
```

Example AI verdicts:

- **Critical**: "CVE-2018-3721 (prototype pollution in `lodash.merge`): The function `_.merge()` is called at `src/api/users.js:42` with `req.body` as the second argument. This is user-controlled HTTP input, making prototype pollution directly exploitable. An attacker can inject `__proto__` properties via a crafted JSON body. **Immediate action required.**"
- **No impact**: "`css-what@3.2.1` (CVE-2021-33587, ReDoS) is only present as a transitive dependency of `svgo`, which is used by `@svgr/webpack` -- a build tool that processes SVG files at compile time. Your source code does not import or call `css-what` directly. The vulnerable CSS selector parser only runs on developer-controlled SVG files during the build. **No impact on your application.**"
- **Low impact**: "`minimist@1.2.5` (CVE-2021-44906, prototype pollution) is imported at `scripts/seed.js:3` but only used in a development seed script that is never executed in production. **No production impact.**"

### Scanner Interface (in `internal/domain/scanner.go`)

```go
type ScanType string

const (
    ScanTypeSCA  ScanType = "sca"
    ScanTypeSAST ScanType = "sast"
    ScanTypeDAST ScanType = "dast"
    ScanTypeIaC  ScanType = "iac"
)

type ScanTarget struct {
    ProjectID      uuid.UUID
    ScanType       ScanType
    Ecosystem      string              // e.g. "npm", "pip", "gomod"
    Files          map[string][]byte   // filename -> content (lockfiles, configs)
    SourceFiles    map[string][]byte   // filename -> content (user source code for AI analysis)
    RepoURL        string
    RegistryConfig *RegistryConfig     // private registry auth (if configured)
    PreviousHash   string              // lockfile hash from last scan (for incremental mode)
    IncludeDevDeps bool                // whether to scan dev deps (default: false)
    AIAnalysis     bool                // whether to run AI impact analysis (default: true, disable with --no-ai)
    Config         map[string]string
}

type ScanOutput struct {
    Findings []FindingData
    Summary  ScanSummary
}

type Scanner interface {
    Type() ScanType
    Scan(ctx context.Context, target *ScanTarget) (*ScanOutput, error)
}
```

### Dependency Graph (SCA-specific)

- **DependencySnapshot**: `id`, `projectID`, `scanID`, `lockfileHash` (SHA-256, for incremental scan matching), `createdAt`
- **Dependency**: `id`, `snapshotID`, `name`, `version`, `isDirect` (bool), `scope` (production/development/optional/peer -- parsed from lockfile `dev`/`optional` flags), `parentName` (nullable), `parentVersion` (nullable), `depth` (int)

This models the full transitive tree with scope awareness. When a transitive dep has a CVE, we walk up the graph to find which direct dep to upgrade. The `scope` propagates from the root: if a direct dep is `development`, all its transitive children inherit `development` scope (even if they also appear as transitive deps of production packages -- in that case, the finding is scoped to `production` since the package is reachable via a production path).

### Supply Chain Signals (SCA-specific)

- **SupplyChainSignal**: `id`, `snapshotID`, `packageName`, `packageVersion`, `signalType` (postinstall_hook/new_package/maintainer_change/typosquat/obfuscated_code), `severity` (high/medium/low), `detail` (JSONB: the specific evidence, e.g. the postinstall script content, the Levenshtein match, days since publish), `createdAt`

These are surfaced as findings with `source=supply_chain` and `scanType=sca`. They do not have CVE IDs but use fingerprint = hash(signalType+packageName+packageVersion).

### Git Integration

- **GitCredential**: see Core GORM Models above. Credentials encrypted with AES-256-GCM at rest using a server-side `ENCRYPTION_KEY` env var. Decrypted only in memory during clone operations. Never returned in API responses (write-only via API, read returns `****` masked).

- **RepoCloneAuditLog**: `id`, `orgID`, `projectID`, `gitCredentialID`, `repoURL`, `action` (clone/webhook_triggered/ownership_verified/clone_failed/access_denied), `detail` (JSONB: clone duration, repo size, files extracted count, error message if failed), `ipAddress`, `userID` (who triggered), `createdAt`

Every clone operation is logged for audit. Admins can view the log in the dashboard. Suspicious patterns (e.g., many access_denied in short time) trigger alerts.

- **WebhookRegistration**: `id`, `orgID`, `projectID`, `provider` (github/gitlab/bitbucket), `webhookSecret` (per-registration HMAC secret for payload verification), `events` (JSONB: ["push", "pull_request"]), `targetBranch` (nullable -- only trigger on pushes to this branch, e.g. "main"), `enabled`, `createdAt`, `updatedAt`

#### Git Clone Security Model

The clone pipeline has multiple layers of protection against abuse:

```go
type CloneRequest struct {
    RepoURL        string
    CredentialID   *uuid.UUID      // nil for public repos
    ProjectID      uuid.UUID
    OrgID          uuid.UUID
    RequestedBy    uuid.UUID       // user who triggered the scan
    Branch         string          // specific branch to clone
}

type CloneResult struct {
    TempDir        string          // path to temp directory (cleaned up after scan)
    LockFiles      map[string][]byte
    SourceFiles    map[string][]byte
    RepoSizeBytes  int64
    FileCount      int
    CloneDurationMs int64
}

type CloneLimits struct {
    MaxRepoSizeBytes    int64         // default: 500MB
    MaxFileCount        int           // default: 50,000
    MaxSourceFileSizeBytes int64      // default: 1MB per file (skip larger files)
    CloneTimeoutSeconds int           // default: 300 (5 min)
    MaxConcurrentClones int           // per org, default: 3
    MaxScansPerHour     int           // per org, default: 30
    MaxMonitoredRepos   int           // per org, depends on plan
}
```

**Security checks executed in order:**

1. **Rate limit check** -- Org hasn't exceeded `MaxScansPerHour`. Returns 429 if exceeded.
2. **Ownership verification** -- Before cloning, call the provider API (GitHub/GitLab/Bitbucket) with the user's credential to verify the authenticated user has at least read access to the repo. If the credential's `ownerUsername` doesn't match any collaborator/member of the repo, reject with 403. This prevents scanning repos you don't own even if you somehow know the URL.
3. **URL validation** -- RepoURL must match `https://(github.com|gitlab.com|bitbucket.org|self-hosted-domain)/owner/repo` pattern. No `file://`, `ssh://` (unless explicitly configured), no IP addresses, no localhost. Reject anything else.
4. **Safe clone** -- Execute git clone with:
   - `--depth 1` (shallow, no history)
   - `--single-branch --branch <branch>` (only the requested branch)
   - `--no-tags` (skip tags)
   - `--config core.hooksPath=/dev/null` (disable all git hooks -- prevents malicious pre-checkout hooks)
   - `--config core.symlinks=false` (don't follow symlinks -- prevents symlink attacks)
   - Environment: `GIT_TERMINAL_PROMPT=0` (never prompt for credentials interactively)
   - Timeout: kill the process after `CloneTimeoutSeconds`
5. **Size check** -- After clone, check total size. If exceeds `MaxRepoSizeBytes`, delete and reject with 413.
6. **File extraction** -- Only extract files matching allowed patterns:
   - Lockfiles (SCA):
     - **Node.js**: `package-lock.json`, `yarn.lock`, `pnpm-lock.yaml`
     - **Python**: `requirements.txt`, `Pipfile.lock`, `poetry.lock`, `setup.py`, `setup.cfg`, `pyproject.toml`
     - **.NET**: `packages.lock.json`, `*.csproj`, `packages.config`, `Directory.Packages.props`
     - **Go**: `go.mod`, `go.sum`
   - Source code (for AI analysis + SAST): `.js`, `.ts`, `.jsx`, `.tsx`, `.mjs`, `.cjs`, `.py`, `.cs`, `.go`, `.java`, `.rb`, `.php`, `.vb`, `.fs`
   - Config: `package.json`, `.vulncheck`, `tsconfig.json`, `nuget.config`, `.npmrc`, `Pipfile`
   - Skip: `node_modules/`, `vendor/`, `dist/`, `bin/`, `obj/`, `.git/`, `__pycache__/`, `.venv/`, `venv/`, binary files, images, videos, files > `MaxSourceFileSizeBytes`
7. **Cleanup** -- Temp directory is deleted immediately after file extraction (before scan starts). Source files are held in memory only for the duration of the scan. No repo data persists on disk.
8. **Audit log** -- Log the clone with all metadata (who, what repo, size, duration, result).

#### Webhook Security

Incoming webhooks (push/PR events from GitHub/GitLab/Bitbucket) are verified:
- **GitHub**: HMAC-SHA256 signature in `X-Hub-Signature-256` header, verified against the per-registration `webhookSecret`.
- **GitLab**: Secret token in `X-Gitlab-Token` header.
- **Bitbucket**: No native signing -- use IP allowlisting (Bitbucket Cloud IPs) + a secret query parameter.
- **Replay protection**: Reject webhooks with timestamps older than 5 minutes (where supported).
- **Event filtering**: Only process events matching the `WebhookRegistration.events` and `targetBranch`. Ignore all other events silently.
- **Dedup**: If a scan is already running for the same project+branch, skip the duplicate trigger (return 200 OK but don't start a new scan).

### License Compliance

- **LicensePolicy**: `id`, `orgID`, `action` (allow/deny/warn), `licenseID` (SPDX identifier, e.g. "GPL-3.0-only"), `reason` (required for deny), `createdBy`, `createdAt`, `updatedAt`
- **LicenseFinding**: `id`, `scanID`, `projectID`, `packageName`, `packageVersion`, `detectedLicense` (SPDX expression, e.g. "MIT", "GPL-2.0-only OR Apache-2.0"), `policyAction` (allow/deny/warn/unknown -- result of policy evaluation), `isDirect` (bool), `dependencyPath` (JSONB: path from direct dep to this package), `source` (lockfile/registry_api/package_json), `status` (open/acknowledged/resolved), `createdAt`

License data is extracted from: (1) lockfile metadata (if available), (2) registry API (`npm view <pkg> license`, PyPI JSON API `license` field, NuGet nuspec `licenseUrl`), (3) `package.json` / `pyproject.toml` / `.nuspec` `license` field in the downloaded package. Dual-licensing (`MIT OR GPL-2.0`) is parsed using SPDX expression parser; the policy evaluator checks if ANY branch of an OR expression satisfies the policy.

### SBOM (Software Bill of Materials)

- **SBOMExport**: `id`, `projectID`, `scanID`, `format` (spdx-json/spdx-tv/cyclonedx-json/cyclonedx-xml), `includeVulns` (bool -- include vulnerability status in SBOM), `fileURL` (path to generated file), `createdAt`

Generated from `DependencySnapshot` + `Dependency` records. Maps our internal dependency model to SPDX/CycloneDX schemas: package name -> `SPDXID`/`bom-ref`, version, license (from LicenseFinding), supplier, dependency relationships (uses `Graph`). Vuln overlay (optional): embeds VEX (Vulnerability Exploitability eXchange) data for each known CVE.

### Automated Fix PRs

- **FixPR**: `id`, `projectID`, `scanID`, `findingFingerprints` (JSONB: list of finding fingerprints this PR addresses), `provider` (github/gitlab/bitbucket), `prURL` (external PR URL after creation), `prNumber` (int), `status` (pending/created/merged/closed/failed), `branchName` (e.g. "vulncheck/fix-CVE-2024-1234"), `changes` (JSONB: list of {file, oldContent, newContent}), `safetyCheckPassed` (bool -- no new vulns introduced), `createdAt`, `updatedAt`

The fix PR pipeline: (1) remediation engine computes upgrade paths, (2) filter to patch/minor risk level only (unless org config allows major), (3) safety check: simulate the upgrade and run SCA on the updated lockfile to verify no new vulns >= current severity, (4) create branch, apply changes to lockfile + manifest, push, create PR via provider API, (5) store `FixPR` record for tracking. Configurable per org: `autoFixPR` (disabled/manual/auto).

### Secrets Detection

Secrets findings use the same `Finding` model with `scanType=secrets`. Additional metadata stored in `raw` JSONB:
- `secretType`: aws_access_key / gcp_service_account / github_token / slack_token / jwt_secret / private_key / database_url / generic_high_entropy
- `maskedValue`: first 4 + "..." + last 4 characters (e.g. "ghp_xxxx...xxxx")
- `entropy`: Shannon entropy score of the detected string
- `confidence`: pattern_match (regex hit) / high_entropy (>4.5 bits/char) / both
- `inGitHistory`: bool (if git history scanning enabled -- was this secret ever committed?)

### Policies (Dual System)

- **IgnorePolicy** (DB): `id`, `orgID`, `projectID` (nullable = org-wide), `matchType` (package/cve/package+cve), `packageName`, `cveID`, `reason` (required), `ignoreUntil` (required), `createdBy`, `createdAt`, `updatedAt`

- **.vulncheck file** (repo, YAML):

```yaml
version: 1
ignore:
  CVE-2021-23337:
    package: lodash
    reason: "No fix available, low risk in our usage context"
    expires: 2026-05-01
  CVE-2024-1234:
    package: "*"
    reason: "False positive confirmed by security team"
    expires: 2026-06-15
```

Policy evaluation merges both sources: DB policies (org-wide + project) + `.vulncheck` file from the scan target. DB policies take precedence on conflicts.

### Jira Integration

- **IntegrationConfig**: `id`, `orgID`, `provider` (jira/github/linear), `baseURL`, `credentials` (encrypted JSONB: API token, project key, etc.), `autoCreateThreshold` (severity: critical/high/medium/low), `enabled` (bool), `createdAt`, `updatedAt`
- **TicketMapping**: `id`, `orgID`, `projectID`, `findingFingerprint`, `provider`, `externalTicketID` (e.g. "VULN-123"), `externalTicketURL`, `externalStatus`, `lastSyncedAt`, `createdAt`, `updatedAt`

Dedup logic: before creating a ticket, check `TicketMapping` for an existing mapping with the same `findingFingerprint` + `projectID`. If found, update instead of creating.

### Monitor Mode

- **MonitoredProject**: `id`, `projectID`, `snapshotID` (FK to latest DependencySnapshot), `checkInterval` (duration), `lastCheckedAt`, `nextCheckAt`, `enabled`, `createdAt`, `updatedAt`

River periodic job queries `MonitoredProject WHERE nextCheckAt <= now() AND enabled`, re-runs vuln lookups against the stored snapshot, and creates new findings if new CVEs are discovered since last check.

### Notifications

- **NotificationConfig**: `id`, `orgID`, `channel` (email/slack), `target` (email address or Slack webhook URL), `events` (JSONB array: scan_completed, critical_finding, monitor_alert), `enabled`, `createdAt`, `updatedAt`

### GitHub PR Checks

- **PRCheckConfig**: `id`, `orgID`, `projectID`, `githubInstallationID` (or PAT), `repoOwner`, `repoName`, `failThreshold` (severity level that fails the check), `commentFindings` (bool), `enabled`, `createdAt`, `updatedAt`

### Reports

- **ReportRequest**: `id`, `orgID`, `projectID` (nullable = org-wide), `format` (pdf/csv), `status` (pending/generating/ready/failed), `filters` (JSONB: date range, severity, scan type), `fileURL` (path to generated file), `requestedBy`, `createdAt`, `completedAt`

### Multi-Ecosystem Parser Interface

```go
// pkg/ecosystem/ecosystem.go
type Ecosystem string

const (
    EcosystemNPM    Ecosystem = "npm"      // Phase E
    EcosystemPip    Ecosystem = "pip"      // Phase E3 (covers pip, pipenv, poetry)
    EcosystemDotNet Ecosystem = "dotnet"   // Phase E4 (NuGet)
    EcosystemGoMod  Ecosystem = "gomod"    // Phase E4
    EcosystemMaven  Ecosystem = "maven"    // future
    EcosystemRuby   Ecosystem = "ruby"     // future
    EcosystemCargo  Ecosystem = "cargo"    // future
)

type DepScope string

const (
    ScopeProduction  DepScope = "production"
    ScopeDevelopment DepScope = "development"
    ScopeOptional    DepScope = "optional"
    ScopePeer        DepScope = "peer"
)

type PackageRef struct {
    Name     string
    Version  string
    IsDirect bool
    Scope    DepScope   // parsed from lockfile dev/optional/peer flags
}

type DependencyTree struct {
    Ecosystem    Ecosystem
    Packages     []PackageRef
    Graph        map[string][]string // package@version -> children
    LockfileHash string              // SHA-256 of parsed lockfile (for incremental scan)
}

type EcosystemParser interface {
    Ecosystem() Ecosystem
    DetectFiles() []string                             // e.g. ["package-lock.json", "yarn.lock"]
    Parse(files map[string][]byte) (*DependencyTree, error)
}
```

The SCA scanner calls the parser registry, which selects the correct parser based on detected files. Ecosystem delivery order:
- **Phase E**: npm (`package-lock.json`, `yarn.lock`, `pnpm-lock.yaml`)
- **Phase E3**: Python (`requirements.txt`, `Pipfile.lock`, `poetry.lock`)
- **Phase E4**: .NET (`packages.lock.json`, `*.csproj`, `packages.config`) + Go (`go.mod`, `go.sum`)
- **Future**: Maven (`pom.xml`), Ruby (`Gemfile.lock`), Rust (`Cargo.lock`)

### Private Registry Config

- **RegistryConfig**: `id`, `orgID`, `ecosystem` (npm/pip/etc.), `registryURL`, `authType` (token/basic/none), `credentials` (encrypted JSONB), `scope` (npm scope, e.g. "@mycompany"), `createdAt`, `updatedAt`

Used by the SCA scanner when resolving packages and querying advisories for private packages.

### Behavioral Package Analysis

- **PackageBehaviorProfile**: `id`, `ecosystem`, `packageName`, `packageVersion`, `tarballURL`, `tarballHash` (SHA-256 of the analyzed tarball), `capabilities` (JSONB: `{network: bool, filesystem: bool, shell: bool, env: bool, unsafe_eval: bool, native_binding: bool, obfuscated: bool, install_hook_shell: bool}`), `capabilityDetails` (JSONB: evidence per capability -- file:line references, offending AST snippets), `installScriptAnalysis` (JSONB: extracted install hook content + semantic analysis verdict), `analyzedAt`, `analyzerVersion` (for re-analysis when analyzer improves)
- **CapabilityExpansion**: `id`, `packageName`, `fromVersion`, `toVersion`, `addedCapabilities` (JSONB array), `removedCapabilities` (JSONB array), `severity` (info/low/medium/high/critical derived from which capabilities changed), `surfacedAsFindingID` (FK to Finding if promoted), `detectedAt`
- **HallucinationCandidate**: `id`, `ecosystem`, `packageName`, `source` (huntr/snyk_advisor/internal_heuristic/user_report), `confidence` (0.0-1.0), `evidence` (JSONB), `firstSeenAt`, `refreshedAt`

Profiles are built asynchronously via `BehaviorProfileJob` when a new package version is encountered. Cached aggressively (keyed by tarballHash). Used by the SCA scanner after lockfile parsing: for each (pkg, version) the scanner looks up the profile, compares to the previous version, and emits `capability_expansion` findings on diff.

### Call-Graph Reachability

- **CallGraph**: `id`, `scanID`, `ecosystem`, `rootPath`, `nodeCount`, `edgeCount`, `entryPointCount`, `builtAt`, `builderVersion`, `partial` (bool -- true if the build hit timeout/memory limits and the graph is incomplete)
- **CallGraphNode**: `id`, `callGraphID`, `packageName` (empty string for user code), `packageVersion`, `filePath`, `functionName`, `isEntryPoint` (bool), `entryPointKind` (http_route / cli_main / message_handler / exported_api / test / unknown)
- **ReachabilityPath**: `id`, `findingID`, `callGraphID`, `pathNodes` (JSONB ordered list of CallGraphNode IDs), `reachable` (bool), `confidence` (direct/indirect/partial -- partial when graph is incomplete), `analyzedAt`

The call-graph is built once per scan, per ecosystem. For each SCA finding with a known vulnerable function (extracted from CVE metadata or GHSA details), the tracer BFS-searches for a path from any entry point. Output drives the `reachability` field on Finding: `confirmed` (path exists from entry point via user code), `potentially_reachable` (path exists but graph is partial), `unreachable` (no path, graph is complete). When the call-graph engine cannot analyze a package (dynamic dispatch heavy, obfuscated, missing types), it returns `partial=true` and the AI Impact Analyzer is invoked as fallback.

### Container / OCI Scanning

- **ContainerImage**: `id`, `projectID`, `scanID`, `registry` (e.g. "docker.io", "ghcr.io"), `repository`, `tag`, `digest` (sha256:...), `os` (alpine/debian/ubuntu/distroless/scratch), `osVersion`, `size` (bytes), `layerCount`, `baseImage` (nullable -- detected base), `recommendedBaseImage` (nullable -- suggested minimal alternative), `cosignVerified` (bool), `slsaProvenance` (JSONB: SLSA attestation if available), `createdAt`
- **ContainerLayer**: `id`, `containerImageID`, `layerDigest`, `layerIndex`, `command` (Dockerfile command that produced the layer), `size`, `osPackageCount`, `createdAt`
- **ContainerFinding**: ContainerFindings extend the regular `Finding` model with additional fields stored in `raw` JSONB: `layerIndex` (which layer introduced the vuln), `osPackage` (apt/apk/rpm package name), `osPackageVersion`, `fixedInLayer` (nullable -- if a later layer removed/upgraded it), `dockerfileLine` (for Dockerfile SAST findings)

### SLSA / Sigstore Verification

- **SLSAAttestation**: `id`, `ecosystem`, `packageName`, `packageVersion`, `provenanceStatus` (verified/missing/mismatch/invalid_signature), `builderID` (e.g. "https://github.com/actions/runner@sha256:..."), `sourceRepo` (claimed git source), `sourceCommit` (claimed commit sha), `buildConfigSource` (claimed workflow file), `sigstoreRekorEntry` (URL to Rekor transparency log entry), `verifiedAt`, `rawAttestation` (JSONB)

Surfaced as supply-chain findings with `signal_type` in {`missing_provenance`, `provenance_mismatch`, `unsigned_artifact`}.

### Upstream VEX Consumer

- **VEXFeed**: `id`, `orgID` (nullable -- null for platform-wide default feeds), `name`, `url`, `format` (openvex/csaf), `enabled`, `lastFetchedAt`, `lastFetchStatus`, `signatureRequired` (bool -- only accept signed statements), `trustedSigners` (JSONB array of sigstore identities / x509 subjects)
- **VEXStatement**: `id`, `feedID`, `statementID` (upstream-assigned ID), `cveID`, `ghsaID`, `packageName`, `packageVersion` (specific version or range), `status` (not_affected / affected / fixed / under_investigation), `justification` (upstream text: e.g. "vulnerable_code_not_in_execute_path"), `sourceURL`, `publishedAt`, `signatureValid` (bool, nullable), `signerIdentity` (nullable), `consumedAt`
- **VEXApplication**: `id`, `scanID`, `findingID`, `vexStatementID`, `action` (auto_dismissed / downgraded / flagged), `appliedAt`

VEX statements that resolve to `not_affected` or `fixed` automatically dismiss the corresponding finding as `source=upstream_vex` ignore-policy. Provenance is preserved -- users see which VEX statement and feed source caused the dismissal.

### ASPM Attack Stories

- **AttackStory**: `id`, `scanID`, `projectID`, `title` (AI-generated short summary), `description` (AI-generated narrative of the attack chain), `combinedSeverity` (max of constituent findings, boosted by chain length), `findingIDs` (JSONB array), `entities` (JSONB: shared entities that link the findings -- files, packages, services, routes), `edges` (JSONB: typed edges between findings: reaches / uses / exposes / chains_into), `status` (open / acknowledged / mitigated), `createdAt`, `updatedAt`

Built by the post-scan `AspmCorrelationJob`. Surfaces in dashboard as a prioritized "Attack Stories" tab on top of the raw findings list.

### AI Verdict Feedback Loop

- **AIVerdictFeedback**: `id`, `findingID`, `orgID`, `userID`, `originalVerdict` (JSONB snapshot of the AIVerdict at feedback time), `correctedImpactLevel` (AIImpactLevel), `correctedExplanation` (nullable -- user's own explanation), `feedbackKind` (wrong_impact / wrong_explanation / missed_usage / spurious_usage / other), `note` (nullable -- free-text comment), `createdAt`
- **PromptQualityMetric**: `id`, `period` (e.g. "2026-04" monthly bucket), `model` (e.g. "gpt-4o-mini"), `promptTemplate` (e.g. "prototype_pollution_v3"), `totalVerdicts`, `correctedCount`, `accuracyRate`, `perCWEAccuracy` (JSONB), `perLanguageAccuracy` (JSONB), `computedAt`

Aggregated daily by a background job. Exported as anonymized eval set (opt-in per org) for prompt-tuning iterations.

## API Endpoints

### Auth
- `POST /api/v1/auth/register` -- Register org + admin user
- `POST /api/v1/auth/login` -- Login, returns JWT
- `POST /api/v1/auth/invite` -- Invite user to org

### Projects
- `POST /api/v1/projects` -- Create project
- `GET /api/v1/projects` -- List projects
- `GET /api/v1/projects/:id` -- Get project with latest scan summary
- `PUT /api/v1/projects/:id` -- Update project
- `DELETE /api/v1/projects/:id` -- Delete project

### Git Credentials
- `POST /api/v1/git-credentials` -- Add git credential (PAT or OAuth token). Server validates token against provider API (checks scope, fetches username). Credentials are encrypted at rest.
- `GET /api/v1/git-credentials` -- List credentials (tokens masked as `ghp_****`). Returns provider, label, ownerUsername, valid, lastValidatedAt.
- `PUT /api/v1/git-credentials/:id` -- Update (rotate token)
- `DELETE /api/v1/git-credentials/:id` -- Revoke and delete credential
- `POST /api/v1/git-credentials/:id/validate` -- Re-validate token against provider API (check if still valid, update scope)
- `GET /api/v1/git-credentials/:id/repos` -- List repos accessible with this credential (proxied from provider API, paginated). Allows user to pick a repo when creating a project.

### Webhooks
- `POST /api/v1/projects/:id/webhooks` -- Register webhook for a project. Returns the webhook URL and secret to configure in GitHub/GitLab/Bitbucket.
- `GET /api/v1/projects/:id/webhooks` -- List webhook registrations
- `PUT /api/v1/webhooks/:id` -- Update (change target branch, toggle events)
- `DELETE /api/v1/webhooks/:id` -- Remove webhook registration
- `POST /api/v1/webhooks/github` -- GitHub webhook receiver (push + PR events, HMAC-SHA256 verified)
- `POST /api/v1/webhooks/gitlab` -- GitLab webhook receiver (push + MR events, token verified)
- `POST /api/v1/webhooks/bitbucket` -- Bitbucket webhook receiver (push + PR events)

### Audit Log
- `GET /api/v1/audit-log` -- List clone audit log entries (admin only, paginated, filterable by project/user/action)

### Scans
- `POST /api/v1/projects/:id/scans` -- Trigger scan. Modes: (1) multipart file upload, (2) clone from project's configured repoURL + gitCredential. If project has repoURL configured, clone is the default -- no upload needed.
- `GET /api/v1/projects/:id/scans` -- List scans
- `GET /api/v1/scans/:id` -- Scan detail + summary
- `GET /api/v1/scans/:id/findings` -- Findings (policies applied, pagination, severity filter)
- `GET /api/v1/scans/:id/findings/:findingId/remediation` -- Upgrade path for a specific finding

### AI Impact Analysis
- `GET /api/v1/scans/:id/findings/:findingId/ai-verdict` -- Get AI impact verdict for a specific finding
- `POST /api/v1/scans/:id/findings/:findingId/ai-analyze` -- Re-run AI analysis for a specific finding (useful if code changed)
- `POST /api/v1/scans/:id/ai-analyze` -- Run AI analysis for all findings in a scan (async, returns job ID)
- `GET /api/v1/scans/:id/ai-status` -- AI analysis job status (pending/running/completed)

### License Compliance
- `POST /api/v1/license-policies` -- Create license policy (allow/deny/warn for a specific SPDX license)
- `GET /api/v1/license-policies` -- List license policies
- `PUT /api/v1/license-policies/:id` -- Update
- `DELETE /api/v1/license-policies/:id` -- Remove
- `GET /api/v1/scans/:id/licenses` -- List license findings for a scan (with policy evaluation results)
- `GET /api/v1/projects/:id/license-summary` -- Aggregate license breakdown for latest scan (pie chart data)

### SBOM
- `POST /api/v1/projects/:id/sbom` -- Generate SBOM (format: spdx-json/spdx-tv/cyclonedx-json/cyclonedx-xml, includeVulns: bool)
- `GET /api/v1/sbom/:id` -- SBOM status + download URL
- `GET /api/v1/sbom/:id/download` -- Download SBOM file

### Fix PRs
- `POST /api/v1/scans/:id/fix-prs` -- Create fix PRs for all fixable findings in a scan (async)
- `POST /api/v1/findings/:findingId/fix-pr` -- Create fix PR for a specific finding
- `GET /api/v1/projects/:id/fix-prs` -- List fix PRs for a project (with status: pending/created/merged/closed)
- `GET /api/v1/fix-prs/:id` -- Fix PR detail (changes, safety check result, PR URL)

### Policies
- `POST /api/v1/policies` -- Create ignore policy (DB)
- `GET /api/v1/policies` -- List policies (active/expired filter)
- `GET /api/v1/policies/:id` -- Get policy
- `PUT /api/v1/policies/:id` -- Update (extend expiry, edit reason)
- `DELETE /api/v1/policies/:id` -- Remove policy
- `POST /api/v1/projects/:id/policies/import` -- Import from `.vulncheck` file
- `GET /api/v1/projects/:id/policies/export` -- Export to `.vulncheck` format

### Integrations
- `POST /api/v1/integrations` -- Configure integration (Jira credentials, project, threshold)
- `GET /api/v1/integrations` -- List configured integrations
- `PUT /api/v1/integrations/:id` -- Update integration config
- `DELETE /api/v1/integrations/:id` -- Remove integration
- `POST /api/v1/integrations/:id/test` -- Test connection
- `GET /api/v1/integrations/:id/tickets` -- List ticket mappings

### Monitor
- `POST /api/v1/projects/:id/monitor` -- Enable monitoring (snapshots current deps)
- `DELETE /api/v1/projects/:id/monitor` -- Disable monitoring
- `GET /api/v1/projects/:id/monitor` -- Monitor status + last check + new findings

### Notifications
- `POST /api/v1/notifications` -- Configure notification channel (email/Slack + events)
- `GET /api/v1/notifications` -- List notification configs
- `PUT /api/v1/notifications/:id` -- Update
- `DELETE /api/v1/notifications/:id` -- Remove
- `POST /api/v1/notifications/:id/test` -- Send test notification

### GitHub PR Checks
- `POST /api/v1/pr-checks` -- Configure PR check integration (GitHub PAT, repo, threshold)
- `GET /api/v1/pr-checks` -- List PR check configs
- `PUT /api/v1/pr-checks/:id` -- Update
- `DELETE /api/v1/pr-checks/:id` -- Remove
- `POST /api/v1/pr-checks/webhook` -- GitHub webhook receiver (PR opened/synchronized events)

### Reports
- `POST /api/v1/reports` -- Request report generation (async, returns report ID)
- `GET /api/v1/reports` -- List reports
- `GET /api/v1/reports/:id` -- Report status + download URL
- `GET /api/v1/reports/:id/download` -- Download PDF/CSV file

### Private Registries
- `POST /api/v1/registries` -- Configure private registry
- `GET /api/v1/registries` -- List registries
- `PUT /api/v1/registries/:id` -- Update
- `DELETE /api/v1/registries/:id` -- Remove
- `POST /api/v1/registries/:id/test` -- Test connectivity

### Behavioral Analysis
- `GET /api/v1/packages/:ecosystem/:name/:version/behavior` -- Get PackageBehaviorProfile (capabilities + evidence)
- `GET /api/v1/packages/:ecosystem/:name/capability-history` -- Timeline of capability changes across versions
- `POST /api/v1/packages/:ecosystem/:name/:version/rebuild-profile` -- Force rebuild (admin, after analyzer upgrade)
- `GET /api/v1/hallucination-candidates` -- List hallucination candidate packages (admin)
- `POST /api/v1/hallucination-candidates/refresh` -- Trigger feed refresh (admin)

### Call-Graph Reachability
- `GET /api/v1/scans/:id/callgraph` -- Call-graph metadata (node/edge counts, entry-point count, partial flag)
- `GET /api/v1/scans/:id/findings/:findingId/reachability-path` -- Detailed reachability path from entry point to vulnerable function
- `POST /api/v1/scans/:id/callgraph/rebuild` -- Rebuild call-graph (after code changes)

### Container / OCI Scanning
- `POST /api/v1/projects/:id/container-scans` -- Trigger container image scan (body: `{registry, repository, tag, credentialID?}`)
- `GET /api/v1/container-scans/:id` -- Scan status + summary
- `GET /api/v1/container-scans/:id/findings` -- Container findings (OS packages + Dockerfile rules)
- `GET /api/v1/container-scans/:id/layers` -- Layer-by-layer breakdown
- `GET /api/v1/container-scans/:id/base-image-recommendation` -- Minimal-image recommendation with size estimate
- `GET /api/v1/container-scans/:id/sbom` -- Container-specific SBOM (CycloneDX with layer breakdown)

### SLSA / Sigstore
- `GET /api/v1/scans/:id/slsa-attestations` -- SLSA provenance verification results for all deps in scan
- `GET /api/v1/packages/:ecosystem/:name/:version/slsa` -- SLSA attestation for a specific package

### Upstream VEX
- `POST /api/v1/vex-feeds` -- Configure VEX feed (URL, format, signature requirement)
- `GET /api/v1/vex-feeds` -- List configured feeds
- `PUT /api/v1/vex-feeds/:id` -- Update feed config
- `DELETE /api/v1/vex-feeds/:id` -- Remove feed
- `POST /api/v1/vex-feeds/:id/refresh` -- Trigger immediate feed refresh
- `GET /api/v1/scans/:id/vex-applications` -- List VEX dismissals applied in this scan (with provenance)

### Rego Policy
- `POST /api/v1/rego-policies` -- Upload Rego policy file (body: `{name, regoContent}`)
- `GET /api/v1/rego-policies` -- List policies
- `PUT /api/v1/rego-policies/:id` -- Update
- `DELETE /api/v1/rego-policies/:id` -- Remove
- `POST /api/v1/rego-policies/:id/evaluate` -- Dry-run evaluate against a specific scan (for testing)
- `GET /api/v1/scans/:id/rego-evaluations` -- Rego evaluation results for a scan (block/warn/notify outcomes)

### PR Baseline Diff
- `GET /api/v1/projects/:id/pr-scans/:pr/diff` -- Get baseline diff for a PR (NEW findings + RESOLVED findings + lockfile diff)

### ASPM Attack Stories
- `GET /api/v1/scans/:id/attack-stories` -- List correlated attack stories (with combined severity, entities, chain)
- `GET /api/v1/attack-stories/:id` -- Attack story detail (full chain graph + findings)
- `PUT /api/v1/attack-stories/:id/status` -- Update status (acknowledged/mitigated)

### AI Verdict Feedback
- `POST /api/v1/findings/:findingId/ai-feedback` -- Submit feedback on AI verdict (body: `{feedbackKind, correctedImpactLevel, note?}`)
- `GET /api/v1/ai-feedback/metrics` -- Prompt quality metrics (admin only, filter by period/model)
- `POST /api/v1/ai-feedback/export-eval-set` -- Export anonymized eval set (admin only, org opt-in required)

### IDE / LSP
- The LSP server is a local process (`vulncheck lsp`), not an HTTP endpoint. The extension communicates with the remote API via the user's authenticated CLI token for heavy operations (AI verdicts, full scans).
- `POST /api/v1/ide/quick-scan` -- Lightweight scan endpoint optimized for IDE save-time hints (returns cached findings by lockfile hash, skips heavy AI analysis by default)

### CLI Commands (Snyk-like, with context awareness)
- `vulncheck test --path .` -- Scan local directory. Default: production-scope only. Exit code 1 if effective severity >= threshold (configurable).
- `vulncheck test --repo https://github.com/owner/repo` -- Scan a remote repo (clones via server-side git integration). Uses configured git credentials.
- `vulncheck test --repo https://github.com/owner/repo --branch develop` -- Scan a specific branch
- `vulncheck test --path . --include-dev` -- Include development dependencies in scan output
- `vulncheck test --path . --all` -- Show all findings including unreachable/informational
- `vulncheck test --path . --json` -- Machine-readable JSON output for CI pipelines
- `vulncheck test --path . --fail-on high` -- Set minimum effective severity that triggers exit code 1
- `vulncheck fix --path .` -- Apply safe (patch/minor) fixes. Never applies major upgrades without `--allow-major`.
- `vulncheck fix --path . --dry-run` -- Show what would be changed without applying
- `vulncheck fix --path . --allow-major` -- Allow major version upgrades (with confirmation prompt)
- `vulncheck monitor --path .` -- Snapshot deps for continuous monitoring (incremental by default)
- `vulncheck ignore --cve CVE-2021-xxx --reason "..." --until 2026-05-01` -- Add to `.vulncheck` file
- `vulncheck policy list` -- Show active policies
- `vulncheck auth login` -- Authenticate CLI with server
- `vulncheck config set api-url https://...` -- Configure server URL
- `vulncheck test --path . --no-ai` -- Skip AI impact analysis (faster, scope-based reachability only)
- `vulncheck test --path . --ai-model claude-sonnet` -- Override LLM model for this scan
- `vulncheck explain --finding <id>` -- Re-run AI analysis for a specific finding and print detailed explanation
- `vulncheck license --path .` -- Run license compliance scan, show violations against org policy
- `vulncheck license --path . --policy deny=GPL-3.0-only,AGPL-3.0-only` -- Inline license policy (without server)
- `vulncheck sbom --path . --format spdx-json` -- Generate SBOM to stdout or file (spdx-json/spdx-tv/cyclonedx-json/cyclonedx-xml)
- `vulncheck sbom --path . --format cyclonedx-json --include-vulns` -- SBOM with vulnerability overlay (VEX)
- `vulncheck fix-pr --path . --repo https://github.com/owner/repo` -- Create fix PRs for all safe fixes (patch/minor)
- `vulncheck secrets --path .` -- Run dedicated secrets scan (100+ patterns + entropy analysis)
- `vulncheck secrets --path . --scan-history` -- Also scan git history for previously committed secrets
- `vulncheck supply-chain --path .` -- Run supply chain analysis only (hooks, typosquat, new packages)
- `vulncheck behavior --package lodash@4.17.22` -- Show behavioral capability profile for a specific package
- `vulncheck behavior --path . --diff-against-prev` -- Show capability expansions in current lockfile vs previous scan
- `vulncheck reachability --path .` -- Build call-graph and show reachability verdict per finding (no LLM cost)
- `vulncheck reachability --path . --finding <id> --trace` -- Print the full path from entry point to vulnerable function for a specific finding
- `vulncheck container scan <image:tag>` -- Scan a container image (pull + layers + OS packages + Dockerfile + Sigstore)
- `vulncheck container scan <image:tag> --recommend-base` -- Include minimal-image recommendation
- `vulncheck container sbom <image:tag> --format cyclonedx-json` -- Generate container SBOM with layer breakdown
- `vulncheck verify-provenance --path .` -- Check SLSA provenance + Sigstore signatures for all deps
- `vulncheck verify-provenance <image:tag>` -- Same for container image
- `vulncheck vex consume --feed https://kubernetes.io/security/vex.json` -- Test-consume a VEX feed, show applied statements
- `vulncheck policy eval --rego policies/my_gate.rego --scan <id>` -- Dry-run a Rego policy against a scan
- `vulncheck pr-diff --base <sha> --head <sha>` -- Show PR-scoped findings (NEW + RESOLVED) as machine-readable JSON for custom CI integration
- `vulncheck attack-stories --scan <id>` -- Print correlated attack stories for a scan
- `vulncheck feedback --finding <id> --corrected-impact no_impact --note "..."` -- Submit AI verdict feedback from CLI
- `vulncheck lsp` -- Run as a Language Server (used by VSCode/JetBrains extensions, not typically invoked by humans)

## Cross-Cutting Concerns

- **Environment variables**: `DATABASE_URL`, `JWT_SECRET`, `JWT_EXPIRY`, `NVD_API_KEY` (optional, higher rate limits), `OSV_API_URL`, `SERVER_PORT`, `LOG_LEVEL`, `ENCRYPTION_KEY` (AES-256-GCM key for git credential encryption), `GIT_CLONE_TIMEOUT` (default: 300s), `GIT_MAX_REPO_SIZE` (default: 500MB), `GIT_MAX_SCANS_PER_HOUR` (default: 30), `GIT_MAX_CONCURRENT_CLONES` (default: 3), `GIT_ALLOWED_HOSTS` (default: github.com,gitlab.com,bitbucket.org -- comma-separated, add self-hosted domains here), `SMTP_HOST`, `SMTP_PORT`, `SMTP_USER`, `SMTP_PASS`, `SLACK_WEBHOOK_URL`, `GITHUB_TOKEN` (also used for GHSA GraphQL queries -- 5000 req/hour), `VULNCHECK_API_KEY` (optional, commercial -- enables exploit intelligence, EPSS scores, KEV matching), `TRIVY_DB_PATH` (optional -- path to offline Trivy DB, auto-downloaded if not set), `TRIVY_DB_UPDATE_INTERVAL` (default: 24h -- how often to refresh Trivy DB), `SEMGREP_RULES_DIR` (default: `rules/semgrep/` -- path to custom Semgrep rules), `LLM_PROVIDER` (openai/anthropic/custom), `LLM_API_KEY`, `LLM_MODEL` (default: gpt-4o-mini), `LLM_BASE_URL` (for self-hosted), `AI_ENABLED` (default: true), `BEHAVIOR_ANALYSIS_ENABLED` (default: true -- enables behavioral package analysis), `BEHAVIOR_TARBALL_CACHE_DIR` (default: `/var/cache/vulncheck/tarballs`), `BEHAVIOR_MAX_TARBALL_SIZE` (default: 100MB -- skip analysis of huge packages), `BEHAVIOR_SANDBOX_ENABLED` (default: true -- unpack tarballs in sandboxed temp dir with no network), `HALLUCINATION_FEED_URLS` (comma-separated feed URLs for hallucination candidates), `HALLUCINATION_FEED_INTERVAL` (default: 6h), `CALLGRAPH_ENABLED` (default: true), `CALLGRAPH_TIMEOUT` (default: 120s per ecosystem), `CALLGRAPH_MAX_MEMORY_MB` (default: 2048), `CALLGRAPH_NODE_JS_BINARY` (path to node for ts-morph subprocess), `CALLGRAPH_PYTHON_BINARY`, `CALLGRAPH_DOTNET_BINARY`, `CONTAINER_SCAN_ENABLED` (default: true), `CONTAINER_REGISTRY_AUTH` (JSONB: per-registry auth config), `COSIGN_PUBLIC_KEYS_DIR` (for verifying container signatures with static keys alongside Sigstore keyless), `SIGSTORE_REKOR_URL` (default: https://rekor.sigstore.dev), `SLSA_ENABLED` (default: true), `VEX_FEED_URLS` (comma-separated default platform-wide VEX feeds), `VEX_FEED_REFRESH_INTERVAL` (default: 12h), `VEX_SIGNATURE_REQUIRED` (default: false -- true means reject unsigned VEX statements), `REGO_POLICIES_DIR` (path where uploaded .rego files are materialized), `REGO_EVAL_TIMEOUT` (default: 30s), `PR_DIFF_ENABLED` (default: true), `ASPM_CORRELATION_ENABLED` (default: true), `ASPM_MAX_CHAIN_DEPTH` (default: 6 -- bound attack-graph traversal), `AI_FEEDBACK_EVAL_EXPORT_ENABLED` (default: false -- org opt-in), `LSP_PORT` (default: stdio -- LSP uses stdio transport by default), `IDE_QUICK_SCAN_CACHE_TTL` (default: 5m -- how long to cache quick-scan results by lockfile hash)
- **Constants**: Scan types, severities (raw + effective), scan statuses, policy match types, notification event types, ecosystem types, **dependency scopes** (production/development/optional/peer), **reachability levels** (confirmed/potentially_reachable/unreachable), **fix risk levels** (patch/minor/major/no_fix), **supply chain signal types** -- all in `internal/domain/`
- **Shared types**: `ErrorResponse`, `PaginatedResponse`, `SuccessResponse` -- in `pkg/httputil/`
- **Config**: All tunables (API timeouts, scan concurrency limits, vuln cache TTLs, monitor check intervals, ticket sync intervals, upload size limits) in `internal/config/`
- **Migrations**: Initial Goose migration covering all tables (org, user, project, scan, finding, dependency_snapshot, dependency, supply_chain_signal, ignore_policy, integration_config, ticket_mapping, monitored_project, notification_config, pr_check_config, report_request, registry_config, **git_credential**, **repo_clone_audit_log**, **webhook_registration**, **package_behavior_profile**, **capability_expansion**, **hallucination_candidate**, **call_graph**, **call_graph_node**, **reachability_path**, **container_image**, **container_layer**, **slsa_attestation**, **vex_feed**, **vex_statement**, **vex_application**, **attack_story**, **ai_verdict_feedback**, **prompt_quality_metric**, **rego_policy**)
- **Background jobs** (River): `ScanJob`, `AIAnalysisJob`, `MonitorJob`, `TicketSyncJob`, `NotificationJob`, `ReportJob`, `GitCredentialValidationJob` (periodic: re-validate all git credentials against provider APIs, mark expired tokens), `BehaviorProfileJob`, `HallucinationFeedJob`, `VEXFeedJob`, `CallGraphJob`, `ContainerScanJob`, `AspmCorrelationJob`, `PromptQualityAggregationJob`

## Scan + Remediation Flow (SCA)

```mermaid
flowchart TD
  Trigger["Scan triggered (API / webhook / CLI)"] --> Source{"Source?"}
  Source -->|"File upload"| Input
  Source -->|"Git repo"| RateLimit["Rate limit check (org quota)"]
  RateLimit --> Ownership["Verify ownership via provider API"]
  Ownership --> Clone["Safe clone (shallow, no hooks, timeout)"]
  Clone --> SizeCheck{"Repo size OK?"}
  SizeCheck -->|No| Reject["Reject 413 + audit log"]
  SizeCheck -->|Yes| Extract["Extract lockfiles + source code"]
  Extract --> Cleanup["Delete temp clone dir"]
  Cleanup --> Input
  Input["package-lock.json + source code"] --> HashCheck{"Lockfile hash changed?"}
  HashCheck -->|No| IncrOnly["Incremental: check only new CVEs since last scan"]
  HashCheck -->|Yes| Parse["Parse lockfile with scope tags"]
  IncrOnly --> Lookup
  Parse --> Graph["Build dependency graph (direct + transitive + scope)"]
  Graph --> Snapshot["Save DependencySnapshot with lockfileHash"]
  Snapshot --> Behavior["Behavioral analysis (capabilities + diff + hallucination check)"]
  Behavior --> SupplyChain["Supply chain analysis (hooks, typosquat, new pkg, maintainer, SLSA/Sigstore)"]
  SupplyChain --> Lookup["Lookup vulns: npm + OSV + NVD + GHSA + VulnCheck + Trivy"]
  Lookup --> Dedup["Deduplicate across sources"]
  Dedup --> VEXApply["Apply upstream VEX statements (auto-dismiss not_affected)"]
  VEXApply --> Context["Context analysis: assign scope + reachability (scope-based)"]
  Context --> CGEnabled{"Call-graph enabled?"}
  CGEnabled -->|Yes| CGBuild["Build call-graph + trace reachability paths"]
  CGEnabled -->|No| AIEnabled
  CGBuild --> AIEnabled{"AI analysis enabled?"}
  AIEnabled -->|Yes| AIIndex["Index source code: find imports/calls of vulnerable packages"]
  AIEnabled -->|No| EffSev
  AIIndex --> AIExtract["Extract code snippets + repo-level context (framework, routes, DI)"]
  AIExtract --> AIPrompt["Send CVE + code context to LLM"]
  AIPrompt --> AIVerdict["Parse AI verdict: impact level + explanation"]
  AIVerdict --> EffSev["Compute effectiveSeverity (scope + reachability + AI verdict + exploit intel)"]
  EffSev --> Remediate["Compute upgrade paths + fix risk level"]
  Remediate --> PolicyFilter["Apply policies (DB + .vulncheck + Rego)"]
  PolicyFilter --> Findings["Store findings with AI verdicts + reachability paths"]
  Findings --> ASPM["ASPM correlation: build attack stories across scanners"]
  ASPM --> PRDiff{"Triggered by PR?"}
  PRDiff -->|Yes| Baseline["Diff against last defaultBranch scan -> post NEW/RESOLVED only"]
  PRDiff -->|No| Threshold
  Baseline --> Threshold
  Threshold{"effectiveSeverity >= threshold?"}
  Threshold -->|Yes| TicketCheck{"Ticket exists?"}
  Threshold -->|No| Done["Done"]
  TicketCheck -->|No| CreateTicket["Create Jira ticket (includes AI verdict + attack story if any)"]
  TicketCheck -->|Yes| UpdateTicket["Update existing ticket"]
  CreateTicket --> Done
  UpdateTicket --> Done
```

### Context Analysis (addresses npm audit's biggest flaw)

After deduplication, each finding goes through a two-stage context analysis:

**Stage 1 -- Scope-Based Analysis (fast, always runs):**
1. **Scope assignment**: Walk the dependency graph from the vulnerable package up to all root paths. If ALL paths go through dev-only direct deps, scope = `development`. If ANY path goes through a production direct dep, scope = `production`.
2. **Reachability classification** (heuristic): `unreachable` (dev/build-only scope), `potentially_reachable` (production scope, call path not verified).

This alone eliminates the "browserslist ReDoS" class of false positives that plagues npm audit.

**Stage 2 -- AI Impact Analysis (deep, optional, runs when source code is available):**
3. **Code indexing**: Scan all user source files (`.js`, `.ts`, `.jsx`, `.tsx`, `.mjs`, `.cjs`) for imports/requires of the vulnerable package. Build an import map: which files import which exports from the package.
4. **Snippet extraction**: For each usage site, extract 10-15 lines of context around the call. Trace the data flow: where does the input to the vulnerable function come from? (HTTP request body, query params, database, hardcoded value, config file, etc.)
5. **LLM analysis**: Construct a prompt containing: (a) the CVE description and affected function, (b) the user's code snippets, (c) the data flow context. Send to the configured LLM. Parse the structured response.
6. **Verdict**: The AI returns an `AIVerdict` with impact level, explanation, and confidence score. This overrides the scope-based reachability:
   - `no_impact` -> reachability becomes `unreachable` regardless of scope
   - `critical_impact` -> reachability becomes `confirmed`
   - `moderate_impact` / `low_impact` -> reachability stays `potentially_reachable` but effective severity is adjusted
7. **Effective severity**: Apply downgrade/upgrade rules combining scope, reachability, AND AI verdict (see Domain Models section above).

**When AI analysis is disabled** (via `--no-ai` flag or org config), the system falls back to Stage 1 only. This keeps CI/CD fast when speed is preferred over depth.

### Supply Chain Analysis (NEW -- catches what CVE databases miss)

Before vulnerability lookup, the dependency tree is scanned for suspicious patterns:

1. **Postinstall hooks**: Check `package.json` of each dependency for `preinstall`/`install`/`postinstall` scripts that execute shell commands (`sh`, `bash`, `node -e`, `eval`). Flag as `high` if the script downloads external resources or accesses env vars.
2. **New package detection**: Query npm registry for publish date. Packages published <7 days ago are flagged as `medium` (new package risk).
3. **Maintainer changes**: Compare current maintainer list against known-good state (from previous snapshot). Flag if maintainer changed since last scan.
4. **Typosquatting**: Compute Levenshtein distance between each dependency name and the top-1000 npm packages. Flag if distance <= 2 and the package is not in the top-1000 list itself.
5. **Obfuscated code**: (Future) Download and scan tarball for common obfuscation patterns.

These signals are stored as `SupplyChainSignal` records and also surfaced as findings with `source=supply_chain`.

### Upgrade Path Remediation (ENHANCED -- safe fixes)

When a transitive dependency (e.g., `qs@6.5.2` at depth 3) has a CVE:
1. Find the fix version from advisory data (`qs@6.5.3`)
2. Walk the dependency graph upward to find which **direct** dependency pulls in `qs`
3. Check if a newer version of that direct dependency uses the fixed `qs` version
4. **Classify fix risk**: Compare current version to fix version using semver:
   - `patch`: 4.17.1 -> 4.17.3 (safe, auto-applicable)
   - `minor`: 4.17.1 -> 4.18.0 (likely safe, review recommended)
   - `major`: 4.17.1 -> 5.0.0 (breaking changes likely, manual review required)
   - `no_fix`: no version of the direct dep resolves the transitive vuln
5. Report: "Upgrade `express` from `4.17.1` to `4.18.2` (minor, low risk) to fix CVE-2022-24999 in transitive dep `qs`"
6. If no direct dep upgrade fixes it: report "No remediation path available" and suggest ignoring with policy
7. **Never auto-apply major upgrades** -- CLI requires explicit `--allow-major` flag (unlike npm audit fix --force)

## Two-Way Jira Sync Flow

```mermaid
sequenceDiagram
  participant Scan
  participant TicketService
  participant Jira
  participant SyncJob

  Scan->>TicketService: New finding (severity >= threshold)
  TicketService->>TicketService: Check TicketMapping for fingerprint
  alt No existing ticket
    TicketService->>Jira: Create issue (summary, description, severity label)
    Jira-->>TicketService: Issue key (VULN-123)
    TicketService->>TicketService: Save TicketMapping
  else Existing ticket
    TicketService->>Jira: Add comment (new scan found same vuln)
  end

  Note over SyncJob: Periodic River job (every 15min)
  SyncJob->>Jira: Fetch status for all mapped tickets
  alt Ticket closed in Jira
    SyncJob->>TicketService: Mark finding as resolved
  end
  SyncJob->>SyncJob: Check findings resolved in latest scan
  alt Finding no longer present
    SyncJob->>Jira: Transition ticket to Done + comment
  end
```

## Dependencies and Build Order

```mermaid
flowchart LR
  A["Phase A: Foundation"] --> B["Phase B: Auth"]
  B --> C["Phase C: Projects"]
  C --> C2["Phase C2: Git Integration"]
  C2 --> D["Phase D: Scan Engine"]
  D --> E["Phase E: Node.js + SAST Engine + AI/LLM rules"]
  E --> EBis["Phase E-bis: Behavioral + Hallucination"]
  EBis --> E2["Phase E2: AI Impact Analyzer"]
  E2 --> E2Bis["Phase E2-bis: Call-Graph Reachability"]
  E2Bis --> F["Phase F: Remediation"]
  F --> F2["Phase F2: License + SBOM + Secrets"]
  F2 --> F3["Phase F3: Automated Fix PRs"]
  F3 --> G["Phase G: Dual Policies"]
  G --> Q["Phase Q: Rego + VEX Consumer + SLSA/Sigstore"]
  Q --> H["Phase H: Jira"]
  H --> I["Phase I: Monitor"]
  I --> E3["Phase E3: Python Ecosystem"]
  E3 --> J["Phase J: Notifications"]
  J --> K["Phase K: PR Checks + PR Baseline Diff"]
  K --> L["Phase L: Reports"]
  L --> E4["Phase E4: .NET + Go Ecosystems"]
  E4 --> M["Phase M: Scaffold DAST/IaC"]
  M --> M2["Phase M2: Container / OCI Scanning"]
  M2 --> N["Phase N: CLI"]
  N --> O["Phase O: Web Dashboard + AI Feedback"]
  O --> P["Phase P: IDE Plugin (LSP)"]
  P --> P2["Phase P2: ASPM Attack Stories"]
  P2 --> R["Phase R (future): Runtime Reachability eBPF Agent"]
```

## Implementation Phases

### Phase A -- Foundation (S)
- `go.mod` with all deps
- `internal/config/`, `internal/domain/`, `pkg/httputil/`
- `internal/middleware/` (auth, logger, recovery, cors)
- `cmd/server/main.go` bootstrap
- `Makefile`, `.env.example`, `Dockerfile`, `docker-compose.yml`
- Initial Goose migration (all tables)

### Phase B -- Auth (M)
- `internal/auth/` full stack
- Org + User GORM models, JWT, bcrypt
- Register, Login, Invite endpoints

### Phase C -- Projects (S)
- `internal/project/` full stack
- CRUD scoped to org from JWT

### Phase C2 -- Git Integration + Repo Security (L)
- `internal/git/model.go` -- `GitCredential`, `RepoCloneAuditLog`, `WebhookRegistration` models + DTOs
- `internal/git/repository.go` -- CRUD for git credentials (encrypted storage), audit log writes, webhook registrations
- `internal/git/cloner.go` -- Safe cloner: `git clone --depth 1 --single-branch --config core.hooksPath=/dev/null --config core.symlinks=false`, timeout kill, size check post-clone
- `internal/git/provider_github.go` -- GitHub REST/GraphQL: validate token, list repos, check repo access (collaborator check), fetch repo metadata (size, default branch)
- `internal/git/provider_gitlab.go` -- GitLab API v4: same as above for GitLab
- `internal/git/provider_bitbucket.go` -- Bitbucket API v2: same as above for Bitbucket
- `internal/git/extractor.go` -- Walk cloned directory, match allowed file patterns, skip binaries/node_modules/vendor, enforce per-file size limit, load into `map[string][]byte`
- `internal/git/security.go` -- Rate limiter (per-org scans/hour, concurrent clones), ownership verification (call provider API before clone), URL validation (allowlist of domains), audit log writer
- `internal/git/webhook.go` -- Webhook receivers: HMAC-SHA256 verification (GitHub), token verification (GitLab), event filtering, dedup (skip if scan already running for same project+branch)
- `internal/git/handler.go` -- API endpoints: credential CRUD, webhook registration, repo listing, audit log
- `internal/git/service.go` -- Orchestrator: validate -> rate limit -> verify ownership -> clone -> extract -> cleanup. Returns `CloneResult` to scan orchestrator.
- Config: `ENCRYPTION_KEY` (for AES-256-GCM credential encryption), `GIT_CLONE_TIMEOUT` (default: 300s), `GIT_MAX_REPO_SIZE` (default: 500MB), `GIT_MAX_SCANS_PER_HOUR` (default: 30), `GIT_MAX_CONCURRENT_CLONES` (default: 3)

### Phase D -- Scan Engine (M)
- `internal/scan/` full stack
- Scan + Finding models (include `effectiveSeverity`, `reachability`, `dependencyScope`, `fixRiskLevel`, `lockfileHash`)
- Scanner registry + orchestrator
- River `ScanJob` for async execution
- `internal/vuln/` for vuln cache layer (TTL-based cache to avoid redundant API calls, supports incremental scan mode)

### Phase E -- Node.js Ecosystem + SAST Engine (L)

**First ecosystem: Node.js (npm/yarn/pnpm)**

#### SCA (dependency analysis)
- `pkg/ecosystem/ecosystem.go` -- `EcosystemParser` interface + `DependencyTree` type + `DepScope` enum (multi-ecosystem design with scope awareness)
- `pkg/ecosystem/npm/parser.go` -- npm lockfile parser (v1/v2/v3), yarn.lock, pnpm-lock.yaml. Extracts `dev`/`optional`/`peer` flags per dependency
- `pkg/ecosystem/registry.go` -- parser registry, auto-detects ecosystem from files
- `internal/sca/graph.go` -- build + traverse dependency graph with **scope propagation** (dev deps and all their transitive children tagged as `development`), store DependencySnapshot + Dependency records with `lockfileHash`
- `internal/sca/context.go` -- **context analysis engine**: for each finding, determine `dependencyScope` by walking all graph paths, assign `reachability` classification
- `internal/sca/effective_severity.go` -- compute `effectiveSeverity` from raw severity + scope + reachability using downgrade rules
- `internal/sca/supply_chain.go` -- supply chain signal detector (npm-specific: postinstall hooks, new package detection <7 days, maintainer change tracking, typosquatting via Levenshtein distance to top-1000 npm packages). Designed with ecosystem-agnostic interface so Python/NuGet/Go supply chain checks plug in later.
- `internal/sca/incremental.go` -- incremental scan: compare `lockfileHash` with previous DependencySnapshot, if unchanged check only CVEs published since last scan timestamp
- `internal/sca/npm.go`, `osv.go`, `nvd.go` -- vuln source clients (support private registry auth via `RegistryConfig`)
- `internal/sca/ghsa.go` -- **GitHub Advisory Database** client (GraphQL API via `securityAdvisories` query). Covers npm, pip, go, maven, nuget, rubygems. Provides GHSA IDs alongside CVE IDs. Higher data quality than raw NVD for ecosystem-specific vulns. Uses `GITHUB_TOKEN` for auth (5000 req/hour). Supports filtering by ecosystem and severity.
- `internal/sca/vulncheck.go` -- **VulnCheck API** client (optional, commercial). Fastest CVE enrichment feed: includes exploit intelligence (is there a known exploit in the wild?), EPSS scores (probability of exploitation), KEV catalog matching (CISA Known Exploited Vulnerabilities). Adds `exploitAvailable` (bool), `epssScore` (float), `kevListed` (bool) fields to findings. Activated when `VULNCHECK_API_KEY` is configured. Falls back gracefully when not configured.
- `internal/sca/trivy.go` -- **Trivy DB** integration. Uses Trivy's vulnerability database (downloaded offline, updated via `trivy db update`). Provides an alternative/supplementary source to NVD with better coverage for OS packages and container images. Also detects Dockerfile/Kubernetes misconfigurations when IaC files are present. Used as a cross-validation source: findings confirmed by 2+ sources get higher confidence.
- `internal/sca/dedup.go` -- cross-source dedup by fingerprint. Enhanced: when the same CVE is found in multiple sources (npm + OSV + NVD + GHSA + VulnCheck + Trivy), merge metadata -- take the richest description, earliest published date, and aggregate exploit intelligence from all sources. Track `confirmedBySources` ([]string) on each finding for transparency.
- `internal/sca/scanner.go` -- tie it all together: parse -> graph -> supply chain -> lookup -> dedup -> context -> effective severity -> remediate

#### SAST Engine (Semgrep infrastructure -- used by all ecosystems)
- `internal/sast/scanner.go` -- Semgrep wrapper: exec `semgrep --config <rules-dir> --json <target>`, parse SARIF/JSON output. Supports both built-in Semgrep rules (`--config auto`) and custom rules from `rules/semgrep/`. Sandboxed execution with timeout (60s/file) and memory limit (1GB).
- `internal/sast/rules.go` -- Rule management: load custom `.yaml` rule files, validate rule syntax, merge with built-in rules. Supports per-org rule overrides and per-ecosystem rule selection.
- `internal/sast/mapper.go` -- Map Semgrep SARIF findings to our Finding model (severity, file, line, description, rule ID).

#### Node.js/Express/NestJS SAST Rules (`rules/semgrep/node-express-nest.yaml`, 20+ patterns)

**Injection & XSS:**
- `express-sqli-raw-query` -- Detect string concatenation in `sequelize.query()`, `knex.raw()`, `pg.query()` with user input
- `express-nosql-injection` -- Detect unsanitized user input in MongoDB `find()`, `findOne()`, `aggregate()` queries
- `express-xss-response` -- Detect `res.send(req.body.*)`, `res.send(req.query.*)` without sanitization -- reflected XSS
- `express-template-injection` -- Detect `res.render()` with user-controlled template name or unescaped variables
- `nestjs-raw-query` -- Detect raw SQL in NestJS `@InjectRepository` / TypeORM `query()` with string interpolation

**Prototype Pollution:**
- `lodash-merge-user-input` -- Detect `_.merge(target, req.body)`, `_.defaultsDeep(target, userInput)` -- prototype pollution
- `object-assign-user-input` -- Detect `Object.assign({}, req.body)` spread into sensitive objects
- `recursive-merge-pattern` -- Detect custom recursive merge functions processing untrusted input

**Auth & Session:**
- `jwt-none-algorithm` -- Detect `jwt.verify()` without `algorithms` option -- algorithm confusion attack
- `express-session-insecure` -- Detect `express-session` without `secure: true`, `httpOnly: true`, `sameSite` flags
- `hardcoded-jwt-secret` -- Detect `jwt.sign(payload, "hardcoded-string")` -- secret in source

**Command & Path Injection:**
- `child-process-injection` -- Detect `exec()`, `execSync()`, `spawn()` with user-controlled arguments
- `path-traversal` -- Detect `path.join(baseDir, req.params.file)` without sanitization -- directory traversal
- `eval-user-input` -- Detect `eval()`, `Function()`, `vm.runInContext()` with any external input

**SSRF & Network:**
- `ssrf-fetch-user-url` -- Detect `fetch(req.body.url)`, `axios.get(userInput)` -- SSRF
- `ssrf-redirect-follow` -- Detect HTTP clients following redirects on user-supplied URLs without policy

**NestJS-Specific:**
- `nestjs-no-validation-pipe` -- Detect controllers without `@UsePipes(ValidationPipe)` -- missing input validation
- `nestjs-expose-internal` -- Detect `@Res()` usage that bypasses NestJS interceptors and exception filters
- `nestjs-cors-wildcard` -- Detect `app.enableCors()` without origin restriction

**Secrets & Logging:**
- `console-log-secrets` -- Detect `console.log(password)`, `console.log(token)`, `console.log(req.headers.authorization)`
- `env-in-response` -- Detect `res.json(process.env)` or leaking env vars in API responses

`rules/semgrep/generic.yaml` -- Language-agnostic rules (also shipped in Phase E): hardcoded secrets regex (`ghp_`, `sk-`, `AKIA`, `-----BEGIN PRIVATE KEY-----`), TODO/FIXME security markers, commented-out auth checks

#### AI / LLM Security Rules (`rules/semgrep/ai-llm.yaml`, 15-20 patterns, Phase E)

New attack class that classical rule packs miss. Rules are ecosystem-agnostic and apply to Node.js, Python, and .NET codebases.

**Prompt Injection:**
- `llm-prompt-injection-user-input` -- Detect user-controlled input (`req.body`, `request.json`, HTTP params) flowing directly into `openai.chat.completions.create()`, `anthropic.messages.create()`, LangChain `LLMChain.run()` without sanitization
- `llm-prompt-concatenation` -- Detect f-strings / template literals that concatenate system prompts with user input in the same message role
- `llm-missing-input-guardrails` -- Detect LLM calls without any input validation / classifier / rate limiter in the middleware chain

**Unsafe LLM Output Handling:**
- `llm-output-to-eval` -- Detect LLM response piped into `eval()` / `exec()` / `Function()` / `vm.runInContext()`
- `llm-output-to-dangerouslysethtml` -- Detect LLM response used as `dangerouslySetInnerHTML` in React or `v-html` in Vue
- `llm-output-to-raw-sql` -- Detect LLM response concatenated into raw SQL queries
- `llm-output-to-shell` -- Detect LLM response passed to `child_process.exec`, `subprocess.call(shell=True)`, etc.
- `llm-output-unbounded-file-write` -- Detect LLM response written to arbitrary file paths

**Secret-Bearing System Prompts:**
- `llm-system-prompt-with-secret` -- Detect `system` role content containing API keys, DB connection strings, or anything matching secrets patterns
- `llm-tool-description-leaks-creds` -- Detect function-calling tool descriptions that expose internal credentials or infra details

**MCP / Agent Trust Boundary:**
- `mcp-tool-unbounded-shell` -- Detect MCP server tools exposing `shell(command: string)` / `exec(cmd: string)` to the LLM without allowlist
- `mcp-tool-arbitrary-file-write` -- Detect MCP server tools that let the LLM write arbitrary filesystem paths
- `langchain-agent-shell-tool` -- Detect LangChain agents with `ShellTool` / `PythonREPLTool` in production without sandboxing
- `llamaindex-query-unbounded` -- Detect LlamaIndex `QueryEngine` without output parser / response validator

**Denial of Service:**
- `llm-no-rate-limit` -- Detect HTTP endpoints that call LLM without any rate limit middleware -- cost-DoS risk
- `llm-no-token-budget` -- Detect LLM calls with user-controlled `max_tokens` or no cap on input tokens
- `tokenizer-dos-input` -- Detect user input passed to tokenizer (`tiktoken.encode()`, `transformers.AutoTokenizer`) without length cap

**Agentic AI Risks:**
- `agent-autonomous-git-push` -- Detect agent tools that push to git remotes without human-in-the-loop approval
- `agent-autonomous-payment` -- Detect agent tools that trigger payments/refunds without confirmation step

### Phase E-bis -- Behavioral Package Analysis + Hallucination Detection (L)

Socket-style behavioral defense + 2026 slopsquatting defense. Runs BEFORE vulnerability lookup, catches malicious packages that don't yet have CVEs.

- `internal/behavior/fetcher.go` -- Fetches published tarballs from registries (npm `registry.npmjs.org/{pkg}/-/{pkg}-{version}.tgz`, PyPI `files.pythonhosted.org`, NuGet `api.nuget.org/v3-flatcontainer/{pkg}/{version}/{pkg}.{version}.nupkg`, Go module proxy). Sandboxed unpack in `/tmp/vulncheck-behavior-{hash}` with no network access during unpack. Tarball size limit (`BEHAVIOR_MAX_TARBALL_SIZE`, default 100MB). Integrity hash verification against registry metadata.
- `internal/behavior/ast_analyzer.go` -- AST-based capability extractor. Uses tree-sitter (`github.com/smacker/go-tree-sitter`) for uniform parsing across JS/TS/Python/C#/Go. Per-language rules identify capability evidence:
  - `network`: `fetch`, `axios`, `http.get`, `net.Socket`, `requests.*`, `urllib.*`, `HttpClient`, `net/http`
  - `filesystem`: `fs.write*`, `fs.read*`, `open()`, `File.WriteAllText`, `os.Create`, path traversal patterns
  - `shell`: `child_process.exec`, `subprocess.*`, `os.system`, `os/exec.Command`, `Process.Start`
  - `env`: `process.env.*`, `os.environ.*`, `Environment.GetEnvironmentVariable`, `os.Getenv`
  - `unsafe_eval`: `eval`, `Function()`, `vm.runIn*`, `exec()`, `compile()`, `Expression.Compile`
  - `native_binding`: `require('bindings')`, ctypes, `[DllImport]`, `cgo` directives, `.node` file loads
  - `obfuscated`: min-file detection (>300 chars/line, very high entropy identifiers, base64-encoded string literals, hex-encoded strings of suspicious length)
  - `install_hook_shell`: `package.json` scripts with `sh -c` / `bash -c` / `node -e` / remote-download patterns; setup.py cmdclass overrides
- `internal/behavior/capability_diff.go` -- For each (pkg, new version) find the most recent analyzed version and diff capabilities. Emit `CapabilityExpansion` with severity: any new `unsafe_eval` / `shell` / `install_hook_shell` = critical; new `network` or `filesystem` on a lib that never had them = high; new `env` = medium; `obfuscated` regression = critical (legitimate libs don't typically flip to obfuscated).
- `internal/behavior/install_script.go` -- Semantic analysis of `preinstall` / `install` / `postinstall` hooks (and Python setup.py cmdclass). Beyond regex: base64 decoding of literal strings, remote-download commands (`curl`, `wget`, `Invoke-WebRequest`), process-spawn chains, env-var exfiltration patterns (`echo $SECRET | curl ...`). Severity mapping: remote download + shell exec = critical.
- `internal/behavior/hallucination.go` -- Hallucinated / slopsquat detector. Queries three predicates:
  1. Is the pkg in the hallucination candidate feed (seeded from public research + internal heuristic pipelines)?
  2. Was it published <30 days ago AND has <5 weekly downloads AND <2 maintainer-history entries?
  3. Does the name match a "plausible-but-made-up" pattern (common AI mistakes: `@org/internal-utils`, `my-company-private-*`, typo-combinations of popular libs)?
  If any two predicates match -> emit `HallucinationCandidate` finding at `critical` severity. Emits an auto-generated ignore-policy blocking the package (`action: block`, `expires: never`) unless explicitly overridden.
- `internal/behavior/feed.go` -- Periodic River `HallucinationFeedJob` refreshes the feed from configured sources (`HALLUCINATION_FEED_URLS`). Internal heuristic feed: scrape recent deletions on npm/PyPI (attackers register + take-down rapidly), package names appearing in popular AI-coding forums without matching real packages.
- `internal/behavior/service.go` -- Orchestrator wired into SCA pipeline. For each (pkg, version) in DependencySnapshot: lookup profile -> if missing, enqueue `BehaviorProfileJob` (async) -> return capability diff + hallucination check -> emit findings.
- `internal/jobs/behavior_profile_job.go` -- Async profile build. Deduplicated by (ecosystem, pkg, version). Cached aggressively.

### Phase E2 -- AI Impact Analyzer (L)
- `internal/ai/llm.go` -- LLM client abstraction: supports OpenAI (`gpt-4o`, `gpt-4o-mini`), Anthropic (`claude-sonnet-4-20250514`, `claude-haiku`), and self-hosted models via OpenAI-compatible API. Retry logic, rate limiting, timeout handling. API key from config/env.
- `internal/ai/code_indexer.go` -- **Multi-ecosystem source code indexer** with pluggable language strategies:
  - **Node.js (Phase E)**: scans `.js`/`.ts`/`.jsx`/`.tsx`/`.mjs`/`.cjs` for `import ... from 'package'`, `require('package')`, `import('package')`. Handles re-exports, barrel files, TypeScript path aliases.
  - **Python (Phase E3)**: scans `.py` for `import package`, `from package import ...`, `__import__('package')`, `importlib.import_module('package')`. Handles relative imports, `__init__.py` re-exports, conditional imports.
  - **.NET (Phase E4)**: scans `.cs`/`.vb`/`.fs` for `using Package.Namespace;`, `global using`, NuGet package references in code. Handles namespace aliases (`using Alias = Namespace`), static imports.
  - **Go (Phase E4)**: scans `.go` for `import "package"`, `import ( ... )` blocks. Handles aliased imports, blank imports (`_ "package"`), dot imports.
  - Common: builds a map `package -> [{file, line, importedNames, ecosystem}]`. Respects `.gitignore` and ecosystem-specific exclusion dirs (`node_modules/`, `__pycache__/`, `venv/`, `bin/`, `obj/`, `vendor/`).
- `internal/ai/snippet_extractor.go` -- Given a usage location, extracts 10-15 lines of context. Traces simple data flow: if the vulnerable function is called with a variable, follow that variable back to its source (function parameter, HTTP request, database query, hardcoded value). Annotates the snippet with data-flow tags. Language-aware: understands Python decorators, C# attributes, Go defer patterns.
- `internal/ai/prompt.go` -- Prompt templates, organized by vulnerability category (prototype_pollution, redos, xss, rce, sqli, path_traversal, ssrf, deserialization, generic) AND by ecosystem (Node.js, Python, .NET, Go) for ecosystem-specific context. Each template includes: CVE description, affected function signature, user's code snippet with data-flow annotation, ecosystem-specific context (e.g., "this is a Django view" or "this is an Express middleware"), and structured output format instructions (JSON). Prompts are designed to minimize hallucination: include explicit "if you cannot determine the impact, say so" instructions.
- `internal/ai/verdict.go` -- Parse LLM JSON response into typed `AIVerdict` struct. Validation: reject verdicts with confidence < 0.3 (fall back to scope-based analysis). Normalize impact levels.
- `internal/ai/analyzer.go` -- Orchestrator: for each finding from SCA, (1) check if package is imported via code_indexer, (2) if not imported -> immediate `no_impact` verdict without LLM call, (3) if imported -> extract snippets -> send to LLM -> parse verdict. Batching: group findings by package to avoid redundant indexing. Concurrency: analyze up to N findings in parallel (configurable). Caching: cache AI verdicts by fingerprint+source_hash to avoid re-analyzing unchanged code.
- River `AIAnalysisJob` for async analysis of large projects (>50 findings)
- Config: `LLM_PROVIDER` (openai/anthropic/custom), `LLM_API_KEY`, `LLM_MODEL`, `LLM_BASE_URL` (for self-hosted), `AI_MAX_CONCURRENT` (default: 5), `AI_ENABLED` (default: true)

Note: The code indexer is built with an extensible `LanguageStrategy` interface in Phase E (Node.js only). Python strategy is added in Phase E3, .NET/Go strategies in Phase E4. The LLM client, snippet extractor, verdict parser, and analyzer orchestrator are ecosystem-agnostic from day one.

**Multi-file Repository-level Context (added in Phase E2):**

- `internal/ai/repo_context.go` -- Builds a compact "repo envelope" injected into every LLM prompt:
  - **Framework detector**: scans `package.json` / `requirements.txt` / `go.mod` / `.csproj` for framework signals (express, nestjs, next, django, flask, fastapi, aspnet-core, gin). Returns ordered list with confidence.
  - **Entry-point + route extractor**: parses routing configs (Express `app.get/post/...`, NestJS `@Controller` + `@Get/@Post`, Django `urls.py`, Flask `@app.route`, FastAPI `@router`, ASP.NET `[Route]` / minimal API, Gin `r.GET/POST`). Returns `[{method, path, handler_file:line}]`.
  - **DI binding resolver**: for NestJS, ASP.NET Core, Spring-like Python containers -- extracts registered providers/services so the LLM understands which class implements which interface.
  - **Environment classifier**: reads `package.json` `engines`, `Dockerfile` base, env-var usage patterns to tag "this is a Next.js API route", "this is a background worker", "this is a CLI tool", "this is a library".
- `internal/ai/token_budget.go` -- Token-budget manager: truncates / summarizes repo context when prompt exceeds model limit. Prioritizes: (1) actual usage snippet, (2) immediate caller chain, (3) framework envelope, (4) route table. Uses cheap LLM (gpt-4o-mini / claude-haiku) to summarize large sections.

### Phase E2-bis -- Function-Level Call-Graph Reachability (L)

Deterministic reachability analysis that runs BEFORE AI to cheaply dismiss unreachable findings. Achieves 95%+ noise reduction comparable to Endor Labs / Socket+Coana.

- `internal/callgraph/engine.go` -- Orchestrator with `LanguageStrategy` interface matching the AI code indexer's. Per-scan lifecycle: select strategies for detected ecosystems -> build graphs in parallel -> merge into unified `CallGraph` -> discover entry points -> for each finding with a known vulnerable function, run tracer.
- `internal/callgraph/entrypoints.go` -- Entry-point discovery strategies:
  - **HTTP routes**: reuses framework detector from Phase E2 (Express / NestJS / Django / Flask / FastAPI / ASP.NET / Gin)
  - **CLI main**: `func main()`, `if __name__ == '__main__'`, `Program.Main`, `package.json` `bin`
  - **Message handlers**: `@OnMessage`, `@RabbitListener`, Kafka consumer registrations, NATS subscriptions
  - **Exported APIs**: `module.exports`, named exports in TS, `__all__` in Python, `public` symbols in C#/Go
  - **Test files**: `*.test.ts`, `*_test.go`, `test_*.py` (tagged as `test_entrypoint`, excluded from prod reachability by default)
- `internal/callgraph/node_js.go` -- Node.js/TypeScript strategy. Spawns a subprocess running `ts-morph` (TypeScript compiler API) via Node.js child process; communicates over JSON-RPC on stdio. Handles: `import`/`require`/`dynamic import`, TypeScript path aliases, barrel re-exports, class inheritance, prototype dispatch, async/await control flow. Memory-capped at `CALLGRAPH_MAX_MEMORY_MB`. Timeout at `CALLGRAPH_TIMEOUT`.
- `internal/callgraph/python.go` -- Python strategy. Subprocess using `jedi` / `scalpel`. Handles: decorators, class inheritance, `importlib`, `*` imports (best-effort), monkey-patching patterns. Partial-graph flag set when dynamic dispatch dominates.
- `internal/callgraph/go.go` -- Native Go implementation using `golang.org/x/tools/go/ssa` + `golang.org/x/tools/go/callgraph/cha` (Class Hierarchy Analysis). Faster and more accurate than subprocess-based strategies. Handles: interfaces, function values, closures, goroutines.
- `internal/callgraph/dotnet.go` -- .NET strategy. Subprocess using Roslyn (`Microsoft.CodeAnalysis`). Handles: generics, extension methods, virtual dispatch, async state machines.
- `internal/callgraph/tracer.go` -- Path tracer. For each finding with a known vulnerable function `{pkg, version, func}`:
  1. Find all `CallGraphNode` matching the target (may match multiple overloads)
  2. BFS backwards along edges to find any entry-point ancestor
  3. If found: emit `ReachabilityPath` with full path, set `confidence=direct` (complete graph) or `partial` (incomplete graph)
  4. If not found AND graph is complete: `reachability=unreachable` on the finding
  5. If not found AND graph is partial: `reachability=potentially_reachable` (do not dismiss)
- `internal/jobs/callgraph_job.go` -- Async graph build for large repos (>10k source files). Small repos build synchronously in the scan pipeline.

**Orchestration**: Call-graph runs AFTER scope-based context analysis and BEFORE AI Impact Analyzer. Decision matrix:
- Call-graph says `unreachable` (confidence=direct) -> set `reachability=unreachable`, skip AI entirely (big cost saving)
- Call-graph says `reachable` -> pass to AI Impact Analyzer for deeper data-flow verdict
- Call-graph says `partial` or extraction failed -> fall back to scope heuristic + AI

**Graceful degradation**: If call-graph extraction times out or the language subprocess crashes, the scan continues with scope-based reachability. A warning is logged + surfaced in scan summary. Never hard-fails.

### Phase F -- Remediation Engine (M)
- `internal/sca/remediation.go` -- walk dep graph upward to find upgrade paths. Ecosystem-agnostic interface with ecosystem-specific implementations:
  - **Node.js (Phase E)**: Query npm registry for available versions of direct deps. Compute minimum upgrade that resolves the transitive vuln.
  - **Python (Phase E3)**: Query PyPI for versions. Handle pinned versions in `requirements.txt` vs ranges in `pyproject.toml`.
  - **.NET (Phase E4)**: Query NuGet registry. Handle `Directory.Packages.props` centralized versions.
  - **Go (Phase E4)**: Query pkg.go.dev / module proxy. Handle `go get -u` semantics and minimum version selection (MVS).
- **Classify fix risk level**: compare current vs fix version using semver (patch/minor/major/no_fix)
- Populate `upgradePath` JSONB + `fixRiskLevel` on Finding
- Safe-fix logic: CLI `vulncheck fix` only applies `patch`/`minor` by default, `major` requires `--allow-major`

### Phase F2 -- License Compliance + SBOM + Secrets (L)

#### License Compliance
- `internal/license/service.go` -- License detection pipeline: for each dependency in DependencySnapshot, extract license from (1) lockfile metadata, (2) registry API call (npm `license` field, PyPI `license` field, NuGet `licenseUrl`), (3) fallback: parse `LICENSE` / `LICENSE.md` files from tarball if available
- `internal/license/spdx.go` -- SPDX expression parser: handles `MIT`, `GPL-2.0-only OR Apache-2.0`, `Apache-2.0 WITH LLVM-exception`, `(MIT AND BSD-3-Clause)`. Uses SPDX license list for validation.
- `internal/license/detector.go` -- Ecosystem-specific license extractors: npm (`package.json.license`), PyPI (JSON API `classifiers` + `license` field), NuGet (`nuspec` `license` element), Go (`go.mod` + module proxy metadata)
- `internal/license/handler.go` -- License policy CRUD + per-scan license findings + project license summary (for dashboard pie chart)
- Policy evaluation: for each detected license, check against `LicensePolicy` rules. If license matches a `deny` rule -> `LicenseFinding` with `policyAction=deny`. For dual-licensing (OR), if ANY option satisfies policy -> pass.

#### SBOM Generation
- `internal/sbom/spdx_writer.go` -- Generate SPDX 2.3 document: `SPDXRef-DOCUMENT`, one `SPDXRef-Package` per dependency, `DESCRIBES`/`DEPENDS_ON` relationships from dep graph, license info from LicenseFindings, supplier info from registry metadata
- `internal/sbom/cyclonedx_writer.go` -- Generate CycloneDX 1.5 BOM: `metadata.component` (root project), `components` array (all deps), `dependencies` (graph edges), `vulnerabilities` (optional VEX overlay with exploit status)
- `internal/sbom/vex.go` -- VEX (Vulnerability Exploitability eXchange) statements: for each Finding, generate VEX status (exploitable/not_affected/under_investigation) based on AI verdict + scope
- Endpoint: `POST /api/v1/projects/:id/sbom` triggers async generation via River `SBOMJob`

#### Secrets Scanning
- `internal/secrets/scanner.go` -- Implements `Scanner` interface (`scanType=secrets`). Runs on source code files extracted from git clone or file upload.
- `internal/secrets/patterns.go` -- 100+ regex patterns organized by provider/type: AWS (AKIA, secret key), GCP (service account JSON, API key), Azure (connection strings, client secrets), GitHub (ghp_, gho_, ghs_), GitLab (glpat-), Slack (xoxb-, xoxp-), Stripe (sk_live_, pk_live_), Twilio, SendGrid, database URLs (postgres://, mongodb://), JWT secrets, private keys (RSA/EC PEM headers), generic API keys
- `internal/secrets/entropy.go` -- Shannon entropy calculator: flag strings with >4.5 bits/char entropy as potential secrets (catches custom tokens without known patterns)
- `internal/secrets/validator.go` -- False positive reduction: validate format (e.g. AWS key must be exactly 20 chars AKIA prefix), check against known false positive patterns (e.g. placeholder "xxxx", test tokens, example values from docs)
- `internal/secrets/git_history.go` -- Optional git log scan: `git log --all --diff-filter=D -p` to find secrets in deleted files / previous commits. Slower, behind `--scan-history` flag.
- Registered in scan orchestrator alongside SCA and SAST. Secrets scan runs in parallel with other scanners.

### Phase F3 -- Automated Fix PRs (M)
- `internal/fixpr/service.go` -- Fix PR orchestrator:
  1. Take findings with `fixAvailable=true` and `fixRiskLevel` in [patch, minor]
  2. Group by direct dependency (one PR per direct dep upgrade, may fix multiple CVEs)
  3. Compute lockfile changes: update version in manifest (`package.json`, `requirements.txt`, `go.mod`) + regenerate lockfile
  4. Safety check: run SCA on the modified lockfile to verify no new vulns >= current severity are introduced
  5. Create branch (`vulncheck/fix-<package>-<version>`), commit changes, create PR via provider API
  6. Store `FixPR` record, set status to `created`
  7. Monitor: periodic job checks if PR was merged/closed and updates status
- `internal/fixpr/lockfile_updater.go` -- Ecosystem-specific lockfile update logic:
  - npm: update `package.json` version range, run `npm install --package-lock-only` to regenerate lockfile
  - Python: update `requirements.txt` pin or `pyproject.toml` version, regenerate with `pip-compile` or `poetry lock`
  - Go: update `go.mod` requires, run `go mod tidy`
  - .NET: update `Directory.Packages.props` or `.csproj` version, run `dotnet restore`
- `internal/fixpr/github_pr.go` / `gitlab_mr.go` / `bitbucket_pr.go` -- PR creation via provider API (reuses git credentials from Phase C2)
- River `FixPRJob` for async PR creation + `FixPRMonitorJob` for tracking merge status
- Config per org: `FIX_PR_MODE` (disabled/manual/auto), `FIX_PR_MAX_RISK_LEVEL` (patch/minor/major)

### Phase G -- Dual Policy System (M)
- `internal/policy/` full stack (DB policies)
- `internal/policy/file.go` -- parse/write `.vulncheck` YAML files
- Merge logic: evaluate DB policies + file policies, DB wins on conflict
- Import/export endpoints
- Expiration check: `WHERE ignoreUntil > now()`

### Phase Q -- Rego Policy-as-Code + Upstream VEX + SLSA/Sigstore (L)

Enterprise-grade controls layered on top of the dual policy system.

#### Policy-as-Code (Rego / OPA)

- `internal/policy/rego.go` -- Opt-in Rego engine using `github.com/open-policy-agent/opa/rego`. Loads `.rego` files from `REGO_POLICIES_DIR` (org uploads via API, files materialized to disk). Eval input = unified `PolicyInput` struct with finding + dependency + license + VEX + AI verdict + exploit intel + attack story.
- `internal/policy/rego_schema.go` -- Defines `PolicyInput` Go struct + matching Rego input documentation. Ensures forward compatibility when new fields are added (versioned schema).
- Three policy outcomes are supported, evaluated on every scan:
  - `block`: finding contributes to non-zero CI exit code
  - `warn`: finding is still visible but marked as "warned" (not blocking)
  - `notify`: finding triggers a notification (email/Slack) even if below normal threshold
- Evaluation order: DB ignore policies -> `.vulncheck` YAML file policies -> Rego policies. DB + YAML produce ignore actions; Rego produces block/warn/notify actions.
- Example policy shipped in `policies/rego-examples/no_high_epss_patch.rego`:
  ```rego
  package vulncheck.block
  default block_reason = ""
  block_reason := reason if {
      input.finding.epss_score > 0.7
      input.finding.fix_risk_level == "patch"
      input.finding.dependency_scope == "production"
      not input.finding.has_vex_not_affected
      reason := sprintf("High-EPSS patch-fix finding %s must be fixed (EPSS=%v)", [input.finding.cve_id, input.finding.epss_score])
  }
  ```
- Rego evaluation has a hard timeout (`REGO_EVAL_TIMEOUT`, default 30s) and is sandboxed (no `http.send`, no filesystem access).
- API for Rego policy management: `POST /api/v1/rego-policies`, etc. Dry-run endpoint (`POST /api/v1/rego-policies/:id/evaluate`) lets users test a policy against a historical scan before enabling it.

#### Upstream VEX Consumer

- `internal/vex/consumer.go` -- Orchestrator for upstream VEX ingestion. Runs on periodic River `VEXFeedJob` (default 12h interval).
- `internal/vex/openvex.go` -- OpenVEX JSON parser. Validates against official OpenVEX schema (`https://openvex.dev/ns/v0.2.0`). Extracts statements with `vuln.name` (CVE/GHSA), `products[].@id` (package PURLs), `status` (not_affected / affected / fixed / under_investigation), `justification`, `impact_statement`.
- `internal/vex/csaf.go` -- OASIS CSAF 2.0 JSON parser. Maps CSAF `vulnerabilities[].product_status` to our `VEXStatement.status`.
- `internal/vex/feed.go` -- Configurable feed list. Platform-wide default feeds shipped: Kubernetes security, Red Hat security, Node.js security WG, Python security WG, PyPA advisories, Rust security advisory DB. Orgs can add custom feeds (vendor pages, internal feeds).
- `internal/vex/applier.go` -- Apply VEX statements to findings during the scan pipeline:
  - `not_affected` + matches finding (CVE + package + version in scope) -> auto-dismiss with `ignore_source=upstream_vex`, provenance linked to `VEXStatement`
  - `fixed` + version matches -> mark as resolved
  - `affected` -> upgrade confidence (counter-signal if previously dismissed)
  - `under_investigation` -> warn in UI, do not auto-dismiss
- Signature validation: if `VEXFeed.signatureRequired=true`, statements must be signed by a trusted signer (sigstore identity, x509 subject, or PGP key). Unsigned statements are stored but not applied.
- Every applied dismissal is transparently shown in the UI: "dismissed per VEX statement from kubernetes.io/security, published 2026-01-15, signed by sigstore identity `spiffe://kubernetes.io/vex-signer`".

#### SLSA Provenance + Sigstore Verification

- `internal/slsa/verifier.go` -- For each dep in a DependencySnapshot:
  - **npm**: Query registry for `provenance` field on the version; if present, download the attestation bundle (in-toto statement + signatures). Verify via `sigstore-go` against Sigstore Rekor transparency log.
  - **PyPI**: Query `/pypi/{pkg}/{version}/json` for `attestations` (PEP 740). Verify similarly.
  - **GitHub artifacts**: For Go modules hosted on GitHub, query GitHub Artifact Attestations API.
  - **Container images**: Verified by Phase M2 scanner via cosign (see below).
- `internal/slsa/sigstore.go` -- Sigstore Rekor client. Looks up transparency log entries, verifies Merkle inclusion proofs, validates signer identity against allowlist.
- Supply-chain signals emitted:
  - `missing_provenance` (info, by default): package has no SLSA attestation at all
  - `provenance_mismatch` (high): attestation's declared source commit does not match the published artifact's claimed git tag
  - `unsigned_artifact` (medium for public packages, info for private)
  - `provenance_verified` (positive signal, shown as green badge in UI; not a finding)
- Config: `SLSA_ENABLED`, `SIGSTORE_REKOR_URL`, trusted-signers allowlist per org.

### Phase H -- Jira Integration (L)
- `internal/integration/` full stack
- `internal/integration/jira.go` -- Jira Cloud REST v3 client (create issue, transition, add comment, fetch status)
- `IntegrationConfig` + `TicketMapping` models
- Auto-create on scan completion (above threshold, dedup check)
- River `TicketSyncJob` (periodic, two-way sync)
- GitHub Issues + Linear clients scaffolded

### Phase I -- Monitor Mode (M)
- `internal/monitor/` full stack
- Snapshot storage (reuses DependencySnapshot from Phase E)
- River `MonitorJob` -- periodic re-check, configurable interval
- Diff logic: compare new vuln results against previous check, surface only NEW findings
- Alert via existing ticket integration

### Phase E3 -- Python Ecosystem (L)

**Second ecosystem: Python (pip/pipenv/poetry)**

#### SCA (dependency analysis)
- `pkg/ecosystem/python/parser.go` -- Lockfile parsers:
  - `requirements.txt`: parse `package==version` lines, handle `-r` includes, `--hash` mode, extras (`package[extra]`). No native scope flags -- all deps are production unless file is named `requirements-dev.txt` or `dev-requirements.txt` (convention-based scope detection).
  - `Pipfile.lock`: JSON parser, native `[packages]` (production) vs `[dev-packages]` (development) scope distinction. Includes hashes for integrity.
  - `poetry.lock`: TOML parser, scope from `[tool.poetry.dependencies]` vs `[tool.poetry.group.dev.dependencies]` in `pyproject.toml`. Full transitive graph available.
  - `pyproject.toml` (PEP 621): parse `[project.dependencies]` and `[project.optional-dependencies]` for scope.
- `internal/sca/pypi.go` -- PyPI advisory client: query PyPI JSON API (`/pypi/{package}/json`) for vulnerability info + use OSV.dev with `ecosystem=PyPI` filter (richer than PyPI's own data). Map PyPI package names to CVE/GHSA IDs.
- `internal/sca/supply_chain_python.go` -- Python-specific supply chain detection:
  - `setup.py` / `setup.cfg` hooks: detect `cmdclass` overrides that execute shell commands during `pip install`
  - PyPI typosquatting: Levenshtein distance against top-1000 PyPI packages (by downloads)
  - New package detection: packages published <7 days on PyPI
  - Maintainer changes: compare against previous snapshot
  - `__init__.py` code execution: detect packages with executable code in `__init__.py` that runs on import (network calls, file writes, env var reads)
- AI code indexer: add Python `LanguageStrategy` to `internal/ai/code_indexer.go` (see Phase E2)

#### Python/Django/Flask/FastAPI SAST Rules (`rules/semgrep/python-django-flask.yaml`, 20+ patterns)

**SQL Injection:**
- `django-raw-sql` -- Detect `cursor.execute(f"SELECT ... {user_input}")`, `RawSQL()` with string formatting
- `sqlalchemy-text-injection` -- Detect `db.execute(text(f"...{input}"))` -- must use `text("...").bindparams()`
- `flask-sqli` -- Detect `db.engine.execute("..." + request.args.get(...))` in Flask-SQLAlchemy

**Command Injection:**
- `python-subprocess-shell` -- Detect `subprocess.call(user_input, shell=True)`, `os.system(user_input)`, `os.popen(user_input)`
- `python-eval-exec` -- Detect `eval(request.data)`, `exec(user_input)`, `compile()` with untrusted code

**Auth & Session:**
- `django-csrf-exempt` -- Detect `@csrf_exempt` decorator on state-changing views without justification
- `flask-secret-key-hardcoded` -- Detect `app.secret_key = "hardcoded-string"` or `SECRET_KEY` in source
- `fastapi-no-auth-dependency` -- Detect endpoints without `Depends(get_current_user)` or similar auth dependency
- `django-debug-true` -- Detect `DEBUG = True` in production settings (settings.py / .env)

**Deserialization:**
- `python-pickle-untrusted` -- Detect `pickle.loads(request.data)`, `pickle.load(network_source)` -- arbitrary code execution
- `python-yaml-unsafe` -- Detect `yaml.load(data)` without `Loader=yaml.SafeLoader` -- code execution via YAML

**Path Traversal & File Handling:**
- `python-path-traversal` -- Detect `open(os.path.join(base, user_input))` without sanitization
- `django-file-upload-unrestricted` -- Detect `FileField` / `ImageField` without `upload_to` + content type validation
- `flask-send-file-traversal` -- Detect `send_file(request.args.get('path'))` -- directory traversal

**SSRF:**
- `python-ssrf-requests` -- Detect `requests.get(user_input)`, `urllib.request.urlopen(user_input)` -- SSRF
- `python-ssrf-httpx` -- Same for `httpx.get()`, `aiohttp.ClientSession().get()` with user URLs

**Django-Specific:**
- `django-xss-mark-safe` -- Detect `mark_safe(user_input)` or `|safe` filter on user data in templates
- `django-mass-assignment` -- Detect `Model.objects.create(**request.POST)` without explicit field list
- `django-open-redirect` -- Detect `HttpResponseRedirect(request.GET.get('next'))` without URL validation

**FastAPI-Specific:**
- `fastapi-response-model-missing` -- Detect endpoints returning raw dicts/objects without `response_model` -- may leak internal fields
- `fastapi-cors-wildcard` -- Detect `CORSMiddleware(allow_origins=["*"])` in production

**Secrets & Logging:**
- `python-secrets-in-logs` -- Detect `logging.info(f"token={token}")`, `print(password)` -- secrets in logs
- `python-hardcoded-credentials` -- Detect `password = "..."`, `api_key = "..."`, `token = "sk-..."` in source

### Phase J -- Notifications (M)
- `internal/notify/` full stack
- Email sender (SMTP via `net/smtp` or a library like `gomail`)
- Slack sender (incoming webhook + optional Slack API for richer messages)
- `NotificationConfig` model, CRUD endpoints
- River `NotificationJob` for async dispatch
- Triggered by: scan completion, critical finding, monitor new-vuln alert

### Phase K -- GitHub PR Checks + PR Baseline Diff (M)
- `internal/prcheck/` service + GitHub client
- GitHub Checks API: create check run with pass/fail based on effective severity threshold
- PR comment: post a findings summary table including **AI verdict summary** per finding (one-line: "No impact -- not imported" or "Critical -- called with user input at handler.js:42")
- Webhook receiver: listen for `pull_request` events, trigger scan on new/updated PRs
- `PRCheckConfig` model, CRUD endpoints
- **PR Baseline Diff (new in this phase)**:
  - `internal/prcheck/baseline.go` -- On PR webhook, identify the last successful scan on `defaultBranch` (the baseline). Run full scan on PR head. Compute three sets: NEW findings (in PR head, not in baseline), RESOLVED findings (in baseline, not in PR head), UNCHANGED findings (in both).
  - PR comment structure has three sections:
    - **Introduced by this PR** (highlighted red): lists NEW findings with full AI verdict and fix-PR trigger button
    - **Fixed by this PR** (highlighted green): lists RESOLVED findings
    - **Lockfile diff**: table of packages added/removed/bumped between baseline and PR head (extracted from `DependencySnapshot` delta)
  - Check-run pass/fail is based ONLY on NEW findings >= threshold. Legacy findings do not block unchanged PRs.
  - Config toggle `PR_DIFF_ENABLED` lets orgs opt into diff mode; default behavior without the flag is the classical "report all findings on PR" mode for backward compat.
  - When the baseline is older than `PR_DIFF_BASELINE_MAX_AGE` (default 14 days), fall back to classical full-scan mode and emit a warning in the PR comment ("baseline too stale, showing full scan").

### Phase L -- Compliance Reports (M)
- `internal/report/` full stack
- PDF generation (using a Go PDF library like `gofpdf` or `maroto`)
- CSV export (standard `encoding/csv`)
- Report content: scan history, open findings sorted by severity, policy justifications, remediation status, dependency graph summary
- River `ReportJob` for async generation (can take time for large projects)
- Download endpoint serves the generated file

### Phase E4 -- .NET + Go Ecosystems (L)

**Third wave: .NET (NuGet) + Go (go.mod)**

#### .NET SCA (NuGet)
- `pkg/ecosystem/dotnet/parser.go` -- Lockfile parsers:
  - `packages.lock.json` (NuGet lock file, .NET 6+): JSON parser, extracts direct vs transitive deps, target framework, content hash. Scope from `type` field: `Direct` vs `Transitive`.
  - `.csproj` / `.vbproj` / `.fsproj`: XML parser, extract `<PackageReference>` elements with version, `<PrivateAssets>all</PrivateAssets>` (dev-only). Supports `Directory.Packages.props` for centralized package management.
  - `packages.config` (legacy .NET Framework): XML parser, basic flat dependency list.
- `internal/sca/nuget.go` -- NuGet advisory client: query NuGet V3 API (`/registration/{package}/index.json`) for vulnerability info. Also use GHSA with `ecosystem=NuGet` filter and OSV.dev with `ecosystem=NuGet`.
- `internal/sca/supply_chain_dotnet.go` -- .NET supply chain detection:
  - MSBuild targets/props injection: detect packages with `.targets` / `.props` files that execute build-time commands
  - NuGet typosquatting: Levenshtein against top-500 NuGet packages
  - Package author/owner changes since last snapshot
- AI code indexer: add .NET `LanguageStrategy` (C#/VB/F# `using` statements, see Phase E2)

#### Go SCA (go.mod)
- `pkg/ecosystem/gomod/parser.go` -- Parse `go.mod` (require/replace/exclude directives) + `go.sum` (integrity hashes). Direct vs indirect deps from `// indirect` comment. No native dev/prod scope -- Go compiles everything, but test-only imports (files ending `_test.go`) can be treated as dev scope.
- `internal/sca/govuln.go` -- Go Vulnerability Database client: query `vuln.go.dev` API (same DB as `govulncheck`). Also use GHSA with `ecosystem=Go` and OSV.dev with `ecosystem=Go`.
- `internal/sca/supply_chain_go.go` -- Go supply chain detection:
  - Suspicious `replace` directives pointing to non-standard module paths
  - New modules published <7 days on pkg.go.dev
  - Module path typosquatting against top Go modules
- AI code indexer: add Go `LanguageStrategy` (Go import blocks, see Phase E2)

#### ASP.NET Core SAST Rules (`rules/semgrep/dotnet-aspnet-core.yaml`, 20+ patterns)

**SQL Injection:**
- `efcore-raw-sql` -- Detect `context.Database.ExecuteSqlRaw($"SELECT ... {input}")` -- must use `ExecuteSqlInterpolated` or parameterized
- `efcore-fromraw` -- Detect `context.Set<T>().FromSqlRaw($"... {input}")` without parameters
- `dapper-string-concat` -- Detect `connection.Query("..." + userInput)` in Dapper -- must use parameterized queries
- `ado-net-sqli` -- Detect `new SqlCommand($"... {input}", connection)` -- use `SqlParameter`

**Auth & Identity:**
- `aspnet-authorize-missing` -- Detect controllers/actions without `[Authorize]` attribute (when not explicitly `[AllowAnonymous]`)
- `aspnet-jwt-no-validation` -- Detect `TokenValidationParameters` with `ValidateIssuer = false` or `ValidateAudience = false`
- `aspnet-antiforgery-missing` -- Detect POST/PUT/DELETE actions without `[ValidateAntiForgeryToken]` in MVC (not needed in API-only apps)
- `identity-weak-password` -- Detect `PasswordOptions` with `RequiredLength < 8` or `RequireDigit = false`

**XSS & Output Encoding:**
- `aspnet-html-raw` -- Detect `@Html.Raw(userInput)` in Razor views -- XSS
- `aspnet-response-write` -- Detect `Response.WriteAsync(userInput)` without encoding
- `blazor-markup-string` -- Detect `MarkupString(userInput)` in Blazor components

**Deserialization:**
- `dotnet-binaryformatter` -- Detect `BinaryFormatter.Deserialize(stream)` -- critical RCE, banned in .NET 9+
- `dotnet-json-typehandling` -- Detect `JsonSerializerSettings { TypeNameHandling = TypeNameHandling.All }` -- deserialization RCE
- `dotnet-xml-external-entity` -- Detect `XmlReader.Create()` / `XmlDocument.Load()` without `DtdProcessing = DtdProcessing.Prohibit` -- XXE

**SSRF & Network:**
- `aspnet-ssrf-httpclient` -- Detect `HttpClient.GetAsync(userInput)` -- SSRF
- `aspnet-open-redirect` -- Detect `Redirect(returnUrl)` without `Url.IsLocalUrl(returnUrl)` check

**Configuration & Secrets:**
- `aspnet-dev-exception-page` -- Detect `app.UseDeveloperExceptionPage()` without environment check -- stack trace leak in prod
- `aspnet-cors-wildcard` -- Detect `.AllowAnyOrigin()` in CORS policy for production
- `aspnet-hardcoded-connection` -- Detect connection strings with passwords in `appsettings.json` or source code
- `aspnet-https-not-enforced` -- Detect missing `app.UseHttpsRedirection()` or `[RequireHttps]`

**Logging & Data Protection:**
- `aspnet-log-pii` -- Detect `_logger.LogInformation($"User email: {email}")` -- PII in logs
- `aspnet-data-protection-weak` -- Detect `DataProtectionProvider` with weak algorithms or no key rotation

#### Go/Gin/GORM SAST Rules (`rules/semgrep/go-gin-gorm.yaml`, 25+ patterns)

**Injection & Input Validation:**
- `gorm-raw-string-concat` -- Detect `db.Raw(fmt.Sprintf(...))`, `db.Where(fmt.Sprintf(...))`, `db.Exec(fmt.Sprintf(...))` -- SQL injection via GORM
- `gorm-raw-string-plus` -- Detect `db.Raw("..." + userInput)` string concatenation in GORM queries
- `gin-open-redirect` -- Detect `c.Redirect(...)` with user-controlled URL
- `gin-header-injection` -- Detect `c.Header(key, userInput)` where value comes from request params
- `gin-template-injection` -- Detect `text/template` (not `html/template`) used with user input

**Cryptographic Misuse:**
- `crypto-math-rand` -- Detect `math/rand` in security-sensitive contexts -- must use `crypto/rand`
- `crypto-weak-hash` -- Detect `md5.New()`, `sha1.New()` for password hashing -- must use SHA-256+ or bcrypt/argon2
- `crypto-weak-tls` -- Detect `tls.Config{MinVersion: tls.VersionTLS10}` or TLS 1.1 -- enforce TLS 1.2+
- `crypto-hardcoded-key` -- Detect hardcoded encryption keys, JWT secrets, or API keys

**Unsafe Deserialization:**
- `unsafe-gob-decode` -- Detect `gob.NewDecoder(untrustedSource).Decode(...)` -- arbitrary type instantiation
- `unsafe-xml-decode` -- Detect `xml.NewDecoder(untrustedSource)` without entity expansion limits -- XXE
- `unsafe-json-number` -- Detect `json.Number` without bounds checking -- integer overflow

**Race Conditions & Resource Safety:**
- `toctou-file-race` -- Detect `os.Stat(path)` followed by `os.Open(path)` -- TOCTOU race
- `unsafe-int-conversion` -- Detect `int32(int64Var)` without bounds checking -- integer overflow/truncation
- `goroutine-no-cancel` -- Detect `go func(){}()` without context cancellation -- goroutine leak
- `unclosed-body` -- Detect `http.Get()` without `defer resp.Body.Close()` -- resource leak

**SSRF & Network Safety:**
- `ssrf-user-url` -- Detect `http.Get(userInput)` where URL comes from request -- SSRF
- `ssrf-dns-rebinding` -- Detect URL validation followed by HTTP request -- DNS rebinding
- `gin-ssrf-redirect` -- Detect following redirects on user-supplied URLs

**Secrets & Logging:**
- `secrets-in-logs` -- Detect `logger.Info/Error/Warn(...)` with sensitive field names -- PII/secrets leak
- `secrets-in-error` -- Detect `fmt.Errorf("... %s", password)` -- secrets in error chains
- `hardcoded-credentials` -- Detect string literals matching secret patterns: `ghp_`, `sk-`, `AKIA`, Bearer tokens

**Gin-Specific Security:**
- `gin-debug-mode-prod` -- Detect missing `gin.SetMode(gin.ReleaseMode)` -- stack trace leak
- `gin-no-trusted-proxies` -- Detect missing `SetTrustedProxies()` -- IP spoofing
- `gin-cors-wildcard` -- Detect CORS `AllowOrigins: ["*"]` in production

### Phase M -- Scaffold DAST/IaC (S)

- `internal/dast/scanner.go` -- ZAP interface, returns "not implemented" error
- `internal/iac/scanner.go` -- Checkov/Trivy misconfig interface, returns "not implemented" error
- Register in scan orchestrator alongside SCA and SAST scanners

Note: SAST engine (`internal/sast/scanner.go`, `rules.go`, `mapper.go`) is built in Phase E. Ecosystem-specific SAST rules ship with their ecosystem phases (Node.js in E, Python in E3, .NET/Go in E4). Phase M only scaffolds DAST and IaC.

### Phase M2 -- Container / OCI Image Scanning (L)

Full container image scanning to close the prod-deployment gap vs Snyk Container / Aqua / Trivy / Wiz.

- `internal/container/scanner.go` -- Implements `Scanner` interface (`scanType=container`). Accepts image reference (`registry/repo:tag` or `@digest`). Orchestrates: pull -> layer extraction -> OS package scan -> Dockerfile SAST -> base-image analysis -> signature/provenance verification -> SBOM output.
- `internal/container/image.go` -- Uses `github.com/google/go-containerregistry` (`crane` library) to pull images without requiring a Docker daemon. Supports: Docker Hub, GHCR, ECR, GAR, Harbor, self-hosted registries. Credentials resolved from `CONTAINER_REGISTRY_AUTH` config (per-registry). Layers materialized to a sandboxed temp dir for inspection.
- `internal/container/os_packages.go` -- Per-layer OS package extraction:
  - **Debian/Ubuntu**: parse `/var/lib/dpkg/status`
  - **Alpine**: parse `/lib/apk/db/installed`
  - **RHEL/CentOS/Fedora**: parse RPM DB (use `rpm-go` library)
  - **Distroless / scratch**: no OS packages (flag all findings as "base image is minimal -- any lang-level deps scanned separately")
  Check each package against Trivy DB (reuse `internal/sca/trivy.go` with `scanType=container` tagging). Emit `ContainerFinding` records with `layerIndex` showing which layer introduced the vuln.
- `internal/container/dockerfile.go` -- Dockerfile SAST. Parses Dockerfile from the image's config history OR a supplied `Dockerfile` file. Rules from `rules/semgrep/dockerfile.yaml` (shipped in this phase, 10+ patterns):
  - `dockerfile-root-user` -- no `USER` directive, or `USER root` (MEDIUM)
  - `dockerfile-add-http` -- `ADD http://...` (allows remote code + SSL-less) (HIGH)
  - `dockerfile-latest-tag` -- `FROM image:latest` (reproducibility + supply chain) (MEDIUM)
  - `dockerfile-missing-healthcheck` -- no `HEALTHCHECK` (INFO)
  - `dockerfile-copy-parent-dir` -- `COPY .. /app` (possible context leak) (MEDIUM)
  - `dockerfile-sensitive-files` -- `.ssh/`, `.aws/`, `.env*` copied into image (HIGH)
  - `dockerfile-apt-no-cleanup` -- `apt-get install` without `apt-get clean` / `rm -rf /var/lib/apt/lists/*` (LOW, image size)
  - `dockerfile-curl-pipe-sh` -- `curl ... | sh` (supply-chain trust) (HIGH)
  - `dockerfile-hardcoded-secret` -- `ENV` / `ARG` matching secrets regex (CRITICAL)
  - `dockerfile-exposed-dangerous-port` -- `EXPOSE 22` (SSH in container) (MEDIUM)
- `internal/container/base_image.go` -- Base-image detection and minimal-alternative recommendation. Strategy:
  - Detect base by: Dockerfile `FROM` (if available), layer-0 fingerprint matching against known base-image database (Debian, Ubuntu, Alpine, distroless/*, chainguard/*).
  - For each detected base, query a curated recommendation table: "debian:bullseye-slim -> gcr.io/distroless/base-debian11 (-80MB, -62 CVEs)", "node:18 -> chainguard/node:latest (-120MB, -240 CVEs)".
  - Output: `recommendedBaseImage` + estimated size delta + estimated CVE delta.
- `internal/container/cosign.go` -- Sigstore signature verification. Uses `github.com/sigstore/cosign` as a library (not shelling out to `cosign` binary). Checks keyless OIDC signatures via Rekor transparency log + static public-key signatures when configured (`COSIGN_PUBLIC_KEYS_DIR`). Emits `unsigned_image` (medium for public registries) / `signature_invalid` (critical) / `signature_verified` (positive badge).
- `internal/container/slsa_attestation.go` -- SLSA provenance attestation verification for images. Similar approach to Phase Q SLSA verification but for container references. Verifies attestation bundle is signed by a trusted build identity (GitHub Actions SLSA generator, self-hosted SLSA-compliant builders).
- `internal/container/sbom_container.go` -- Generate container-specific CycloneDX SBOM with layer breakdown. Each component includes `properties` with `container:layer_index`, `container:layer_digest`, `container:installed_by_command`.
- API: `POST /api/v1/projects/:id/container-scans`, `GET /api/v1/container-scans/:id/...`. Async via `ContainerScanJob`.
- CLI: `vulncheck container scan <image:tag>`, `vulncheck container sbom <image:tag>`.

### Phase P -- IDE Plugin (LSP) (L)

Developer adoption driver. Ships alongside CLI.

- `cmd/cli/main.go` -- Adds `vulncheck lsp` subcommand that starts a Language Server over stdio transport.
- `internal/lsp/server.go` -- LSP server implementation using `github.com/tliron/glsp` (Go LSP framework). Implements:
  - `initialize` / `shutdown` / `exit` lifecycle
  - `textDocument/didOpen`, `didChange`, `didSave` -- triggers save-time debounced lightweight scan on the specific file (lockfile or source)
  - `textDocument/hover` -- on hover over a version string in `package.json` / `requirements.txt` / `go.mod` / `.csproj`, return tooltip with: CVE count, top AI verdict one-liner, capability profile badge, fix-risk
  - `textDocument/codeAction` -- offer quick fixes: "Trigger Fix PR for this dep", "Upgrade to vX.Y.Z (patch)", "Add `.vulncheck` ignore with reason template", "Explain with AI"
  - `textDocument/publishDiagnostics` -- inline diagnostics for findings (yellow squiggles under vulnerable versions)
  - `workspace/executeCommand` -- command hooks for CLI operations (`vulncheck.fix`, `vulncheck.explain`, `vulncheck.auth`)
- `internal/lsp/quick_scan.go` -- Uses `POST /api/v1/ide/quick-scan` endpoint for heavy work (full AI analysis runs server-side), but local lockfile parsing + scope classification runs in-process for instant feedback.
- `ide/vscode/` -- VSCode extension (TypeScript). `package.json` + `extension.ts` minimal wrapper that spawns `vulncheck lsp` and connects via `vscode-languageclient`. Bundled as `.vsix` via `vsce`. Published to VSCode Marketplace + Open VSX.
- `ide/jetbrains/` -- JetBrains plugin (Kotlin, Gradle-built). Uses IntelliJ LSP4IJ plugin to talk to our LSP server. Published to JetBrains Marketplace.
- Config: `LSP_PORT` (default stdio), authentication reuses the user's `~/.vulncheck/credentials` CLI token.

### Phase P2 -- ASPM Attack-Story Correlation (M)

- `internal/aspm/correlator.go` -- Post-scan correlator. Runs after all scanners complete. Strategy:
  1. Collect all findings across SCA / SAST / Secrets / IaC / Container for the scan.
  2. Extract "entities" per finding: file paths mentioned, package names, service names (from framework detector), route patterns, K8s resources, DB connection strings, S3 bucket names.
  3. Build a bipartite graph: findings <-> entities. Edges are typed (`mentions`, `uses`, `exposes`, `depends_on`).
  4. Find connected components where at least 2 findings share entities AND the combination forms a plausible attack chain (pre-defined chain templates + heuristics).
- `internal/aspm/attack_graph.go` -- Chain templates:
  - `secret + sast_injection + sca_in_path`: hardcoded credential + injection sink + vulnerable lib on the path
  - `iac_public + sast_missing_auth + sca_kev`: public resource + missing auth + KEV-listed CVE
  - `dockerfile_root + container_cve_rce`: container runs as root + RCE CVE in base image
  - `dep_confusion + behavioral_expansion`: internal package name found on public registry + behavioral capability expansion
- `internal/aspm/story.go` -- For each chain, generate a human-readable narrative using an LLM (optional; template-based fallback without LLM). Compute combined severity (max constituent + chain-length bonus). Store as `AttackStory`.
- `internal/jobs/aspm_correlation_job.go` -- Async. Runs after scan completion. Typical runtime <10s.
- Dashboard surfaces a top-level "Attack Stories" tab with prioritized cards. Clicking a story shows the chain graph visualization (using `reactflow` in the React frontend) with each finding as a node, edges labeled with relationship type.
- `cmd/cli/main.go` -- Cobra
- `vulncheck test` (scan + exit code based on **effective severity**, default: production-scope only), `vulncheck test --include-dev`, `vulncheck test --all`, `vulncheck test --fail-on <level>`
- `vulncheck fix` (safe auto-fix: patch/minor only), `vulncheck fix --dry-run`, `vulncheck fix --allow-major`
- `vulncheck supply-chain` (standalone supply chain analysis)
- `vulncheck monitor`, `vulncheck ignore`, `vulncheck auth login`, `vulncheck policy list`, `vulncheck report generate`
- `.vulncheck.yaml` config file for API URL + token
- Colorized terminal output: **effective severity** shown alongside raw severity, scope badge (PROD/DEV), reachability indicator
- `--json` flag for machine-readable output in CI pipelines
- Exit code semantics: 0 = no actionable findings, 1 = findings above threshold in production scope, 2 = scan error

### Phase O -- Web Dashboard + AI Feedback Loop (L)
- `web/` -- React + TypeScript (Vite or Next.js)
- Pages: Login, Dashboard (org overview with **noise metrics** -- action rate, findings trend, **AI verdict accuracy** from feedback loop), Projects list, Project detail (scans timeline + **scope breakdown chart**), Scan results (findings table with **scope/reachability/effective severity/AI verdict/reachability-path/capability-expansion filters**, default view = "Actionable Only" showing production-scope findings with available fixes), **AI Verdict Card** per finding (expandable: shows impact level badge, explanation text, code snippets with highlighted usage sites, data flow diagram, confidence score, **call-graph reachability path visualization**, "re-analyze" button, **"Submit Feedback" form**), Upgrade path details with **fix risk badges** (patch/minor/major) + **"Create Fix PR" button**, **Behavioral Analysis tab** (capability profile per package, capability-expansion history, hallucination candidates), **Call-Graph tab** (entry-point coverage, partial-graph warnings, reachability statistics), **Attack Stories tab** (ASPM-correlated attack chains with interactive graph visualization using reactflow), **License Compliance tab** (license breakdown pie chart, policy violations list, allow/deny/warn badges, license policy management), **SBOM tab** (generate + download SPDX/CycloneDX, format picker, include-vulns toggle, VEX overlay toggle), **Secrets findings tab** (secret type badges, masked values, file:line links, git history indicator), **Fix PRs tab** (list of auto-created PRs with status: pending/created/merged/closed, safety check results), **Container Images tab** (scanned images with layer breakdown, base-image recommendations, signature/provenance badges), **VEX Consumer tab** (configured feeds, recent statements applied, provenance trail per dismissal), **Rego Policies tab** (upload/edit/test Rego policies with dry-run UI), **AI Feedback Dashboard** (admin-only: prompt-quality metrics over time, per-model accuracy, per-CWE FP rate, eval-set export button), Policy management, Integration settings (Jira/Slack/GitHub/email), Notification preferences, Report generation + download
- **Default view philosophy**: show the minimum a developer needs to act on. Everything else is one click away but not in-your-face.
- **AI Verdict UX**: Findings with AI verdict show a colored badge (red = critical impact, yellow = moderate, gray = no impact) and a one-line summary. Click to expand full explanation with code snippets + call-graph reachability path if available. Findings without AI analysis (--no-ai mode) show "AI analysis not run" with a button to trigger it.
- **Feedback Form UX**: On every AI verdict card, a "Was this verdict correct?" prompt with thumbs up/down. Thumbs-down opens a structured feedback form: feedbackKind (wrong_impact / wrong_explanation / missed_usage / spurious_usage / other), correctedImpactLevel (dropdown), optional note. Submission creates `AIVerdictFeedback`. UI shows aggregated accuracy for the current model/prompt in a subtle footer.
- **Attack Story UX**: Each attack story card shows: combined severity badge, one-line narrative, list of constituent findings with their severities, "View chain" button that opens an interactive graph (reactflow) with nodes (findings) and edges (typed relationships). Node click drills into the finding detail.
- **Reachability path UX**: For findings with call-graph paths, an expandable tree view shows the path `entry-point -> func -> func -> ... -> vulnerable function`, each node clickable to jump to the source line in the (read-only) code viewer tab.
- API client generated from swagger spec

### Phase R (FUTURE, opt-in) -- Runtime Reachability eBPF Agent

Deferred until core platform is stable. Delivers production-runtime prioritization signal.

- `cmd/agent/main.go` -- Standalone Go binary deployable as DaemonSet in Kubernetes or sidecar in Docker Compose. Requires CAP_BPF / CAP_SYS_ADMIN.
- `internal/agent/ebpf.go` -- Uses `github.com/cilium/ebpf` to attach kprobes/uprobes. Observes:
  - `mmap`/`openat` calls to identify loaded shared objects + JS/Python module files per process
  - Function-level syscall sampling for vulnerable-function signatures (optional, higher overhead)
- `internal/agent/reporter.go` -- Periodically pushes a diff of loaded-module-state to the platform API at `POST /api/v1/runtime-observations` with project-binding token.
- Platform side: each finding gets a `runtimeObserved: true/false/unknown` flag. Effective severity Stage 4 (added post-exploit-intel): `runtimeObserved=false` + complete graph = extra downgrade; `runtimeObserved=true` + KEV = maximum priority.
- Privacy / security: agent only reports aggregated module paths + function symbol names, never request data or user content. Auth via short-lived project tokens. All communication over mTLS.

## CI/CD Security Pipeline (Dogfooding)

ArtFend VulnCheck should eat its own dog food -- the platform's own codebase is scanned by its own tools in CI.

### GitHub Actions Workflow (`.github/workflows/security.yml`)

```yaml
name: Security Scan
on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]
  schedule:
    - cron: '0 6 * * 1'  # Weekly Monday 6 AM UTC

jobs:
  go-security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'

      # Supply chain: verify module integrity
      - name: Verify Go modules
        run: go mod verify

      # Known CVEs in Go dependencies
      - name: govulncheck
        run: |
          go install golang.org/x/vuln/cmd/govulncheck@latest
          govulncheck ./...

      # Static security analysis (Go-specific patterns)
      - name: gosec
        run: |
          go install github.com/securego/gosec/v2/cmd/gosec@latest
          gosec -fmt sarif -out gosec-results.sarif ./...

      # Multi-source vulnerability scan (deps + Dockerfile + K8s)
      - name: Trivy filesystem scan
        uses: aquasecurity/trivy-action@master
        with:
          scan-type: 'fs'
          severity: 'HIGH,CRITICAL'
          exit-code: '1'

      # OSV scanner (Google's vuln DB -- broader than Go vuln DB)
      - name: OSV Scanner
        uses: google/osv-scanner-action/osv-scanner-action@v2
        with:
          scan-args: |-
            --recursive
            .

      # Custom SAST rules (our Semgrep rules from rules/semgrep/)
      - name: Semgrep (custom + auto rules)
        uses: semgrep/semgrep-action@v1
        with:
          config: >-
            rules/semgrep/go-gin-gorm.yaml
            rules/semgrep/generic.yaml
            p/golang

      # Secret detection (prevent accidental credential commits)
      - name: Gitleaks
        uses: gitleaks/gitleaks-action@v2
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

      # Upload SARIF results to GitHub Security tab
      - name: Upload SARIF
        if: always()
        uses: github/codeql-action/upload-sarif@v3
        with:
          sarif_file: gosec-results.sarif

  docker-security:
    runs-on: ubuntu-latest
    if: github.event_name == 'push' && github.ref == 'refs/heads/main'
    steps:
      - uses: actions/checkout@v4
      - name: Build Docker image
        run: docker build -t artfend-vulncheck:ci .
      - name: Trivy image scan
        uses: aquasecurity/trivy-action@master
        with:
          image-ref: 'artfend-vulncheck:ci'
          severity: 'HIGH,CRITICAL'
          exit-code: '1'
```

### Makefile Security Targets

```makefile
.PHONY: security security-quick security-full

# Quick scan: local developer workflow (< 30 seconds)
security-quick:
	go mod verify
	govulncheck ./...
	gosec ./...

# Full scan: all tools (CI equivalent, 2-3 minutes)
security-full: security-quick
	trivy fs --severity HIGH,CRITICAL .
	osv-scanner --recursive .
	semgrep --config rules/semgrep/ --config p/golang .
	gitleaks detect --source .

# Alias for quick scan
security: security-quick
```

### Pre-Commit Hook (optional, via `.pre-commit-config.yaml`)

Developers can opt into running security checks locally before commit:
- `gosec` on changed `.go` files only (fast, <5s)
- `gitleaks` on staged changes (prevent secret commits)
- Full suite runs in CI regardless

## Risks and Open Questions

- **npm audit API**: Bulk advisory endpoint requires specific POST format. No official Go SDK -- raw HTTP. Private registry advisory endpoints may differ from public npm.
- **NVD rate limits**: 5 req/30s without API key. Vuln cache in `internal/vuln/` mitigates this. Users should configure `NVD_API_KEY` for production.
- **Jira Cloud auth**: Requires API token + email (basic auth) or OAuth 2.0. Start with API token, add OAuth later for marketplace compatibility.
- **package-lock.json versions**: v1 (npm 5-6), v2 (npm 7-8), v3 (npm 9+) have different structures. Parser must handle all three. All versions encode `dev`/`optional` flags differently -- parser must normalize to our `DepScope` enum.
- **Scope propagation edge cases**: A package can appear as both a dev transitive dep AND a prod transitive dep (diamond dependency). The correct scope is `production` because it IS reachable via a production path. Graph traversal must check ALL paths, not just the first one found.
- **Effective severity tuning**: The downgrade rules (dev scope = -2 levels) are opinionated. Need to be configurable per org, with sensible defaults. Some security teams may want to see dev deps at full severity.
- **Supply chain false positives**: Typosquat detection via Levenshtein can flag legitimate packages with short names. Need a whitelist of known-good packages and a configurable similarity threshold.
- **Top-1000 npm packages list**: Needed for typosquat detection. Must be refreshed periodically. Can be fetched from npm registry sorted by weekly downloads, cached with 7-day TTL.
- **Postinstall hook analysis**: Parsing `package.json` scripts of ALL transitive deps requires downloading package metadata. For large trees (1000+ deps), this can be slow. Consider lazy analysis: only check direct deps + new deps since last scan.
- **Git clone abuse -- scanning repos you don't own**: Mitigated by ownership verification: before cloning, the server calls the provider API with the user's credential to confirm the authenticated user is a collaborator/member of the repo. Without this check, someone with a valid PAT could scan any public repo and generate findings (potentially using our AI analysis as a free vulnerability research tool on arbitrary projects). Ownership check ensures you can only scan what you have access to.
- **Git clone abuse -- repo bombs**: A malicious repo could contain huge binary blobs, `.git` directory tricks, or symlinks to host filesystem. Mitigated by: (1) `--depth 1` prevents history-based bloat, (2) post-clone size check with hard limit, (3) `core.symlinks=false` disables symlinks, (4) file extractor only reads known extension patterns, ignores everything else.
- **Git clone abuse -- malicious git hooks**: A repo's `.git/hooks/` or `core.hooksPath` config could execute arbitrary code on clone. Mitigated by `--config core.hooksPath=/dev/null` which disables all hooks during clone.
- **Git clone abuse -- SSRF via clone URL**: An attacker could provide a repo URL pointing to an internal service (`http://169.254.169.254/` AWS metadata, `http://localhost:8080/`). Mitigated by URL validation: only allow HTTPS to hosts in `GIT_ALLOWED_HOSTS` allowlist, reject raw IPs, reject private network ranges, reject non-443 ports.
- **Git credential security**: Tokens stored encrypted (AES-256-GCM). Never returned via API. Decrypted only in memory during clone. `ENCRYPTION_KEY` must be rotated periodically; plan for key rotation (decrypt with old, re-encrypt with new).
- **Webhook replay attacks**: An attacker could replay a legitimate webhook payload to trigger excess scans. Mitigated by timestamp checking (reject >5 min old), dedup (skip if scan already running), and rate limiting.
- **Self-hosted git providers**: Orgs using GitLab CE/EE or Bitbucket Server need to add their domain to `GIT_ALLOWED_HOSTS`. The provider detection logic should infer the provider from the API response (GitLab returns `X-Gitlab-Meta` header, etc.) rather than relying solely on domain name.
- **OAuth token expiry**: OAuth tokens expire. The `GitCredential.valid` flag + `lastValidatedAt` field allow periodic re-validation. A River job should check all credentials periodically and mark expired ones as `valid=false`, notifying the user.
- **Reachability accuracy (scope-only mode)**: Without AI or SAST call-graph analysis, reachability is scope-based only (dev vs prod), not function-level. This still eliminates the majority of false positives.
- **AI Impact Analysis -- LLM hallucination**: LLMs can hallucinate that a function is/isn't called. Mitigations: (1) always include actual code snippets in the prompt (grounded analysis, not guessing), (2) require confidence score >= 0.3, (3) never auto-resolve a finding based solely on AI verdict -- AI downgrades effective severity but the finding remains visible, (4) provide "re-analyze" button for manual re-check.
- **AI Impact Analysis -- code privacy**: User source code is sent to external LLM APIs. Mitigations: (1) support self-hosted LLMs via OpenAI-compatible API (e.g. vLLM, Ollama), (2) only send relevant snippets (10-15 lines per usage site), not entire files, (3) configurable: orgs can disable AI analysis entirely or restrict to self-hosted models only, (4) no source code is stored on the server beyond the analysis session.
- **AI Impact Analysis -- cost**: GPT-4o / Claude at scale can be expensive. Mitigations: (1) default to `gpt-4o-mini` / `claude-haiku` for cost efficiency, (2) skip AI for findings already classified as `unreachable` by scope (dev-only deps don't need LLM), (3) cache verdicts by finding fingerprint + source code hash (reuse across scans if code and vuln haven't changed), (4) batch API calls where possible.
- **AI Impact Analysis -- latency**: LLM calls add 2-10 seconds per finding. Mitigations: (1) run AI analysis in parallel (configurable concurrency), (2) for CI/CD, offer `--no-ai` flag for fast scope-only mode, (3) AI analysis can run as a separate async job after the initial scan returns results -- findings are first shown with scope-based reachability, then enriched with AI verdicts when ready.
- **AI Impact Analysis -- dynamic imports**: `require(variable)`, `import()` with expressions, and eval-based loading cannot be statically indexed. The code indexer should flag these as "dynamic import detected, AI analysis may be incomplete" and default to `potentially_reachable`.
- **AI Impact Analysis -- monorepos**: Large monorepos may have thousands of source files. The code indexer should respect project boundaries (scan only the relevant package/workspace, not the entire repo). Use `package.json` workspaces or project config to determine scope.
- **License detection accuracy**: Not all packages declare licenses correctly. Some use non-standard SPDX expressions ("BSD" instead of "BSD-3-Clause"), custom licenses, or no license at all. Parser must handle fuzzy matching (e.g. "MIT License" -> "MIT", "BSD" -> "BSD-2-Clause"), and flag `unknown` when detection fails. For `unknown`, suggest manual review.
- **License dual-licensing complexity**: Expressions like `(MIT OR GPL-2.0-only)` require understanding that the user can CHOOSE either license. Policy evaluation must check if ANY branch of an OR satisfies the policy. AND expressions require ALL licenses to satisfy. Nested expressions `(MIT OR (GPL-2.0-only AND ClasspathException))` need recursive evaluation.
- **SBOM regulatory requirements**: SPDX 2.3 and CycloneDX 1.5 have strict schema requirements. Generated SBOMs must pass validation tools (`spdx-tools`, `cyclonedx-cli validate`). Consider adding schema validation as part of generation pipeline.
- **SBOM VEX accuracy**: VEX statements (exploitable/not_affected) should align with AI verdicts and effective severity. When AI verdict changes after re-analysis, cached SBOMs become stale. Consider adding SBOM generation timestamp and "regenerate" option.
- **Automated Fix PRs -- lockfile regeneration**: Updating `package.json` version and regenerating `package-lock.json` requires running `npm install` in a sandboxed environment. This means the server needs npm/pip/go/dotnet runtime available (in Docker or as binaries). Sandbox must prevent network access to arbitrary URLs during install. Consider pre-built Docker images per ecosystem.
- **Automated Fix PRs -- merge conflicts**: If the user has local changes to `package.json` or lockfile, the fix PR will conflict. The PR description should warn about potential conflicts and include manual fix instructions as fallback.
- **Automated Fix PRs -- monorepo challenges**: In monorepos with npm workspaces, upgrading one dep may affect multiple packages. The fix PR must update the root lockfile and all affected workspace `package.json` files.
- **Automated Fix PRs -- safety check re-scan cost**: Each fix PR requires a full SCA re-scan on the modified lockfile to verify no new vulns. For projects with many fixable findings, this means many re-scans. Consider batch mode: group all safe fixes into one PR, re-scan once.
- **Secrets scanning -- false positive rate**: High-entropy string detection catches many non-secrets (UUIDs, base64-encoded non-sensitive data, test fixtures). The validator must be aggressive about filtering: check known false positive patterns, require minimum length, and allow per-org suppression rules. Track false positive rate per pattern and auto-disable patterns with >40% FP rate.
- **Secrets scanning -- git history performance**: Scanning full git history (`git log --all -p`) can be extremely slow for repos with long history. Limit to last N commits (configurable, default 1000) or scan only diff between branches. Consider using `gitleaks` as an alternative backend for history scanning.
- **Secrets scanning -- live validation**: Some tools (e.g. TruffleHog) can verify if detected secrets are still active by making API calls. This is powerful but risky (could trigger alerts on the secret owner's side). Consider as opt-in feature with explicit warning.
- **Upgrade path accuracy**: Requires querying npm registry for version metadata of direct deps. May be slow for large dep trees -- consider caching/parallel requests.
- **Fix risk level and semver**: Some packages don't follow semver strictly. A "patch" bump could still break things. The risk level is advisory, not a guarantee. CLI should always show what would change before applying.
- **Incremental scan cache invalidation**: If vuln databases update but lockfile hasn't changed, we still need to check for new CVEs. Incremental mode must track the timestamp of the last full CVE check, not just the lockfile hash.
- **GitHub App vs PAT**: PR checks work best with a GitHub App (installation tokens, fine-grained permissions). Start with PAT for simplicity, migrate to GitHub App later.
- **Email deliverability**: Self-hosted SMTP may land in spam. Consider supporting third-party providers (SendGrid, SES) via config.
- **PDF generation**: Go PDF libraries have limitations for complex layouts. If reports get complex, consider generating HTML first and converting to PDF.
- **Private registry diversity**: Artifactory, GitHub Packages, Verdaccio, and others have slightly different API shapes. Start with npm token-based auth (covers most), extend per-registry later.
- **File upload size**: `package-lock.json` can be 10+ MB for large projects. Configure multipart upload with size limits in Gin middleware.
- **Concurrent scans**: River handles job queuing, but max concurrent scan workers should be configurable in `internal/config/` to prevent resource exhaustion.
- **Web dashboard scope**: Large. Start with read-only views (dashboard, scans, findings with scope/reachability filters) + settings pages, defer advanced UX to later iteration.
- **Two-way Jira sync**: Requires polling Jira (no webhooks in initial version). Webhook support can be added later to reduce latency.
- **Alert fatigue metrics**: Tracking noise ratio (findings generated vs findings acted on) requires storing user actions on findings. Start simple: count findings by status (open vs ignored vs resolved) per project over time.
- **Multi-source data consistency**: The same CVE may have different severity ratings across NVD, GHSA, OSV, and VulnCheck. Strategy: prefer GHSA severity for ecosystem-specific packages (more accurate), fall back to NVD CVSS, and use VulnCheck EPSS/KEV as severity modifiers rather than overrides. Document the precedence logic clearly.
- **GHSA GraphQL rate limits**: GitHub API has 5000 points/hour limit. SecurityAdvisories queries cost 1 point each but pagination can consume more. For large scans (1000+ deps), batch queries by ecosystem and cache aggressively (12h TTL for advisory data).
- **VulnCheck API cost**: VulnCheck is a commercial API with usage-based pricing. Must degrade gracefully when not configured -- exploit intelligence fields (`exploitAvailable`, `epssScore`, `kevListed`) default to `null`/`false` and effective severity Stage 3 is skipped entirely. Never hard-fail a scan because a paid data source is unavailable.
- **Trivy DB freshness**: Trivy's offline DB requires periodic updates (`trivy db update`). If the DB is stale (>48h), log a warning but don't fail the scan. Consider auto-update in a background River job.
- **Semgrep custom rules maintenance**: Custom rules need ongoing tuning. False positives (e.g., `math/rand` used in non-security contexts like test data generation) require exceptions. Ship rules with severity and confidence levels; allow per-org rule overrides. Track FP rate per rule and auto-disable rules with >30% FP rate.
- **Semgrep execution security**: Semgrep runs on user-submitted code. While Semgrep itself is safe (it doesn't execute the code), ensure it runs in a sandboxed environment with resource limits (timeout: 60s per file, memory: 1GB max). Large files (>100KB) should be skipped.
- **EPSS score volatility**: EPSS scores are recalculated daily and can change significantly. Cache scores with 24h TTL and log score changes >0.1 for auditability. Don't send alerts based solely on EPSS changes -- only alert when a finding crosses the effective severity threshold.
- **KEV catalog as a forcing function**: CISA KEV inclusion is a strong signal but is US-centric. International teams may want to configure alternative catalogs or disable KEV-based severity upgrades. Make it configurable per org.
- **Source cross-validation confidence**: When 3+ sources confirm a CVE, confidence is high. When only 1 source reports it, it may be a false positive or a very new CVE not yet propagated. Consider showing a "confirmation status" badge: "confirmed" (3+ sources), "likely" (2 sources), "unconfirmed" (1 source).
- **Python scope detection inconsistency**: `requirements.txt` has no native dev/prod distinction -- relies on file naming conventions (`dev-requirements.txt`). `Pipfile.lock` and `poetry.lock` have explicit scope. Parser must handle all cases and clearly document assumptions. Consider adding a manual scope override in `.vulncheck` policy file.
- **Python virtual environments**: Python projects may have multiple `requirements.txt` files (per-service in monorepo). Parser must handle `constraints.txt`, `-c constraints.txt` directives, and `-r` file includes. `pyproject.toml` (PEP 621) is becoming the standard -- prioritize it over legacy formats.
- **.NET target framework matrix**: NuGet packages can have different vulnerabilities per target framework (`net6.0` vs `net8.0`). The parser must extract `TargetFramework` from `.csproj` and filter advisories accordingly. A package may be vulnerable in `net6.0` but patched in `net8.0`.
- **.NET Central Package Management**: When `Directory.Packages.props` is used, individual `.csproj` files don't specify versions. Parser must resolve versions from the central file, which can be at any parent directory level.
- **Go test-only scope**: Go doesn't have a native dev/prod split -- all dependencies are compiled. However, imports only from `_test.go` files can be treated as dev-scope (test-only). The parser should analyze import graphs to determine which modules are only used in test code.
- **Go replace directives**: `replace` directives in `go.mod` point to local paths or alternative modules. These bypass the standard module proxy and vuln DB. The scanner should resolve `replace` targets and check them against the vuln DB using the replacement module path, while flagging `replace` to local paths as a supply chain signal.
- **85+ SAST rules across 4 ecosystems**: Maintaining rules at this scale requires a testing strategy. Each rule should have positive and negative test cases (vulnerable code + safe code). Consider a `rules/semgrep/tests/` directory with test fixtures and a CI job that runs `semgrep --test`.
- **Ecosystem-specific remediation**: npm has `npm audit fix`, Python has `pip install --upgrade`, .NET has `dotnet outdated`, Go has `go get -u`. The remediation engine needs ecosystem-specific strategies for suggesting and applying fixes. Start with advisory-based upgrade suggestions, defer auto-fix to CLI phase (N).
- **AI code indexer language coverage**: Import pattern recognition varies in complexity: Python has `import x`, `from x import y`, `importlib`, conditional imports; C# has `using`, `global using`, namespace aliases; Go has import blocks with aliases. Each needs dedicated regex/AST logic. Consider using tree-sitter for more robust parsing if regex becomes fragile.
- **Behavioral analysis tarball storage**: Analyzing tarballs requires downloading them. For large dep trees (1000+ packages), the total download size can be GBs. Mitigations: (1) cache profiles aggressively keyed by tarball hash, (2) only analyze new/updated packages per scan (skip already-profiled versions), (3) configurable max-tarball-size (default 100MB), (4) parallel downloads with connection pooling.
- **Behavioral analysis sandboxing**: Unpacking untrusted tarballs could trigger zip-bomb / symlink attacks / path-traversal. Mitigations: (1) unpack in a temp dir on a size-limited tmpfs, (2) reject tarballs with absolute paths, path-traversal (`../`), or symlinks to outside the temp dir, (3) hard cap on extracted file count + total size, (4) no network access during analysis (seccomp/egress firewall if running in container).
- **Behavioral capability false positives**: Many legitimate libraries have `network` / `filesystem` / `shell` capabilities by design (http clients, logging libs, build tools). The `capability_expansion` signal -- which fires only on DIFF -- is more actionable than raw capability detection. Document clearly that capabilities alone are informational; only expansions are findings.
- **Hallucinated package feed quality**: The hallucination feed depends on public research + internal heuristics. Risk of false positives (legitimate new packages flagged). Mitigations: (1) require two predicates to match before flagging (not just feed membership), (2) allow per-org overrides, (3) quarantine policies are time-limited (default 14 days) during which users can mark as legitimate, auto-removing the block.
- **Call-graph build time**: Full call-graph analysis on large monorepos (10k+ files) can take minutes and consume gigabytes of RAM. Mitigations: (1) per-ecosystem timeouts + memory limits enforced via subprocess sandboxing, (2) incremental call-graph caching keyed by source-file hashes (skip unchanged files), (3) partial-graph mode: analyze only files that transitively import any CVE-affected package (prunes ~80% of work for most repos), (4) fallback to scope-based reachability if call-graph fails.
- **Call-graph language limitations**: Dynamic dispatch (JS reflection, Python `getattr`, Go reflect, C# dynamic) defeats static call-graph analysis. For these cases the graph is marked `partial` and we conservatively set `reachability=potentially_reachable`. Never claim `unreachable` on a partial graph.
- **Call-graph subprocess management**: Node.js / Python / .NET analyzers run as external processes (ts-morph via node, jedi via python, Roslyn via dotnet). Requires these runtimes to be available in the deployment. Mitigations: (1) ship Docker image with all three pre-installed, (2) make each analyzer optional (skip gracefully if runtime missing), (3) Go analyzer is always available (native), so Go projects always get call-graph.
- **Container scan image size**: Some images (CUDA, ML, heavy-OS) are multi-GB. Pulling + extracting layers consumes disk and time. Mitigations: (1) pull only the necessary layers (skip unused for our analysis -- we need root-fs layers, not runtime-only), (2) stream-extract instead of materializing full image to disk, (3) hard cap on image size (default 10GB), (4) reuse `CONTAINER_IMAGE_CACHE_DIR` across scans keyed by digest.
- **Cosign / Sigstore keyless trust model**: Keyless signing trusts the OIDC issuer (GitHub Actions, GitLab CI, etc.) and Sigstore Rekor transparency log. Compromised OIDC issuers or Rekor would undermine the model. Mitigations: (1) allow orgs to configure strict signer allowlist by OIDC subject regex (e.g. only accept signatures from `https://github.com/myorg/*/.github/workflows/build.yml@refs/heads/main`), (2) support static public-key signatures as belt-and-suspenders.
- **SLSA provenance availability**: Only a small fraction of npm/PyPI packages currently publish SLSA attestations (mostly major orgs with GitHub Actions). The `missing_provenance` signal will fire on MOST packages. Severity defaults to INFO, not blocking. As the ecosystem adopts SLSA, severity can be raised per org config.
- **Rego policy complexity**: Powerful, but users can write policies that never terminate, consume excess memory, or have subtle logic errors. Mitigations: (1) hard timeout (default 30s), (2) memory limit on OPA evaluation, (3) dry-run API endpoint for testing against historical scans before enabling, (4) ship well-commented example policies, (5) policy-linter CLI (`vulncheck policy eval --rego <file>`) catches common mistakes.
- **Upstream VEX trust model**: Accepting VEX statements from upstream maintainers means trusting their correctness. A compromised maintainer could push `not_affected` statements to hide exploitable CVEs. Mitigations: (1) support signed VEX via sigstore/PGP, optionally require signatures per feed, (2) log every auto-dismissal with full provenance so security teams can audit, (3) configurable feed allowlist per org, (4) expire VEX applications after N days and re-validate.
- **VEX feed format fragmentation**: OpenVEX, CSAF 2.0, CSAF 1.x, VDR, and vendor-specific formats all exist. Start with OpenVEX + CSAF 2.0 which cover >80% of serious feeds. Add parsers for other formats incrementally based on user demand.
- **PR baseline diff requires baseline**: On a brand-new project with no prior scan, there is no baseline. First PR should fall back to full scan with a notice ("first scan -- showing all findings"). Force-push scenarios that rewrite history can invalidate the baseline; handle by rescanning.
- **PR baseline diff in monorepos**: Scope baseline and PR-head scans to the changed paths in monorepos with many unrelated projects. Without scoping, even tiny PRs trigger full-repo scans. Use `paths-changed` from the Git diff to scope.
- **ASPM false attack stories**: Chain templates may fire on coincidental entity overlaps (e.g. two findings happen to mention `db` in their entities). Mitigations: (1) require at least one "strong" entity match (package name, service name, DB connection string) -- weak matches (generic file paths like `utils.js`) alone don't trigger chains, (2) confidence score on each story, (3) users can dismiss spurious stories which trains the heuristic.
- **AI feedback data privacy**: User corrections may leak code context. Mitigations: (1) eval-set export is org-opt-in only, (2) exported sets strip source code snippets by default (only CVE ID + user's impact label), (3) orgs can choose "code-included" export only with explicit legal review.
- **LSP performance on huge lockfiles**: Hovering over a version in a 5MB `package-lock.json` shouldn't freeze the IDE. LSP server maintains an in-memory index of the parsed lockfile per buffer, re-indexes on save (debounced 500ms), and quick-hover uses the cached index without re-parsing.
- **LSP authentication**: The LSP server needs to call the remote API on the user's behalf. Reuses `~/.vulncheck/credentials` written by `vulncheck auth login`. Refresh token flow for OAuth providers. Never embeds tokens in extension packages.
- **IDE extension cross-platform**: VSCode extension works on VSCode, VSCodium, Cursor, Windsurf (all fork-compatible). JetBrains plugin must target multiple IDE families (IntelliJ IDEA, PyCharm, GoLand, Rider, WebStorm). Shared LSP server is the right abstraction.
- **Runtime agent performance overhead**: eBPF probes have <1% CPU overhead typically but kprobes on hot paths can degrade. Mitigations: (1) default to coarse-grained (loaded-modules only) mode, (2) fine-grained syscall sampling is opt-in per workload, (3) allow sampling rate tuning.
- **Runtime agent deployment complexity**: Requires CAP_BPF/CAP_SYS_ADMIN in cluster. Not all ops teams will allow this. Document trade-offs clearly; make the whole feature optional.

## Summary

- **New packages**: 33 (config, domain, middleware, auth, project, **git**, scan, policy, sca, **ai**, **behavior**, **callgraph**, **container**, **slsa**, **vex**, **aspm**, **feedback**, **license**, **sbom**, **fixpr**, **secrets**, sast, dast, iac, **lsp**, integration, monitor, notify, prcheck, report, vuln, ecosystem) + 4 ecosystem parsers (npm, python, dotnet, gomod) + 2 entry points + jobs + httputil + web frontend + **IDE plugins (VSCode + JetBrains)**
- **Ecosystems (4, progressive)**: Node.js (npm/yarn/pnpm), Python (pip/pipenv/poetry), .NET (NuGet), Go (go.mod). Each with full SCA + SAST + supply chain + **behavioral package analysis** + **call-graph reachability** + AI code indexer + license detection. Future: Java, Ruby, PHP, Rust via same `EcosystemParser` interface.
- **Vulnerability data sources (6)**: npm audit API, OSV.dev, NVD, **GitHub Advisory Database (GHSA)**, **VulnCheck API** (exploit intelligence, EPSS, KEV), **Trivy DB** (offline, misconfigs). Multi-source dedup with cross-validation confidence scoring. Sources are ecosystem-aware (GHSA/OSV filter by ecosystem). **Plus upstream VEX consumption** (OpenVEX + CSAF feeds) to auto-dismiss not-affected CVEs.
- **SAST rules (100+)**: Custom Semgrep rules across 4 ecosystems + generic + AI/LLM + Dockerfile: Node.js/Express/NestJS (20+), Python/Django/Flask/FastAPI (20+), ASP.NET Core (20+), Go/Gin/GORM (25+), **AI/LLM security rules (15-20: prompt injection, unsafe LLM output handling, MCP trust boundaries, LangChain/LlamaIndex misconfigs, LLM rate-limit DoS, agentic AI risks)**, **Dockerfile rules (10+)**, language-agnostic secrets/credentials detection. Covering injection, XSS, auth, deserialization, race conditions, SSRF, crypto misuse, secrets in logs, framework-specific patterns, AI-integration risks, container hardening.
- **Secrets scanning**: Dedicated engine with 100+ regex patterns (AWS, GCP, Azure, GitHub, Slack, Stripe, DB URLs, private keys) + Shannon entropy analysis (>4.5 bits/char) + false positive validation + git history scanning. Runs as a separate scanner alongside SCA/SAST.
- **Behavioral Package Analysis (NEW, Socket-style)**: AST-based capability extraction (network/filesystem/shell/env/unsafe_eval/native_binding/obfuscated/install_hook_shell) per published tarball + capability diff between versions + semantic install-script analysis + hallucinated/slopsquat package detector with auto-quarantine. Catches malicious packages **before CVEs exist**.
- **Call-Graph Reachability (NEW, Endor-style)**: Deterministic function-level call-graph analysis per ecosystem (ts-morph/TypeScript, jedi/Python, golang.org/x/tools native/Go, Roslyn/.NET). Entry-point discovery (HTTP routes, CLI main, message handlers, exported APIs) + BFS path tracing from entry points to vulnerable functions. Runs BEFORE AI Impact Analyzer to cheaply dismiss unreachable findings -- 95%+ deterministic noise reduction without LLM cost.
- **Container / OCI Image Scanning (NEW)**: Full image scanning via go-containerregistry + Trivy DB. Per-layer OS package extraction (apt/apk/rpm) + Dockerfile SAST (10+ rules) + base-image detection with minimal-alternative recommendations (distroless, chainguard) + Sigstore/cosign signature verification + SLSA provenance attestation verification + container-specific SBOM with layer breakdown.
- **SLSA / Sigstore Provenance (NEW)**: For each dep (+ containers): query registry for SLSA attestations, verify Sigstore signatures via Rekor transparency log, emit supply-chain signals (`missing_provenance`, `provenance_mismatch`, `unsigned_artifact`, `provenance_verified`). Compliance tier for EU CRA / US EO 14028.
- **Upstream VEX Consumer (NEW)**: Periodic feed ingestion (OpenVEX + CSAF 2.0) from configurable sources (Kubernetes, Red Hat, Node.js security WG, Python security WG, PyPA, Rust advisory DB, org-specific feeds). Signature-verified statements auto-dismiss findings with full provenance tracking. Drastically reduces false positives without heuristics.
- **Policy-as-Code (Rego/OPA) (NEW)**: Opt-in Rego engine for composite enterprise gates against full schema (Finding + Dependency + License + VEX + AIVerdict + ExploitIntel + AttackStory). Three outcomes: block/warn/notify. Dry-run API for testing against historical scans. Shipped example policies.
- **PR Baseline Diff (NEW)**: On PR webhooks, diff findings against last successful defaultBranch scan. Post only NEW findings as blocking + RESOLVED in green + inline lockfile diff. Eliminates "legacy finding dump" fatigue that plagues Snyk/npm audit PR comments.
- **ASPM Attack-Story Correlation (NEW)**: Post-scan correlator joins findings across SCA/SAST/Secrets/IaC/Container by shared entities. Builds typed attack graph, surfaces prioritized "Attack Story" cards with interactive chain visualization. Pre-defined chain templates (secret+injection+SCA-in-path, IaC-public+missing-auth+KEV, container-root+RCE, dep-confusion+behavioral-expansion).
- **AI Verdict Feedback Loop (NEW)**: Structured feedback form on every AI verdict card -> `AIVerdictFeedback` records -> aggregated prompt-quality metrics dashboard (per-model accuracy, per-CWE FP rate, per-language confidence) -> opt-in anonymized eval-set export for prompt-tuning. Compounding advantage over closed-box ML in Snyk.
- **Multi-file Repository-level AI Context (NEW)**: Framework detector + entry-point/route extractor + DI binding resolver + environment classifier -> token-budgeted repo envelope injected into LLM prompts -> drastically improves accuracy on middleware chains, DI-based wiring, framework-specific patterns.
- **IDE Plugin via LSP (NEW)**: Go-based Language Server (`vulncheck lsp`) + VSCode extension + JetBrains plugin. Hover hints on package versions (CVE count, AI verdict one-liner, capability badge, fix-risk), save-time debounced scan, quick-fix code actions (trigger Fix PR, apply safe upgrade, add ignore with template, explain with AI).
- **Runtime Reachability Agent (FUTURE, opt-in)**: Light eBPF-based Go agent deployable as DaemonSet/sidecar. Observes loaded modules + function-level syscall sampling per process. Findings gain `runtimeObserved` flag driving production prioritization.
- **Key features**: **Multi-ecosystem** SCA with full transitive dependency graph **with scope awareness** (dev/prod/optional/peer), **context-based effective severity** (not just raw CVSS -- also **exploit intelligence**: EPSS score, KEV listing, known exploit availability), **reachability classification** (scope-based + **deterministic function-level call-graph** + AI-powered), **AI-powered impact analysis** (LLM analyzes user's actual code to determine if vulnerability is exploitable -- supports JS/TS/Python/C#/Go import patterns + repo-level framework context), **behavioral package analysis** (catches malicious packages pre-CVE), **hallucinated/slopsquat detection**, **supply chain signal detection** per ecosystem (postinstall hooks, setup.py hooks, typosquatting, new packages, maintainer changes, SLSA provenance, Sigstore signatures), **safe remediation with fix risk levels** (patch/minor/major/no_fix), upgrade path remediation, **automated fix PRs** (auto-create PRs for safe upgrades with safety check re-scan), **license compliance** (SPDX expression parsing, org-level allow/deny/warn policies, dual-licensing support), **SBOM generation** (SPDX 2.3 + CycloneDX 1.5 with VEX vulnerability overlay, plus container SBOM with layer breakdown), **upstream VEX consumption** (OpenVEX + CSAF), **secrets scanning** (100+ patterns + entropy + git history), **triple policy system** (.vulncheck file + DB + Rego/OPA), **incremental scanning** (lockfile hash + delta CVE check), **PR baseline diff** (only NEW findings blocked on PRs), **container / OCI image scanning**, **ASPM attack-story correlation**, **AI verdict feedback loop**, **IDE plugin (VSCode + JetBrains via LSP)**, continuous monitoring, two-way Jira sync, email+Slack notifications, GitHub PR checks, PDF/CSV compliance reports, private registry support, **CI/CD security pipeline** (dogfooding -- platform scans itself)
- **Core differentiators vs npm audit / Snyk / Socket / Endor / Aikido**: **Multi-ecosystem from day one** (not Node.js only); **AI answers "is this dangerous for MY code?"** by analyzing actual import/call sites and data flow (Snyk lacks this); **deterministic function-level call-graph reachability** runs before AI for cheap 95% noise reduction (Endor Labs parity); **behavioral package analysis catches malicious packages pre-CVE** (Socket.dev parity); **hallucinated/slopsquat package detection** (new 2026 attack class); **AI/LLM security rule pack** (prompt injection, unsafe LLM output, MCP trust boundaries -- none of the incumbents have this fully); **6 vuln data sources + upstream VEX consumption** with cross-validation (unique combination); **exploit intelligence** (EPSS + KEV) prioritizes CVEs actually being exploited in the wild; **100+ custom SAST rules** catch code patterns that no CVE database covers; **automated fix PRs with safety check** (Snyk creates fix PRs but doesn't verify the fix doesn't introduce new vulns); **license compliance with SPDX expression parsing** (handles complex dual-licensing); **SBOM generation + VEX bidirectional** (generate AND consume, SPDX + CycloneDX, regulatory compliance ready); **container / OCI image scanning with base-image recommendations + cosign/SLSA verification**; **secrets scanning** (dedicated engine, not just regex); **Rego/OPA policy-as-code for enterprise gates** (Endor Labs parity); **PR baseline diff** (show only what THIS PR added); **ASPM attack-story correlation** (Apiiro/OX parity); **AI verdict feedback loop** compounding advantage over closed-box ML; **IDE plugin (VSCode + JetBrains via unified LSP)** matches Snyk's adoption driver; context-aware scope analysis eliminates ~90% false positives; supply chain detection catches what CVE databases miss; safe remediation never breaks builds by default; incremental scans keep CI fast; effective severity ensures only actionable findings block pipelines
- **Overall complexity**: XXL -- this is an enterprise-grade AppSec / ASPM product. **28+ implementation phases** (core A-O plus E-bis/E2-bis/F2/F3/Q/M2/P/P2 plus future R), phased delivery keeps each step manageable.
- **Phase 1 deliverable** (A-G + E-bis + E2-bis): Working **Node.js** SCA scanner with **git repo integration** (GitHub/GitLab/Bitbucket, public+private), dep graph, **SAST engine + Node.js/Express/NestJS + AI/LLM SAST rules**, **context analysis, call-graph reachability, AI impact analysis with repo-level context, effective severity, supply chain detection, behavioral package analysis, hallucination detection**, remediation with risk levels, **license compliance + SBOM + secrets scanning + automated fix PRs**, dual policies -- the core product with all primary differentiators
- **Phase 2 deliverable** (Q + H-I + E3): **Rego policy-as-code + upstream VEX consumer + SLSA/Sigstore verification**, Jira integration, monitoring (incremental), **Python ecosystem** (pip/pipenv/poetry SCA, Django/Flask/FastAPI SAST rules, PyPI supply chain, Python call-graph + behavioral analysis)
- **Phase 3 deliverable** (J-L + E4): Notifications, **PR checks with baseline diff (only NEW findings block)**, reports + **.NET + Go ecosystems** (NuGet/go.mod SCA, ASP.NET Core + Go/Gin/GORM SAST rules, .NET/Go call-graph + behavioral analysis)
- **Phase 4 deliverable** (M-O + M2 + P + P2): DAST/IaC scaffolds, **container/OCI image scanning (Phase M2)**, CLI (scope-aware output, AI explain command, safe fix command, license audit, SBOM export, secrets scan, fix-pr command, behavior command, reachability command, container command, verify-provenance command, policy eval command, pr-diff command, attack-stories command, multi-ecosystem auto-detection), web dashboard with **AI verdict feedback + attack-story tab + behavioral analysis tab + call-graph tab + container images tab + VEX consumer tab + Rego policies tab + AI feedback metrics** (Phase O), **IDE plugin (VSCode + JetBrains via LSP) (Phase P)**, **ASPM attack-story correlation (Phase P2)**
- **Phase 5 (future, opt-in)** (R): Runtime reachability eBPF agent for production prioritization
