/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 标签信息
type CreateLabel struct {
	LabelName            *string          `json:"label_name,omitempty"`
	FirstLabelLevelName  *string          `json:"first_label_level_name,omitempty"`
	SecondLabelLevelName *string          `json:"second_label_level_name,omitempty"`
	BusinessScenario     BusinessScenario `json:"business_scenario,omitempty"`
}