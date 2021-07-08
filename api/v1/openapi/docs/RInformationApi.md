# \RInformationApi

All URIs are relative to *http://localhost/__api__*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRInformation**](RInformationApi.md#GetRInformation) | **Get** /v1/server_settings/r | Get R Information



## GetRInformation

> RInstallations GetRInformation(ctx).Execute()

Get R Information



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RInformationApi.GetRInformation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RInformationApi.GetRInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRInformation`: RInstallations
    fmt.Fprintf(os.Stdout, "Response from `RInformationApi.GetRInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRInformationRequest struct via the builder pattern


### Return type

[**RInstallations**](RInstallations.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

