/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 归因
type CompeteAttributionStruct struct {
	AttributeCode CompeteAttribution `json:"attribute_code,omitempty"`
	AttributeName string             `json:"attribute_name,omitempty"`
	Ratio         float64            `json:"ratio,omitempty"`
}
