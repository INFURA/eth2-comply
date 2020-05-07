# \MinimalValidatorApiApi

All URIs are relative to *http://public-mainnet-node.ethereum.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGenesis**](MinimalValidatorApiApi.md#GetGenesis) | **Get** /v1/beacon/genesis | Retrieve details of the chain&#39;s genesis.
[**GetStateFork**](MinimalValidatorApiApi.md#GetStateFork) | **Get** /v1/beacon/states/{state_id}/fork | Get Fork object for requested state
[**GetStateValidator**](MinimalValidatorApiApi.md#GetStateValidator) | **Get** /v1/beacon/states/{state_id}/validators/{validator_id} | Get validator from state by id
[**GetSyncingStatus**](MinimalValidatorApiApi.md#GetSyncingStatus) | **Get** /v1/node/syncing | Get node syncing status
[**SubmitPoolAttestations**](MinimalValidatorApiApi.md#SubmitPoolAttestations) | **Post** /v1/beacon/pool/attestations | Submit Attestation object to node



## GetGenesis

> InlineResponse200 GetGenesis(ctx, )

Retrieve details of the chain's genesis.

Retrieve details of the chain's genesis which can be used to identify chain.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStateFork

> InlineResponse2002 GetStateFork(ctx, stateId)

Get Fork object for requested state

Returns [Fork](https://github.com/ethereum/eth2.0-specs/blob/v0.11.1/specs/phase0/beacon-chain.md#fork) object for state with given 'stateId'.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateId** | **string**| State identifier. Can be one of: \&quot;head\&quot; (canonical head in node&#39;s view), \&quot;genesis\&quot;, \&quot;finalized\&quot;, \&quot;justified\&quot;, \\&lt;slot\\&gt;, \\&lt;hex encoded stateRoot with 0x prefix\\&gt;.  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStateValidator

> InlineResponse2005 GetStateValidator(ctx, stateId, validatorId)

Get validator from state by id

Returns validator specified by state and id or public key along with status and balance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateId** | **string**| State identifier. Can be one of: \&quot;head\&quot; (canonical head in node&#39;s view), \&quot;genesis\&quot;, \&quot;finalized\&quot;, \&quot;justified\&quot;, \\&lt;slot\\&gt;, \\&lt;hex encoded stateRoot with 0x prefix\\&gt;.  | 
**validatorId** | **string**| Either hex encoded public key (with 0x prefix) or validator index | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncingStatus

> InlineResponse20019 GetSyncingStatus(ctx, )

Get node syncing status

Requests the beacon node to describe if it's currently syncing or not, and if it is, what block it is up to.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitPoolAttestations

> SubmitPoolAttestations(ctx, inlineObject)

Submit Attestation object to node

Submits Attestation object to node. If attestation passes all validation constraints, node MUST publish attestation on appropriate subnet.

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

