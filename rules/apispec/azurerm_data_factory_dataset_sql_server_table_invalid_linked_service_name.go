// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule checks the pattern is valid
type AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule returns new rule with default attributes
func NewAzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule() *AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule {
	return &AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule{
		resourceType:  "azurerm_data_factory_dataset_sql_server_table",
		attributeName: "linked_service_name",
		pattern:       regexp.MustCompile(`^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`),
	}
}

// Name returns the rule name
func (r *AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule) Name() string {
	return "azurerm_data_factory_dataset_sql_server_table_invalid_linked_service_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
