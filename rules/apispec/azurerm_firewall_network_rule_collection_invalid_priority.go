// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermFirewallNetworkRuleCollectionInvalidPriorityRule checks the pattern is valid
type AzurermFirewallNetworkRuleCollectionInvalidPriorityRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAzurermFirewallNetworkRuleCollectionInvalidPriorityRule returns new rule with default attributes
func NewAzurermFirewallNetworkRuleCollectionInvalidPriorityRule() *AzurermFirewallNetworkRuleCollectionInvalidPriorityRule {
	return &AzurermFirewallNetworkRuleCollectionInvalidPriorityRule{
		resourceType:  "azurerm_firewall_network_rule_collection",
		attributeName: "priority",
		max:           65000,
		min:           100,
	}
}

// Name returns the rule name
func (r *AzurermFirewallNetworkRuleCollectionInvalidPriorityRule) Name() string {
	return "azurerm_firewall_network_rule_collection_invalid_priority"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermFirewallNetworkRuleCollectionInvalidPriorityRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermFirewallNetworkRuleCollectionInvalidPriorityRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermFirewallNetworkRuleCollectionInvalidPriorityRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermFirewallNetworkRuleCollectionInvalidPriorityRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val int
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if val > r.max {
				runner.EmitIssueOnExpr(
					r,
					"priority must be 65000 or less",
					attribute.Expr,
				)
			}
			if val < r.min {
				runner.EmitIssueOnExpr(
					r,
					"priority must be 100 or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
