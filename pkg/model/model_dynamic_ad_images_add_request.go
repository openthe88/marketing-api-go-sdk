/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type DynamicAdImagesAddRequest struct {
	AccountId             int64                 `json:"account_id,omitempty"`
	ProductCatalogId      int64                 `json:"product_catalog_id,omitempty"`
	ProductMode           ProductMode           `json:"product_mode,omitempty"`
	ProductSource         string                `json:"product_source,omitempty"`
	DynamicAdTemplateSize DynamicAdTemplateSize `json:"dynamic_ad_template_size,omitempty"`
	DynamicAdTemplateId   int64                 `json:"dynamic_ad_template_id,omitempty"`
	ImageMattingEnabled   bool                  `json:"image_matting_enabled,omitempty"`
}
