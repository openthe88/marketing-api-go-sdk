/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// VisibilityCode : 不支持沉淀的状态码
type VisibilityCode string

// List of VisibilityCode
const (
	VisibilityCode_UNKNOWN               VisibilityCode = "VISIBILITY_CODE_UNKNOWN"
	VisibilityCode_INSERT_READY          VisibilityCode = "VISIBILITY_CODE_INSERT_READY"
	VisibilityCode_INSERTED              VisibilityCode = "VISIBILITY_CODE_INSERTED"
	VisibilityCode_NOT_APPROVED          VisibilityCode = "VISIBILITY_CODE_NOT_APPROVED"
	VisibilityCode_NOT_BEGINTIME         VisibilityCode = "VISIBILITY_CODE_NOT_BEGINTIME"
	VisibilityCode_DIFFERENT_CORPORATION VisibilityCode = "VISIBILITY_CODE_DIFFERENT_CORPORATION"
)