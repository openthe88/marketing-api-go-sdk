/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 获取兴趣定向标签的条件，type 为 INTEREST 时必填
type InterestTargetingTagSpec struct {
	QueryMode TargetingTagQueryMode `json:"query_mode,omitempty"`
	QuerySpec QuerySpec             `json:"query_spec,omitempty"`
}
