// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermCdnProfileInvalidResourceGroupNameRule checks the pattern is valid
type AzurermCdnProfileInvalidResourceGroupNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermCdnProfileInvalidResourceGroupNameRule returns new rule with default attributes
func NewAzurermCdnProfileInvalidResourceGroupNameRule() *AzurermCdnProfileInvalidResourceGroupNameRule {
	return &AzurermCdnProfileInvalidResourceGroupNameRule{
		resourceType:  "azurerm_cdn_profile",
		attributeName: "resource_group_name",
		pattern:       regexp.MustCompile(`^[-\w\._\(\)]+$`),
	}
}

// Name returns the rule name
func (r *AzurermCdnProfileInvalidResourceGroupNameRule) Name() string {
	return "azurerm_cdn_profile_invalid_resource_group_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermCdnProfileInvalidResourceGroupNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermCdnProfileInvalidResourceGroupNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermCdnProfileInvalidResourceGroupNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermCdnProfileInvalidResourceGroupNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[-\w\._\(\)]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
