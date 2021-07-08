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

// Task The task tracks the output and status of some operation being performed by RStudio Connect. It is most commonly used to contain details about the progress of a deployment operation. 
type Task struct {
	// Numeric indication as to the cause of an error. Non-zero when an error has occured. 
	Code *int32 `json:"code,omitempty"`
	// The total number of output lines produced so far. Use as the value to `first` in the next request to only fetch output lines beyond what you have already received. 
	Last *int32 `json:"last,omitempty"`
	// Indicates that a task has completed. 
	Finished *bool `json:"finished,omitempty"`
	// Description of the error. An empty string when no error has occurred. 
	Error *string `json:"error,omitempty"`
	// An array containing lines of output produced by the task. The initial line of output is dictated by the `first` input parameter. The offset of the last output line is indicated by the `last` response field. 
	Output *[]string `json:"output,omitempty"`
	// The identifier for this task. 
	Id *string `json:"id,omitempty"`
}

// NewTask instantiates a new Task object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTask() *Task {
	this := Task{}
	return &this
}

// NewTaskWithDefaults instantiates a new Task object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Task) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Task) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Task) SetCode(v int32) {
	o.Code = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *Task) GetLast() int32 {
	if o == nil || o.Last == nil {
		var ret int32
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetLastOk() (*int32, bool) {
	if o == nil || o.Last == nil {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *Task) HasLast() bool {
	if o != nil && o.Last != nil {
		return true
	}

	return false
}

// SetLast gets a reference to the given int32 and assigns it to the Last field.
func (o *Task) SetLast(v int32) {
	o.Last = &v
}

// GetFinished returns the Finished field value if set, zero value otherwise.
func (o *Task) GetFinished() bool {
	if o == nil || o.Finished == nil {
		var ret bool
		return ret
	}
	return *o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetFinishedOk() (*bool, bool) {
	if o == nil || o.Finished == nil {
		return nil, false
	}
	return o.Finished, true
}

// HasFinished returns a boolean if a field has been set.
func (o *Task) HasFinished() bool {
	if o != nil && o.Finished != nil {
		return true
	}

	return false
}

// SetFinished gets a reference to the given bool and assigns it to the Finished field.
func (o *Task) SetFinished(v bool) {
	o.Finished = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *Task) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *Task) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *Task) SetError(v string) {
	o.Error = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *Task) GetOutput() []string {
	if o == nil || o.Output == nil {
		var ret []string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetOutputOk() (*[]string, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *Task) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given []string and assigns it to the Output field.
func (o *Task) SetOutput(v []string) {
	o.Output = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Task) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Task) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Task) SetId(v string) {
	o.Id = &v
}

func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Last != nil {
		toSerialize["last"] = o.Last
	}
	if o.Finished != nil {
		toSerialize["finished"] = o.Finished
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableTask struct {
	value *Task
	isSet bool
}

func (v NullableTask) Get() *Task {
	return v.value
}

func (v *NullableTask) Set(val *Task) {
	v.value = val
	v.isSet = true
}

func (v NullableTask) IsSet() bool {
	return v.isSet
}

func (v *NullableTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTask(val *Task) *NullableTask {
	return &NullableTask{value: val, isSet: true}
}

func (v NullableTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


