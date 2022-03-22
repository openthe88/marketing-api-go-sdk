/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 修改广告组深度优化价值的强化 ROI
type UpdateAdgroupDeepConversionWorthAdvancedRateItem struct {
	AdgroupId                       *int64   `json:"adgroup_id,omitempty"`
	DeepConversionWorthAdvancedRate *float64 `json:"deep_conversion_worth_advanced_rate,omitempty"`
}