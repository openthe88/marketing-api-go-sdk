/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 微信搜一搜品牌专区-doc区块-卡片信息
type WxSearchBrandAreaDocActivityInfoActivityBlock struct {
	Title    *string                                              `json:"title,omitempty"`
	DescList *[]string                                            `json:"desc_list,omitempty"`
	DescSpec *string                                              `json:"desc_spec,omitempty"`
	Button   *WxSearchBrandAreaDocActivityInfoActivityBlockButton `json:"button,omitempty"`
}
