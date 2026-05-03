package sast

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed rules/llm/prompt_injection.yml
var promptInjectionYML []byte

//go:embed rules/llm/missing_guardrail.yml
var missingGuardrailYML []byte

//go:embed rules/llm/rag_unsanitized.yml
var ragUnsanitizedYML []byte

var llmRuleFiles = map[string][]byte{
	"prompt_injection.yml":  promptInjectionYML,
	"missing_guardrail.yml": missingGuardrailYML,
	"rag_unsanitized.yml":   ragUnsanitizedYML,
}

type Rule struct {
	ID       string
	Filename string
}

type ruleFile struct {
	Rules []struct {
		ID string `yaml:"id"`
	} `yaml:"rules"`
}

func LoadLLMRules() ([]Rule, error) {
	var rules []Rule
	for name, data := range llmRuleFiles {
		var rf ruleFile
		if err := yaml.Unmarshal(data, &rf); err != nil {
			return nil, fmt.Errorf("parse %s: %w", name, err)
		}
		for _, r := range rf.Rules {
			rules = append(rules, Rule{ID: r.ID, Filename: name})
		}
	}
	return rules, nil
}

type SemgrepFinding struct {
	RuleID  string
	Path    string
	Line    int
	Message string
}

type semgrepOutput struct {
	Results []struct {
		CheckID string `json:"check_id"`
		Path    string `json:"path"`
		Start   struct {
			Line int `json:"line"`
		} `json:"start"`
		Extra struct {
			Message string `json:"message"`
		} `json:"extra"`
	} `json:"results"`
}

func RunLLMRulesOnFile(path string) ([]SemgrepFinding, error) {
	tmpDir, err := os.MkdirTemp("", "vulncheck-rules-*")
	if err != nil {
		return nil, fmt.Errorf("mkdirtemp: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	for name, data := range llmRuleFiles {
		dest := filepath.Join(tmpDir, name)
		if err := os.WriteFile(dest, data, 0600); err != nil {
			return nil, fmt.Errorf("write %s: %w", name, err)
		}
	}

	semgrepBin, err := exec.LookPath("semgrep")
	if err != nil {
		return nil, fmt.Errorf("semgrep not found: %w", err)
	}

	cmd := exec.Command(semgrepBin, "--config", tmpDir, "--json", path)
	out, runErr := cmd.Output()

	if runErr != nil {
		if exitErr, ok := runErr.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 1 {
				// findings present — parse below
			} else if exitErr.ExitCode() >= 2 {
				return nil, fmt.Errorf("semgrep error (exit %d): %s", exitErr.ExitCode(), string(exitErr.Stderr))
			}
		} else {
			return nil, fmt.Errorf("semgrep exec: %w", runErr)
		}
	}

	var result semgrepOutput
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, fmt.Errorf("parse semgrep JSON: %w", err)
	}

	findings := make([]SemgrepFinding, 0, len(result.Results))
	for _, r := range result.Results {
		findings = append(findings, SemgrepFinding{
			RuleID:  r.CheckID,
			Path:    r.Path,
			Line:    r.Start.Line,
			Message: r.Extra.Message,
		})
	}
	return findings, nil
}

