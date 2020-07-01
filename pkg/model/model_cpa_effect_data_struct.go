/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 效果数据（成本）
type CpaEffectDataStruct struct {
	Rank                 int64   `json:"rank,omitempty"`
	ConversionCount      int64   `json:"conversion_count,omitempty"`
	Cvr                  float64 `json:"cvr,omitempty"`
	TargetCpa            float64 `json:"target_cpa,omitempty"`
	RealCpa              float64 `json:"real_cpa,omitempty"`
	CpaBias              float64 `json:"cpa_bias,omitempty"`
	IndustryTopTargetcpa int64   `json:"industry_top_targetcpa,omitempty"`
	IndustryAvgTargetcpa int64   `json:"industry_avg_targetcpa,omitempty"`
}
