/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type RealtimeCostGetListStruct struct {
	CampaignId int64 `json:"campaign_id,omitempty"`
	AdgroupId  int64 `json:"adgroup_id,omitempty"`
	Cost       int64 `json:"cost,omitempty"`
}
