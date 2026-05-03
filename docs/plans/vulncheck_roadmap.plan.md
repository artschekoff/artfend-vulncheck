---
name: VulnCheck Delivery Roadmap
overview: Top-level staging roadmap for ArtFend VulnCheck. Splits the full master plan ([vulncheck_snyk-like_platform_41691b25.plan.md](vulncheck_snyk-like_platform_41691b25.plan.md)) into 5 progressive, independently-shippable stages -- from a 1-week MVP that beats npm audit on noise, to a full enterprise ASPM platform with ASPM correlation, AI feedback loop, and eBPF runtime agent. Each stage ends with a demoable / sellable artifact and defers the next wave of features cleanly.
todos:
  - id: stage-1-mvp
    content: "Stage 1 -- MVP (5-7 days): Working CLI + npm lockfile SCA with scope awareness. Beats npm audit on noise. See [vulncheck_stage1_mvp.plan.md](vulncheck_stage1_mvp.plan.md)"
    status: pending
  - id: stage-2-alpha
    content: "Stage 2 -- Alpha (2 weeks): + Git integration (GitHub) + AI Impact Analyzer (OpenAI, Node.js) + Remediation + Dual Policies + supply chain + NVD/GHSA + minimal SAST. See [vulncheck_stage2_alpha.plan.md](vulncheck_stage2_alpha.plan.md)"
    status: pending
  - id: stage-3-beta
    content: "Stage 3 -- Beta (~1 month): Full Node.js product -- behavioral analysis + call-graph + License/SBOM/Secrets + Fix PRs + Jira + Monitor + Notifications + PR baseline diff + Reports + all 6 vuln sources + minimal dashboard. See [vulncheck_stage3_beta.plan.md](vulncheck_stage3_beta.plan.md)"
    status: pending
  - id: stage-4-v1
    content: "Stage 4 -- V1 (~2 months): Multi-ecosystem (Python/.NET/Go) + Rego + VEX consumer + SLSA/Sigstore + Container scanning + full Web Dashboard + IDE plugin (VSCode+JetBrains). See [vulncheck_stage4_v1.plan.md](vulncheck_stage4_v1.plan.md)"
    status: pending
  - id: stage-5-enterprise
    content: "Stage 5 -- V2 / Enterprise (3+ months): ASPM attack-story correlation + AI verdict feedback loop + runtime eBPF agent (opt-in) + hardening / SLO / compliance polish. See [vulncheck_stage5_enterprise.plan.md](vulncheck_stage5_enterprise.plan.md)"
    status: pending
isProject: false
---

# VulnCheck Delivery Roadmap

The full master plan in [vulncheck_snyk-like_platform_41691b25.plan.md](vulncheck_snyk-like_platform_41691b25.plan.md) describes the finished product (28+ phases). This roadmap splits that into **5 staged deliveries**, each shippable and demoable on its own.

## Staging Principles

- **Every stage ends with a demoable artifact.** No stage leaves the product in a half-working state.
- **Later stages add without removing.** Stage N+1 extends Stage N; no rewrites.
- **Differentiators land early.** The killer feature (AI verdict + scope-aware severity) is available in Stage 2, so the product has a story from month 1.
- **Ecosystem width comes late.** Node.js is the only ecosystem until Stage 4. Nailing one ecosystem deeply is more valuable than superficial coverage of many.
- **Enterprise controls come after product-market fit.** Rego, VEX, SLSA, ASPM ship in Stages 4-5 because they matter for enterprise deals, not solo developers.

## Timeline

```mermaid
gantt
  title VulnCheck Delivery Stages
  dateFormat  YYYY-MM-DD
  section Stage 1 -- MVP
    Foundation + SCA core (Node.js): s1, 2026-04-22, 7d
  section Stage 2 -- Alpha
    Git + AI + Remediation + Policies: s2, after s1, 14d
  section Stage 3 -- Beta
    Full Node.js product: s3, after s2, 35d
  section Stage 4 -- V1
    Multi-ecosystem + Container + IDE: s4, after s3, 56d
  section Stage 5 -- Enterprise
    ASPM + Feedback + Runtime: s5, after s4, 84d
```

Cumulative effort (solo senior dev):
- End of Stage 1: ~1 week
- End of Stage 2: ~3 weeks
- End of Stage 3: ~8 weeks (~2 months)
- End of Stage 4: ~16 weeks (~4 months)
- End of Stage 5: ~28 weeks (~6-7 months)

With a team of 3: roughly divide by 2 (with coordination overhead). With a team of 6: divide by 3-4.

## What Each Stage Delivers (Elevator Pitch)

| Stage | Duration | Pitch |
| --- | --- | --- |
| **1 -- MVP** | 1 week | "Scan a Node.js lockfile, get 90% fewer false positives than npm audit thanks to dependency-scope awareness." |
| **2 -- Alpha** | 2 weeks | "Point us at your GitHub repo. We tell you which CVEs actually matter to YOUR code using AI. Safe fix suggestions. No alert dump." |
| **3 -- Beta** | 1 month | "Full Node.js AppSec: behavioral malware defense, call-graph reachability, license + SBOM + secrets + auto Fix PRs + Jira sync + monitoring. Customer-ready." |
| **4 -- V1** | 2 months | "All major ecosystems (Node.js + Python + .NET + Go), container scanning, enterprise policies (Rego + VEX + SLSA), full dashboard, VSCode + JetBrains plugins. GA-ready." |
| **5 -- V2** | 3+ months | "Enterprise ASPM: attack-story correlation, self-improving AI via feedback loop, optional eBPF runtime agent. Competitive with Apiiro / OX / Endor Labs." |

## Cross-Cutting Tracks (run in parallel to all stages)

Some work streams are not phase-specific and should run continuously once started:

- **Security self-audit (dogfooding)**: from Stage 1, enable `govulncheck` + `gosec` + `trivy fs` in CI on the platform's own codebase. Expand as new scanners come online.
- **Testing strategy**: unit tests per component from Stage 1; integration tests per phase; end-to-end smoke tests from Stage 2; load tests from Stage 4; fuzz tests for parsers from Stage 3.
- **Documentation**: README + API reference from Stage 1; user guides from Stage 3; enterprise deployment guide + compliance mapping from Stage 4.
- **Observability**: structured logs from Stage 1; Prometheus metrics from Stage 2; distributed tracing from Stage 3; SLO dashboards from Stage 4.
- **Rule maintenance**: SAST rules ship with positive + negative test cases from Stage 2 onwards; CI job runs `semgrep --test` on every PR; track per-rule FP rate from Stage 3.

## Risk-Based Ordering Rationale

Why this order and not another? Key decisions:

1. **AI in Stage 2, not Stage 1.** AI is the product differentiator but adding it to MVP blurs the core thesis ("scope awareness beats raw CVSS"). We want Stage 1 to prove the simpler thesis first.
2. **Behavioral analysis in Stage 3, not earlier.** It requires mature package tarball fetching + sandboxing infrastructure. Rushing it risks a security incident (sandboxed code escaping). Stage 3 gives us time to do it right.
3. **Call-graph before AI would be purist but wrong.** Call-graph is deterministic and cheaper but complex to build per-language. AI works day one with the LLM API. Ship AI first (Stage 2) for the demo, add call-graph later (Stage 3) as a cheaper replacement layer.
4. **Multi-ecosystem in Stage 4, not Stage 3.** Adding Python / .NET / Go too early dilutes the Node.js story. Nail Node.js first. Every additional ecosystem in Stage 4 gets to reuse mature core infrastructure.
5. **Container scanning in Stage 4 with the other ecosystem work.** It's a large scope increase, but lives in its own subsystem -- good fit with the ecosystem-expansion stage.
6. **ASPM last.** Correlation only works if you have good signals first. Stages 1-4 build the signals; Stage 5 joins them.

## How to Use These Plans

- Each stage plan has its own frontmatter with **todos** that track execution.
- File references use markdown links to the master plan sections. Example: [Phase E in master](vulncheck_snyk-like_platform_41691b25.plan.md) describes the Node.js SCA scanner in full; Stage 1 implements a minimal slice of it.
- When starting a stage, read that stage's plan file first, then jump to the referenced master-plan sections for depth.
- When a stage is done, the "Definition of Done" checklist must pass. Only then advance to the next stage.
- If a feature appears in multiple stages (e.g. AI runs Stage 2 -> 3 -> 4), each stage adds a specific slice defined in the stage plan.

## Quick Decision Tree

"I want to ship as fast as possible, what do I do?"
- Read [vulncheck_stage1_mvp.plan.md](vulncheck_stage1_mvp.plan.md). Execute. Ship in a week.

"I want the AI differentiator for a demo in 3 weeks."
- Stages 1 + 2. Execute both. Demo after ~3 weeks.

"I want customer-ready Node.js product."
- Stages 1 + 2 + 3. ~8 weeks.

"I want GA-ready multi-ecosystem product."
- Stages 1 + 2 + 3 + 4. ~16 weeks.

"I want enterprise ASPM competitor."
- All 5 stages. ~28 weeks.
