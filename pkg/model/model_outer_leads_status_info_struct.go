/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 回传信息结构
type OuterLeadsStatusInfoStruct struct {
	OuterLeadsId             *string `json:"outer_leads_id,omitempty"`
	OuterLeadsConvertType    *string `json:"outer_leads_convert_type,omitempty"`
	OuterLeadsIneffectReason *string `json:"outer_leads_ineffect_reason,omitempty"`
}
