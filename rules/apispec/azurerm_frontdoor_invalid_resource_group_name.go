// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermFrontdoorInvalidResourceGroupNameRule checks the pattern is valid
type AzurermFrontdoorInvalidResourceGroupNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermFrontdoorInvalidResourceGroupNameRule returns new rule with default attributes
func NewAzurermFrontdoorInvalidResourceGroupNameRule() *AzurermFrontdoorInvalidResourceGroupNameRule {
	return &AzurermFrontdoorInvalidResourceGroupNameRule{
		resourceType:  "azurerm_frontdoor",
		attributeName: "resource_group_name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`),
	}
}

// Name returns the rule name
func (r *AzurermFrontdoorInvalidResourceGroupNameRule) Name() string {
	return "azurerm_frontdoor_invalid_resource_group_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermFrontdoorInvalidResourceGroupNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermFrontdoorInvalidResourceGroupNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermFrontdoorInvalidResourceGroupNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermFrontdoorInvalidResourceGroupNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
