/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type QualificationsDeleteRequest struct {
	QualificationType QualificationType `json:"qualification_type,omitempty"`
	QualificationId   int64             `json:"qualification_id,omitempty"`
	AccountId         int64             `json:"account_id,omitempty"`
}
