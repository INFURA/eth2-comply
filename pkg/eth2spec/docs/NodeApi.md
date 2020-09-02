# \NodeApi

All URIs are relative to *http://public-mainnet-node.ethereum.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHealth**](NodeApi.md#GetHealth) | **Get** /eth/v1/node/health | Get health check
[**GetNetworkIdentity**](NodeApi.md#GetNetworkIdentity) | **Get** /eth/v1/node/identity | Get node network identity
[**GetNodeVersion**](NodeApi.md#GetNodeVersion) | **Get** /eth/v1/node/version | Get version string of the running beacon node.
[**GetPeer**](NodeApi.md#GetPeer) | **Get** /eth/v1/node/peers/{peer_id} | Get peer
[**GetPeers**](NodeApi.md#GetPeers) | **Get** /eth/v1/node/peers | Get node network peers
[**GetSyncingStatus**](NodeApi.md#GetSyncingStatus) | **Get** /eth/v1/node/syncing | Get node syncing status



## GetHealth

> GetHealth(ctx, )

Get health check

Returns node health status in http status codes. Useful for load balancers.

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkIdentity

> GetNetworkIdentityResponse GetNetworkIdentity(ctx, )

Get node network identity

Retrieves data about the node's network presence

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetNetworkIdentityResponse**](GetNetworkIdentityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeVersion

> GetVersionResponse GetNodeVersion(ctx, )

Get version string of the running beacon node.

Requests that the beacon node identify information about its implementation in a format similar to a  [HTTP User-Agent](https://tools.ietf.org/html/rfc7231#section-5.5.3) field.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetVersionResponse**](GetVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeer

> GetPeerResponse GetPeer(ctx, peerId)

Get peer

Retrieves data about the given peer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**peerId** | **string**|  | 

### Return type

[**GetPeerResponse**](GetPeerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeers

> GetPeersResponse GetPeers(ctx, )

Get node network peers

Retrieves data about the node's network peers

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetPeersResponse**](GetPeersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncingStatus

> GetSyncingStatusResponse GetSyncingStatus(ctx, )

Get node syncing status

Requests the beacon node to describe if it's currently syncing or not, and if it is, what block it is up to.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetSyncingStatusResponse**](GetSyncingStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

