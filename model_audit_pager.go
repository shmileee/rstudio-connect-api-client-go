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
	"encoding/json"
)

// AuditPager Paging object that can be used for navigation 
type AuditPager struct {
	// A full URL of the next page of results when the current request was made.  It will be `null` if the current response represents the last page of results. 
	Next NullableString `json:"next"`
	// A full URL of the previous page of results when the curent request was made.  It will be `null` if the current response represents the first page of results. 
	Previous NullableString `json:"previous"`
	// A full URL of the last page of results.  It will be `null` if the current response represents the last page of results. 
	Last NullableString `json:"last"`
	Cursors AuditPagerCursors `json:"cursors"`
	// A full URL of the first page of results.  It will be `null` if the current response represents the first page of results. 
	First NullableString `json:"first"`
}

// NewAuditPager instantiates a new AuditPager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditPager(next NullableString, previous NullableString, last NullableString, cursors AuditPagerCursors, first NullableString) *AuditPager {
	this := AuditPager{}
	this.Next = next
	this.Previous = previous
	this.Last = last
	this.Cursors = cursors
	this.First = first
	return &this
}

// NewAuditPagerWithDefaults instantiates a new AuditPager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditPagerWithDefaults() *AuditPager {
	this := AuditPager{}
	return &this
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuditPager) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditPager) GetNextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *AuditPager) SetNext(v string) {
	o.Next.Set(&v)
}

// GetPrevious returns the Previous field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuditPager) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}

	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditPager) GetPreviousOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// SetPrevious sets field value
func (o *AuditPager) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// GetLast returns the Last field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuditPager) GetLast() string {
	if o == nil || o.Last.Get() == nil {
		var ret string
		return ret
	}

	return *o.Last.Get()
}

// GetLastOk returns a tuple with the Last field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditPager) GetLastOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Last.Get(), o.Last.IsSet()
}

// SetLast sets field value
func (o *AuditPager) SetLast(v string) {
	o.Last.Set(&v)
}

// GetCursors returns the Cursors field value
func (o *AuditPager) GetCursors() AuditPagerCursors {
	if o == nil {
		var ret AuditPagerCursors
		return ret
	}

	return o.Cursors
}

// GetCursorsOk returns a tuple with the Cursors field value
// and a boolean to check if the value has been set.
func (o *AuditPager) GetCursorsOk() (*AuditPagerCursors, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cursors, true
}

// SetCursors sets field value
func (o *AuditPager) SetCursors(v AuditPagerCursors) {
	o.Cursors = v
}

// GetFirst returns the First field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuditPager) GetFirst() string {
	if o == nil || o.First.Get() == nil {
		var ret string
		return ret
	}

	return *o.First.Get()
}

// GetFirstOk returns a tuple with the First field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditPager) GetFirstOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.First.Get(), o.First.IsSet()
}

// SetFirst sets field value
func (o *AuditPager) SetFirst(v string) {
	o.First.Set(&v)
}

func (o AuditPager) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["next"] = o.Next.Get()
	}
	if true {
		toSerialize["previous"] = o.Previous.Get()
	}
	if true {
		toSerialize["last"] = o.Last.Get()
	}
	if true {
		toSerialize["cursors"] = o.Cursors
	}
	if true {
		toSerialize["first"] = o.First.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAuditPager struct {
	value *AuditPager
	isSet bool
}

func (v NullableAuditPager) Get() *AuditPager {
	return v.value
}

func (v *NullableAuditPager) Set(val *AuditPager) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditPager) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditPager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditPager(val *AuditPager) *NullableAuditPager {
	return &NullableAuditPager{value: val, isSet: true}
}

func (v NullableAuditPager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditPager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


