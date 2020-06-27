// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermVirtualWanInvalidOffice365LocalBreakoutCategoryRule checks the pattern is valid
type AzurermVirtualWanInvalidOffice365LocalBreakoutCategoryRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermVirtualWanInvalidOffice365LocalBreakoutCategoryRule returns new rule with default attributes
func NewAzurermVirtualWanInvalidOffice365LocalBreakoutCategoryRule() *AzurermVirtualWanInvalidOffice365LocalBreakoutCategoryRule {
	return &AzurermVirtualWanInvalidOffice365LocalBreakoutCategoryRule{
		resourceType:  "azurerm_virtual_wan",
		attributeName: "office365_local_breakout_category",
		enum: []string{
			"Optimize",
			"OptimizeAndAllow",
			"All",
			"None",
		},
	}
}

// Name returns the rule name
func (r *AzurermVirtualWanInvalidOffice365LocalBreakoutCategoryRule) Name() string {
	return "azurerm_virtual_wan_invalid_office365_local_breakout_category"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermVirtualWanInvalidOffice365LocalBreakoutCategoryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermVirtualWanInvalidOffice365LocalBreakoutCategoryRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermVirtualWanInvalidOffice365LocalBreakoutCategoryRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermVirtualWanInvalidOffice365LocalBreakoutCategoryRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as office365_local_breakout_category`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
