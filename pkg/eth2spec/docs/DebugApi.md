# \DebugApi

All URIs are relative to *http://public-mainnet-node.ethereum.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDebugChainHeads**](DebugApi.md#GetDebugChainHeads) | **Get** /eth/v1/debug/beacon/heads | Get fork choice leaves
[**GetState**](DebugApi.md#GetState) | **Get** /eth/v1/debug/beacon/states/{state_id} | Get full BeaconState object



## GetDebugChainHeads

> GetDebugChainHeadsResponse GetDebugChainHeads(ctx, )

Get fork choice leaves

Retrieves all possible chain heads (leaves of fork choice tree).

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetDebugChainHeadsResponse**](GetDebugChainHeadsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetState

> GetStateResponse GetState(ctx, stateId)

Get full BeaconState object

Returns full BeaconState object for given stateId.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateId** | **string**| State identifier. Can be one of: \&quot;head\&quot; (canonical head in node&#39;s view), \&quot;genesis\&quot;, \&quot;finalized\&quot;, \&quot;justified\&quot;, \\&lt;slot\\&gt;, \\&lt;hex encoded stateRoot with 0x prefix\\&gt;.  | 

### Return type

[**GetStateResponse**](GetStateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

