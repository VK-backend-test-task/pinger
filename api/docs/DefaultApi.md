# {{classname}}

All URIs are relative to *http://backend*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PingsAggregateGet**](DefaultApi.md#PingsAggregateGet) | **Get** /pings/aggregate | Получить информацию о контейнерах и их последнем пинге
[**PingsGet**](DefaultApi.md#PingsGet) | **Get** /pings | Получить информацию о пингах
[**PingsPut**](DefaultApi.md#PingsPut) | **Put** /pings | Загрузить информацию о пингах

# **PingsAggregateGet**
> []ContainerInfo PingsAggregateGet(ctx, optional)
Получить информацию о контейнерах и их последнем пинге

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiPingsAggregateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiPingsAggregateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oldestFirst** | **optional.Bool**| Сортировать начиная со старых записей | 
 **limit** | **optional.Int32**| Максимальное количество выдаваемых результатов | 
 **offset** | **optional.Int32**| Количество результатов, которое необходимо пропустить | 

### Return type

[**[]ContainerInfo**](ContainerInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PingsGet**
> []PingInfo PingsGet(ctx, optional)
Получить информацию о пингах

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiPingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiPingsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerIp** | **optional.String**| IP контейнера, для которого выдавать список пингов | 
 **oldestFirst** | **optional.Bool**| Сортировать начиная со старых записей | 
 **success** | **optional.Bool**| Возвращать только успешные или неудавшиеся записи | 
 **limit** | **optional.Int32**| Максимальное количество выдаваемых результатов | 
 **offset** | **optional.Int32**| Количество результатов, которое необходимо пропустить | 

### Return type

[**[]PingInfo**](PingInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PingsPut**
> PingsPut(ctx, body)
Загрузить информацию о пингах

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PingInfo**](PingInfo.md)| Список результатов пинга | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

