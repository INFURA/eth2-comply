# \ConfigApi

All URIs are relative to *http://public-mainnet-node.ethereum.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDepositContract**](ConfigApi.md#GetDepositContract) | **Get** /eth/v1/config/deposit_contract | Get deposit contract address.
[**GetForkSchedule**](ConfigApi.md#GetForkSchedule) | **Get** /eth/v1/config/fork_schedule | Get scheduled upcoming forks.
[**GetSpec**](ConfigApi.md#GetSpec) | **Get** /eth/v1/config/spec | Get spec params.



## GetDepositContract

> GetDepositContractResponse GetDepositContract(ctx, )

Get deposit contract address.

Retrieve deposit contract address and genesis fork version.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetDepositContractResponse**](GetDepositContractResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForkSchedule

> GetForkScheduleResponse GetForkSchedule(ctx, )

Get scheduled upcoming forks.

Retrieve all scheduled upcoming forks this node is aware of.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetForkScheduleResponse**](GetForkScheduleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpec

> GetSpecResponse GetSpec(ctx, )

Get spec params.

Retrieve specification configuration (without Phase 1 params) used on this node. [Specification params list](https://github.com/ethereum/eth2.0-specs/blob/v0.11.1/configs/mainnet.yaml)  Values are returned with following format:   - any value starting with 0x in the spec is returned as a hex string   - all other values are returned as number 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetSpecResponse**](GetSpecResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

