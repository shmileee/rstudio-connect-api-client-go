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
	"time"
)

// ContentVisit struct for ContentVisit
type ContentVisit struct {
	// The GUID, in [RFC4122](https://tools.ietf.org/html/rfc4122) format, of the content this information pertains to. 
	ContentGuid string `json:"content_guid"`
	// The ID of the rendering the user visited.  This will be `null` for static content. 
	RenderingId *int32 `json:"rendering_id,omitempty"`
	// The key of the variant the user visited.  This will be `null` for static content. 
	VariantKey *string `json:"variant_key,omitempty"`
	// The data version the record was recorded with.  The [Rendered and Static Content Visit Events](../admin/historical-information/#rendered-and-static-content-visit-events) section of the RStudio Connect Admin Guide explains how to interpret `data_version` values. 
	DataVersion int32 `json:"data_version"`
	// The timestamp, in [RFC3339](https://tools.ietf.org/html/rfc3339) format, when the user visited the content. 
	Time time.Time `json:"time"`
	// The ID of the particular bundle used. 
	BundleId int32 `json:"bundle_id"`
	// The GUID, in [RFC4122](https://tools.ietf.org/html/rfc4122) format, of the user that visited the content. 
	UserGuid string `json:"user_guid"`
}

// NewContentVisit instantiates a new ContentVisit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentVisit(contentGuid string, dataVersion int32, time time.Time, bundleId int32, userGuid string) *ContentVisit {
	this := ContentVisit{}
	this.ContentGuid = contentGuid
	this.DataVersion = dataVersion
	this.Time = time
	this.BundleId = bundleId
	this.UserGuid = userGuid
	return &this
}

// NewContentVisitWithDefaults instantiates a new ContentVisit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentVisitWithDefaults() *ContentVisit {
	this := ContentVisit{}
	return &this
}

// GetContentGuid returns the ContentGuid field value
func (o *ContentVisit) GetContentGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentGuid
}

// GetContentGuidOk returns a tuple with the ContentGuid field value
// and a boolean to check if the value has been set.
func (o *ContentVisit) GetContentGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentGuid, true
}

// SetContentGuid sets field value
func (o *ContentVisit) SetContentGuid(v string) {
	o.ContentGuid = v
}

// GetRenderingId returns the RenderingId field value if set, zero value otherwise.
func (o *ContentVisit) GetRenderingId() int32 {
	if o == nil || o.RenderingId == nil {
		var ret int32
		return ret
	}
	return *o.RenderingId
}

// GetRenderingIdOk returns a tuple with the RenderingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentVisit) GetRenderingIdOk() (*int32, bool) {
	if o == nil || o.RenderingId == nil {
		return nil, false
	}
	return o.RenderingId, true
}

// HasRenderingId returns a boolean if a field has been set.
func (o *ContentVisit) HasRenderingId() bool {
	if o != nil && o.RenderingId != nil {
		return true
	}

	return false
}

// SetRenderingId gets a reference to the given int32 and assigns it to the RenderingId field.
func (o *ContentVisit) SetRenderingId(v int32) {
	o.RenderingId = &v
}

// GetVariantKey returns the VariantKey field value if set, zero value otherwise.
func (o *ContentVisit) GetVariantKey() string {
	if o == nil || o.VariantKey == nil {
		var ret string
		return ret
	}
	return *o.VariantKey
}

// GetVariantKeyOk returns a tuple with the VariantKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentVisit) GetVariantKeyOk() (*string, bool) {
	if o == nil || o.VariantKey == nil {
		return nil, false
	}
	return o.VariantKey, true
}

// HasVariantKey returns a boolean if a field has been set.
func (o *ContentVisit) HasVariantKey() bool {
	if o != nil && o.VariantKey != nil {
		return true
	}

	return false
}

// SetVariantKey gets a reference to the given string and assigns it to the VariantKey field.
func (o *ContentVisit) SetVariantKey(v string) {
	o.VariantKey = &v
}

// GetDataVersion returns the DataVersion field value
func (o *ContentVisit) GetDataVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DataVersion
}

// GetDataVersionOk returns a tuple with the DataVersion field value
// and a boolean to check if the value has been set.
func (o *ContentVisit) GetDataVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DataVersion, true
}

// SetDataVersion sets field value
func (o *ContentVisit) SetDataVersion(v int32) {
	o.DataVersion = v
}

// GetTime returns the Time field value
func (o *ContentVisit) GetTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *ContentVisit) GetTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *ContentVisit) SetTime(v time.Time) {
	o.Time = v
}

// GetBundleId returns the BundleId field value
func (o *ContentVisit) GetBundleId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value
// and a boolean to check if the value has been set.
func (o *ContentVisit) GetBundleIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BundleId, true
}

// SetBundleId sets field value
func (o *ContentVisit) SetBundleId(v int32) {
	o.BundleId = v
}

// GetUserGuid returns the UserGuid field value
func (o *ContentVisit) GetUserGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserGuid
}

// GetUserGuidOk returns a tuple with the UserGuid field value
// and a boolean to check if the value has been set.
func (o *ContentVisit) GetUserGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserGuid, true
}

// SetUserGuid sets field value
func (o *ContentVisit) SetUserGuid(v string) {
	o.UserGuid = v
}

func (o ContentVisit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content_guid"] = o.ContentGuid
	}
	if o.RenderingId != nil {
		toSerialize["rendering_id"] = o.RenderingId
	}
	if o.VariantKey != nil {
		toSerialize["variant_key"] = o.VariantKey
	}
	if true {
		toSerialize["data_version"] = o.DataVersion
	}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["bundle_id"] = o.BundleId
	}
	if true {
		toSerialize["user_guid"] = o.UserGuid
	}
	return json.Marshal(toSerialize)
}

type NullableContentVisit struct {
	value *ContentVisit
	isSet bool
}

func (v NullableContentVisit) Get() *ContentVisit {
	return v.value
}

func (v *NullableContentVisit) Set(val *ContentVisit) {
	v.value = val
	v.isSet = true
}

func (v NullableContentVisit) IsSet() bool {
	return v.isSet
}

func (v *NullableContentVisit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentVisit(val *ContentVisit) *NullableContentVisit {
	return &NullableContentVisit{value: val, isSet: true}
}

func (v NullableContentVisit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentVisit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


