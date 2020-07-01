/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 悬浮组件按钮
type FloatButtonSpec struct {
	LinkSpec            LinkSpec            `json:"link_spec,omitempty"`
	AppDownloadSpec     AppDownloadSpec     `json:"app_download_spec,omitempty"`
	MiniProgramSpec     MiniProgramSpec     `json:"mini_program_spec,omitempty"`
	MiniGameProgramSpec MiniGameProgramSpec `json:"mini_game_program_spec,omitempty"`
	FengyeSpec          FengyeSpec          `json:"fengye_spec,omitempty"`
	CardSpec            CardSpec            `json:"card_spec,omitempty"`
	FollowSpec          FollowSpec          `json:"follow_spec,omitempty"`
	ServiceSpec         ServiceSpec         `json:"service_spec,omitempty"`
}
