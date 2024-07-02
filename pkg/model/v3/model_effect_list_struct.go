/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type EffectListStruct struct {
	EffectDate *string `json:"effect_date,omitempty"`
	ExpireDate *string `json:"expire_date,omitempty"`
	Amount     *int64  `json:"amount,omitempty"`
}
