// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermLbRuleInvalidProtocolRule checks the pattern is valid
type AzurermLbRuleInvalidProtocolRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermLbRuleInvalidProtocolRule returns new rule with default attributes
func NewAzurermLbRuleInvalidProtocolRule() *AzurermLbRuleInvalidProtocolRule {
	return &AzurermLbRuleInvalidProtocolRule{
		resourceType:  "azurerm_lb_rule",
		attributeName: "protocol",
		enum: []string{
			"Udp",
			"Tcp",
			"All",
		},
	}
}

// Name returns the rule name
func (r *AzurermLbRuleInvalidProtocolRule) Name() string {
	return "azurerm_lb_rule_invalid_protocol"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermLbRuleInvalidProtocolRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermLbRuleInvalidProtocolRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermLbRuleInvalidProtocolRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermLbRuleInvalidProtocolRule) Check(runner tflint.Runner) error {
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
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as protocol`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
