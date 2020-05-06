// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermDataFactoryIntegrationRuntimeManagedInvalidLicenseTypeRule checks the pattern is valid
type AzurermDataFactoryIntegrationRuntimeManagedInvalidLicenseTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermDataFactoryIntegrationRuntimeManagedInvalidLicenseTypeRule returns new rule with default attributes
func NewAzurermDataFactoryIntegrationRuntimeManagedInvalidLicenseTypeRule() *AzurermDataFactoryIntegrationRuntimeManagedInvalidLicenseTypeRule {
	return &AzurermDataFactoryIntegrationRuntimeManagedInvalidLicenseTypeRule{
		resourceType:  "azurerm_data_factory_integration_runtime_managed",
		attributeName: "license_type",
		enum: []string{
			"BasePrice",
			"LicenseIncluded",
		},
	}
}

// Name returns the rule name
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidLicenseTypeRule) Name() string {
	return "azurerm_data_factory_integration_runtime_managed_invalid_license_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidLicenseTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidLicenseTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidLicenseTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermDataFactoryIntegrationRuntimeManagedInvalidLicenseTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as license_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
