// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermMapsAccountInvalidSkuNameRule checks the pattern is valid
type AzurermMapsAccountInvalidSkuNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermMapsAccountInvalidSkuNameRule returns new rule with default attributes
func NewAzurermMapsAccountInvalidSkuNameRule() *AzurermMapsAccountInvalidSkuNameRule {
	return &AzurermMapsAccountInvalidSkuNameRule{
		resourceType:  "azurerm_maps_account",
		attributeName: "sku_name",
		enum: []string{
			"S0",
			"S1",
			"G2",
		},
	}
}

// Name returns the rule name
func (r *AzurermMapsAccountInvalidSkuNameRule) Name() string {
	return "azurerm_maps_account_invalid_sku_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermMapsAccountInvalidSkuNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermMapsAccountInvalidSkuNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermMapsAccountInvalidSkuNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermMapsAccountInvalidSkuNameRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as sku_name`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
