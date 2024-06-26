/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type WechatChannelsAdAccountCertificationFileAddRequest struct {
	AccountId *int64   `json:"account_id,omitempty"`
	Signature *string  `json:"signature,omitempty"`
	FileBytes *string  `json:"file_bytes,omitempty"`
	FileName  *string  `json:"file_name,omitempty"`
	FileType  FileType `json:"file_type,omitempty"`
}
