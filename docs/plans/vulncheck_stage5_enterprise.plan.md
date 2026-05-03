---
name: VulnCheck Stage 5 -- V2 / Enterprise
overview: 3+ months (12+ weeks) extension. Takes the GA V1 into enterprise ASPM territory. Adds ASPM Attack-Story Correlation (graphs plausible kill chains across SCA + SAST + IaC + container + secret findings sharing entities), AI Verdict Feedback Loop with user-feedback UI + prompt quality metrics dashboard + periodic prompt regeneration from aggregated feedback, and the optional eBPF Runtime Reachability Agent for production workloads (observes loaded modules + syscalls, reports back real reachability signals). Polish layer adds SCIM provisioning, private registry depth, compliance artifacts (SOC2 / ISO27001 mappings + evidence exports), and full production hardening (multi-region HA, DR runbook, rate-limit tiers, cost observability). Delivers a credible competitor to Apiiro / OX / Endor Labs in the ASPM space.
todos:
  - id: s5-aspm-correlation
    content: "ASPM Attack-Story Correlation: entity extraction from every finding (file, fn, pkg, host, secret, cve); bipartite graph; community detection -> attack stories; per-story severity + suggested remediation priority."
    status: pending
  - id: s5-ai-feedback
    content: "AI Verdict Feedback Loop: UI + CLI + API for user to thumbs-up/down AI verdict with reason taxonomy; prompt_quality_aggregation_job processes feedback; metrics dashboard; automated prompt regeneration proposals gated by human review."
    status: pending
  - id: s5-runtime-agent
    content: "eBPF Runtime Reachability Agent (opt-in): Go agent + bcc / cilium-ebpf probes; module-load + syscall sampling; secure reporting protocol; control-plane endpoint; runtime reachability signal merged into effective severity."
    status: pending
  - id: s5-scim
    content: "SCIM 2.0 provisioning: /Users and /Groups endpoints; bearer-token auth; just-in-time user deprovisioning; audit log entries."
    status: pending
  - id: s5-private-registry
    content: "Private registry depth: Artifactory, Nexus, Azure Artifacts, CodeArtifact; credential mgmt; network proxy awareness; internal package advisory overlay."
    status: pending
  - id: s5-compliance
    content: "Compliance artifacts: SOC2 Trust Services Criteria mapping; ISO27001 Annex A mapping; NIST SSDF + OWASP ASVS v5 mapping; evidence export job; per-tenant audit report."
    status: pending
  - id: s5-hardening-sre
    content: "Production hardening: multi-region HA (active-passive Postgres + replicated object storage); DR runbook + restore drill; rate-limit tiers; cost observability (per-org OpenAI usage + storage); alert pager integration."
    status: pending
  - id: s5-dast-iac-polish
    content: "DAST + IaC polish: ZAP full scan + auth scripting; Checkov custom rule support; Terraform drift-detection hooks; OSCAL export for IaC findings."
    status: pending
  - id: s5-continuous-learning
    content: "Continuous learning pipelines: advisory feed ingestion SLOs; OSV delta detection; KEV auto-re-triage; AI verdict drift detection on re-runs; customer-submitted false-positive corpus for SAST rule tuning."
    status: pending
isProject: false
---

# Stage 5 -- V2 / Enterprise (3+ months / 12+ weeks)

Parent roadmap: [vulncheck_roadmap.plan.md](vulncheck_roadmap.plan.md). Master plan: [vulncheck_snyk-like_platform_41691b25.plan.md](vulncheck_snyk-like_platform_41691b25.plan.md). Requires Stages 1-4 complete.

## Goal

Turn the multi-ecosystem GA V1 into an enterprise ASPM platform that can win competitive evals against Apiiro / OX / Endor Labs / Snyk Enterprise. By end of Stage 5:

- ASPM Attack-Story Correlation groups individual findings into prioritized kill-chain narratives
- AI verdicts improve automatically via a user-feedback loop with metrics observability
- Optional eBPF runtime agent brings production-context-aware reachability back to the platform (closes the "static says reachable, runtime says dead code" gap)
- Platform is operationally hardened (multi-region HA, DR drill passed, SLO + cost dashboards)
- Compliance artifacts are exportable per tenant
- Enterprise contract-ready

## Scope IN

### ASPM Attack-Story Correlation (master Phase P2)

- `internal/aspm/entities.go` -- canonicalizes entities across findings:
  - File path (normalized to repo-relative)
  - Function (module + name)
  - Package (ecosystem + name + version)
  - Host / container (registry + repo + digest)
  - Secret fingerprint (hashed)
  - CVE / GHSA id
  - Route / endpoint (HTTP method + path pattern)
- `internal/aspm/graph.go` -- builds a bipartite graph: findings -> entities
- `internal/aspm/correlation.go` -- community detection (Louvain) + heuristic rules ("SAST XSS + secret in same file + reachable CVE in response-rendering lib = one story")
- `internal/aspm/story.go` -- story model: id, title (generated), entities, findings, combined severity (max + chain bonus), suggested remediation ordering
- `attack_story` GORM model + persistence
- API: `GET /api/attack-stories`, `GET /api/attack-stories/:id`, `POST /api/attack-stories/:id/resolve`
- CLI: `vulncheck stories list`, `vulncheck stories show <id>`
- Dashboard: "Attack Stories" tab with graph viz (cytoscape.js); per-story timeline; remediation checklist
- River job `aspm_correlation_job.go` runs post-scan + nightly for cross-project stories (within same org)

### AI Verdict Feedback Loop (master Phase new)

- `internal/feedback/ui.go` -- thumbs up / down + reason taxonomy (`false_no_impact`, `missed_path`, `insufficient_context`, `hallucinated_code`, `correct`, `overly_confident`, `underly_confident`)
- Storage: `ai_verdict_feedback` table linked to verdict + user
- Aggregation job `prompt_quality_aggregation_job.go` (weekly):
  - Per-prompt-version metrics: precision, recall proxy (via feedback), avg confidence drift
  - Per-category metrics: which CVE classes have worst verdict quality
  - Surface top 10 problem verdicts for manual review
- Prompt regeneration proposal:
  - Templated improvement suggestions based on aggregated reasons
  - **Human-in-the-loop** -- automatic deployment disabled; engineer reviews proposals in dashboard + merges via PR
- Dashboard: "AI Quality" tab -- per-prompt version, feedback volume, agreement rate, cost / token trend
- Privacy: feedback text scrubbed of code snippets on ingestion; opt-in telemetry per org

### eBPF Runtime Reachability Agent (master Phase R)

- `agent/` Go module (separate binary): `vulncheck-agent`
- Distribution: statically-linked Linux binary + container image + systemd unit + Kubernetes DaemonSet Helm chart
- Auth: mTLS certs provisioned per deployment; agent pulls config from control plane; heartbeat + health
- eBPF probes (via cilium/ebpf or bcc):
  - Module load events (Node.js `require` via libuv, Python import, ELF DT_NEEDED)
  - Syscall sampling (openat, connect, execve) with aggregation
  - HTTP listeners enumerated (from socket map) -> correlate to routes
- Privacy:
  - No payload capture; aggregated telemetry only
  - Per-agent opt-in per-process namespace
  - All sent data signed + encrypted
- Control plane endpoints:
  - `POST /api/agents/heartbeat`
  - `POST /api/agents/modules` (upsert loaded modules)
  - `POST /api/agents/syscalls` (aggregated counters)
- Platform side:
  - `internal/runtime/ingestor.go` -- validates signatures; persists `runtime_module_load` + `runtime_syscall_bucket`
  - `internal/runtime/correlation.go` -- joins runtime module data with package advisories: package version loaded in production -> prioritize findings for that (pkg,version) pair
  - Effective severity Stage 4 formula extended: `runtime_adj` adds +1 level if module is runtime-loaded and vulnerable, irrespective of static-reachability guess
- Dashboard: "Runtime" tab -- fleet of agents, production-loaded modules, runtime-confirmed vulnerabilities
- Strict opt-in per org; documented threat model

### SCIM 2.0 Provisioning

- `/scim/v2/Users` CRUD
- `/scim/v2/Groups` CRUD
- Bearer-token auth with scopes
- Supported attrs: userName, emails, active, name.familyName/name.givenName, groups
- Deactivating in IDP -> just-in-time token revocation in platform
- Audit log entries for every SCIM op

### Private Registry Depth

- Extended registry clients: Artifactory (npm/pypi/nuget/maven endpoints), Nexus Repository (same), Azure Artifacts, AWS CodeArtifact
- Credential rotation + scoped tokens
- Network proxy awareness (HTTP_PROXY/NO_PROXY handling in all fetchers)
- Internal advisory overlay: orgs can publish their own CVEs for private packages; scanner merges as additional source

### Compliance Artifacts

- Mapping tables stored in repo:
  - `compliance/soc2.yaml` -- which scanner capability satisfies which TSC
  - `compliance/iso27001.yaml`
  - `compliance/nist-ssdf.yaml`
  - `compliance/owasp-asvs.yaml`
- Evidence export job: per tenant, per period, bundles reports + scan summaries + audit log + policy evaluations into a signed ZIP
- Dashboard "Compliance" page with per-framework coverage map

### Hardening / SRE

- Multi-region HA:
  - Postgres: streaming replication with managed failover (RDS Multi-AZ or similar); documented failover procedure
  - Object storage (SBOMs, reports, sandboxed tarballs): multi-region replication
  - Control plane stateless; auto-scale group
- DR runbook documented + periodic restore drill (quarterly)
- Rate-limit tiers:
  - Free: 10 scans/day, 10 AI verdicts/scan
  - Pro: 1000 scans/day, 500 AI verdicts/scan
  - Enterprise: negotiated
- Cost observability:
  - Per-org OpenAI / Anthropic spend
  - Storage spend
  - Scan compute spend
- Alert pager: PagerDuty / Opsgenie integration for SLO breaches

### DAST + IaC Polish

- OWASP ZAP full scan profile (authenticated via scripted login)
- ZAP add-ons curated list
- Checkov custom rule support (users can upload rules)
- IaC drift detection hooks (integration with Terraform Cloud / Atlantis)
- OSCAL 1.1 export for IaC + compliance findings

### Continuous Learning Pipelines

- Advisory ingestion SLOs: new CVE in KEV -> platform re-triage within 1h
- OSV delta detection: when an advisory changes `affected.ranges`, re-run impacted findings
- KEV auto-re-triage: severity boost to all findings where CVE entered KEV
- AI verdict drift detection: re-run sampled verdicts on new scans; compare; flag drift > threshold for human review
- Customer false-positive corpus: opt-in submission from ignores; SAST rule owners review quarterly; feeds back into test fixtures

## Scope OUT

Anything not listed above is out of scope for V2. Additional ideas parked in master plan "Future Work" section.

## Milestones

- **Weeks 1-3**: ASPM correlation end-to-end (entity extraction across existing finding types, graph, community detection, dashboard viz, stories API + CLI)
- **Weeks 4-6**: AI feedback loop + prompt quality dashboard + aggregation job + scrubbing + governance workflow
- **Weeks 7-10**: eBPF runtime agent end-to-end (agent binary + probes + signing + control plane ingestion + runtime correlation + dashboard)
- **Weeks 11-12**: SCIM + private registry depth + compliance artifacts + hardening + SRE polish + DAST/IaC polish + continuous learning wiring

## Detailed Checklist

ASPM:
- [ ] Entity extractor per finding type has unit tests
- [ ] Louvain community detection library vetted + wrapped
- [ ] Story generation includes human-readable title via LLM with fallback template
- [ ] Dashboard graph viz performant up to 500 stories
- [ ] API pagination + filters (severity, ecosystem, project)
- [ ] False-positive feedback loop for stories (user marks "not a real story")

AI feedback:
- [ ] UI: thumbs up/down + reason dropdown + optional note
- [ ] API: `POST /api/verdicts/:id/feedback`
- [ ] Aggregation job scheduled weekly; metrics persisted + charted
- [ ] Proposal generator emits markdown "prompt diff suggestion" documents
- [ ] Governance: prompt changes go through PR + regression evaluation set

eBPF agent:
- [ ] Agent builds for x86_64 + aarch64
- [ ] Container + Helm chart + systemd unit
- [ ] Signed binary releases
- [ ] mTLS cert rotation documented
- [ ] Load test: agent reports for 1K pods without >2% CPU overhead
- [ ] Correlation logic merges runtime data into effective severity
- [ ] Dashboard "Runtime" tab usable
- [ ] Threat model published in `docs/runtime-agent-threat-model.md`

SCIM:
- [ ] Conformance tests against known SCIM clients (Okta, Azure AD)
- [ ] Tokens scoped + rotatable
- [ ] Audit log

Private registry:
- [ ] Each registry type has a happy-path integration test
- [ ] Cred rotation UX
- [ ] Proxy awareness tested

Compliance:
- [ ] Each mapping file peer-reviewed by a security engineer
- [ ] Evidence export bundle validates against schema
- [ ] Compliance dashboard shows per-framework % coverage

Hardening:
- [ ] DR drill executed; report in `docs/dr-drill-YYYYMMDD.md`
- [ ] Multi-region config documented in `docs/deployment-multi-region.md`
- [ ] Rate-limit tiers enforced + tested
- [ ] Cost dashboard live per org
- [ ] Pager integration tested

DAST + IaC:
- [ ] ZAP full scan with auth works on a test app
- [ ] Custom Checkov rules uploadable + runnable
- [ ] OSCAL export validates

Continuous learning:
- [ ] KEV re-triage SLO observable
- [ ] OSV delta job tested
- [ ] Drift detection surfaces suspicious verdict changes
- [ ] Corpus submission UI + review queue

## Definition of Done

- Stages 1-4 DoD still passes
- All todos above marked complete
- Enterprise eval checklist (used during sales process) passes end-to-end on a multi-ecosystem multi-region deployment
- Published docs: ASPM guide, AI feedback handbook, runtime agent deployment guide, compliance export guide, DR runbook, multi-region deployment guide
- At least one enterprise customer in production using runtime agent in opt-in mode
- Cost model documented; AI spend predictable per tier

## Key Risks

- **ASPM false-positive stories** -- over-correlating creates confusion. Ship with conservative heuristics + clear "merge / split" UX so users can correct; use corrections as training signal.
- **AI feedback data privacy** -- feedback text may leak code. Scrub aggressively; redact pre-storage; per-org opt-in.
- **Prompt churn** -- auto-regenerated prompts can silently regress verdict quality. Mandatory regression eval set + human PR review.
- **eBPF agent trust + privilege** -- running privileged code in customer production. Strong threat model doc; signed binaries; minimal privilege; opt-in per workload; kill-switch.
- **Kernel compatibility matrix** -- eBPF varies across kernels. Document supported kernels (5.10+); degrade gracefully.
- **Runtime data volume** -- module-load + syscall sampling can be noisy. Client-side aggregation; sample rates configurable; server-side rate limiting.
- **SCIM provider quirks** -- Okta + Azure AD have subtly different SCIM interpretations. Conformance test against both + matrix docs.
- **Private registry credential sprawl** -- one secret per registry per project; rotate reminders.
- **Compliance claims accuracy** -- mapping tables must be reviewed by qualified auditors; disclaim scope in legal review; never auto-generate attestation letters.
- **Multi-region latency** -- cross-region writes expensive. Keep single primary region per tenant; replicas for read + DR only.
- **DR drill risk** -- restore drills are risky. Schedule in maintenance window; rehearse in staging first.
- **Cost of long-running enterprise pilots** -- AI spend can balloon. Strict quota + alerting on outliers.
- **Scope creep** -- V2 is the longest stage. Treat the listed Scope IN as a contract; park new ideas in a "V3 candidate" list without negotiating them into V2.
