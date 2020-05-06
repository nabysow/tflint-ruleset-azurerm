// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermFrontdoorFirewallPolicyInvalidCustomBlockResponseBodyRule checks the pattern is valid
type AzurermFrontdoorFirewallPolicyInvalidCustomBlockResponseBodyRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermFrontdoorFirewallPolicyInvalidCustomBlockResponseBodyRule returns new rule with default attributes
func NewAzurermFrontdoorFirewallPolicyInvalidCustomBlockResponseBodyRule() *AzurermFrontdoorFirewallPolicyInvalidCustomBlockResponseBodyRule {
	return &AzurermFrontdoorFirewallPolicyInvalidCustomBlockResponseBodyRule{
		resourceType:  "azurerm_frontdoor_firewall_policy",
		attributeName: "custom_block_response_body",
		pattern:       regexp.MustCompile(`^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$`),
	}
}

// Name returns the rule name
func (r *AzurermFrontdoorFirewallPolicyInvalidCustomBlockResponseBodyRule) Name() string {
	return "azurerm_frontdoor_firewall_policy_invalid_custom_block_response_body"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermFrontdoorFirewallPolicyInvalidCustomBlockResponseBodyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermFrontdoorFirewallPolicyInvalidCustomBlockResponseBodyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermFrontdoorFirewallPolicyInvalidCustomBlockResponseBodyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermFrontdoorFirewallPolicyInvalidCustomBlockResponseBodyRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$`),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
