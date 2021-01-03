// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDataFactoryLinkedServiceSQLServerInvalidNameRule checks the pattern is valid
type AzurermDataFactoryLinkedServiceSQLServerInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermDataFactoryLinkedServiceSQLServerInvalidNameRule returns new rule with default attributes
func NewAzurermDataFactoryLinkedServiceSQLServerInvalidNameRule() *AzurermDataFactoryLinkedServiceSQLServerInvalidNameRule {
	return &AzurermDataFactoryLinkedServiceSQLServerInvalidNameRule{
		resourceType:  "azurerm_data_factory_linked_service_sql_server",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`),
	}
}

// Name returns the rule name
func (r *AzurermDataFactoryLinkedServiceSQLServerInvalidNameRule) Name() string {
	return "azurerm_data_factory_linked_service_sql_server_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDataFactoryLinkedServiceSQLServerInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDataFactoryLinkedServiceSQLServerInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDataFactoryLinkedServiceSQLServerInvalidNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDataFactoryLinkedServiceSQLServerInvalidNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
