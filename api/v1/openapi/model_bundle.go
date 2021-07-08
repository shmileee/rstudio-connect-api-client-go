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

// Bundle Content published to RStudio Connect is encapsulated in a \"bundle\" that contains the source code and data necessary to execute the content. An application or report is updated by uploading a new bundle. 
type Bundle struct {
	// The identifier of the owning content. 
	ContentGuid *string `json:"content_guid,omitempty"`
	// The version of the R interpreter used when last restoring this bundle. The value `null` represents that an R interpreter is not used by this bundle or that the R package environment has not been successfully restored.  R version is not disclosed to users with a \"viewer\" role and returned with the value `null`. 
	RVersion NullableString `json:"r_version,omitempty"`
	// Indicates if this bundle is active for the owning content. 
	Active *bool `json:"active,omitempty"`
	// The timestamp (RFC3339) of when this bundle was created. 
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// The version of the Python interpreter used when last restoring this bundle.  The value `null` represents that a Python interpreter is not used by this bundle or that the Python package environment has not been successfully restored.  Python version is not disclosed to users with a \"viewer\" role and returned with the value `null`. 
	PyVersion NullableString `json:"py_version,omitempty"`
	// The identifier for this bundle. 
	Id *string `json:"id,omitempty"`
	// On-disk size in bytes of the tar.gz file associated with this bundle. Zero when there is no on-disk file. 
	Size *float32 `json:"size,omitempty"`
}

// NewBundle instantiates a new Bundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundle() *Bundle {
	this := Bundle{}
	return &this
}

// NewBundleWithDefaults instantiates a new Bundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleWithDefaults() *Bundle {
	this := Bundle{}
	return &this
}

// GetContentGuid returns the ContentGuid field value if set, zero value otherwise.
func (o *Bundle) GetContentGuid() string {
	if o == nil || o.ContentGuid == nil {
		var ret string
		return ret
	}
	return *o.ContentGuid
}

// GetContentGuidOk returns a tuple with the ContentGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetContentGuidOk() (*string, bool) {
	if o == nil || o.ContentGuid == nil {
		return nil, false
	}
	return o.ContentGuid, true
}

// HasContentGuid returns a boolean if a field has been set.
func (o *Bundle) HasContentGuid() bool {
	if o != nil && o.ContentGuid != nil {
		return true
	}

	return false
}

// SetContentGuid gets a reference to the given string and assigns it to the ContentGuid field.
func (o *Bundle) SetContentGuid(v string) {
	o.ContentGuid = &v
}

// GetRVersion returns the RVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bundle) GetRVersion() string {
	if o == nil || o.RVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.RVersion.Get()
}

// GetRVersionOk returns a tuple with the RVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bundle) GetRVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RVersion.Get(), o.RVersion.IsSet()
}

// HasRVersion returns a boolean if a field has been set.
func (o *Bundle) HasRVersion() bool {
	if o != nil && o.RVersion.IsSet() {
		return true
	}

	return false
}

// SetRVersion gets a reference to the given NullableString and assigns it to the RVersion field.
func (o *Bundle) SetRVersion(v string) {
	o.RVersion.Set(&v)
}
// SetRVersionNil sets the value for RVersion to be an explicit nil
func (o *Bundle) SetRVersionNil() {
	o.RVersion.Set(nil)
}

// UnsetRVersion ensures that no value is present for RVersion, not even an explicit nil
func (o *Bundle) UnsetRVersion() {
	o.RVersion.Unset()
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Bundle) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Bundle) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Bundle) SetActive(v bool) {
	o.Active = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Bundle) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Bundle) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *Bundle) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetPyVersion returns the PyVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bundle) GetPyVersion() string {
	if o == nil || o.PyVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.PyVersion.Get()
}

// GetPyVersionOk returns a tuple with the PyVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bundle) GetPyVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PyVersion.Get(), o.PyVersion.IsSet()
}

// HasPyVersion returns a boolean if a field has been set.
func (o *Bundle) HasPyVersion() bool {
	if o != nil && o.PyVersion.IsSet() {
		return true
	}

	return false
}

// SetPyVersion gets a reference to the given NullableString and assigns it to the PyVersion field.
func (o *Bundle) SetPyVersion(v string) {
	o.PyVersion.Set(&v)
}
// SetPyVersionNil sets the value for PyVersion to be an explicit nil
func (o *Bundle) SetPyVersionNil() {
	o.PyVersion.Set(nil)
}

// UnsetPyVersion ensures that no value is present for PyVersion, not even an explicit nil
func (o *Bundle) UnsetPyVersion() {
	o.PyVersion.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Bundle) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Bundle) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Bundle) SetId(v string) {
	o.Id = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Bundle) GetSize() float32 {
	if o == nil || o.Size == nil {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetSizeOk() (*float32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Bundle) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *Bundle) SetSize(v float32) {
	o.Size = &v
}

func (o Bundle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentGuid != nil {
		toSerialize["content_guid"] = o.ContentGuid
	}
	if o.RVersion.IsSet() {
		toSerialize["r_version"] = o.RVersion.Get()
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.CreatedTime != nil {
		toSerialize["created_time"] = o.CreatedTime
	}
	if o.PyVersion.IsSet() {
		toSerialize["py_version"] = o.PyVersion.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableBundle struct {
	value *Bundle
	isSet bool
}

func (v NullableBundle) Get() *Bundle {
	return v.value
}

func (v *NullableBundle) Set(val *Bundle) {
	v.value = val
	v.isSet = true
}

func (v NullableBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundle(val *Bundle) *NullableBundle {
	return &NullableBundle{value: val, isSet: true}
}

func (v NullableBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


