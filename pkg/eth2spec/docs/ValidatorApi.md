# \ValidatorApi

All URIs are relative to *http://public-mainnet-node.ethereum.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregatedAttestation**](ValidatorApi.md#GetAggregatedAttestation) | **Get** /eth/v1/validator/aggregate_attestation | Get aggregated attestation
[**GetAttesterDuties**](ValidatorApi.md#GetAttesterDuties) | **Get** /eth/v1/validator/duties/attester/{epoch} | Get attester duties
[**GetProposerDuties**](ValidatorApi.md#GetProposerDuties) | **Get** /eth/v1/validator/duties/proposer/{epoch} | Get block proposers duties
[**PrepareBeaconCommitteeSubnet**](ValidatorApi.md#PrepareBeaconCommitteeSubnet) | **Post** /eth/v1/validator/beacon_committee_subscriptions | Signal beacon node to prepare for a committee subnet
[**ProduceAttestationData**](ValidatorApi.md#ProduceAttestationData) | **Get** /eth/v1/validator/attestation_data | Produce an attestation data
[**ProduceBlock**](ValidatorApi.md#ProduceBlock) | **Get** /eth/v1/validator/blocks/{slot} | Produce a new block, without signature.
[**PublishAggregateAndProof**](ValidatorApi.md#PublishAggregateAndProof) | **Post** /eth/v1/validator/aggregate_and_proofs | Publish aggregate and proof



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

