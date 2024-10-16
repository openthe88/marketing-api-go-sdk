/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type AdgroupsGetListStruct struct {
	CampaignId                        *int64                          `json:"campaign_id,omitempty"`
	AdgroupId                         *int64                          `json:"adgroup_id,omitempty"`
	AdgroupName                       *string                         `json:"adgroup_name,omitempty"`
	SiteSet                           *[]string                       `json:"site_set,omitempty"`
	AutomaticSiteEnabled              *bool                           `json:"automatic_site_enabled,omitempty"`
	ExplorationStrategy               SiteSetExplorationStrategy      `json:"exploration_strategy,omitempty"`
	PrioritySiteSet                   *[]string                       `json:"priority_site_set,omitempty"`
	OptimizationGoal                  OptimizationGoal                `json:"optimization_goal,omitempty"`
	BillingEvent                      BillingEvent                    `json:"billing_event,omitempty"`
	BidAmount                         *int64                          `json:"bid_amount,omitempty"`
	TotalBudget                       *int64                          `json:"total_budget,omitempty"`
	DailyBudget                       *int64                          `json:"daily_budget,omitempty"`
	PromotedObjectType                PromotedObjectType              `json:"promoted_object_type,omitempty"`
	PromotedObjectId                  *string                         `json:"promoted_object_id,omitempty"`
	AppAndroidChannelPackageId        *string                         `json:"app_android_channel_package_id,omitempty"`
	MiniGameProgramId                 *string                         `json:"mini_game_program_id,omitempty"`
	TargetingId                       *int64                          `json:"targeting_id,omitempty"`
	Targeting                         *ReadTargetingSettingForAdgroup `json:"targeting,omitempty"`
	TargetingTranslation              *string                         `json:"targeting_translation,omitempty"`
	IsIncludeUnsupportedTargeting     *bool                           `json:"is_include_unsupported_targeting,omitempty"`
	SceneSpec                         *SceneTargeting                 `json:"scene_spec,omitempty"`
	FlowOptimizationEnabled           *bool                           `json:"flow_optimization_enabled,omitempty"`
	BeginDate                         *string                         `json:"begin_date,omitempty"`
	FirstDayBeginTime                 *string                         `json:"first_day_begin_time,omitempty"`
	EndDate                           *string                         `json:"end_date,omitempty"`
	TimeSeries                        *string                         `json:"time_series,omitempty"`
	ConfiguredStatus                  AdStatus                        `json:"configured_status,omitempty"`
	CustomizedCategory                *string                         `json:"customized_category,omitempty"`
	CreatedTime                       *int64                          `json:"created_time,omitempty"`
	LastModifiedTime                  *int64                          `json:"last_modified_time,omitempty"`
	AdCount                           *int64                          `json:"ad_count,omitempty"`
	DynamicAdSpec                     *DynamicAdSpec                  `json:"dynamic_ad_spec,omitempty"`
	UserActionSets                    *[]UserActionSetStructDn        `json:"user_action_sets,omitempty"`
	AdditionalUserActionSets          *[]UserActionSetStruct          `json:"additional_user_action_sets,omitempty"`
	IsDeleted                         *bool                           `json:"is_deleted,omitempty"`
	DynamicCreativeId                 *int64                          `json:"dynamic_creative_id,omitempty"`
	IsRewardedVideoAd                 *bool                           `json:"is_rewarded_video_ad,omitempty"`
	CostGuaranteeMessage              *string                         `json:"cost_guarantee_message,omitempty"`
	CostGuaranteeStatus               CostGuaranteeStatus             `json:"cost_guarantee_status,omitempty"`
	BidStrategy                       BidStrategy                     `json:"bid_strategy,omitempty"`
	ColdStartAudience                 *[]int64                        `json:"cold_start_audience,omitempty"`
	AutoAudience                      *bool                           `json:"auto_audience,omitempty"`
	ExpandEnabled                     *bool                           `json:"expand_enabled,omitempty"`
	ExpandTargeting                   *[]string                       `json:"expand_targeting,omitempty"`
	DeepConversionSpec                *DeepConversionSpec             `json:"deep_conversion_spec,omitempty"`
	DeepOptimizationActionType        DeepOptimizationActionType      `json:"deep_optimization_action_type,omitempty"`
	PoiList                           *[]string                       `json:"poi_list,omitempty"`
	ConversionId                      *int64                          `json:"conversion_id,omitempty"`
	DeepConversionBehaviorBid         *int64                          `json:"deep_conversion_behavior_bid,omitempty"`
	DeepConversionWorthRate           *float64                        `json:"deep_conversion_worth_rate,omitempty"`
	DeepConversionWorthAdvancedRate   *float64                        `json:"deep_conversion_worth_advanced_rate,omitempty"`
	DeepConversionBehaviorAdvancedBid *int64                          `json:"deep_conversion_behavior_advanced_bid,omitempty"`
	AndroidChannelPackageAuditMessage *string                         `json:"android_channel_package_audit_message,omitempty"`
	SystemStatus                      AdgroupSysStatus                `json:"system_status,omitempty"`
	BidMode                           BidMode                         `json:"bid_mode,omitempty"`
	Status                            CalcStatus                      `json:"status,omitempty"`
	BidAdjustment                     *BidAdjustment                  `json:"bid_adjustment,omitempty"`
	AutoAcquisitionEnabled            *bool                           `json:"auto_acquisition_enabled,omitempty"`
	AutoAcquisitionBudget             *int64                          `json:"auto_acquisition_budget,omitempty"`
	CreativeDisplayType               CreativeDisplayType             `json:"creative_display_type,omitempty"`
	AutoDerivedCreativeEnabled        *bool                           `json:"auto_derived_creative_enabled,omitempty"`
	SmartBidType                      SmartBidType                    `json:"smart_bid_type,omitempty"`
	SmartCostCap                      *int64                          `json:"smart_cost_cap,omitempty"`
	MarketingScene                    MarketingScene                  `json:"marketing_scene,omitempty"`
	CustomAdgroupTag                  *[]string                       `json:"custom_adgroup_tag,omitempty"`
	SmartTargeting                    *SmartTargeting                 `json:"smart_targeting,omitempty"`
	DynamicCreativeIdSet              *[]int64                        `json:"dynamic_creative_id_set,omitempty"`
	SystemStatusExplanation           *string                         `json:"system_status_explanation,omitempty"`
	AutoDerivedLandingPageSwitch      *bool                           `json:"auto_derived_landing_page_switch,omitempty"`
	SearchExpandTargetingSwitch       SearchExpandTargetingSwitch     `json:"search_expand_targeting_switch,omitempty"`
	AutoAcquisitionStatus             AutoAcquisitionStatus           `json:"auto_acquisition_status,omitempty"`
	EcomPkamSwitch                    EcomPkamSwitch                  `json:"ecom_pkam_switch,omitempty"`
	BidScene                          BidScene                        `json:"bid_scene,omitempty"`
	SearchIntelligentExtension        ModelSwitch                     `json:"search_intelligent_extension,omitempty"`
	ForwardLinkAssist                 OptimizationGoal                `json:"forward_link_assist,omitempty"`
	ShortPlayPayType                  ShortPlayPayType                `json:"short_play_pay_type,omitempty"`
	SellStrategyId                    *int64                          `json:"sell_strategy_id,omitempty"`
	FeedbackId                        *int64                          `json:"feedback_id,omitempty"`
}
