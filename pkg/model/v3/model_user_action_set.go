/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// UserActionSet返回结构
type UserActionSet struct {
	UserActionSetId       *int64                 `json:"user_action_set_id,omitempty"`
	Type_                 AmUserActionSetType    `json:"type,omitempty"`
	MobileAppId           *int64                 `json:"mobile_app_id,omitempty"`
	Name                  *string                `json:"name,omitempty"`
	Description           *string                `json:"description,omitempty"`
	ActivateStatus        *bool                  `json:"activate_status,omitempty"`
	CreatedTime           *string                `json:"created_time,omitempty"`
	AccessWay             ActionSetAccessWayType `json:"access_way,omitempty"`
	Usages                *[]string              `json:"usages,omitempty"`
	EnableConversionClaim *bool                  `json:"enable_conversion_claim,omitempty"`
	Permission            *Permission            `json:"permission,omitempty"`
}
