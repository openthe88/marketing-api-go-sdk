/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 游标分页信息
type CursorPageInfoStruct struct {
	PageSize       *int64  `json:"page_size,omitempty"`
	TotalNumber    *int64  `json:"total_number,omitempty"`
	NextCursor     *string `json:"next_cursor,omitempty"`
	PreviousCursor *string `json:"previous_cursor,omitempty"`
}