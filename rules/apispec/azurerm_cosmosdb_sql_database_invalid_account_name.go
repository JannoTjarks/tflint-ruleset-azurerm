// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermCosmosdbSQLDatabaseInvalidAccountNameRule checks the pattern is valid
type AzurermCosmosdbSQLDatabaseInvalidAccountNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermCosmosdbSQLDatabaseInvalidAccountNameRule returns new rule with default attributes
func NewAzurermCosmosdbSQLDatabaseInvalidAccountNameRule() *AzurermCosmosdbSQLDatabaseInvalidAccountNameRule {
	return &AzurermCosmosdbSQLDatabaseInvalidAccountNameRule{
		resourceType:  "azurerm_cosmosdb_sql_database",
		attributeName: "account_name",
		pattern:       regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*`),
	}
}

// Name returns the rule name
func (r *AzurermCosmosdbSQLDatabaseInvalidAccountNameRule) Name() string {
	return "azurerm_cosmosdb_sql_database_invalid_account_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermCosmosdbSQLDatabaseInvalidAccountNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermCosmosdbSQLDatabaseInvalidAccountNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermCosmosdbSQLDatabaseInvalidAccountNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermCosmosdbSQLDatabaseInvalidAccountNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z0-9]+(-[a-z0-9]+)*`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
