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
type WechatFundsGetListStruct struct {
	FundType        WechatMpOpenFundType `json:"fund_type,omitempty"`
	Balance         int64                `json:"balance,omitempty"`
	CreditRollSpec  CreditRollSpec       `json:"credit_roll_spec,omitempty"`
	MiniprogramSpec MiniprogramAmount    `json:"miniprogram_spec,omitempty"`
}
