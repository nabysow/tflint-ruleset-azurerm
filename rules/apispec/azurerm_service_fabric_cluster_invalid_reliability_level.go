// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermServiceFabricClusterInvalidReliabilityLevelRule checks the pattern is valid
type AzurermServiceFabricClusterInvalidReliabilityLevelRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermServiceFabricClusterInvalidReliabilityLevelRule returns new rule with default attributes
func NewAzurermServiceFabricClusterInvalidReliabilityLevelRule() *AzurermServiceFabricClusterInvalidReliabilityLevelRule {
	return &AzurermServiceFabricClusterInvalidReliabilityLevelRule{
		resourceType:  "azurerm_service_fabric_cluster",
		attributeName: "reliability_level",
		enum: []string{
			"None",
			"Bronze",
			"Silver",
			"Gold",
			"Platinum",
		},
	}
}

// Name returns the rule name
func (r *AzurermServiceFabricClusterInvalidReliabilityLevelRule) Name() string {
	return "azurerm_service_fabric_cluster_invalid_reliability_level"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermServiceFabricClusterInvalidReliabilityLevelRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermServiceFabricClusterInvalidReliabilityLevelRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermServiceFabricClusterInvalidReliabilityLevelRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermServiceFabricClusterInvalidReliabilityLevelRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as reliability_level`, truncateLongMessage(val)),
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
