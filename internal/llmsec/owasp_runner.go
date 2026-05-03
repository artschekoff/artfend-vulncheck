package llmsec

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/artfend/vulncheck/internal/llmsec/tests"
)

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatRequest struct {
	Model    string        `json:"model"`
	Messages []chatMessage `json:"messages"`
}

type chatChoice struct {
	Message      chatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

type chatResponse struct {
	Choices []chatChoice `json:"choices"`
}

type OWASPRunner struct {
	cfg    RunnerConfig
	client *http.Client
}

func NewOWASPRunner(cfg RunnerConfig) *OWASPRunner {
	if cfg.Timeout == 0 {
		cfg.Timeout = 30
	}
	return &OWASPRunner{
		cfg: cfg,
		client: &http.Client{
			Timeout: time.Duration(cfg.Timeout) * time.Second,
		},
	}
}

func (r *OWASPRunner) callLLM(ctx context.Context, payload string) (string, error) {
	body := chatRequest{
		Model: r.cfg.Model,
		Messages: []chatMessage{
			{Role: "user", Content: payload},
		},
	}
	b, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.cfg.Endpoint, bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.cfg.APIKey)

	resp, err := r.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var cr chatResponse
	if err := json.NewDecoder(resp.Body).Decode(&cr); err != nil {
		return "", err
	}
	if len(cr.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}
	return cr.Choices[0].Message.Content, nil
}

func convertResult(r tests.Result) TestResult {
	return TestResult{
		Category:    r.Category,
		Name:        r.Name,
		Risk:        Risk(r.Risk),
		Passed:      r.Passed,
		Description: r.Description,
		Evidence:    r.Evidence,
		Payload:     r.Payload,
		Remediation: r.Remediation,
	}
}

func (r *OWASPRunner) Run(ctx context.Context) ([]TestResult, error) {
	callFn := func(ctx context.Context, payload string) (string, error) {
		return r.callLLM(ctx, payload)
	}

	t1 := &tests.PromptInjectionTest{CallLLM: callFn}
	t2 := &tests.InsecureOutputTest{CallLLM: callFn}
	t6 := &tests.SensitiveInfoTest{CallLLM: callFn}

	results := []TestResult{
		convertResult(t1.Run(ctx)),
		convertResult(t2.Run(ctx)),
		convertResult(t6.Run(ctx)),
	}
	return results, nil
}
