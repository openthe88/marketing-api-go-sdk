/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

//
type LiveNoticeRecord struct {
	NoticeId     *string `json:"notice_id,omitempty"`
	Status       *int64  `json:"status,omitempty"`
	StartTime    *int64  `json:"start_time,omitempty"`
	Introduction *string `json:"introduction,omitempty"`
}
