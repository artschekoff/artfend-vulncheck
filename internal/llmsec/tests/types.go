package tests

type Risk string

const (
	RiskCritical Risk = "critical"
	RiskHigh     Risk = "high"
	RiskMedium   Risk = "medium"
	RiskLow      Risk = "low"
	RiskInfo     Risk = "info"
	RiskPass     Risk = "pass"
)

type Result struct {
	Category    string
	Name        string
	Risk        Risk
	Passed      bool
	Description string
	Evidence    string
	Payload     string
	Remediation string
}
