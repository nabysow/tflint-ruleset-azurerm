// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermCosmosdbAccountInvalidOfferTypeRule checks the pattern is valid
type AzurermCosmosdbAccountInvalidOfferTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermCosmosdbAccountInvalidOfferTypeRule returns new rule with default attributes
func NewAzurermCosmosdbAccountInvalidOfferTypeRule() *AzurermCosmosdbAccountInvalidOfferTypeRule {
	return &AzurermCosmosdbAccountInvalidOfferTypeRule{
		resourceType:  "azurerm_cosmosdb_account",
		attributeName: "offer_type",
		enum: []string{
			"Standard",
		},
	}
}

// Name returns the rule name
func (r *AzurermCosmosdbAccountInvalidOfferTypeRule) Name() string {
	return "azurerm_cosmosdb_account_invalid_offer_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermCosmosdbAccountInvalidOfferTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermCosmosdbAccountInvalidOfferTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermCosmosdbAccountInvalidOfferTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermCosmosdbAccountInvalidOfferTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as offer_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
