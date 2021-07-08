# \ContentApi

All URIs are relative to *http://localhost/__api__*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContent**](ContentApi.md#CreateContent) | **Post** /v1/experimental/content | Create content item
[**DeleteContent**](ContentApi.md#DeleteContent) | **Delete** /v1/experimental/content/{guid} | Delete content
[**DeployContentBundle**](ContentApi.md#DeployContentBundle) | **Post** /v1/experimental/content/{guid}/deploy | Deploy deployment bundle
[**GetContent**](ContentApi.md#GetContent) | **Get** /v1/experimental/content/{guid} | Get content details
[**UpdateContent**](ContentApi.md#UpdateContent) | **Post** /v1/experimental/content/{guid} | Update content
[**UploadContentBundle**](ContentApi.md#UploadContentBundle) | **Post** /v1/experimental/content/{guid}/upload | Upload deployment bundle



## CreateContent

> Content CreateContent(ctx).Content(content).Execute()

Create content item



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
    content := *openapiclient.NewContent("quarterly-analysis") // Content | The request body required when creating a new content item. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentApi.CreateContent(context.Background()).Content(content).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.CreateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContent`: Content
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.CreateContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content** | [**Content**](Content.md) | The request body required when creating a new content item.  | 

### Return type

[**Content**](Content.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContent

> DeleteContent(ctx, guid).Execute()

Delete content



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentApi.DeleteContent(context.Background(), guid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.DeleteContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | [**string**](.md) | The unique identifier of the desired content item.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContentRequest struct via the builder pattern


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


## DeployContentBundle

> ContentDeploymentTask DeployContentBundle(ctx, guid).Instructions(instructions).Execute()

Deploy deployment bundle



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
    instructions := *openapiclient.NewContentDeploymentInstructions() // ContentDeploymentInstructions | The request body when requesting the deployment of a bundle.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentApi.DeployContentBundle(context.Background(), guid).Instructions(instructions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.DeployContentBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployContentBundle`: ContentDeploymentTask
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.DeployContentBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | [**string**](.md) | The unique identifier of the desired content item.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployContentBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instructions** | [**ContentDeploymentInstructions**](ContentDeploymentInstructions.md) | The request body when requesting the deployment of a bundle.  | 

### Return type

[**ContentDeploymentTask**](ContentDeploymentTask.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContent

> Content GetContent(ctx, guid).Execute()

Get content details



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentApi.GetContent(context.Background(), guid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.GetContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContent`: Content
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.GetContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | [**string**](.md) | The unique identifier of the desired content item.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Content**](Content.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContent

> Content UpdateContent(ctx, guid).Content(content).Execute()

Update content



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
    content := *openapiclient.NewContent("quarterly-analysis") // Content | The request body required when creating a new content item. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentApi.UpdateContent(context.Background(), guid).Content(content).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.UpdateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContent`: Content
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.UpdateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | [**string**](.md) | The unique identifier of the desired content item.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | [**Content**](Content.md) | The request body required when creating a new content item.  | 

### Return type

[**Content**](Content.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadContentBundle

> ContentUploadBundle UploadContentBundle(ctx, guid).Archive(archive).XContentChecksum(xContentChecksum).Execute()

Upload deployment bundle



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
    archive := os.NewFile(1234, "some_file") // *os.File | A `gzip` compressed `tar` archive file. 
    xContentChecksum := "xContentChecksum_example" // string | The MD5 sum of the archive file.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentApi.UploadContentBundle(context.Background(), guid).Archive(archive).XContentChecksum(xContentChecksum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentApi.UploadContentBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadContentBundle`: ContentUploadBundle
    fmt.Fprintf(os.Stdout, "Response from `ContentApi.UploadContentBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | [**string**](.md) | The unique identifier of the desired content item.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadContentBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archive** | ***os.File** | A &#x60;gzip&#x60; compressed &#x60;tar&#x60; archive file.  | 
 **xContentChecksum** | **string** | The MD5 sum of the archive file.  | 

### Return type

[**ContentUploadBundle**](ContentUploadBundle.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

