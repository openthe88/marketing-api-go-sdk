/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告诊断信息列表
type DiagnosisGetListStruct struct {
	AdgroupId                 int64                     `json:"adgroup_id,omitempty"`
	WechatDiagnosisResultSpec WechatDiagnosisResultSpec `json:"wechat_diagnosis_result_spec,omitempty"`
}
