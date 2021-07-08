# \TasksApi

All URIs are relative to *http://localhost/__api__*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTask**](TasksApi.md#GetTask) | **Get** /v1/experimental/tasks/{id} | Get task details



## GetTask

> Task GetTask(ctx, id).First(first).Wait(wait).Execute()

Get task details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The identifier of the desired task. 
    first := int32(56) // int32 | The first line of output to include in the response. The value `0` indicates that all lines should be returned.  Values less than `0` are not permitted. Values greater than the total number of lines produced by the task are not permitted.  (optional) (default to 0)
    wait := int32(56) // int32 | The number of seconds to wait for task completion before responding. The current state of the task is returned once the task completes or this time elapses.  Values less than `0` or greater than `20` are not permitted.  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TasksApi.GetTask(context.Background(), id).First(first).Wait(wait).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GetTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTask`: Task
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identifier of the desired task.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **first** | **int32** | The first line of output to include in the response. The value &#x60;0&#x60; indicates that all lines should be returned.  Values less than &#x60;0&#x60; are not permitted. Values greater than the total number of lines produced by the task are not permitted.  | [default to 0]
 **wait** | **int32** | The number of seconds to wait for task completion before responding. The current state of the task is returned once the task completes or this time elapses.  Values less than &#x60;0&#x60; or greater than &#x60;20&#x60; are not permitted.  | [default to 0]

### Return type

[**Task**](Task.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

