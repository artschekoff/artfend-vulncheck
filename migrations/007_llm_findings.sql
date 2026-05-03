CREATE TABLE llm_test_results (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id  UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    scan_id     UUID REFERENCES scans(id) ON DELETE SET NULL,
    category    TEXT NOT NULL,
    name        TEXT NOT NULL,
    risk        TEXT NOT NULL,
    passed      BOOLEAN NOT NULL,
    description TEXT,
    evidence    TEXT,
    payload     TEXT,
    remediation TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_llm_test_results_project ON llm_test_results(project_id);
CREATE INDEX idx_llm_test_results_risk ON llm_test_results(risk) WHERE NOT passed;

CREATE TABLE model_inventory (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id  UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    provider    TEXT NOT NULL,
    model_name  TEXT NOT NULL,
    model_id    TEXT,
    detected_in TEXT[],
    known_vuln  BOOLEAN NOT NULL DEFAULT false,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE(project_id, model_name, provider)
);
