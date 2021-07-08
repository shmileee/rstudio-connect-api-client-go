# \UsersApi

All URIs are relative to *http://localhost/__api__*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePullUser**](UsersApi.md#CreatePullUser) | **Put** /v1/users | Create a user using details from a remote authentication provider (LDAP, OAuth2) 
[**CreatePushUser**](UsersApi.md#CreatePushUser) | **Post** /v1/users | Create a user from caller-supplied details (SAML, password, PAM, proxied) 
[**GetUser**](UsersApi.md#GetUser) | **Get** /v1/users/{guid} | Get user details
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /v1/users | List or search for user details
[**LockUser**](UsersApi.md#LockUser) | **Post** /v1/users/{guid}/lock | Lock a user
[**SearchRemoteUsers**](UsersApi.md#SearchRemoteUsers) | **Get** /v1/users/remote | Search for user details from a remote provider
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /v1/users/{guid} | Update a user



## CreatePullUser

> User CreatePullUser(ctx).Request(request).Execute()

Create a user using details from a remote authentication provider (LDAP, OAuth2) 



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
    request := *openapiclient.NewInlineObject("TempTicket_example") // InlineObject | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.CreatePullUser(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreatePullUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePullUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreatePullUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePullUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePushUser

> User CreatePushUser(ctx).User(user).Execute()

Create a user from caller-supplied details (SAML, password, PAM, proxied) 



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
    user := *openapiclient.NewInlineObject1("john_doe") // InlineObject1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.CreatePushUser(context.Background()).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreatePushUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePushUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreatePushUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePushUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, guid).Execute()

Get user details



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
    guid := TODO // string | The user's GUID, or unique identifier  Example value = `6f300623-1e0c-48e6-a473-ddf630c0c0c3` using UUID [RFC4122](https://tools.ietf.org/html/rfc4122) format. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetUser(context.Background(), guid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | [**string**](.md) | The user&#39;s GUID, or unique identifier  Example value &#x3D; &#x60;6f300623-1e0c-48e6-a473-ddf630c0c0c3&#x60; using UUID [RFC4122](https://tools.ietf.org/html/rfc4122) format.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> Users GetUsers(ctx).Prefix(prefix).UserRole(userRole).AccountStatus(accountStatus).PageNumber(pageNumber).PageSize(pageSize).AscOrder(ascOrder).Execute()

List or search for user details



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
    prefix := "prefix_example" // string | Filters users by prefix (username, first name, or last name). The filter is case insensitive.  (optional)
    userRole := "userRole_example" // string | Filter by user role. \"|\" represents logical OR. For example, `user_role=viewer|publisher` means users who are either a viewer or a publisher will be included in the result.  Note that for `user_role`, logical AND is also supported but always returns no results. For example, `user_role=viewer&user_role=publisher` tries to return users who are both viewers and publishers.  (optional)
    accountStatus := "accountStatus_example" // string | Filter by account status. \"|\" represents logical OR. For example, `account_status=locked|licensed` means users who are either locked or licensed. - `locked` - Users with a locked account. - `licensed` - Users regarded as licensed (unlocked and   recently active). - `inactive` - Users not locked and not recently active.  Note that for `account_status`, logical AND is also supported but always returns no results. For example, `account_status=locked&account_status=licensed` tries to return users who are both locked and licensed.  (optional)
    pageNumber := int32(56) // int32 | The page to return relative to the given `page_size`. If `page_number` is 0 or negative, an error will be returned.  (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of items to include in each page. This parameter is \"best effort\" since there may not be enough users to honor the request. If `page_size` is less than 1 or greater than 500, an error will be returned.  (optional) (default to 20)
    ascOrder := true // bool | Whether or not to return the users in ascending order, sorted by first name, last name, username, and email.  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetUsers(context.Background()).Prefix(prefix).UserRole(userRole).AccountStatus(accountStatus).PageNumber(pageNumber).PageSize(pageSize).AscOrder(ascOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: Users
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** | Filters users by prefix (username, first name, or last name). The filter is case insensitive.  | 
 **userRole** | **string** | Filter by user role. \&quot;|\&quot; represents logical OR. For example, &#x60;user_role&#x3D;viewer|publisher&#x60; means users who are either a viewer or a publisher will be included in the result.  Note that for &#x60;user_role&#x60;, logical AND is also supported but always returns no results. For example, &#x60;user_role&#x3D;viewer&amp;user_role&#x3D;publisher&#x60; tries to return users who are both viewers and publishers.  | 
 **accountStatus** | **string** | Filter by account status. \&quot;|\&quot; represents logical OR. For example, &#x60;account_status&#x3D;locked|licensed&#x60; means users who are either locked or licensed. - &#x60;locked&#x60; - Users with a locked account. - &#x60;licensed&#x60; - Users regarded as licensed (unlocked and   recently active). - &#x60;inactive&#x60; - Users not locked and not recently active.  Note that for &#x60;account_status&#x60;, logical AND is also supported but always returns no results. For example, &#x60;account_status&#x3D;locked&amp;account_status&#x3D;licensed&#x60; tries to return users who are both locked and licensed.  | 
 **pageNumber** | **int32** | The page to return relative to the given &#x60;page_size&#x60;. If &#x60;page_number&#x60; is 0 or negative, an error will be returned.  | [default to 1]
 **pageSize** | **int32** | The number of items to include in each page. This parameter is \&quot;best effort\&quot; since there may not be enough users to honor the request. If &#x60;page_size&#x60; is less than 1 or greater than 500, an error will be returned.  | [default to 20]
 **ascOrder** | **bool** | Whether or not to return the users in ascending order, sorted by first name, last name, username, and email.  | [default to true]

### Return type

[**Users**](Users.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockUser

> LockUser(ctx, guid).Request(request).Execute()

Lock a user



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
    guid := TODO // string | The user's GUID, or unique identifier  Example value = `6f300623-1e0c-48e6-a473-ddf630c0c0c3` using UUID [RFC4122](https://tools.ietf.org/html/rfc4122) format. 
    request := *openapiclient.NewInlineObject3(false) // InlineObject3 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.LockUser(context.Background(), guid).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.LockUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | [**string**](.md) | The user&#39;s GUID, or unique identifier  Example value &#x3D; &#x60;6f300623-1e0c-48e6-a473-ddf630c0c0c3&#x60; using UUID [RFC4122](https://tools.ietf.org/html/rfc4122) format.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**InlineObject3**](InlineObject3.md) |  | 

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


## SearchRemoteUsers

> RemoteSearchResults SearchRemoteUsers(ctx).Prefix(prefix).Limit(limit).Execute()

Search for user details from a remote provider



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
    prefix := "prefix_example" // string | Search users by prefix (username, first name, or last name). `prefix` is case insensitive. 
    limit := int32(56) // int32 | The maximum number of users to include in the results. If `limit` is less than 1 or greater than 500, an error will be returned.  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.SearchRemoteUsers(context.Background()).Prefix(prefix).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.SearchRemoteUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchRemoteUsers`: RemoteSearchResults
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.SearchRemoteUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRemoteUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** | Search users by prefix (username, first name, or last name). &#x60;prefix&#x60; is case insensitive.  | 
 **limit** | **int32** | The maximum number of users to include in the results. If &#x60;limit&#x60; is less than 1 or greater than 500, an error will be returned.  | [default to 20]

### Return type

[**RemoteSearchResults**](RemoteSearchResults.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> EditableUser UpdateUser(ctx, guid).User(user).Execute()

Update a user



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
    guid := TODO // string | The user's GUID, or unique identifier  Example value = `6f300623-1e0c-48e6-a473-ddf630c0c0c3` using UUID [RFC4122](https://tools.ietf.org/html/rfc4122) format. 
    user := *openapiclient.NewInlineObject2() // InlineObject2 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UpdateUser(context.Background(), guid).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: EditableUser
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | [**string**](.md) | The user&#39;s GUID, or unique identifier  Example value &#x3D; &#x60;6f300623-1e0c-48e6-a473-ddf630c0c0c3&#x60; using UUID [RFC4122](https://tools.ietf.org/html/rfc4122) format.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

[**EditableUser**](EditableUser.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

