/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 行为兴趣意向定向的兴趣部分
type Interest struct {
	TargetingTags  []string `json:"targeting_tags,omitempty"`
	CategoryIdList []int64  `json:"category_id_list,omitempty"`
	KeywordList    []string `json:"keyword_list,omitempty"`
}
