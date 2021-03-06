# \ValidatorRequiredApiApi

All URIs are relative to *http://public-mainnet-node.ethereum.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregatedAttestation**](ValidatorRequiredApiApi.md#GetAggregatedAttestation) | **Get** /eth/v1/validator/aggregate_attestation | Get aggregated attestation
[**GetAttesterDuties**](ValidatorRequiredApiApi.md#GetAttesterDuties) | **Get** /eth/v1/validator/duties/attester/{epoch} | Get attester duties
[**GetGenesis**](ValidatorRequiredApiApi.md#GetGenesis) | **Get** /eth/v1/beacon/genesis | Retrieve details of the chain&#39;s genesis.
[**GetProposerDuties**](ValidatorRequiredApiApi.md#GetProposerDuties) | **Get** /eth/v1/validator/duties/proposer/{epoch} | Get block proposers duties
[**GetStateFork**](ValidatorRequiredApiApi.md#GetStateFork) | **Get** /eth/v1/beacon/states/{state_id}/fork | Get Fork object for requested state
[**GetStateValidator**](ValidatorRequiredApiApi.md#GetStateValidator) | **Get** /eth/v1/beacon/states/{state_id}/validators/{validator_id} | Get validator from state by id
[**GetSyncingStatus**](ValidatorRequiredApiApi.md#GetSyncingStatus) | **Get** /eth/v1/node/syncing | Get node syncing status
[**PrepareBeaconCommitteeSubnet**](ValidatorRequiredApiApi.md#PrepareBeaconCommitteeSubnet) | **Post** /eth/v1/validator/beacon_committee_subscriptions | Signal beacon node to prepare for a committee subnet
[**ProduceAttestationData**](ValidatorRequiredApiApi.md#ProduceAttestationData) | **Get** /eth/v1/validator/attestation_data | Produce an attestation data
[**ProduceBlock**](ValidatorRequiredApiApi.md#ProduceBlock) | **Get** /eth/v1/validator/blocks/{slot} | Produce a new block, without signature.
[**PublishAggregateAndProof**](ValidatorRequiredApiApi.md#PublishAggregateAndProof) | **Post** /eth/v1/validator/aggregate_and_proofs | Publish aggregate and proof
[**PublishBlock**](ValidatorRequiredApiApi.md#PublishBlock) | **Post** /eth/v1/beacon/blocks | Publish a signed block.
[**SubmitPoolAttestations**](ValidatorRequiredApiApi.md#SubmitPoolAttestations) | **Post** /eth/v1/beacon/pool/attestations | Submit Attestation object to node



## GetAggregatedAttestation

> GetAggregatedAttestationResponse GetAggregatedAttestation(ctx, attestationDataRoot, slot)

Get aggregated attestation

Aggregates all attestations matching given attestation data root and slot

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attestationDataRoot** | **string**| HashTreeRoot of AttestationData that validator want&#39;s aggregated | 
**slot** | **string**|  | 

### Return type

[**GetAggregatedAttestationResponse**](GetAggregatedAttestationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttesterDuties

> GetAttesterDutiesResponse GetAttesterDuties(ctx, epoch, index)

Get attester duties

Requests the beacon node to provide a set of attestation duties, which should be performed by validators, for a particular epoch. Duties should only need to be checked once per epoch, however a chain reorganization (of > MIN_SEED_LOOKAHEAD epochs) could occur, resulting in a change of duties. For full safety, you should monitor chain reorganizations events.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epoch** | **string**| Should only be allowed 1 epoch ahead | 
**index** | [**[]string**](string.md)| Validator index | 

### Return type

[**GetAttesterDutiesResponse**](GetAttesterDutiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## GetProposerDuties

> GetProposerDutiesResponse GetProposerDuties(ctx, epoch)

Get block proposers duties

Request beacon node to provide all validators that are suppose to propose block in given epoch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epoch** | **string**|  | 

### Return type

[**GetProposerDutiesResponse**](GetProposerDutiesResponse.md)

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


## PrepareBeaconCommitteeSubnet

> PrepareBeaconCommitteeSubnet(ctx, optional)

Signal beacon node to prepare for a committee subnet

After beacon node receives this request, search using discv5 for peers related to this subnet and replace current peers with those ones if necessary If validator `is_aggregator`, beacon node must: - announce subnet topic subscription on gossipsub - aggregate attestations received on that subnet 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrepareBeaconCommitteeSubnetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PrepareBeaconCommitteeSubnetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**optional.Interface of []InlineObject**](InlineObject.md)|  | 

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


## ProduceAttestationData

> ProduceAttestationDataResponse ProduceAttestationData(ctx, slot, committeeIndex)

Produce an attestation data

Requests that the beacon node produce an AttestationData.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slot** | **string**| The slot for which an attestation data should be created. | 
**committeeIndex** | **string**| The committee index for which an attestation data should be created. | 

### Return type

[**ProduceAttestationDataResponse**](ProduceAttestationDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProduceBlock

> ProduceBlockResponse ProduceBlock(ctx, slot, randaoReveal, optional)

Produce a new block, without signature.

Requests a beacon node to produce a valid block, which can then be signed by a validator.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slot** | **string**| The slot for which the block should be proposed. | 
**randaoReveal** | **string**| The validator&#39;s randao reveal value. | 
 **optional** | ***ProduceBlockOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProduceBlockOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **graffiti** | **optional.String**| Arbitrary data validator wants to include in block. | 

### Return type

[**ProduceBlockResponse**](ProduceBlockResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishAggregateAndProof

> PublishAggregateAndProof(ctx, optional)

Publish aggregate and proof

Verifies given aggregate and proof and publishes it on appropriate gossipsub topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublishAggregateAndProofOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PublishAggregateAndProofOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject5** | [**optional.Interface of InlineObject5**](InlineObject5.md)|  | 

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


## SubmitPoolAttestations

> SubmitPoolAttestations(ctx, inlineObject1)

Submit Attestation object to node

Submits Attestation object to node. If attestation passes all validation constraints, node MUST publish attestation on appropriate subnet.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

