/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type QualificationStructureGetResponseData struct {
	QualificationCode     *string                  `json:"qualification_code,omitempty"`
	QualificationName     *string                  `json:"qualification_name,omitempty"`
	ExpandFieldDefinition *[]ExpandFieldDefinition `json:"expand_field_definition,omitempty"`
}