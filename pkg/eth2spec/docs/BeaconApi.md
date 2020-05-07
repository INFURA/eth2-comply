# \BeaconApi

All URIs are relative to *http://public-mainnet-node.ethereum.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlock**](BeaconApi.md#GetBlock) | **Get** /v1/beacon/blocks/{block_id} | Get block
[**GetBlockAttestations**](BeaconApi.md#GetBlockAttestations) | **Get** /v1/beacon/blocks/{block_id}/attestations | Get block attestations
[**GetBlockHeader**](BeaconApi.md#GetBlockHeader) | **Get** /v1/beacon/headers/{block_id} | Get block header
[**GetBlockHeaders**](BeaconApi.md#GetBlockHeaders) | **Get** /v1/beacon/headers | Get block headers
[**GetBlockRoot**](BeaconApi.md#GetBlockRoot) | **Get** /v1/beacon/blocks/{block_id}/root | Get block root
[**GetEpochCommittees**](BeaconApi.md#GetEpochCommittees) | **Get** /v1/beacon/states/{state_id}/committees/{epoch} | Get all committees for epoch
[**GetGenesis**](BeaconApi.md#GetGenesis) | **Get** /v1/beacon/genesis | Retrieve details of the chain&#39;s genesis.
[**GetPoolAttestations**](BeaconApi.md#GetPoolAttestations) | **Get** /v1/beacon/pool/attestations | Get Attestations from operations pool
[**GetPoolAttesterSlashings**](BeaconApi.md#GetPoolAttesterSlashings) | **Get** /v1/beacon/pool/atttester_slashings | Get AttesterSlashings from operations pool
[**GetPoolProposerSlashings**](BeaconApi.md#GetPoolProposerSlashings) | **Get** /v1/beacon/pool/proposer_slashings | Get ProposerSlashings from operations pool
[**GetPoolVoluntaryExits**](BeaconApi.md#GetPoolVoluntaryExits) | **Get** /v1/beacon/pool/voluntary_exits | Get SignedVoluntaryExit from operations pool
[**GetStateFinalityCheckpoints**](BeaconApi.md#GetStateFinalityCheckpoints) | **Get** /v1/beacon/states/{state_id}/finality_checkpoints | Get state finality checkpoints
[**GetStateFork**](BeaconApi.md#GetStateFork) | **Get** /v1/beacon/states/{state_id}/fork | Get Fork object for requested state
[**GetStateRoot**](BeaconApi.md#GetStateRoot) | **Get** /v1/beacon/states/{state_id}/root | Get state SSZ HashTreeRoot
[**GetStateValidator**](BeaconApi.md#GetStateValidator) | **Get** /v1/beacon/states/{state_id}/validators/{validator_id} | Get validator from state by id
[**GetStateValidators**](BeaconApi.md#GetStateValidators) | **Get** /v1/beacon/states/{state_id}/validators | Get validators from state
[**SubmitPoolAttestations**](BeaconApi.md#SubmitPoolAttestations) | **Post** /v1/beacon/pool/attestations | Submit Attestation object to node
[**SubmitPoolAttesterSlashings**](BeaconApi.md#SubmitPoolAttesterSlashings) | **Post** /v1/beacon/pool/atttester_slashings | Submit AttesterSlashing object to node&#39;s pool
[**SubmitPoolProposerSlashings**](BeaconApi.md#SubmitPoolProposerSlashings) | **Post** /v1/beacon/pool/proposer_slashings | Submit ProposerSlashing object to node&#39;s pool
[**SubmitPoolVoluntaryExit**](BeaconApi.md#SubmitPoolVoluntaryExit) | **Post** /v1/beacon/pool/voluntary_exits | Submit SignedVoluntaryExit object to node&#39;s pool



## GetBlock

> InlineResponse2009 GetBlock(ctx, blockId)

Get block

Retrieves block details for given block id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string**| Block identifier. Can be one of: \&quot;head\&quot; (canonical head in node&#39;s view), \&quot;genesis\&quot;, \&quot;finalized\&quot;, \\&lt;slot\\&gt;, \\&lt;hex encoded blockRoot with 0x prefix\\&gt;.  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockAttestations

> InlineResponse20011 GetBlockAttestations(ctx, blockId)

Get block attestations

Retrieves attestation included in requested block.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string**| Block identifier. Can be one of: \&quot;head\&quot; (canonical head in node&#39;s view), \&quot;genesis\&quot;, \&quot;finalized\&quot;, \\&lt;slot\\&gt;, \\&lt;hex encoded blockRoot with 0x prefix\\&gt;.  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockHeader

> InlineResponse2008 GetBlockHeader(ctx, blockId)

Get block header

Retrieves block header for given block id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string**| Block identifier. Can be one of: \&quot;head\&quot; (canonical head in node&#39;s view), \&quot;genesis\&quot;, \&quot;finalized\&quot;, \\&lt;slot\\&gt;, \\&lt;hex encoded blockRoot with 0x prefix\\&gt;.  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockHeaders

> InlineResponse2007 GetBlockHeaders(ctx, optional)

Get block headers

Retrieves block headers matching given query. By default it will fetch current head slot blocks.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBlockHeadersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBlockHeadersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slot** | [**optional.Interface of string**](.md)|  | 
 **parentRoot** | [**optional.Interface of string**](.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockRoot

> InlineResponse20010 GetBlockRoot(ctx, blockId)

Get block root

Retrieves hashTreeRoot of BeaconBlock/BeaconBlockHeader

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string**| Block identifier. Can be one of: \&quot;head\&quot; (canonical head in node&#39;s view), \&quot;genesis\&quot;, \&quot;finalized\&quot;, \\&lt;slot\\&gt;, \\&lt;hex encoded blockRoot with 0x prefix\\&gt;.  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEpochCommittees

> InlineResponse2006 GetEpochCommittees(ctx, stateId, epoch, optional)

Get all committees for epoch

Retrieves the committees for the given state at the given epoch.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateId** | **string**| State identifier. Can be one of: \&quot;head\&quot; (canonical head in node&#39;s view), \&quot;genesis\&quot;, \&quot;finalized\&quot;, \&quot;justified\&quot;, \\&lt;slot\\&gt;, \\&lt;hex encoded stateRoot with 0x prefix\\&gt;.  | 
**epoch** | [**string**](.md)| Epoch for which to calculate committees. Defaults to beacon state epoch. | 
 **optional** | ***GetEpochCommitteesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEpochCommitteesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **index** | [**optional.Interface of string**](.md)| Committee index | 
 **slot** | [**optional.Interface of string**](.md)|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## GetPoolAttestations

> InlineResponse20011 GetPoolAttestations(ctx, optional)

Get Attestations from operations pool

Retrieves attestations known by the node but not necessarily incorporated into any block

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPoolAttestationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPoolAttestationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slot** | **optional.String**|  | 
 **committeeIndex** | **optional.String**|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolAttesterSlashings

> InlineResponse20012 GetPoolAttesterSlashings(ctx, )

Get AttesterSlashings from operations pool

Retrieves attester slashings known by the node but not necessarily incorporated into any block

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolProposerSlashings

> InlineResponse20013 GetPoolProposerSlashings(ctx, )

Get ProposerSlashings from operations pool

Retrieves proposer slashings known by the node but not necessarily incorporated into any block

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolVoluntaryExits

> InlineResponse20014 GetPoolVoluntaryExits(ctx, )

Get SignedVoluntaryExit from operations pool

Retrieves voluntary exits known by the node but not necessarily incorporated into any block

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStateFinalityCheckpoints

> InlineResponse2003 GetStateFinalityCheckpoints(ctx, stateId)

Get state finality checkpoints

Returns finality checkpoints for state with given 'stateId'. In case finality is not yet achieved, checkpoint should return epoch 0 and ZERO_HASH as root. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateId** | **string**| State identifier. Can be one of: \&quot;head\&quot; (canonical head in node&#39;s view), \&quot;genesis\&quot;, \&quot;finalized\&quot;, \&quot;justified\&quot;, \\&lt;slot\\&gt;, \\&lt;hex encoded stateRoot with 0x prefix\\&gt;.  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

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


## GetStateRoot

> InlineResponse2001 GetStateRoot(ctx, stateId)

Get state SSZ HashTreeRoot

Calculates HashTreeRoot for state with given 'stateId'. If stateId is root, same value will be returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateId** | **string**| State identifier. Can be one of: \&quot;head\&quot; (canonical head in node&#39;s view), \&quot;genesis\&quot;, \&quot;finalized\&quot;, \&quot;justified\&quot;, \\&lt;slot\\&gt;, \\&lt;hex encoded stateRoot with 0x prefix\\&gt;.  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

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


## GetStateValidators

> InlineResponse2004 GetStateValidators(ctx, stateId, optional)

Get validators from state

Returns filterable list of validators with their balance, status and index.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateId** | **string**| State identifier. Can be one of: \&quot;head\&quot; (canonical head in node&#39;s view), \&quot;genesis\&quot;, \&quot;finalized\&quot;, \&quot;justified\&quot;, \\&lt;slot\\&gt;, \\&lt;hex encoded stateRoot with 0x prefix\\&gt;.  | 
 **optional** | ***GetStateValidatorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetStateValidatorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | [**optional.Interface of []string**](string.md)| Either hex encoded public key (with 0x prefix) or validator index | 
 **status** | [**optional.Interface of []string**](string.md)| [Validator status specification](https://hackmd.io/ofFJ5gOmQpu1jjHilHbdQQ) | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

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


## SubmitPoolAttesterSlashings

> SubmitPoolAttesterSlashings(ctx, inlineObject1)

Submit AttesterSlashing object to node's pool

Submits AttesterSlashing object to node's pool and if passes validation node MUST broadcast it to network.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject1** | [**InlineObject1**](InlineObject1.md)|  | 

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


## SubmitPoolProposerSlashings

> SubmitPoolProposerSlashings(ctx, inlineObject2)

Submit ProposerSlashing object to node's pool

Submits ProposerSlashing object to node's pool and if passes validation  node MUST broadcast it to network.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject2** | [**InlineObject2**](InlineObject2.md)|  | 

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


## SubmitPoolVoluntaryExit

> SubmitPoolVoluntaryExit(ctx, inlineObject3)

Submit SignedVoluntaryExit object to node's pool

Submits SignedVoluntaryExit object to node's pool and if passes validation node MUST broadcast it to network.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject3** | [**InlineObject3**](InlineObject3.md)|  | 

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

