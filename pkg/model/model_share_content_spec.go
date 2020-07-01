/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 分享信息，仅在部分朋友圈创意规格下设置有效，具体可通过创意规格查询工具或 adcreative_templates/get 接口进行查询
type ShareContentSpec struct {
	ShareTitle       string `json:"share_title,omitempty"`
	ShareDescription string `json:"share_description,omitempty"`
}
