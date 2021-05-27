/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdParamGetRequest struct {
	Uid                   *int64             `json:"uid,omitempty"`
	CampaignType          CampaignType       `json:"campaign_type,omitempty"`
	CreativeTemplateId    *int64             `json:"creative_template_id,omitempty"`
	MediaPlacementGroupId *int64             `json:"media_placement_group_id,omitempty"`
	DynamicAdCategory     *int64             `json:"dynamic_ad_category,omitempty"`
	PromotedObjectType    PromotedObjectType `json:"promoted_object_type,omitempty"`
}