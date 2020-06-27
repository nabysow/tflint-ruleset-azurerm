// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermStorageAccountNetworkRulesInvalidDefaultActionRule checks the pattern is valid
type AzurermStorageAccountNetworkRulesInvalidDefaultActionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermStorageAccountNetworkRulesInvalidDefaultActionRule returns new rule with default attributes
func NewAzurermStorageAccountNetworkRulesInvalidDefaultActionRule() *AzurermStorageAccountNetworkRulesInvalidDefaultActionRule {
	return &AzurermStorageAccountNetworkRulesInvalidDefaultActionRule{
		resourceType:  "azurerm_storage_account_network_rules",
		attributeName: "default_action",
		enum: []string{
			"Allow",
			"Deny",
		},
	}
}

// Name returns the rule name
func (r *AzurermStorageAccountNetworkRulesInvalidDefaultActionRule) Name() string {
	return "azurerm_storage_account_network_rules_invalid_default_action"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermStorageAccountNetworkRulesInvalidDefaultActionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermStorageAccountNetworkRulesInvalidDefaultActionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermStorageAccountNetworkRulesInvalidDefaultActionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermStorageAccountNetworkRulesInvalidDefaultActionRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as default_action`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
