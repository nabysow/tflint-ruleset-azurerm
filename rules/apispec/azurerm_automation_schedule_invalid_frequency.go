// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermAutomationScheduleInvalidFrequencyRule checks the pattern is valid
type AzurermAutomationScheduleInvalidFrequencyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermAutomationScheduleInvalidFrequencyRule returns new rule with default attributes
func NewAzurermAutomationScheduleInvalidFrequencyRule() *AzurermAutomationScheduleInvalidFrequencyRule {
	return &AzurermAutomationScheduleInvalidFrequencyRule{
		resourceType:  "azurerm_automation_schedule",
		attributeName: "frequency",
		enum: []string{
			"OneTime",
			"Day",
			"Hour",
			"Week",
			"Month",
			"Minute",
		},
	}
}

// Name returns the rule name
func (r *AzurermAutomationScheduleInvalidFrequencyRule) Name() string {
	return "azurerm_automation_schedule_invalid_frequency"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermAutomationScheduleInvalidFrequencyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermAutomationScheduleInvalidFrequencyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermAutomationScheduleInvalidFrequencyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermAutomationScheduleInvalidFrequencyRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as frequency`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
