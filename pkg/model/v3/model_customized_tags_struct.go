/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 自定义标签集合
type CustomizedTagsStruct struct {
	TagGroupName *string   `json:"tag_group_name,omitempty"`
	TagNameList  *[]string `json:"tag_name_list,omitempty"`
}
