// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermMysqlFirewallRuleInvalidEndIPAddressRule checks the pattern is valid
type AzurermMysqlFirewallRuleInvalidEndIPAddressRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermMysqlFirewallRuleInvalidEndIPAddressRule returns new rule with default attributes
func NewAzurermMysqlFirewallRuleInvalidEndIPAddressRule() *AzurermMysqlFirewallRuleInvalidEndIPAddressRule {
	return &AzurermMysqlFirewallRuleInvalidEndIPAddressRule{
		resourceType:  "azurerm_mysql_firewall_rule",
		attributeName: "end_ip_address",
		pattern:       regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`),
	}
}

// Name returns the rule name
func (r *AzurermMysqlFirewallRuleInvalidEndIPAddressRule) Name() string {
	return "azurerm_mysql_firewall_rule_invalid_end_ip_address"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermMysqlFirewallRuleInvalidEndIPAddressRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermMysqlFirewallRuleInvalidEndIPAddressRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermMysqlFirewallRuleInvalidEndIPAddressRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermMysqlFirewallRuleInvalidEndIPAddressRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
