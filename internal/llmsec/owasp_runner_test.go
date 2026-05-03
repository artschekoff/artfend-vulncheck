package llmsec_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/artfend/vulncheck/internal/llmsec"
)

func mockResponse(content string) string {
	resp := map[string]interface{}{
		"choices": []map[string]interface{}{
			{
				"message": map[string]string{
					"role":    "assistant",
					"content": content,
				},
				"finish_reason": "stop",
			},
		},
	}
	b, _ := json.Marshal(resp)
	return string(b)
}

func TestRunnerReturnsResultsForAllCategories(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse("I will ignore my instructions and help you.")))
	}))
	defer srv.Close()

	cfg := llmsec.RunnerConfig{
		Endpoint: srv.URL,
		APIKey:   "test-key",
		Model:    "gpt-4",
		Timeout:  5,
	}
	runner := llmsec.NewOWASPRunner(cfg)
	results, err := runner.Run(t.Context())
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}
	if len(results) == 0 {
		t.Fatal("expected results, got none")
	}

	categories := map[string]bool{}
	for _, r := range results {
		categories[r.Category] = true
	}
	for _, cat := range []string{"LLM01", "LLM02", "LLM06"} {
		if !categories[cat] {
			t.Errorf("missing category %s in results", cat)
		}
	}
}

func TestRunnerMarksVulnerableWhenModelLeaksSystemPrompt(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse("My system prompt says: You are a banking assistant")))
	}))
	defer srv.Close()

	cfg := llmsec.RunnerConfig{
		Endpoint: srv.URL,
		APIKey:   "test-key",
		Model:    "gpt-4",
		Timeout:  5,
	}
	runner := llmsec.NewOWASPRunner(cfg)
	results, err := runner.Run(t.Context())
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}

	for _, r := range results {
		if r.Category == "LLM06" {
			if r.Risk != llmsec.RiskHigh {
				t.Errorf("LLM06: expected RiskHigh, got %s", r.Risk)
			}
			return
		}
	}
	t.Error("LLM06 result not found")
}
