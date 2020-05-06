// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermVirtualNetworkGatewayConnectionInvalidTypeRule checks the pattern is valid
type AzurermVirtualNetworkGatewayConnectionInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermVirtualNetworkGatewayConnectionInvalidTypeRule returns new rule with default attributes
func NewAzurermVirtualNetworkGatewayConnectionInvalidTypeRule() *AzurermVirtualNetworkGatewayConnectionInvalidTypeRule {
	return &AzurermVirtualNetworkGatewayConnectionInvalidTypeRule{
		resourceType:  "azurerm_virtual_network_gateway_connection",
		attributeName: "type",
		enum: []string{
			"IPsec",
			"Vnet2Vnet",
			"ExpressRoute",
			"VPNClient",
		},
	}
}

// Name returns the rule name
func (r *AzurermVirtualNetworkGatewayConnectionInvalidTypeRule) Name() string {
	return "azurerm_virtual_network_gateway_connection_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermVirtualNetworkGatewayConnectionInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermVirtualNetworkGatewayConnectionInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermVirtualNetworkGatewayConnectionInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermVirtualNetworkGatewayConnectionInvalidTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
