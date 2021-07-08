# \BundlesApi

All URIs are relative to *http://localhost/__api__*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBundle**](BundlesApi.md#DeleteBundle) | **Delete** /v1/experimental/bundles/{id} | Delete bundle
[**DownloadBundle**](BundlesApi.md#DownloadBundle) | **Get** /v1/experimental/bundles/{id}/download | Download the bundle archive
[**GetBundle**](BundlesApi.md#GetBundle) | **Get** /v1/experimental/bundles/{id} | Get bundle details
[**GetBundles**](BundlesApi.md#GetBundles) | **Get** /v1/experimental/content/{guid}/bundles | List bundles



## DeleteBundle

> DeleteBundle(ctx, id).Execute()

Delete bundle



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
    id := "id_example" // string | The unique identifier of the desired bundle. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BundlesApi.DeleteBundle(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.DeleteBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the desired bundle.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadBundle

> *os.File DownloadBundle(ctx, id).Execute()

Download the bundle archive



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
    id := "id_example" // string | The unique identifier of the desired bundle. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BundlesApi.DownloadBundle(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.DownloadBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadBundle`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `BundlesApi.DownloadBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the desired bundle.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBundle

> Bundle GetBundle(ctx, id).Execute()

Get bundle details



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
    id := "id_example" // string | The unique identifier of the desired bundle. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BundlesApi.GetBundle(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.GetBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBundle`: Bundle
    fmt.Fprintf(os.Stdout, "Response from `BundlesApi.GetBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the desired bundle.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Bundle**](Bundle.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBundles

> Bundles GetBundles(ctx, guid).PageNumber(pageNumber).PageSize(pageSize).Execute()

List bundles



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
    guid := TODO // string | The unique identifier of the desired content item. 
    pageNumber := int32(56) // int32 | The page to return relative to the given `page_size`. If `page_number` is 0 or negative, an error will be returned.  (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to include in each page. This parameter is \"best effort\" since there may not be enough bundles to honor the request. If `page_size` is less than 1 or greater than 500, an error will be returned.  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BundlesApi.GetBundles(context.Background(), guid).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.GetBundles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBundles`: Bundles
    fmt.Fprintf(os.Stdout, "Response from `BundlesApi.GetBundles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | [**string**](.md) | The unique identifier of the desired content item.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page to return relative to the given &#x60;page_size&#x60;. If &#x60;page_number&#x60; is 0 or negative, an error will be returned.  | [default to 1]
 **pageSize** | **int32** | The number of items to include in each page. This parameter is \&quot;best effort\&quot; since there may not be enough bundles to honor the request. If &#x60;page_size&#x60; is less than 1 or greater than 500, an error will be returned.  | [default to 20]

### Return type

[**Bundles**](Bundles.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

