// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDataLakeAnalyticsAccountInvalidTierRule checks the pattern is valid
type AzurermDataLakeAnalyticsAccountInvalidTierRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermDataLakeAnalyticsAccountInvalidTierRule returns new rule with default attributes
func NewAzurermDataLakeAnalyticsAccountInvalidTierRule() *AzurermDataLakeAnalyticsAccountInvalidTierRule {
	return &AzurermDataLakeAnalyticsAccountInvalidTierRule{
		resourceType:  "azurerm_data_lake_analytics_account",
		attributeName: "tier",
		enum: []string{
			"Consumption",
			"Commitment_100AUHours",
			"Commitment_500AUHours",
			"Commitment_1000AUHours",
			"Commitment_5000AUHours",
			"Commitment_10000AUHours",
			"Commitment_50000AUHours",
			"Commitment_100000AUHours",
			"Commitment_500000AUHours",
		},
	}
}

// Name returns the rule name
func (r *AzurermDataLakeAnalyticsAccountInvalidTierRule) Name() string {
	return "azurerm_data_lake_analytics_account_invalid_tier"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDataLakeAnalyticsAccountInvalidTierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDataLakeAnalyticsAccountInvalidTierRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDataLakeAnalyticsAccountInvalidTierRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDataLakeAnalyticsAccountInvalidTierRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as tier`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
