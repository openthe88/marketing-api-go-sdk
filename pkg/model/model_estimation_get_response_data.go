/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type EstimationGetResponseData struct {
	MaxActiveUserCount      int64           `json:"max_active_user_count,omitempty"`
	ApproximateCount        int64           `json:"approximate_count,omitempty"`
	Impression              int64           `json:"impression,omitempty"`
	MinBidAmount            int64           `json:"min_bid_amount,omitempty"`
	MaxBidAmount            int64           `json:"max_bid_amount,omitempty"`
	SuggestMinBidAmount     int64           `json:"suggest_min_bid_amount,omitempty"`
	SuggestMaxBidAmount     int64           `json:"suggest_max_bid_amount,omitempty"`
	SuggestBidContentOcpa   string          `json:"suggest_bid_content_ocpa,omitempty"`
	MinUsersDaily           int64           `json:"min_users_daily,omitempty"`
	MaxUsersDaily           int64           `json:"max_users_daily,omitempty"`
	MinExposureDaily        int64           `json:"min_exposure_daily,omitempty"`
	MaxExposureDaily        int64           `json:"max_exposure_daily,omitempty"`
	TargetingStatus         TargetingStatus `json:"targeting_status,omitempty"`
	SuggestTargeting        []string        `json:"suggest_targeting,omitempty"`
	IsRealExposureSupported bool            `json:"is_real_exposure_supported,omitempty"`
}
