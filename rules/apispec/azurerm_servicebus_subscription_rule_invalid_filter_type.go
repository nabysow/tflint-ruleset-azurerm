// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermServicebusSubscriptionRuleInvalidFilterTypeRule checks the pattern is valid
type AzurermServicebusSubscriptionRuleInvalidFilterTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermServicebusSubscriptionRuleInvalidFilterTypeRule returns new rule with default attributes
func NewAzurermServicebusSubscriptionRuleInvalidFilterTypeRule() *AzurermServicebusSubscriptionRuleInvalidFilterTypeRule {
	return &AzurermServicebusSubscriptionRuleInvalidFilterTypeRule{
		resourceType:  "azurerm_servicebus_subscription_rule",
		attributeName: "filter_type",
		enum: []string{
			"SqlFilter",
			"CorrelationFilter",
		},
	}
}

// Name returns the rule name
func (r *AzurermServicebusSubscriptionRuleInvalidFilterTypeRule) Name() string {
	return "azurerm_servicebus_subscription_rule_invalid_filter_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermServicebusSubscriptionRuleInvalidFilterTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermServicebusSubscriptionRuleInvalidFilterTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermServicebusSubscriptionRuleInvalidFilterTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermServicebusSubscriptionRuleInvalidFilterTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as filter_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
