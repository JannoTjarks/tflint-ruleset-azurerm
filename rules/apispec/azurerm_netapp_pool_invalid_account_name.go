// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermNetappPoolInvalidAccountNameRule checks the pattern is valid
type AzurermNetappPoolInvalidAccountNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermNetappPoolInvalidAccountNameRule returns new rule with default attributes
func NewAzurermNetappPoolInvalidAccountNameRule() *AzurermNetappPoolInvalidAccountNameRule {
	return &AzurermNetappPoolInvalidAccountNameRule{
		resourceType:  "azurerm_netapp_pool",
		attributeName: "account_name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9\-_]{0,127}$`),
	}
}

// Name returns the rule name
func (r *AzurermNetappPoolInvalidAccountNameRule) Name() string {
	return "azurerm_netapp_pool_invalid_account_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermNetappPoolInvalidAccountNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermNetappPoolInvalidAccountNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermNetappPoolInvalidAccountNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermNetappPoolInvalidAccountNameRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}
		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9][a-zA-Z0-9\-_]{0,127}$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}