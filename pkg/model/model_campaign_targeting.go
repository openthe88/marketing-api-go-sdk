/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 推广计划信息
type CampaignTargeting struct {
	CampaignId   int64        `json:"campaign_id,omitempty"`
	CampaignType CampaignType `json:"campaign_type,omitempty"`
	DailyBudget  int64        `json:"daily_budget,omitempty"`
}
