/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type MarketingTargetAssetsAddRequest struct {
	OrganizationId      *int64              `json:"organization_id,omitempty"`
	MarketingAssetName  *string             `json:"marketing_asset_name,omitempty"`
	MarketingTargetType MarketingTargetType `json:"marketing_target_type,omitempty"`
	MarketingAssetType  MarketingAssetType  `json:"marketing_asset_type,omitempty"`
	Properties          *[]PropertyStruct   `json:"properties,omitempty"`
}
