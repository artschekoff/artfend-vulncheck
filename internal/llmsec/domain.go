package llmsec

type Risk string

const (
	RiskCritical Risk = "critical"
	RiskHigh     Risk = "high"
	RiskMedium   Risk = "medium"
	RiskLow      Risk = "low"
	RiskInfo     Risk = "info"
	RiskPass     Risk = "pass"
)

type TestResult struct {
	Category    string
	Name        string
	Risk        Risk
	Passed      bool
	Description string
	Evidence    string
	Payload     string
	Remediation string
}

type RunnerConfig struct {
	Endpoint string
	APIKey   string
	Model    string
	Timeout  int
}
