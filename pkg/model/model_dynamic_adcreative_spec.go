/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 动态商品广告属性
type DynamicAdcreativeSpec struct {
	ProductCatalogId int64       `json:"product_catalog_id,omitempty"`
	ProductMode      ProductMode `json:"product_mode,omitempty"`
	ProductSource    string      `json:"product_source,omitempty"`
}
