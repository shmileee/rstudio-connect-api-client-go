# \AuditLogsApi

All URIs are relative to *http://localhost/__api__*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogs**](AuditLogsApi.md#GetAuditLogs) | **Get** /v1/audit_logs | Get audit logs



## GetAuditLogs

> AuditLogs GetAuditLogs(ctx).Limit(limit).Previous(previous).Next(next).AscOrder(ascOrder).Execute()

Get audit logs



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
    limit := int32(56) // int32 | Number of logs to return. The minimum supported value is 1 and maximum supported value is 500. Note that `limit` is a \"best effort\" request since there may not be enough logs to satisfy the limit.  (optional) (default to 20)
    previous := "previous_example" // string | Gets the previous page of audit logs relative to the given id.  (optional)
    next := "next_example" // string | Gets the next page of audit logs relative to the given id.  (optional)
    ascOrder := true // bool | Whether the audit logs should be listed in ascending order  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsApi.GetAuditLogs(context.Background()).Limit(limit).Previous(previous).Next(next).AscOrder(ascOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsApi.GetAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditLogs`: AuditLogs
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsApi.GetAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of logs to return. The minimum supported value is 1 and maximum supported value is 500. Note that &#x60;limit&#x60; is a \&quot;best effort\&quot; request since there may not be enough logs to satisfy the limit.  | [default to 20]
 **previous** | **string** | Gets the previous page of audit logs relative to the given id.  | 
 **next** | **string** | Gets the next page of audit logs relative to the given id.  | 
 **ascOrder** | **bool** | Whether the audit logs should be listed in ascending order  | [default to true]

### Return type

[**AuditLogs**](AuditLogs.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

