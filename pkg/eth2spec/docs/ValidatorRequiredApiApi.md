# \ValidatorRequiredApiApi

All URIs are relative to *http://public-mainnet-node.ethereum.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGenesis**](ValidatorRequiredApiApi.md#GetGenesis) | **Get** /eth/v1/beacon/genesis | Retrieve details of the chain&#39;s genesis.
[**GetStateFork**](ValidatorRequiredApiApi.md#GetStateFork) | **Get** /eth/v1/beacon/states/{state_id}/fork | Get Fork object for requested state
[**GetStateValidator**](ValidatorRequiredApiApi.md#GetStateValidator) | **Get** /eth/v1/beacon/states/{state_id}/validators/{validator_id} | Get validator from state by id
[**GetSyncingStatus**](ValidatorRequiredApiApi.md#GetSyncingStatus) | **Get** /eth/v1/node/syncing | Get node syncing status
[**PublishBlock**](ValidatorRequiredApiApi.md#PublishBlock) | **Post** /eth/v1/beacon/blocks | Publish a signed block.



## GetGenesis

> GetGenesisResponse GetGenesis(ctx, )

Retrieve details of the chain's genesis.

Retrieve details of the chain's genesis which can be used to identify chain.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetGenesisResponse**](GetGenesisResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStateFork

> GetStateForkResponse GetStateFork(ctx, stateId)

Get Fork object for requested state

Returns [Fork](https://github.com/ethereum/eth2.0-specs/blob/v0.11.1/specs/phase0/beacon-chain.md#fork) object for state with given 'stateId'.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateId** | **string**| State identifier. Can be one of: \&quot;head\&quot; (canonical head in node&#39;s view), \&quot;genesis\&quot;, \&quot;finalized\&quot;, \&quot;justified\&quot;, \\&lt;slot\\&gt;, \\&lt;hex encoded stateRoot with 0x prefix\\&gt;.  | 

### Return type

[**GetStateForkResponse**](GetStateForkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStateValidator

> GetStateValidatorResponse GetStateValidator(ctx, stateId, validatorId)

Get validator from state by id

Returns validator specified by state and id or public key along with status and balance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateId** | **string**| State identifier. Can be one of: \&quot;head\&quot; (canonical head in node&#39;s view), \&quot;genesis\&quot;, \&quot;finalized\&quot;, \&quot;justified\&quot;, \\&lt;slot\\&gt;, \\&lt;hex encoded stateRoot with 0x prefix\\&gt;.  | 
**validatorId** | **string**| Either hex encoded public key (with 0x prefix) or validator index | 

### Return type

[**GetStateValidatorResponse**](GetStateValidatorResponse.md)

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


## PublishBlock

> PublishBlock(ctx, inlineObject)

Publish a signed block.

Instructs the beacon node to broadcast a newly signed beacon block to the beacon network, to be included in the beacon chain. The beacon node is not required to validate the signed `BeaconBlock`, and a successful response (20X) only indicates that the broadcast has been successful. The beacon node is expected to integrate the new block into its state, and therefore validate the block internally, however blocks which fail the validation are still broadcast but a different status code is returned (202)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject** | [**InlineObject**](InlineObject.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

