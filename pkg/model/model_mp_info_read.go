/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// <p>微信广告开户信息，满足如下条件时使用</p><li>输入参数 account_id 不为空时使用</li><li>微信公众号开通广告业务或开通通用账号权限</li>
type MpInfoRead struct {
	WechatAccountId        string             `json:"wechat_account_id,omitempty"`
	WechatAccountName      string             `json:"wechat_account_name,omitempty"`
	SystemStatus           SysStatus          `json:"system_status,omitempty"`
	IndustryName           string             `json:"industry_name,omitempty"`
	ContactPerson          string             `json:"contact_person,omitempty"`
	ContactPersonTelephone string             `json:"contact_person_telephone,omitempty"`
	BusinessType           WechatBusinessType `json:"business_type,omitempty"`
	BusinessContent        string             `json:"business_content,omitempty"`
	RejectMessage          string             `json:"reject_message,omitempty"`
	ProfilePhoto           string             `json:"profile_photo,omitempty"`
	BrandIntroduction      string             `json:"brand_introduction,omitempty"`
	IntroductionUrl        string             `json:"introduction_url,omitempty"`
	SystemIndustryId       int64              `json:"system_industry_id,omitempty"`
}
