/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// data_file返回结构
type DataFile struct {
	AudienceId           *int64            `json:"audience_id,omitempty"`
	CustomAudienceFileId *int64            `json:"custom_audience_file_id,omitempty"`
	Name                 *string           `json:"name,omitempty"`
	UserIdType           UserIdType        `json:"user_id_type,omitempty"`
	OperationType        OperationType     `json:"operation_type,omitempty"`
	OpenAppId            *string           `json:"open_app_id,omitempty"`
	ProcessStatus        FileProcessStatus `json:"process_status,omitempty"`
	ProcessCode          *int64            `json:"process_code,omitempty"`
	ErrorMessage         *string           `json:"error_message,omitempty"`
	LineCount            *int64            `json:"line_count,omitempty"`
	ValidLineCount       *int64            `json:"valid_line_count,omitempty"`
	UserCount            *int64            `json:"user_count,omitempty"`
	Size                 *int64            `json:"size,omitempty"`
	CreatedTime          *string           `json:"created_time,omitempty"`
}
