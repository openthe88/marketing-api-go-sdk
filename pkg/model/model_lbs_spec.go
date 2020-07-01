/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// LBS人群信息
type LbsSpec struct {
	LbsType            LbsType            `json:"lbs_type,omitempty"`
	CrossCityRule      CrossCityRule      `json:"cross_city_rule,omitempty"`
	PoiRule            PoiRule            `json:"poi_rule,omitempty"`
	CustomLocationRule CustomLocationRule `json:"custom_location_rule,omitempty"`
}
