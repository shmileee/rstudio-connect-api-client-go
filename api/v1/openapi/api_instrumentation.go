/*
 * RStudio Connect API Reference
 *
 * The RStudio Connect API can be used to perform certain user actions remotely. You will need to install a tool or library that can make HTTP requests, such as:  - [httr](http://httr.r-lib.org/) (R HTTP library) - [cURL](https://curl.haxx.se/) (Linux tool for making HTTP calls) - [requests](https://requests.readthedocs.io/en/master/) (Python HTTP library)  ## <a name=\"versioning-policy\">Versioning of the API</a>  The RStudio Connect Server API uses a simple, single number versioning scheme as noted as the first part of each endpoint path.  This version number will only be incremented in the event that non-backward compatible changes are made to an existing endpoint. Note that this occurs on a per-endpoint basis; see the section on [deprecation](#deprecation) below for more information.  Changes that are considered backward compatible are:  * New fields in responses. * New non-required fields in requests. * New endpoint behavior which does not violate the current functional intent of the   endpoint.  Changes that are considered non-backward compatible are:  * Removal or rename of request or response fields. * A change of the type or format of one or more request or response fields. * Addition of new required request fields. * A substantial deviation from the current functional intent of the endpoint.  The points relating to functional intent are assumed to be extremely rare as more often such situations will result in a completely new endpoint, which makes the change a backward compatible addition.  ### Experimentation  RStudio Connect labels experimental endpoints in the API by including `/experimental` in the endpoint path immediately after the version indicator.  If an endpoint is noted as experimental, it should not be relied upon for any production work.  These are endpoints that RStudio Connect is making available to our customers to solicit feedback; they are subject to change without notice.  Such changes include anything from altered request/response shapes, to complete abandonment of the endpoint.  This public review of an experimental endpoint will last as long as necessary to either prove its viability or to determine that it’s not really needed.  The time for this will vary based on the intricacies of each endpoint.  When the endpoint is finalized, the next release of RStudio Connect will mark the experimental path as deprecated while adding the endpoint without the `/experimental` prefix. The path with the experimental prefix will be removed six months later.  The documentation for the endpoint will also note, during that time, the original, experimental, path.  All experimental endpoints are clearly marked as such in this documentation.  ### <a name=\"deprecation\">Deprecation and Removal of Old Versions</a>  It is possible that RStudio Connect may decide to deprecate an endpoint.  This will happen if either the endpoint serves no useful purpose because it’s functionality is now handled by a different endpoint or because there is a newer version of the endpoint that should be used.  If a deprecated endpoint is called, the response to it will include an extra HTTP header called, `X-Deprecated-Endpoint` and will have as a value the path of the endpoint that should be used instead.  If the functionality has no direct replacement, the value will be set to `n/a`.  Deprecated versions of an endpoint will be supported for 1 year from the release date of RStudio Connect in which the endpoint was marked as deprecated.  At that time, the endpoint is subject to removal at the discretion of RStudio Connect.  The life cycle of an endpoint will follow these steps.  1. The `/v1/endpoint` is public and in use by RStudio Connect customers. 1. RStudio Connect makes `/v2/experimental/endpoint` available for testing and feedback.    Customers should still use `/v1/endpoint` for production work. 1. RStudio Connect moves version 2 of the endpoint out of experimentation so, all within    the same release:     1. `/v1/endpoint` is marked as deprecated.     1. `/v2/experimental/endpoint` is marked as deprecated.     1. `/v2/endpoint` is made public. 1. Six months later, `/v2/experimental/endpoint` is removed from the product. 1. Twelve months later, `/v1/endpoint` is removed from the product.  Note that it is possible that RStudio Connect may produce a new version of an existing endpoint without making an experimental version of it first.  The same life cycle, without those parts, will still be followed.  ## <a name=\"authentication\">Authentication</a>  Most endpoints require you to identify yourself as a valid RStudio Connect user. You do this by specifying an API Key when you make a call to the server. The [API Keys](../user/api-keys/) chapter of the RStudio Connect User Guide explains how to create an API Key.  ### <a name=\"api-keys\">API Keys</a>  API Keys are managed by each user in the RStudio Connect dashboard. If you ever lose an API Key or otherwise feel it has been compromised, use the dashboard to revoke the key and create another one.  ::alert danger Keep your API Key safe.  If your RStudio Connect server's URL does not begin with `https`, your API Key could be intercepted and used by a malicious actor. ::  Once you have an API Key, you can authenticate by passing the key with a prefix of `\"Key \"` (the space is important) in the Authorization header.  Below are examples of invoking the \"Get R Information\" endpoint.  **cURL** ```bash curl -H \"Authorization: Key XXXXXXXXXXX\" \\      https://rstudioconnect.example.com/__api__/v1/server_settings/r ```  **R** ```r library(httr) apiKey <- \"XXXXXXXXXXX\" result <- GET(\"https://rstudioconnect.example.com/__api__/v1/server_settings/r\",   add_headers(Authorization = paste(\"Key\", apiKey))) ```  **Python** ```python import requests r = requests.get(   'https://rstudioconnect.example.com/__api__/v1/server_settings/r',   headers = { 'Authorization': 'Key XXXXXXXXXXX' } ) ``` 
 *
 * API version: 1.0.1
 * Contact: support@rstudio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// InstrumentationApiService InstrumentationApi service
type InstrumentationApiService service

type ApiGetContentVisitsRequest struct {
	ctx _context.Context
	ApiService *InstrumentationApiService
	contentGuid *string
	minDataVersion *int32
	from *time.Time
	to *time.Time
	limit *int32
	previous *string
	next *string
	ascOrder *bool
}

func (r ApiGetContentVisitsRequest) ContentGuid(contentGuid string) ApiGetContentVisitsRequest {
	r.contentGuid = &contentGuid
	return r
}
func (r ApiGetContentVisitsRequest) MinDataVersion(minDataVersion int32) ApiGetContentVisitsRequest {
	r.minDataVersion = &minDataVersion
	return r
}
func (r ApiGetContentVisitsRequest) From(from time.Time) ApiGetContentVisitsRequest {
	r.from = &from
	return r
}
func (r ApiGetContentVisitsRequest) To(to time.Time) ApiGetContentVisitsRequest {
	r.to = &to
	return r
}
func (r ApiGetContentVisitsRequest) Limit(limit int32) ApiGetContentVisitsRequest {
	r.limit = &limit
	return r
}
func (r ApiGetContentVisitsRequest) Previous(previous string) ApiGetContentVisitsRequest {
	r.previous = &previous
	return r
}
func (r ApiGetContentVisitsRequest) Next(next string) ApiGetContentVisitsRequest {
	r.next = &next
	return r
}
func (r ApiGetContentVisitsRequest) AscOrder(ascOrder bool) ApiGetContentVisitsRequest {
	r.ascOrder = &ascOrder
	return r
}

func (r ApiGetContentVisitsRequest) Execute() (ContentVisitLogs, *_nethttp.Response, error) {
	return r.ApiService.GetContentVisitsExecute(r)
}

/*
 * GetContentVisits Get Content Visits
 * This endpoint returns a portion of the visit (or "hits") information for
renderend and static content.  Rendered content includes R Markdown (both
with and without parameters), Jupyter Notebooks, etc.  Static content
includes plots, spreadsheets, sites, etc.  This endpoint does not return
information on Shiny applications, plumber APIs or TensorFlow models.  The
results returned include paging details that can be used to navigate the
information this endpoint returns.

The information returned is based on data collected by RStudio Connect as
users visit rendered and static content.

::alert warning
Prior to the release of this API, there was an issue with how visits were returned
that caused extra entries to be stored under certain circumstances.  These will
affect analyses that interpret visit counts.  Entries that were recorded before
the issue was addressed are not returned by default.  If you desire these
records, specify the `min_data_version` filter with a value of 0.
::

::alert info
This endpoint requires administrator or publisher access.
::

#### Filtering of results:

There are several ways the result set can be filtered by the server before
being sent back within the API response. If multiple filters are in effect,
they will be logically ANDed together.

##### Implicit Filtering
If the user calling the endpoint is a publisher, the data returned will be
limited to the content owned by the user.

##### Time Windows

This API accepts optional `from` and `to` timestamps to define a window of
interest.  If `from` is not specified, it is assumed to be before the earliest
recorded information.  If `to` is not specified, it is assumed to be "now".

Any visit to content that falls inclusively within the time window will be
part of the result set.

#### Responses

The response of a call will contain zero or more data records representing a visit
by a user to a piece of content.  No more than `limit` records will be returned.
Multiple requests of this endpoint are typically required to retrieve the complete
result set from the server.  To facilitate this, each response includes a paging
object containing full URL links which can be requested to iteratively move forward
or backward through multiple response pages.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetContentVisitsRequest
 */
func (a *InstrumentationApiService) GetContentVisits(ctx _context.Context) ApiGetContentVisitsRequest {
	return ApiGetContentVisitsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ContentVisitLogs
 */
func (a *InstrumentationApiService) GetContentVisitsExecute(r ApiGetContentVisitsRequest) (ContentVisitLogs, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ContentVisitLogs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstrumentationApiService.GetContentVisits")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instrumentation/content/visits"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.contentGuid != nil {
		localVarQueryParams.Add("content_guid", parameterToString(*r.contentGuid, ""))
	}
	if r.minDataVersion != nil {
		localVarQueryParams.Add("min_data_version", parameterToString(*r.minDataVersion, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.previous != nil {
		localVarQueryParams.Add("previous", parameterToString(*r.previous, ""))
	}
	if r.next != nil {
		localVarQueryParams.Add("next", parameterToString(*r.next, ""))
	}
	if r.ascOrder != nil {
		localVarQueryParams.Add("asc_order", parameterToString(*r.ascOrder, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetShinyAppUsageRequest struct {
	ctx _context.Context
	ApiService *InstrumentationApiService
	contentGuid *string
	minDataVersion *int32
	from *time.Time
	to *time.Time
	limit *int32
	previous *string
	next *string
	ascOrder *bool
}

func (r ApiGetShinyAppUsageRequest) ContentGuid(contentGuid string) ApiGetShinyAppUsageRequest {
	r.contentGuid = &contentGuid
	return r
}
func (r ApiGetShinyAppUsageRequest) MinDataVersion(minDataVersion int32) ApiGetShinyAppUsageRequest {
	r.minDataVersion = &minDataVersion
	return r
}
func (r ApiGetShinyAppUsageRequest) From(from time.Time) ApiGetShinyAppUsageRequest {
	r.from = &from
	return r
}
func (r ApiGetShinyAppUsageRequest) To(to time.Time) ApiGetShinyAppUsageRequest {
	r.to = &to
	return r
}
func (r ApiGetShinyAppUsageRequest) Limit(limit int32) ApiGetShinyAppUsageRequest {
	r.limit = &limit
	return r
}
func (r ApiGetShinyAppUsageRequest) Previous(previous string) ApiGetShinyAppUsageRequest {
	r.previous = &previous
	return r
}
func (r ApiGetShinyAppUsageRequest) Next(next string) ApiGetShinyAppUsageRequest {
	r.next = &next
	return r
}
func (r ApiGetShinyAppUsageRequest) AscOrder(ascOrder bool) ApiGetShinyAppUsageRequest {
	r.ascOrder = &ascOrder
	return r
}

func (r ApiGetShinyAppUsageRequest) Execute() (ShinyAppUsageLogs, *_nethttp.Response, error) {
	return r.ApiService.GetShinyAppUsageExecute(r)
}

/*
 * GetShinyAppUsage Get Shiny App Usage
 * This endpoint returns a portion of the Shiny application usage information,
as well as paging details that can be used to navigate that information.

The information returned is based on data collected by RStudio Connect as
users visit Shiny applications.  Because of how visits are detected, end
times will be slightly inflated by a reconnect timeout, generally around
15 seconds.

::alert warning
Prior to the release of this API, there was an issue with how visits were
recorded that caused extra entries to be stored.  These will affect analyses
that interpret visit counts or durations.  Entries that were recorded before
the issue was addressed are not returned by default.  If you desire these
records, specify the `min_data_version` filter with a value of 0.
::

::alert info

- Because of how visits are detected, end times will be slightly inflated by the
  currently configured client reconnect timeout, which defaults to 15 seconds.
  The ending time may also be affected by connect and read timeout
  settings.

  The [Shiny Application
  Events](../admin/historical-information/#shiny-application-events)
  section of the RStudio Connect Admin Guide has more details about
  how these events are collected.
- This endpoint requires administrator or publisher access.

::

#### Filtering of results:

There are several ways the result set can be filtered by the server before
being sent back within the API response. If multiple filters are in effect,
they will be logically ANDed together.

##### Implicit Filtering

If the user calling the endpoint is a publisher, the data returned will be
limited to those applications owned by the user.

##### Time Windows

This API accepts optional `from` and `to` timestamps to define a window of
interest.  If `from` is not specified, it is assumed to be before the earliest
recorded information.  If `to` is not specified, it is assumed to be "now".

Any visit to content that falls inclusively within the time window will be
part of the result set.

#### Responses

The response of a call will contain zero or more data records representing a session
by a user of a Shiny application.  No more than `limit` records will be returned.
Multiple requests of this endpoint are typically required to retrieve the complete
result set from the server.  To facilitate this, each response includes a paging
object containing full URL links which can be requested to iteratively move forward
or backward through multiple response pages.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetShinyAppUsageRequest
 */
func (a *InstrumentationApiService) GetShinyAppUsage(ctx _context.Context) ApiGetShinyAppUsageRequest {
	return ApiGetShinyAppUsageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ShinyAppUsageLogs
 */
func (a *InstrumentationApiService) GetShinyAppUsageExecute(r ApiGetShinyAppUsageRequest) (ShinyAppUsageLogs, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ShinyAppUsageLogs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstrumentationApiService.GetShinyAppUsage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instrumentation/shiny/usage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.contentGuid != nil {
		localVarQueryParams.Add("content_guid", parameterToString(*r.contentGuid, ""))
	}
	if r.minDataVersion != nil {
		localVarQueryParams.Add("min_data_version", parameterToString(*r.minDataVersion, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.previous != nil {
		localVarQueryParams.Add("previous", parameterToString(*r.previous, ""))
	}
	if r.next != nil {
		localVarQueryParams.Add("next", parameterToString(*r.next, ""))
	}
	if r.ascOrder != nil {
		localVarQueryParams.Add("asc_order", parameterToString(*r.ascOrder, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
