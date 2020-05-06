// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermHdinsightHbaseClusterInvalidTierRule checks the pattern is valid
type AzurermHdinsightHbaseClusterInvalidTierRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermHdinsightHbaseClusterInvalidTierRule returns new rule with default attributes
func NewAzurermHdinsightHbaseClusterInvalidTierRule() *AzurermHdinsightHbaseClusterInvalidTierRule {
	return &AzurermHdinsightHbaseClusterInvalidTierRule{
		resourceType:  "azurerm_hdinsight_hbase_cluster",
		attributeName: "tier",
		enum: []string{
			"Standard",
			"Premium",
		},
	}
}

// Name returns the rule name
func (r *AzurermHdinsightHbaseClusterInvalidTierRule) Name() string {
	return "azurerm_hdinsight_hbase_cluster_invalid_tier"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermHdinsightHbaseClusterInvalidTierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermHdinsightHbaseClusterInvalidTierRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermHdinsightHbaseClusterInvalidTierRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermHdinsightHbaseClusterInvalidTierRule) Check(runner tflint.Runner) error {
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
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as tier`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
