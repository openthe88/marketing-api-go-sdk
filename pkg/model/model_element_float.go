/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 悬浮组件
type ElementFloat struct {
	Title           string          `json:"title,omitempty"`
	Desc            string          `json:"desc,omitempty"`
	FloatButtonSpec FloatButtonSpec `json:"float_button_spec,omitempty"`
	ImageIdList     string          `json:"image_id_list,omitempty"`
}
