/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 链路og对象
type ConversionLinkOgItem struct {
	Og                OptimizationGoal        `json:"og,omitempty"`
	DeepOg            OptimizationGoal        `json:"deep_og,omitempty"`
	DeepRoi           DeepConversionWorthGoal `json:"deep_roi,omitempty"`
	AdvancedRoi       DeepConversionWorthGoal `json:"advanced_roi,omitempty"`
	AdvancedOg        OptimizationGoal        `json:"advanced_og,omitempty"`
	ForwardLinkAssist OptimizationGoal        `json:"forward_link_assist,omitempty"`
}
