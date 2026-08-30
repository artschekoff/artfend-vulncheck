![artfend vulncheck](images/hero.jpg)

# artfend-vulncheck

**A Snyk-like security scanner focused on AI / LLM applications.**

[![Go](https://img.shields.io/badge/Go-1.26%2B-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![OWASP LLM Top 10](https://img.shields.io/badge/OWASP-LLM_Top_10-000000)](https://owasp.org/www-project-top-10-for-large-language-model-applications/)

vulncheck scans codebases and dependencies for the risks that matter in LLM-powered apps: unsafe prompt handling, missing guardrails, unsanitized RAG inputs, insecure output handling, and risky AI packages. Findings are normalized and stored so you can track them over time.

## What it does

- **SAST for LLM code** — YAML rule packs detect patterns like prompt injection, missing guardrails, and unsanitized RAG inputs (`internal/sast/rules/llm`).
- **OWASP LLM Top-10 runner** — executable checks for LLM01 (prompt injection), LLM02 (insecure output handling), LLM06 (sensitive-information disclosure), and more (`internal/llmsec`).
- **AI-dependency scanning** — parses Python `requirements.txt` / `pyproject.toml` and flags AI/LLM packages via an ecosystem registry (`pkg/ecosystem`).
- **Findings storage** — PostgreSQL migrations persist findings for history and reporting (`migrations/`).

## Project layout

```
internal/
  llmsec/     OWASP LLM Top-10 runner + test suites (LLM01/02/06…)
  sast/       rule engine + YAML rules (prompt_injection, missing_guardrail, rag_unsanitized)
pkg/
  ecosystem/  dependency parsing (Python) + AI-package registry
migrations/   PostgreSQL schema for findings
docs/plans/   roadmap: MVP → alpha → beta → v1 → enterprise
```

## Roadmap

The staged roadmap toward a full Snyk-like platform lives in [`docs/plans`](docs/plans).

**Tech:** Go 1.26.
