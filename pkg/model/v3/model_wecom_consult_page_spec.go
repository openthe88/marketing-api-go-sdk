/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 企业微信客服落地页
type WecomConsultPageSpec struct {
	PageId    *int64    `json:"page_id,omitempty"`
	GroupType GroupType `json:"group_type,omitempty"`
}
