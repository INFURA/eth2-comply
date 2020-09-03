# \EventsApi

All URIs are relative to *http://public-mainnet-node.ethereum.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Eventstream**](EventsApi.md#Eventstream) | **Get** /eth/v1/events | Subscribe to beacon node events



## Eventstream

> string Eventstream(ctx, topics)

Subscribe to beacon node events

Provides endpoint to subscribe to beacon node Server-Sent-Events stream. Consumers should use [eventsource](https://html.spec.whatwg.org/multipage/server-sent-events.html#the-eventsource-interface) implementation to listen on those events. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topics** | [**[]string**](string.md)| Event types to subscribe to | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

