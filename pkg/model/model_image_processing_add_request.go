/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ImageProcessingAddRequest struct {
	AccountId     int64              `json:"account_id,omitempty"`
	ImageId       string             `json:"image_id,omitempty"`
	OperationType ImageOperationType `json:"operation_type,omitempty"`
	OperationSpec CropOperationSpec  `json:"operation_spec,omitempty"`
	FileSizeMax   int64              `json:"file_size_max,omitempty"`
}
