---
name: VulnCheck Stage 4 -- V1 GA
overview: ~2 months (6-8 weeks) extension. Turns the Node.js-focused Beta into a multi-ecosystem enterprise-ready GA product. Adds Python (pip/pipenv/poetry, Django/Flask/FastAPI SAST, Python call-graph), .NET (NuGet, ASP.NET Core SAST, Roslyn call-graph), Go (go.sum/go.mod, Gin/GORM/stdlib SAST, Go call-graph -- now for external targets not just self-host). Adds enterprise controls: Rego policy engine (OPA), upstream VEX consumer (OpenVEX + CSAF), SLSA provenance + Sigstore/Rekor signature verification. Adds DAST and IaC scaffolds (OWASP ZAP + Checkov/Trivy misconfig). Adds Container / OCI image scanning (layer extraction, OS package scanning, Dockerfile SAST, base-image recommendations). Replaces minimal dashboard with the full Web Dashboard (all tabs, write operations, user/org management, policies, fix PR triggers). Adds IDE Plugin (VSCode + JetBrains via LSP server) with inline hints and quick-fixes.
todos:
  - id: s4-python-ecosystem
    content: "Python ecosystem: pip/pipenv/poetry parsers; Django/Flask/FastAPI SAST rules (25+); Python call-graph via jedi + pyan3 (subprocess or wrapped); pip-compatible behavioral analyzer."
    status: pending
  - id: s4-dotnet-ecosystem
    content: ".NET ecosystem: packages.config / PackageReference / packages.lock.json parsers; ASP.NET Core SAST rules (25+); Roslyn-based call-graph via subprocess analyzer tool."
    status: pending
  - id: s4-go-ecosystem
    content: "Go ecosystem (external targets): go.mod / go.sum parsers; Gin/GORM/stdlib SAST rules (20+); call-graph reused from Stage 3 self-host work; govulncheck integration."
    status: pending
  - id: s4-rego-policies
    content: "Rego policy engine (OPA embedded): policy packages per org; unit-test harness; Rego examples (block-critical-prod, require-signed-images, license-allowlist); Rego-as-alternative to DB+file policies (both still supported, Rego optional advanced mode)."
    status: pending
  - id: s4-vex-consumer
    content: "VEX consumer: OpenVEX + CSAF + CycloneDX VEX parsers; trust-pinned upstream feeds; auto-dismiss findings matching NOT_AFFECTED statements with rationale + source attribution."
    status: pending
  - id: s4-slsa-sigstore
    content: "SLSA provenance + Sigstore / Rekor verification: verify signed artifacts (containers, Go modules, PyPI wheels with PEP 740); validate SLSA provenance chain; flag unsigned / unverifiable."
    status: pending
  - id: s4-container-scanning
    content: "Container / OCI image scanning: pull image, extract layers, scan OS packages (apt/apk/rpm via Trivy DB), Dockerfile SAST, base-image recommendations, Sigstore verification, SBOM for image."
    status: pending
  - id: s4-dast-iac-scaffold
    content: "DAST scaffold (OWASP ZAP wrapper, baseline scan) + IaC scaffold (Checkov + Trivy misconfig for Terraform/K8s/CloudFormation). Minimal but real integrations."
    status: pending
  - id: s4-dashboard-full
    content: "Full Web Dashboard: all tabs (Projects, Scans, Findings, Policies, SBOM, Reports, Integrations, Team, Settings). Write operations (triggers, policies, user mgmt). Audit log. SSO (SAML / OIDC) optional."
    status: pending
  - id: s4-ide-plugin
    content: "IDE plugin: Language Server Protocol server in Go; VSCode extension (TS); JetBrains plugin (Kotlin). Inline findings as diagnostics; quick-fix for safe upgrades; panel for AI verdict + reachability."
    status: pending
  - id: s4-multifile-ai-expand
    content: "Multi-file AI context expanded to Python / .NET / Go. Framework-specific context bundles."
    status: pending
  - id: s4-load-hardening
    content: "Load testing, rate limiting, observability: Prometheus metrics, Grafana dashboard, OpenTelemetry tracing, SLO definitions, alerting rules."
    status: pending
isProject: false
---

# Stage 4 -- V1 GA (~2 months / 6-8 weeks)

Parent roadmap: [vulncheck_roadmap.plan.md](vulncheck_roadmap.plan.md). Master plan: [vulncheck_snyk-like_platform_41691b25.plan.md](vulncheck_snyk-like_platform_41691b25.plan.md). Requires Stages 1-3 complete.

## Goal

Turn the Node.js Beta into a GA-ready multi-ecosystem enterprise platform. By end of Stage 4:

- Product covers 4 ecosystems (Node.js, Python, .NET, Go) with feature parity where applicable (SCA, behavioral, call-graph, SAST, AI verdict)
- Container / OCI image scanning available end-to-end
- Enterprise controls: Rego policy engine, VEX consumer, SLSA + Sigstore
- Full web dashboard with write operations and audit logging
- IDE plugins for VSCode + JetBrains deliver the "shift-left" story
- Platform passes a pilot with an enterprise-scale customer (50+ repos, multi-ecosystem)

## Scope IN

### Python Ecosystem (master Phase E3)

- `pkg/ecosystem/python/parser.go` -- `requirements.txt` (with constraints), `Pipfile.lock`, `poetry.lock`, `pyproject.toml`
- `pkg/ecosystem/python/registry_client.go` -- PyPI
- Vuln sources: OSV covers Python well; add PyPI advisory DB + GHSA Python
- `rules/semgrep/python-django.yaml`, `python-flask.yaml`, `python-fastapi.yaml` -- 25+ rules
- Call-graph: subprocess wrapper around jedi + pyan3 (Python 3.10+)
- Behavioral analyzer: pip sandbox install into throwaway venv inside sandbox; AST capabilities via `ast` module

### .NET Ecosystem (master Phase E4)

- `pkg/ecosystem/dotnet/parser.go` -- `packages.config`, `*.csproj` PackageReference, `packages.lock.json`
- `pkg/ecosystem/dotnet/registry_client.go` -- NuGet
- Vuln sources: GHSA .NET + OSV (limited) + Microsoft Security Response Center feed
- `rules/semgrep/dotnet-aspnet.yaml` -- 25+ rules
- Call-graph: subprocess wrapper around a Roslyn-based analyzer (shipped as `tools/vulncheck-roslyn-cg` .NET console app)
- Behavioral: harder for .NET due to compiled assemblies; implement a lighter "metadata-only" profile for GA (file IO + network imports at declaration level); full sandboxed run deferred to V2 polish

### Go Ecosystem (master Phase E5 new)

- `pkg/ecosystem/go/parser.go` -- `go.mod` + `go.sum`, module graph via `go mod graph` fallback
- `pkg/ecosystem/go/registry_client.go` -- proxy.golang.org
- Vuln sources: govulncheck DB + OSV + GHSA Go
- `rules/semgrep/go-gin-gorm.yaml` -- 20+ rules
- Call-graph: reuse Stage 3 self-host engine
- Behavioral: static-only (Go lacks npm-style install hooks); detect `init()` side effects + unsafe usage

### Rego Policies (master Phase Q)

- `internal/policy/rego.go` -- embed OPA via `github.com/open-policy-agent/opa/rego`
- Policy storage: `rego_policy` table with package name + rego source + version
- Evaluation contract: `data.vulncheck.allow` / `data.vulncheck.violations` with structured violation objects
- Stock examples in `policies/rego-examples/`:
  - `block-critical-prod.rego`
  - `require-signed-containers.rego`
  - `license-allowlist.rego`
  - `no-kev-in-prod.rego`
  - `approved-base-images.rego`
- CLI: `vulncheck policy test --rego <file>` for offline eval
- Rego is **additive** -- DB + file policies (Stage 2) still work; Rego is advanced-mode

### VEX Consumer (master Phase Q)

- `internal/vex/parser.go` -- OpenVEX JSON, CSAF 2.0 JSON, CycloneDX VEX embedded
- `internal/vex/feeds.go` -- pull per-vendor VEX feeds on schedule; store in `vex_statements`
- Trust model: per-org trusted publisher list; match by GPG signature or HTTPS + pinned fingerprint
- Auto-dismiss: if statement matches `(package, vuln, version-range)` with status `not_affected` or `fixed`, finding is auto-dismissed with attribution `{source, statement_id, rationale}`
- Manual override still allowed

### SLSA / Sigstore (master Phase Q)

- `internal/slsa/verifier.go` -- parses SLSA v1.0 provenance; validates chain per policy
- `internal/slsa/sigstore.go` -- verifies signatures via Rekor transparency log API; supports keyless (Fulcio) + key-pair; validates cert-identity claims (e.g. `release-please@github-actions.iam.gserviceaccount.com`)
- Works with: PyPI wheels (PEP 740), Go modules (sum db), npm provenance statements, container images (cosign)
- Per-org Rego policy can require provenance for prod deps

### Container / OCI Image Scanning (master Phase M2)

- `internal/container/puller.go` -- pulls images from Docker Hub / ECR / GCR / Artifact Registry / GHCR with OAuth
- `internal/container/layers.go` -- extracts layers; identifies OS (Debian/Ubuntu/Alpine/RHEL/Distroless) by probing `/etc/os-release`
- `internal/container/os_scanner.go` -- queries Trivy DB (synced in Stage 3) for OS pkg vulns
- `internal/container/dockerfile_sast.go` -- parses Dockerfile from build context; runs rules: `USER 0`, `ADD` with URL, `curl | sh`, no healthcheck, exposed privileged port, world-writable chmod
- `internal/container/base_image.go` -- recommends hardened / smaller base images (distroless, chainguard) with delta CVE counts
- `internal/container/cosign.go` -- wraps Sigstore verification for image signatures
- SBOM for images: CycloneDX + SPDX
- CLI: `vulncheck container scan <image-ref>` + API + dashboard UI

### DAST + IaC Scaffold (master Phase M)

- `internal/dast/zap.go` -- wraps OWASP ZAP baseline scan via its docker container / REST API; imports alerts as findings
- `internal/iac/checkov.go` -- wraps Checkov; supports Terraform, CloudFormation, K8s manifests, Helm
- `internal/iac/trivy_config.go` -- wraps `trivy config` for misconfig
- Scaffold level: working but not deep; polish deferred to V2

### Full Web Dashboard (master Phase O)

- React 18 + Vite + TailwindCSS + React Query + Zustand
- Pages / tabs:
  - Auth: Login, Register, Invite acceptance, SSO (SAML 2.0, OIDC optional; SCIM deferred to V2)
  - Projects: list, create, archive, per-project settings (providers, policies, monitor, notifications)
  - Scans: list, detail, re-run, compare
  - Findings: table + filters + detail view with AI verdict card + reachability path + fix plan + behavioral profile + SBOM context
  - Policies: DB policy builder, `.vulncheck` file preview, Rego editor with lint
  - SBOM: per-scan download + VEX statement viewer
  - Reports: on-demand + scheduled
  - Integrations: Git providers, Jira, Slack, webhooks, SSO
  - Team: users, roles (owner/admin/developer/viewer), invites, audit log
  - Settings: org profile, billing stub, API tokens, AI provider, sandbox hosts
- Write ops gated by role; every write logged to `audit_log`
- E2E Playwright tests

### IDE Plugin (master Phase P)

- `internal/ide/lsp/server.go` -- LSP server in Go; speaks JSON-RPC over stdio and/or TCP
- Capabilities: diagnostics (findings as warnings/errors), code actions (quick-fix: upgrade dep, add ignore), hover (AI verdict summary), document symbols (none), commands (trigger scan, open platform)
- VSCode extension in `internal/ide/vscode/`: TS, consumes LSP; UI panel for findings tree + AI verdict cards + reachability viewer
- JetBrains plugin in `internal/ide/jetbrains/`: Kotlin + Gradle; consumes LSP; matching UI via JetBrains tool window API
- Both published to official marketplaces (extension marketplace for VSCode, JetBrains Marketplace)
- Auth: OAuth device code flow; token stored in OS keychain

### Observability + Hardening

- Prometheus `/metrics` endpoint with RED metrics per handler + scanner stage
- Grafana dashboard JSON in `deploy/grafana/`
- OpenTelemetry tracing: propagate context through River jobs
- SLO targets: scan p95 < 2 min for 1K-dep repo; API p95 < 200ms; AI verdict p95 < 30s
- Alerting rules for error rate, queue depth, external API failure
- Rate limiting per-org (token bucket) on API + CLI

## Scope OUT (deferred to Stage 5)

| Feature | Stage |
| --- | --- |
| ASPM Attack-Story Correlation | 5 |
| AI Verdict Feedback Loop + Prompt Quality Metrics | 5 |
| Runtime eBPF Agent | 5 |
| SCIM provisioning | 5 |
| Compliance attestation polish (SOC2 / ISO27001 artifacts) | 5 |
| Private registry polish (Artifactory / Nexus / Azure Artifacts custom) | 5 |

## Milestones

- **Weeks 1-2**: Python ecosystem end-to-end (parser, SAST, call-graph, behavioral, AI context)
- **Weeks 3-4**: .NET + Go ecosystems; multi-ecosystem dedup / dashboard support
- **Weeks 5-6**: Container scanning + Dockerfile SAST + Sigstore + SLSA + VEX consumer + Rego
- **Weeks 7-8**: Full web dashboard (writes + audit + SSO) + IDE plugins + observability + load testing

## Detailed Checklist

Per-ecosystem (repeat for Python, .NET, Go):

- [ ] Parser with unit test fixtures for each lockfile format
- [ ] Registry client with caching
- [ ] Semgrep rules with positive + negative fixtures, `semgrep --test` green
- [ ] Call-graph subprocess bridge with timeout + fallback
- [ ] Behavioral analyzer (best-effort per language constraints)
- [ ] AI repo-context bundle tuned for language's frameworks
- [ ] Integration test: a seeded sample repo produces expected findings

Enterprise controls:

- [ ] Rego engine wired; 5 stock policies; CLI tester
- [ ] VEX feeds configurable with trust pinning; auto-dismiss tested with fixtures
- [ ] SLSA verification integrated into fix PR + scan pipeline
- [ ] Sigstore verification for containers + npm + PyPI + Go modules

Container:

- [ ] Scans work for Docker Hub, ECR, GCR, GHCR
- [ ] Layer extraction handles multi-stage images
- [ ] OS detection + per-distro pkg vuln queries
- [ ] Dockerfile SAST rules with fixtures
- [ ] Base-image recommendation with measurable CVE delta
- [ ] Cosign signature verification happy path

Dashboard:

- [ ] All tabs navigable with real data
- [ ] Every write path has an audit log entry
- [ ] Role-based route guards
- [ ] OIDC + SAML login tested against one real IDP each (Okta + Google Workspace)
- [ ] Mobile-responsive core pages (findings detail, scans list)

IDE:

- [ ] LSP server starts via `vulncheck ide lsp`; speaks JSON-RPC
- [ ] VSCode extension published (even pre-release) with marketplace listing
- [ ] JetBrains plugin published (even EAP)
- [ ] Both connect to platform, show findings inline
- [ ] Quick-fix applies upgrade + writes ignore rule

Hardening:

- [ ] Prometheus metrics exposed
- [ ] Grafana dashboard JSON in repo
- [ ] OTel traces end-to-end for a scan
- [ ] Rate limits: default 100 req/min per user, 1000/min per org
- [ ] k6 load test script: 50 concurrent scans; p95 SLOs met

## Definition of Done

- Stages 1-3 DoD still passes
- All todos above marked complete
- Multi-ecosystem E2E: a seeded mono-repo with Node.js + Python + .NET + Go + Dockerfile + Terraform + Helm produces findings across all scanners
- Enterprise pilot checklist passes (50+ repos, SSO, audit log, Rego policy, VEX auto-dismiss, Jira sync, PR check)
- All IDE marketplace listings live (even if EAP/pre-release)
- User documentation: multi-ecosystem setup, container scanning guide, Rego cookbook, VEX producer/consumer guide, SLSA/Sigstore guide, IDE plugin quickstart
- Compliance mapping doc: OWASP ASVS v5, NIST SSDF, CIS

## Key Risks

- **Roslyn analyzer sandboxing complexity** -- .NET call-graph runs a .NET subprocess; container image bloats; isolate with dedicated image tier.
- **Python behavioral analysis venv storage** -- caps: per-package max 500MB; aggressive cleanup.
- **OPA evaluation cost** -- cache compiled policies per hash; eval on finding batches not per finding.
- **VEX feed trust** -- document trust model explicitly; any auto-dismiss is reversible via UI; audit log entry.
- **Sigstore verification false negatives from network issues** -- treat transient failures as `unverified` not `invalid`; alert but do not block unless policy forces.
- **Container scanning base-image pulls bandwidth** -- local registry mirror option; caching by digest.
- **Dashboard scope creep** -- tab-by-tab shippable; ruthless cut of non-essential write paths (e.g. defer bulk finding triage to V2 unless pilot demands).
- **IDE marketplace review times** -- begin submission early in the stage; iterate on rejections.
- **Multi-ecosystem findings dedup** -- a Go binary containing embedded JS assets is rare but happens; document scope per scan; don't cross-ecosystem dedup automatically.
- **SSO integration depth** -- ship OIDC + SAML 2.0 only; SCIM and advanced just-in-time provisioning in Stage 5.
- **Rule maintenance burden** -- 100+ Semgrep rules total; introduce a per-rule owner + quarterly FP review.
- **Regression across ecosystems** -- add an E2E matrix CI job: node+python+dotnet+go+container every PR; nightly extended matrix.
