/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// CityLevel : 城市级别，仅当 type=REGION、BUSINESS_DISTRICT 时有效
type CityLevel string

// List of CityLevel
const (
	CityLevel_FIRST  CityLevel = "CITY_LEVEL_FIRST"
	CityLevel_SECOND CityLevel = "CITY_LEVEL_SECOND"
	CityLevel_THIRD  CityLevel = "CITY_LEVEL_THIRD"
	CityLevel_NONE   CityLevel = "CITY_LEVEL_NONE"
)
