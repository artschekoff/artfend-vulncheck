package sast_test

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/artfend/vulncheck/internal/sast"
)

func TestLLMRulesLoad(t *testing.T) {
	rules, err := sast.LoadLLMRules()
	if err != nil {
		t.Fatalf("LoadLLMRules: %v", err)
	}
	if len(rules) < 3 {
		t.Fatalf("expected >= 3 rules, got %d", len(rules))
	}

	want := map[string]bool{
		"llm-prompt-injection":  false,
		"llm-missing-guardrail": false,
		"llm-rag-unsanitized":   false,
	}
	for _, r := range rules {
		if _, ok := want[r.ID]; ok {
			want[r.ID] = true
		}
	}
	for id, found := range want {
		if !found {
			t.Errorf("rule %q not found", id)
		}
	}
}

func semgrepAvailable() bool {
	_, err := exec.LookPath("semgrep")
	return err == nil
}

func testdataPath(name string) string {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	return filepath.Join(dir, "rules", "llm", "testdata", name)
}

func TestPromptInjectionRuleMatchesBadCode(t *testing.T) {
	if !semgrepAvailable() {
		t.Skip("semgrep not found")
	}
	findings, err := sast.RunLLMRulesOnFile(testdataPath("bad_prompt.py"))
	if err != nil {
		t.Fatalf("RunLLMRulesOnFile: %v", err)
	}
	if len(findings) == 0 {
		t.Fatal("expected findings on bad_prompt.py, got none")
	}
}

func TestPromptInjectionRulePassesGoodCode(t *testing.T) {
	if !semgrepAvailable() {
		t.Skip("semgrep not found")
	}
	findings, err := sast.RunLLMRulesOnFile(testdataPath("good_prompt.py"))
	if err != nil {
		t.Fatalf("RunLLMRulesOnFile: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings on good_prompt.py, got %d: %+v", len(findings), findings)
	}
}
