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

// ShinyAppUsage struct for ShinyAppUsage
type ShinyAppUsage struct {
	// The timestamp, in [RFC3339](https://tools.ietf.org/html/rfc3339) format, when the user opened the application. 
	Started time.Time `json:"started"`
	// The GUID, in [RFC4122](https://tools.ietf.org/html/rfc4122) format, of the Shiny application this information pertains to. 
	ContentGuid string `json:"content_guid"`
	// The GUID, in [RFC4122](https://tools.ietf.org/html/rfc4122) format, of the user that visited the application. 
	UserGuid string `json:"user_guid"`
	// The data version the record was recorded with.  The [Shiny Application Events](../admin/historical-information/#shiny-application-events) section of the RStudio Connect Admin Guide explains how to interpret `data_version` values. 
	DataVersion int32 `json:"data_version"`
	// The timestamp, in [RFC3339](https://tools.ietf.org/html/rfc3339) format, when the user left the application. 
	Ended time.Time `json:"ended"`
}

// NewShinyAppUsage instantiates a new ShinyAppUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShinyAppUsage(started time.Time, contentGuid string, userGuid string, dataVersion int32, ended time.Time) *ShinyAppUsage {
	this := ShinyAppUsage{}
	this.Started = started
	this.ContentGuid = contentGuid
	this.UserGuid = userGuid
	this.DataVersion = dataVersion
	this.Ended = ended
	return &this
}

// NewShinyAppUsageWithDefaults instantiates a new ShinyAppUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShinyAppUsageWithDefaults() *ShinyAppUsage {
	this := ShinyAppUsage{}
	return &this
}

// GetStarted returns the Started field value
func (o *ShinyAppUsage) GetStarted() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Started
}

// GetStartedOk returns a tuple with the Started field value
// and a boolean to check if the value has been set.
func (o *ShinyAppUsage) GetStartedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Started, true
}

// SetStarted sets field value
func (o *ShinyAppUsage) SetStarted(v time.Time) {
	o.Started = v
}

// GetContentGuid returns the ContentGuid field value
func (o *ShinyAppUsage) GetContentGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentGuid
}

// GetContentGuidOk returns a tuple with the ContentGuid field value
// and a boolean to check if the value has been set.
func (o *ShinyAppUsage) GetContentGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentGuid, true
}

// SetContentGuid sets field value
func (o *ShinyAppUsage) SetContentGuid(v string) {
	o.ContentGuid = v
}

// GetUserGuid returns the UserGuid field value
func (o *ShinyAppUsage) GetUserGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserGuid
}

// GetUserGuidOk returns a tuple with the UserGuid field value
// and a boolean to check if the value has been set.
func (o *ShinyAppUsage) GetUserGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserGuid, true
}

// SetUserGuid sets field value
func (o *ShinyAppUsage) SetUserGuid(v string) {
	o.UserGuid = v
}

// GetDataVersion returns the DataVersion field value
func (o *ShinyAppUsage) GetDataVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DataVersion
}

// GetDataVersionOk returns a tuple with the DataVersion field value
// and a boolean to check if the value has been set.
func (o *ShinyAppUsage) GetDataVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DataVersion, true
}

// SetDataVersion sets field value
func (o *ShinyAppUsage) SetDataVersion(v int32) {
	o.DataVersion = v
}

// GetEnded returns the Ended field value
func (o *ShinyAppUsage) GetEnded() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Ended
}

// GetEndedOk returns a tuple with the Ended field value
// and a boolean to check if the value has been set.
func (o *ShinyAppUsage) GetEndedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ended, true
}

// SetEnded sets field value
func (o *ShinyAppUsage) SetEnded(v time.Time) {
	o.Ended = v
}

func (o ShinyAppUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["started"] = o.Started
	}
	if true {
		toSerialize["content_guid"] = o.ContentGuid
	}
	if true {
		toSerialize["user_guid"] = o.UserGuid
	}
	if true {
		toSerialize["data_version"] = o.DataVersion
	}
	if true {
		toSerialize["ended"] = o.Ended
	}
	return json.Marshal(toSerialize)
}

type NullableShinyAppUsage struct {
	value *ShinyAppUsage
	isSet bool
}

func (v NullableShinyAppUsage) Get() *ShinyAppUsage {
	return v.value
}

func (v *NullableShinyAppUsage) Set(val *ShinyAppUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableShinyAppUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableShinyAppUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShinyAppUsage(val *ShinyAppUsage) *NullableShinyAppUsage {
	return &NullableShinyAppUsage{value: val, isSet: true}
}

func (v NullableShinyAppUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShinyAppUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


