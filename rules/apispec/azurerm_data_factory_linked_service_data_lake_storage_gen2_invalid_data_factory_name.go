// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDataFactoryLinkedServiceDataLakeStorageGen2InvalidDataFactoryNameRule checks the pattern is valid
type AzurermDataFactoryLinkedServiceDataLakeStorageGen2InvalidDataFactoryNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermDataFactoryLinkedServiceDataLakeStorageGen2InvalidDataFactoryNameRule returns new rule with default attributes
func NewAzurermDataFactoryLinkedServiceDataLakeStorageGen2InvalidDataFactoryNameRule() *AzurermDataFactoryLinkedServiceDataLakeStorageGen2InvalidDataFactoryNameRule {
	return &AzurermDataFactoryLinkedServiceDataLakeStorageGen2InvalidDataFactoryNameRule{
		resourceType:  "azurerm_data_factory_linked_service_data_lake_storage_gen2",
		attributeName: "data_factory_name",
		pattern:       regexp.MustCompile(`^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`),
	}
}

// Name returns the rule name
func (r *AzurermDataFactoryLinkedServiceDataLakeStorageGen2InvalidDataFactoryNameRule) Name() string {
	return "azurerm_data_factory_linked_service_data_lake_storage_gen2_invalid_data_factory_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDataFactoryLinkedServiceDataLakeStorageGen2InvalidDataFactoryNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDataFactoryLinkedServiceDataLakeStorageGen2InvalidDataFactoryNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDataFactoryLinkedServiceDataLakeStorageGen2InvalidDataFactoryNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDataFactoryLinkedServiceDataLakeStorageGen2InvalidDataFactoryNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
