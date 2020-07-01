/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 账户内部竞争激烈度
type CompeteWinAdgroupStruct struct {
	Score      int64                       `json:"score,omitempty"`
	List       []CompeteWinAdgroupListItem `json:"list,omitempty"`
	Conclusion string                      `json:"conclusion,omitempty"`
}
