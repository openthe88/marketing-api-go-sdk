/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type TargetingTagsGetListStruct struct {
	Id         int64     `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	ParentId   int64     `json:"parent_id,omitempty"`
	ParentName string    `json:"parent_name,omitempty"`
	CityLevel  CityLevel `json:"city_level,omitempty"`
	TagClass   TagClass  `json:"tag_class,omitempty"`
}
