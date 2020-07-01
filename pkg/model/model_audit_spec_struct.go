/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 多站点的审核结果信息
type AuditSpecStruct struct {
	SiteSet       SiteSetDefinition `json:"site_set,omitempty"`
	SystemStatus  SysStatus         `json:"system_status,omitempty"`
	RejectMessage string            `json:"reject_message,omitempty"`
}
