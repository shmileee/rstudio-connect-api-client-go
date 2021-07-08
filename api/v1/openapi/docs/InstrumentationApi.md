# \InstrumentationApi

All URIs are relative to *http://localhost/__api__*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContentVisits**](InstrumentationApi.md#GetContentVisits) | **Get** /v1/instrumentation/content/visits | Get Content Visits
[**GetShinyAppUsage**](InstrumentationApi.md#GetShinyAppUsage) | **Get** /v1/instrumentation/shiny/usage | Get Shiny App Usage



## GetContentVisits

> ContentVisitLogs GetContentVisits(ctx).ContentGuid(contentGuid).MinDataVersion(minDataVersion).From(from).To(to).Limit(limit).Previous(previous).Next(next).AscOrder(ascOrder).Execute()

Get Content Visits



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    contentGuid := TODO // string | Filter results by content GUID.  This parameter will limit the results to include only the access records for content matching the content GUID filter value.  Example value = `6f300623-1e0c-48e6-a473-ddf630c0c0c3` using UUID [RFC4122](https://tools.ietf.org/html/rfc4122) format.  Multiple content GUIDs may be specified using the bar character `|` which represents a logical OR across the surounding GUIDs.  For example, `content_guid=6f300623-1e0c-48e6-a473-ddf630c0c0c3|e08a86af-a262-4152-8366-f2d8ec3c54f9` will filter the result set to only the access records for content with a GUID of `6f300623-1e0c-48e6-a473-ddf630c0c0c3` or a GUID of `e08a86af-a262-4152-8366-f2d8ec3c54f9`.  ::alert info  - The GUID associated with published content can be found on the Info panel for   the content within the RStudio Connect Dashboard window. - Note that if you specify the `content_guid` more than once like this,   `content_guid=6f300623-1e0c-48e6-a473-ddf630c0c0c3&content_guid=&e08a86af-a262-4152-8366-f2d8ec3c54f9`   you will receive results for the first GUID only; the 2nd and subsequent `content_guid`   fields are ignored. - Note that there is a practical limit of around 40 to the number of GUIDs that   may be specified due to the length of a html query string.  ::  (optional)
    minDataVersion := int32(56) // int32 | Filter by data version.  Records with a data version lower than the given value will be excluded from the set of results.  The [Rendered and Static Content Visit Events](../admin/historical-information/#rendered-and-static-content-visit-events) section of the RStudio Connect Admin Guide explains how to interpret `data_version` values.  (optional) (default to 1)
    from := time.Now() // time.Time | The timestamp that starts the time window of interest.  Any visit information that happened prior to this timestamp will not be returned.  Example value = `2018-09-15T18:00:00Z` using [RFC3339](https://tools.ietf.org/html/rfc3339) format of (`yyyy-mm-ddThh:mm:ss` followed by either `-##:##`, `.##Z` or `Z`)  (optional)
    to := time.Now() // time.Time | The timestamp that ends the time window of interest.  Any visit information that happened after this timestamp will not be returned.  Example value = `2018-09-15T18:00:00Z` using [RFC3339](https://tools.ietf.org/html/rfc3339) format of (`yyyy-mm-ddThh:mm:ss` followed by either `-##:##`, `.##Z` or `Z`)  (optional)
    limit := int32(56) // int32 | The number of records to return.  The minimum supported value is 1 and maximum supported value is 500.  Note that `limit` is a \"best effort\" request since there may not be enough visit entries to satisfy the limit.  (optional) (default to 20)
    previous := "previous_example" // string | Retrieve the previous page of content visit logs relative to the provided value.  This value corresponds to an internal reference within the server and should be sourced from the appropriate attribute within the paging object of a previous response.  (optional)
    next := "next_example" // string | Retrieve the next page of content visit logs relative to the provided value.  This value corresponds to an internal reference within the server and should be sourced from the appropriate attribute within the paging object of a previous response.  (optional)
    ascOrder := true // bool | Determines if the response records should be listed in ascending or descending order within the response.  Ordering is by the `started` timestamp field.  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstrumentationApi.GetContentVisits(context.Background()).ContentGuid(contentGuid).MinDataVersion(minDataVersion).From(from).To(to).Limit(limit).Previous(previous).Next(next).AscOrder(ascOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstrumentationApi.GetContentVisits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContentVisits`: ContentVisitLogs
    fmt.Fprintf(os.Stdout, "Response from `InstrumentationApi.GetContentVisits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContentVisitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentGuid** | [**string**](string.md) | Filter results by content GUID.  This parameter will limit the results to include only the access records for content matching the content GUID filter value.  Example value &#x3D; &#x60;6f300623-1e0c-48e6-a473-ddf630c0c0c3&#x60; using UUID [RFC4122](https://tools.ietf.org/html/rfc4122) format.  Multiple content GUIDs may be specified using the bar character &#x60;|&#x60; which represents a logical OR across the surounding GUIDs.  For example, &#x60;content_guid&#x3D;6f300623-1e0c-48e6-a473-ddf630c0c0c3|e08a86af-a262-4152-8366-f2d8ec3c54f9&#x60; will filter the result set to only the access records for content with a GUID of &#x60;6f300623-1e0c-48e6-a473-ddf630c0c0c3&#x60; or a GUID of &#x60;e08a86af-a262-4152-8366-f2d8ec3c54f9&#x60;.  ::alert info  - The GUID associated with published content can be found on the Info panel for   the content within the RStudio Connect Dashboard window. - Note that if you specify the &#x60;content_guid&#x60; more than once like this,   &#x60;content_guid&#x3D;6f300623-1e0c-48e6-a473-ddf630c0c0c3&amp;content_guid&#x3D;&amp;e08a86af-a262-4152-8366-f2d8ec3c54f9&#x60;   you will receive results for the first GUID only; the 2nd and subsequent &#x60;content_guid&#x60;   fields are ignored. - Note that there is a practical limit of around 40 to the number of GUIDs that   may be specified due to the length of a html query string.  ::  | 
 **minDataVersion** | **int32** | Filter by data version.  Records with a data version lower than the given value will be excluded from the set of results.  The [Rendered and Static Content Visit Events](../admin/historical-information/#rendered-and-static-content-visit-events) section of the RStudio Connect Admin Guide explains how to interpret &#x60;data_version&#x60; values.  | [default to 1]
 **from** | **time.Time** | The timestamp that starts the time window of interest.  Any visit information that happened prior to this timestamp will not be returned.  Example value &#x3D; &#x60;2018-09-15T18:00:00Z&#x60; using [RFC3339](https://tools.ietf.org/html/rfc3339) format of (&#x60;yyyy-mm-ddThh:mm:ss&#x60; followed by either &#x60;-##:##&#x60;, &#x60;.##Z&#x60; or &#x60;Z&#x60;)  | 
 **to** | **time.Time** | The timestamp that ends the time window of interest.  Any visit information that happened after this timestamp will not be returned.  Example value &#x3D; &#x60;2018-09-15T18:00:00Z&#x60; using [RFC3339](https://tools.ietf.org/html/rfc3339) format of (&#x60;yyyy-mm-ddThh:mm:ss&#x60; followed by either &#x60;-##:##&#x60;, &#x60;.##Z&#x60; or &#x60;Z&#x60;)  | 
 **limit** | **int32** | The number of records to return.  The minimum supported value is 1 and maximum supported value is 500.  Note that &#x60;limit&#x60; is a \&quot;best effort\&quot; request since there may not be enough visit entries to satisfy the limit.  | [default to 20]
 **previous** | **string** | Retrieve the previous page of content visit logs relative to the provided value.  This value corresponds to an internal reference within the server and should be sourced from the appropriate attribute within the paging object of a previous response.  | 
 **next** | **string** | Retrieve the next page of content visit logs relative to the provided value.  This value corresponds to an internal reference within the server and should be sourced from the appropriate attribute within the paging object of a previous response.  | 
 **ascOrder** | **bool** | Determines if the response records should be listed in ascending or descending order within the response.  Ordering is by the &#x60;started&#x60; timestamp field.  | [default to true]

### Return type

[**ContentVisitLogs**](ContentVisitLogs.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShinyAppUsage

> ShinyAppUsageLogs GetShinyAppUsage(ctx).ContentGuid(contentGuid).MinDataVersion(minDataVersion).From(from).To(to).Limit(limit).Previous(previous).Next(next).AscOrder(ascOrder).Execute()

Get Shiny App Usage



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    contentGuid := TODO // string | Filter results by content GUID.  This parameter will limit the results to include only the access records for the Shiny application(s) matching the content GUID filter value.  Example value = `6f300623-1e0c-48e6-a473-ddf630c0c0c3` using UUID [RFC4122](https://tools.ietf.org/html/rfc4122) format.  Multiple content GUIDs may be specified using the bar character `|` which represents a logical OR across the surounding GUIDs.  For example, `content_guid=6f300623-1e0c-48e6-a473-ddf630c0c0c3|e08a86af-a262-4152-8366-f2d8ec3c54f9` will filter the result set to only the access records for the Shiny Applications with a GUID of `6f300623-1e0c-48e6-a473-ddf630c0c0c3` or a GUID of `e08a86af-a262-4152-8366-f2d8ec3c54f9`.  ::alert info  - The GUID associated with a published Shiny application can be found on the   Info panel for the application within the RStudio Connect Dashboard window. - Note that if you specify the `content_guid` more than once like this,   `content_guid=6f300623-1e0c-48e6-a473-ddf630c0c0c3&content_guid=&e08a86af-a262-4152-8366-f2d8ec3c54f9`   you will receive results for the first GUID only; the 2nd and subsequent `content_guid`   fields are ignored. - Note that there is a practical limit of around 40 to the number of GUIDs that   may be specified due to the length of a html query string.  ::  (optional)
    minDataVersion := int32(56) // int32 | Filter by data version.  Records with a data version lower than the given value will be excluded from the set of results.  The [Shiny Application Events](../admin/historical-information/#shiny-application-events) section of the RStudio Connect Admin Guide explains how to interpret `data_version` values.  (optional) (default to 1)
    from := time.Now() // time.Time | The timestamp that starts the time window of interest.  Any usage information that _ends_ prior to this timestamp will not be returned.  Individual records may contain a starting time that is before this if they end after it or have not finished.  Example value = `2018-09-15T18:00:00Z` using [RFC3339](https://tools.ietf.org/html/rfc3339) format of (`yyyy-mm-ddThh:mm:ss` followed by either `-##:##`, `.##Z` or `Z`)  (optional)
    to := time.Now() // time.Time | The timestamp that ends the time window of interest.  Any usage information that _starts_ after this timestamp will not be returned.  Individual records may contain an ending time that is after this (or no ending time) if they start before it.  Example value = `2018-09-15T18:00:00Z` using [RFC3339](https://tools.ietf.org/html/rfc3339) format of (`yyyy-mm-ddThh:mm:ss` followed by either `-##:##`, `.##Z` or `Z`)  (optional)
    limit := int32(56) // int32 | The number of records to return.  The minimum supported value is 1 and maximum supported value is 500.  Note that `limit` is a \"best effort\" request since there may not be enough usage entries to satisfy the limit.  (optional) (default to 20)
    previous := "previous_example" // string | Retrieve the previous page of Shiny application usage logs relative to the provided value.  This value corresponds to an internal reference within the server and should be sourced from the appropriate attribute within the paging object of a previous response.  (optional)
    next := "next_example" // string | Retrieve the next page of Shiny application usage logs relative to the provided value.  This value corresponds to an internal reference within the server and should be sourced from the appropriate attribute within the paging object of a previous response.  (optional)
    ascOrder := true // bool | Determines if the response records should be listed in ascending or descending order within the response.  Ordering is by the `started` timestamp field.  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstrumentationApi.GetShinyAppUsage(context.Background()).ContentGuid(contentGuid).MinDataVersion(minDataVersion).From(from).To(to).Limit(limit).Previous(previous).Next(next).AscOrder(ascOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstrumentationApi.GetShinyAppUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShinyAppUsage`: ShinyAppUsageLogs
    fmt.Fprintf(os.Stdout, "Response from `InstrumentationApi.GetShinyAppUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetShinyAppUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentGuid** | [**string**](string.md) | Filter results by content GUID.  This parameter will limit the results to include only the access records for the Shiny application(s) matching the content GUID filter value.  Example value &#x3D; &#x60;6f300623-1e0c-48e6-a473-ddf630c0c0c3&#x60; using UUID [RFC4122](https://tools.ietf.org/html/rfc4122) format.  Multiple content GUIDs may be specified using the bar character &#x60;|&#x60; which represents a logical OR across the surounding GUIDs.  For example, &#x60;content_guid&#x3D;6f300623-1e0c-48e6-a473-ddf630c0c0c3|e08a86af-a262-4152-8366-f2d8ec3c54f9&#x60; will filter the result set to only the access records for the Shiny Applications with a GUID of &#x60;6f300623-1e0c-48e6-a473-ddf630c0c0c3&#x60; or a GUID of &#x60;e08a86af-a262-4152-8366-f2d8ec3c54f9&#x60;.  ::alert info  - The GUID associated with a published Shiny application can be found on the   Info panel for the application within the RStudio Connect Dashboard window. - Note that if you specify the &#x60;content_guid&#x60; more than once like this,   &#x60;content_guid&#x3D;6f300623-1e0c-48e6-a473-ddf630c0c0c3&amp;content_guid&#x3D;&amp;e08a86af-a262-4152-8366-f2d8ec3c54f9&#x60;   you will receive results for the first GUID only; the 2nd and subsequent &#x60;content_guid&#x60;   fields are ignored. - Note that there is a practical limit of around 40 to the number of GUIDs that   may be specified due to the length of a html query string.  ::  | 
 **minDataVersion** | **int32** | Filter by data version.  Records with a data version lower than the given value will be excluded from the set of results.  The [Shiny Application Events](../admin/historical-information/#shiny-application-events) section of the RStudio Connect Admin Guide explains how to interpret &#x60;data_version&#x60; values.  | [default to 1]
 **from** | **time.Time** | The timestamp that starts the time window of interest.  Any usage information that _ends_ prior to this timestamp will not be returned.  Individual records may contain a starting time that is before this if they end after it or have not finished.  Example value &#x3D; &#x60;2018-09-15T18:00:00Z&#x60; using [RFC3339](https://tools.ietf.org/html/rfc3339) format of (&#x60;yyyy-mm-ddThh:mm:ss&#x60; followed by either &#x60;-##:##&#x60;, &#x60;.##Z&#x60; or &#x60;Z&#x60;)  | 
 **to** | **time.Time** | The timestamp that ends the time window of interest.  Any usage information that _starts_ after this timestamp will not be returned.  Individual records may contain an ending time that is after this (or no ending time) if they start before it.  Example value &#x3D; &#x60;2018-09-15T18:00:00Z&#x60; using [RFC3339](https://tools.ietf.org/html/rfc3339) format of (&#x60;yyyy-mm-ddThh:mm:ss&#x60; followed by either &#x60;-##:##&#x60;, &#x60;.##Z&#x60; or &#x60;Z&#x60;)  | 
 **limit** | **int32** | The number of records to return.  The minimum supported value is 1 and maximum supported value is 500.  Note that &#x60;limit&#x60; is a \&quot;best effort\&quot; request since there may not be enough usage entries to satisfy the limit.  | [default to 20]
 **previous** | **string** | Retrieve the previous page of Shiny application usage logs relative to the provided value.  This value corresponds to an internal reference within the server and should be sourced from the appropriate attribute within the paging object of a previous response.  | 
 **next** | **string** | Retrieve the next page of Shiny application usage logs relative to the provided value.  This value corresponds to an internal reference within the server and should be sourced from the appropriate attribute within the paging object of a previous response.  | 
 **ascOrder** | **bool** | Determines if the response records should be listed in ascending or descending order within the response.  Ordering is by the &#x60;started&#x60; timestamp field.  | [default to true]

### Return type

[**ShinyAppUsageLogs**](ShinyAppUsageLogs.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

