/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 裁剪信息
type CropCustomizedSpec struct {
	Width  int64 `json:"width,omitempty"`
	Height int64 `json:"height,omitempty"`
	AxisX  int64 `json:"axis_x,omitempty"`
	AxisY  int64 `json:"axis_y,omitempty"`
}
