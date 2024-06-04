/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 修改广告组客户设置的状态
type UpdateAdgroupConfiguredStatusItem struct {
	AdgroupId        *int64           `json:"adgroup_id,omitempty"`
	ConfiguredStatus ConfiguredStatus `json:"configured_status,omitempty"`
}