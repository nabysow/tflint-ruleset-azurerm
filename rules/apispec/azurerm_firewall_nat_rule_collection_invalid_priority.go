// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermFirewallNatRuleCollectionInvalidPriorityRule checks the pattern is valid
type AzurermFirewallNatRuleCollectionInvalidPriorityRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAzurermFirewallNatRuleCollectionInvalidPriorityRule returns new rule with default attributes
func NewAzurermFirewallNatRuleCollectionInvalidPriorityRule() *AzurermFirewallNatRuleCollectionInvalidPriorityRule {
	return &AzurermFirewallNatRuleCollectionInvalidPriorityRule{
		resourceType:  "azurerm_firewall_nat_rule_collection",
		attributeName: "priority",
		max:           65000,
		min:           100,
	}
}

// Name returns the rule name
func (r *AzurermFirewallNatRuleCollectionInvalidPriorityRule) Name() string {
	return "azurerm_firewall_nat_rule_collection_invalid_priority"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermFirewallNatRuleCollectionInvalidPriorityRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermFirewallNatRuleCollectionInvalidPriorityRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermFirewallNatRuleCollectionInvalidPriorityRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermFirewallNatRuleCollectionInvalidPriorityRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val int
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if val > r.max {
				runner.EmitIssue(
					r,
					"priority must be 65000 or less",
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			if val < r.min {
				runner.EmitIssue(
					r,
					"priority must be 100 or higher",
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
