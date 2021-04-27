/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdParamGetResponseData struct {
	SystemAdParamList  *[]AdParamParentListStruct `json:"system_ad_param_list,omitempty"`
	ProductAdParamList *[]AdParamParentListStruct `json:"product_ad_param_list,omitempty"`
	UserAdParamList    *[]AdParamParentListStruct `json:"user_ad_param_list,omitempty"`
}
