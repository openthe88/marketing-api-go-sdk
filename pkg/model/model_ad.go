/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告信息
type Ad struct {
	AdName                  string            `json:"ad_name,omitempty"`
	Adcreative              PreviewAdcreative `json:"adcreative,omitempty"`
	FeedsInteractionEnabled bool              `json:"feeds_interaction_enabled,omitempty"`
}
