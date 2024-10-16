/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意组件元数据
type ComponentMetadataStruct struct {
	Name                    *string                        `json:"name,omitempty"`
	ComponentType           ComponentType                  `json:"component_type,omitempty"`
	ComponentSubType        ComponentSubType               `json:"component_sub_type,omitempty"`
	ValueField              *[]ComponentMetadataValueField `json:"value_field,omitempty"`
	ComponentCandidateCount *int64                         `json:"component_candidate_count,omitempty"`
}