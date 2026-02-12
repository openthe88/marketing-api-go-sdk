# TencentAds\AdvertiserConfigApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdvertiserConfigAdd**](AdvertiserConfigApi.md#AdvertiserConfigAdd) | **Post** /advertiser_config/add | 添加广告主配置
[**AdvertiserConfigGet**](AdvertiserConfigApi.md#AdvertiserConfigGet) | **Get** /advertiser_config/get | 查询广告主配置


# **AdvertiserConfigAdd**
> AdvertiserConfigAddResponse AdvertiserConfigAdd(ctx, data)
添加广告主配置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdvertiserConfigAddRequest**](AdvertiserConfigAddRequest.md)|  | 

### Return type

[**AdvertiserConfigAddResponse**](AdvertiserConfigAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdvertiserConfigGet**
> AdvertiserConfigGetResponse AdvertiserConfigGet(ctx, accountId, optional)
查询广告主配置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***AdvertiserConfigApiAdvertiserConfigGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdvertiserConfigApiAdvertiserConfigGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**AdvertiserConfigGetResponse**](AdvertiserConfigGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

