/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type JointBudgetRulesGetListStruct struct {
	JointBudgetRuleId   *int64              `json:"joint_budget_rule_id,omitempty"`
	JointBudgetRuleName *string             `json:"joint_budget_rule_name,omitempty"`
	DailyBudget         *int64              `json:"daily_budget,omitempty"`
	TotalBudget         *int64              `json:"total_budget,omitempty"`
	CreatedTime         *int64              `json:"created_time,omitempty"`
	LastModTime         *int64              `json:"last_mod_time,omitempty"`
	CompletedTime       *int64              `json:"completed_time,omitempty"`
	Bldate              *int64              `json:"bldate,omitempty"`
	AdgroupIdList       *[]int64            `json:"adgroup_id_list,omitempty"`
	JointBudgetRuleType JointBudgetRuleType `json:"joint_budget_rule_type,omitempty"`
}
