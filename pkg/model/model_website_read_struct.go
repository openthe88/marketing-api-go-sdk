/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 推广链接结构
type WebsiteReadStruct struct {
	WebsiteDomain *string       `json:"website_domain,omitempty"`
	IcpImageId    *string       `json:"icp_image_id,omitempty"`
	SystemStatus  WebsiteStatus `json:"system_status,omitempty"`
	RejectMessage *string       `json:"reject_message,omitempty"`
}
