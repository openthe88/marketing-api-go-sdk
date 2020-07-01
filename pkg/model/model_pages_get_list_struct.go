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
type PagesGetListStruct struct {
	PageId           int64               `json:"page_id,omitempty"`
	PageName         string              `json:"page_name,omitempty"`
	PreviewUrl       string              `json:"preview_url,omitempty"`
	CreatedTime      int64               `json:"created_time,omitempty"`
	LastModifiedTime int64               `json:"last_modified_time,omitempty"`
	PromotedObjectId string              `json:"promoted_object_id,omitempty"`
	PageType         DestinationTypeRead `json:"page_type,omitempty"`
}
