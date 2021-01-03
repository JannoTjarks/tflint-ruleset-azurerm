// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermFrontdoorInvalidBackendPoolsSendReceiveTimeoutSecondsRule checks the pattern is valid
type AzurermFrontdoorInvalidBackendPoolsSendReceiveTimeoutSecondsRule struct {
	resourceType  string
	attributeName string
	min           int
}

// NewAzurermFrontdoorInvalidBackendPoolsSendReceiveTimeoutSecondsRule returns new rule with default attributes
func NewAzurermFrontdoorInvalidBackendPoolsSendReceiveTimeoutSecondsRule() *AzurermFrontdoorInvalidBackendPoolsSendReceiveTimeoutSecondsRule {
	return &AzurermFrontdoorInvalidBackendPoolsSendReceiveTimeoutSecondsRule{
		resourceType:  "azurerm_frontdoor",
		attributeName: "backend_pools_send_receive_timeout_seconds",
		min:           16,
	}
}

// Name returns the rule name
func (r *AzurermFrontdoorInvalidBackendPoolsSendReceiveTimeoutSecondsRule) Name() string {
	return "azurerm_frontdoor_invalid_backend_pools_send_receive_timeout_seconds"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermFrontdoorInvalidBackendPoolsSendReceiveTimeoutSecondsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermFrontdoorInvalidBackendPoolsSendReceiveTimeoutSecondsRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermFrontdoorInvalidBackendPoolsSendReceiveTimeoutSecondsRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermFrontdoorInvalidBackendPoolsSendReceiveTimeoutSecondsRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val int
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if val < r.min {
				runner.EmitIssueOnExpr(
					r,
					"backend_pools_send_receive_timeout_seconds must be 16 or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
