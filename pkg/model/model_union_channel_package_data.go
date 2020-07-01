/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 渠道包信息
type UnionChannelPackageData struct {
	AndroidUnionAppId int64     `json:"android_union_app_id,omitempty"`
	PackageName       string    `json:"package_name,omitempty"`
	ChannelPackageId  int64     `json:"channel_package_id,omitempty"`
	VersionCode       int64     `json:"version_code,omitempty"`
	VersionName       string    `json:"version_name,omitempty"`
	CreatedTime       int64     `json:"created_time,omitempty"`
	LastModifiedTime  int64     `json:"last_modified_time,omitempty"`
	SystemStatus      SysStatus `json:"system_status,omitempty"`
	SystemMessage     string    `json:"system_message,omitempty"`
}
