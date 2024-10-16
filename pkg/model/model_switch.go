/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// ModelSwitch : 使用MP广告新特性，OPEN ：使用， CLOSE ：不使用，默认CLOSE
type ModelSwitch string

// List of Switch
const (
	ModelSwitch_CLOSE ModelSwitch = "CLOSE"
	ModelSwitch_OPEN  ModelSwitch = "OPEN"
)