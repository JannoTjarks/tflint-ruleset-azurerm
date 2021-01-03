// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermVirtualNetworkGatewayInvalidTypeRule checks the pattern is valid
type AzurermVirtualNetworkGatewayInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermVirtualNetworkGatewayInvalidTypeRule returns new rule with default attributes
func NewAzurermVirtualNetworkGatewayInvalidTypeRule() *AzurermVirtualNetworkGatewayInvalidTypeRule {
	return &AzurermVirtualNetworkGatewayInvalidTypeRule{
		resourceType:  "azurerm_virtual_network_gateway",
		attributeName: "type",
		enum: []string{
			"Vpn",
			"ExpressRoute",
		},
	}
}

// Name returns the rule name
func (r *AzurermVirtualNetworkGatewayInvalidTypeRule) Name() string {
	return "azurerm_virtual_network_gateway_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermVirtualNetworkGatewayInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermVirtualNetworkGatewayInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermVirtualNetworkGatewayInvalidTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermVirtualNetworkGatewayInvalidTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

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
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
