// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDevTestPolicyInvalidEvaluatorTypeRule checks the pattern is valid
type AzurermDevTestPolicyInvalidEvaluatorTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermDevTestPolicyInvalidEvaluatorTypeRule returns new rule with default attributes
func NewAzurermDevTestPolicyInvalidEvaluatorTypeRule() *AzurermDevTestPolicyInvalidEvaluatorTypeRule {
	return &AzurermDevTestPolicyInvalidEvaluatorTypeRule{
		resourceType:  "azurerm_dev_test_policy",
		attributeName: "evaluator_type",
		enum: []string{
			"AllowedValuesPolicy",
			"MaxValuePolicy",
		},
	}
}

// Name returns the rule name
func (r *AzurermDevTestPolicyInvalidEvaluatorTypeRule) Name() string {
	return "azurerm_dev_test_policy_invalid_evaluator_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDevTestPolicyInvalidEvaluatorTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDevTestPolicyInvalidEvaluatorTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDevTestPolicyInvalidEvaluatorTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDevTestPolicyInvalidEvaluatorTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as evaluator_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
