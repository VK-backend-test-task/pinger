# \DefaultAPI

All URIs are relative to *http://backend*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PingsAggregateGet**](DefaultAPI.md#PingsAggregateGet) | **Get** /pings/aggregate | Получить информацию о контейнерах и их последнем пинге
[**PingsGet**](DefaultAPI.md#PingsGet) | **Get** /pings | Получить информацию о пингах
[**PingsPut**](DefaultAPI.md#PingsPut) | **Put** /pings | Загрузить информацию о пингах



## PingsAggregateGet

> []ContainerInfo PingsAggregateGet(ctx).OldestFirst(oldestFirst).Limit(limit).Offset(offset).Execute()

Получить информацию о контейнерах и их последнем пинге

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	oldestFirst := true // bool | Сортировать начиная со старых записей (optional)
	limit := int32(56) // int32 | Максимальное количество выдаваемых результатов (optional)
	offset := int32(56) // int32 | Количество результатов, которое необходимо пропустить (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PingsAggregateGet(context.Background()).OldestFirst(oldestFirst).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PingsAggregateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PingsAggregateGet`: []ContainerInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PingsAggregateGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPingsAggregateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oldestFirst** | **bool** | Сортировать начиная со старых записей | 
 **limit** | **int32** | Максимальное количество выдаваемых результатов | 
 **offset** | **int32** | Количество результатов, которое необходимо пропустить | 

### Return type

[**[]ContainerInfo**](ContainerInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingsGet

> []PingInfo PingsGet(ctx).ContainerIp(containerIp).OldestFirst(oldestFirst).Success(success).Limit(limit).Offset(offset).Execute()

Получить информацию о пингах

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	containerIp := "127.0.0.1" // string | IP контейнера, для которого выдавать список пингов (optional)
	oldestFirst := true // bool | Сортировать начиная со старых записей (optional)
	success := true // bool | Возвращать только успешные или неудавшиеся записи (optional)
	limit := int32(56) // int32 | Максимальное количество выдаваемых результатов (optional)
	offset := int32(56) // int32 | Количество результатов, которое необходимо пропустить (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PingsGet(context.Background()).ContainerIp(containerIp).OldestFirst(oldestFirst).Success(success).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PingsGet`: []PingInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerIp** | **string** | IP контейнера, для которого выдавать список пингов | 
 **oldestFirst** | **bool** | Сортировать начиная со старых записей | 
 **success** | **bool** | Возвращать только успешные или неудавшиеся записи | 
 **limit** | **int32** | Максимальное количество выдаваемых результатов | 
 **offset** | **int32** | Количество результатов, которое необходимо пропустить | 

### Return type

[**[]PingInfo**](PingInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingsPut

> PingsPut(ctx).PingInfo(pingInfo).Execute()

Загрузить информацию о пингах

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pingInfo := []openapiclient.PingInfo{*openapiclient.NewPingInfo("ContainerIp_example", time.Now(), false)} // []PingInfo | Список результатов пинга

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PingsPut(context.Background()).PingInfo(pingInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pingInfo** | [**[]PingInfo**](PingInfo.md) | Список результатов пинга | 

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

