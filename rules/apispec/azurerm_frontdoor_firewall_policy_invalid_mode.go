// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermFrontdoorFirewallPolicyInvalidModeRule checks the pattern is valid
type AzurermFrontdoorFirewallPolicyInvalidModeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermFrontdoorFirewallPolicyInvalidModeRule returns new rule with default attributes
func NewAzurermFrontdoorFirewallPolicyInvalidModeRule() *AzurermFrontdoorFirewallPolicyInvalidModeRule {
	return &AzurermFrontdoorFirewallPolicyInvalidModeRule{
		resourceType:  "azurerm_frontdoor_firewall_policy",
		attributeName: "mode",
		enum: []string{
			"Prevention",
			"Detection",
		},
	}
}

// Name returns the rule name
func (r *AzurermFrontdoorFirewallPolicyInvalidModeRule) Name() string {
	return "azurerm_frontdoor_firewall_policy_invalid_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermFrontdoorFirewallPolicyInvalidModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermFrontdoorFirewallPolicyInvalidModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermFrontdoorFirewallPolicyInvalidModeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermFrontdoorFirewallPolicyInvalidModeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as mode`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
