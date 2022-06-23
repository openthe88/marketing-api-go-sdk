/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 微信搜一搜品牌专区-通用跳转结构
type WxSearchBrandAreaDocJumpInfo struct {
	JumpType     *int64  `json:"jump_type,omitempty"`
	UserName     *string `json:"user_name,omitempty"`
	WeappUrl     *string `json:"weapp_url,omitempty"`
	JumpUrl      *string `json:"jump_url,omitempty"`
	FeedId       *string `json:"feed_id,omitempty"`
	CommentScene *int64  `json:"comment_scene,omitempty"`
	FeedNonceId  *string `json:"feed_nonce_id,omitempty"`
	ExtInfo      *string `json:"ext_info,omitempty"`
	CanvasId     *string `json:"canvas_id,omitempty"`
}
