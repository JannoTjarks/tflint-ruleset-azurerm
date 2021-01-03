// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermMysqlServerInvalidVersionRule checks the pattern is valid
type AzurermMysqlServerInvalidVersionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermMysqlServerInvalidVersionRule returns new rule with default attributes
func NewAzurermMysqlServerInvalidVersionRule() *AzurermMysqlServerInvalidVersionRule {
	return &AzurermMysqlServerInvalidVersionRule{
		resourceType:  "azurerm_mysql_server",
		attributeName: "version",
		enum: []string{
			"5.6",
			"5.7",
			"8.0",
		},
	}
}

// Name returns the rule name
func (r *AzurermMysqlServerInvalidVersionRule) Name() string {
	return "azurerm_mysql_server_invalid_version"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermMysqlServerInvalidVersionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermMysqlServerInvalidVersionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermMysqlServerInvalidVersionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermMysqlServerInvalidVersionRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as version`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
