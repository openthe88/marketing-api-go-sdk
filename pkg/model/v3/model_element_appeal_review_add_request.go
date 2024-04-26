/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ElementAppealReviewAddRequest struct {
	AccountId                  *int64            `json:"account_id,omitempty"`
	DynamicCreativeId          *int64            `json:"dynamic_creative_id,omitempty"`
	ComponentId                *int64            `json:"component_id,omitempty"`
	ElementId                  *int64            `json:"element_id,omitempty"`
	ElementType                ReviewElementType `json:"element_type,omitempty"`
	ElementValue               *string           `json:"element_value,omitempty"`
	ElementFingerPrint         *string           `json:"element_finger_print,omitempty"`
	AppealDemand               *string           `json:"appeal_demand,omitempty"`
	AppealReason               *string           `json:"appeal_reason,omitempty"`
	HistoryApprovalComponentId *int64            `json:"history_approval_component_id,omitempty"`
	Description                *string           `json:"description,omitempty"`
	ImageList                  *[]string         `json:"image_list,omitempty"`
}
