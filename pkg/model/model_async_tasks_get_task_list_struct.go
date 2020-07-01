/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type AsyncTasksGetTaskListStruct struct {
	TaskId      int64               `json:"task_id,omitempty"`
	TaskName    string              `json:"task_name,omitempty"`
	TaskType    TaskType            `json:"task_type,omitempty"`
	Status      TaskStatus          `json:"status,omitempty"`
	CreatedTime int64               `json:"created_time,omitempty"`
	Result      AsyncTasksGetResult `json:"result,omitempty"`
}
