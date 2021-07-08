# \GroupsApi

All URIs are relative to *http://localhost/__api__*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupMember**](GroupsApi.md#AddGroupMember) | **Post** /v1/groups/{group_guid}/members | Add a group member
[**CreateGroup**](GroupsApi.md#CreateGroup) | **Post** /v1/groups | Create a group from caller-supplied details (Password, PAM, OAuth2, SAML, Proxied) 
[**CreateRemoteGroup**](GroupsApi.md#CreateRemoteGroup) | **Put** /v1/groups | Create a group using details from a remote authentication provider (LDAP) 
[**DeleteGroup**](GroupsApi.md#DeleteGroup) | **Delete** /v1/groups/{guid} | Delete a group
[**GetGroup**](GroupsApi.md#GetGroup) | **Get** /v1/groups/{guid} | Get group details
[**GetGroupMembers**](GroupsApi.md#GetGroupMembers) | **Get** /v1/groups/{group_guid}/members | Get group member details
[**GetGroups**](GroupsApi.md#GetGroups) | **Get** /v1/groups | List or search for group details
[**RemoveGroupMember**](GroupsApi.md#RemoveGroupMember) | **Delete** /v1/groups/{group_guid}/members/{user_guid} | Remove a group member
[**SearchRemoteGroups**](GroupsApi.md#SearchRemoteGroups) | **Get** /v1/groups/remote | Search for group details from a remote provider



## AddGroupMember

> AddGroupMember(ctx, groupGuid).User(user).Execute()

Add a group member



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
    groupGuid := TODO // string | The group's unique identifier 
    user := *openapiclient.NewInlineObject6("6f300623-1e0c-48e6-a473-ddf630c0c0c3") // InlineObject6 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.AddGroupMember(context.Background(), groupGuid).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AddGroupMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | [**string**](.md) | The group&#39;s unique identifier  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**InlineObject6**](InlineObject6.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> Group CreateGroup(ctx).Group(group).Execute()

Create a group from caller-supplied details (Password, PAM, OAuth2, SAML, Proxied) 



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
    group := *openapiclient.NewInlineObject5("marketing") // InlineObject5 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.CreateGroup(context.Background()).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | [**InlineObject5**](InlineObject5.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRemoteGroup

> Group CreateRemoteGroup(ctx).Request(request).Execute()

Create a group using details from a remote authentication provider (LDAP) 



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
    request := *openapiclient.NewInlineObject4("TempTicket_example") // InlineObject4 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.CreateRemoteGroup(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.CreateRemoteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRemoteGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.CreateRemoteGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemoteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InlineObject4**](InlineObject4.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, guid).Execute()

Delete a group



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
    guid := TODO // string | The group's unique identifier 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.DeleteGroup(context.Background(), guid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | [**string**](.md) | The group&#39;s unique identifier  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## GetGroup

> Group GetGroup(ctx, guid).Execute()

Get group details



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
    guid := TODO // string | The group's unique identifier 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetGroup(context.Background(), guid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | [**string**](.md) | The group&#39;s unique identifier  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMembers

> GroupMembers GetGroupMembers(ctx, groupGuid).Execute()

Get group member details



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
    groupGuid := TODO // string | The group's unique identifier 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetGroupMembers(context.Background(), groupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupMembers`: GroupMembers
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | [**string**](.md) | The group&#39;s unique identifier  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupMembers**](GroupMembers.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> Groups GetGroups(ctx).Prefix(prefix).PageNumber(pageNumber).PageSize(pageSize).AscOrder(ascOrder).Execute()

List or search for group details



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
    prefix := "prefix_example" // string | Filters groups by prefix (group name). The filter is case insensitive.  (optional)
    pageNumber := int32(56) // int32 | The page to return relative to the given `page_size`. If `page_number` is 0 or negative, an error will be returned. For a `prefix` search, the first page of results will always be returned.  (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to include in each page. This parameter is \"best effort\" since there may not be enough groups to honor the request. If `page_size` is less than 1 or greater than 500, an error will be returned.  (optional) (default to 20)
    ascOrder := true // bool | Whether or not to return the groups in ascending order, sorted by name. For a `prefix` search, results are sorted first by exact match, then by increasing word length.  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetGroups(context.Background()).Prefix(prefix).PageNumber(pageNumber).PageSize(pageSize).AscOrder(ascOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroups`: Groups
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** | Filters groups by prefix (group name). The filter is case insensitive.  | 
 **pageNumber** | **int32** | The page to return relative to the given &#x60;page_size&#x60;. If &#x60;page_number&#x60; is 0 or negative, an error will be returned. For a &#x60;prefix&#x60; search, the first page of results will always be returned.  | [default to 1]
 **pageSize** | **int32** | The number of items to include in each page. This parameter is \&quot;best effort\&quot; since there may not be enough groups to honor the request. If &#x60;page_size&#x60; is less than 1 or greater than 500, an error will be returned.  | [default to 20]
 **ascOrder** | **bool** | Whether or not to return the groups in ascending order, sorted by name. For a &#x60;prefix&#x60; search, results are sorted first by exact match, then by increasing word length.  | [default to true]

### Return type

[**Groups**](Groups.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveGroupMember

> RemoveGroupMember(ctx, groupGuid, userGuid).Execute()

Remove a group member



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
    groupGuid := TODO // string | The group's unique identifier 
    userGuid := TODO // string | The group member's unique identifier 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.RemoveGroupMember(context.Background(), groupGuid, userGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.RemoveGroupMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | [**string**](.md) | The group&#39;s unique identifier  | 
**userGuid** | [**string**](.md) | The group member&#39;s unique identifier  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupMemberRequest struct via the builder pattern


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


## SearchRemoteGroups

> GroupRemoteSearch SearchRemoteGroups(ctx).Prefix(prefix).Limit(limit).Execute()

Search for group details from a remote provider



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
    prefix := "prefix_example" // string | Search groups by prefix. `prefix` is case insensitive. 
    limit := int32(56) // int32 | The maximum number of groups to include in the results. If `limit` is less than 1 or greater than 500, an error will be returned.  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.SearchRemoteGroups(context.Background()).Prefix(prefix).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.SearchRemoteGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchRemoteGroups`: GroupRemoteSearch
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.SearchRemoteGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRemoteGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** | Search groups by prefix. &#x60;prefix&#x60; is case insensitive.  | 
 **limit** | **int32** | The maximum number of groups to include in the results. If &#x60;limit&#x60; is less than 1 or greater than 500, an error will be returned.  | [default to 20]

### Return type

[**GroupRemoteSearch**](GroupRemoteSearch.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

