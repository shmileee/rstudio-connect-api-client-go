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

// GroupWithTicket struct for GroupWithTicket
type GroupWithTicket struct {
	// This value is for actions that require a `temp_ticket`, such as adding an LDAP group to the Connect server. 
	TempTicket string `json:"temp_ticket"`
	// The group's unique identifier in [RFC4122](https://tools.ietf.org/html/rfc4122) format. When a group does not exist in the RStudio Connect server, this property will be `null`. 
	Guid NullableString `json:"guid"`
	// The group name 
	Name string `json:"name"`
}

// NewGroupWithTicket instantiates a new GroupWithTicket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupWithTicket(tempTicket string, guid NullableString, name string) *GroupWithTicket {
	this := GroupWithTicket{}
	this.TempTicket = tempTicket
	this.Guid = guid
	this.Name = name
	return &this
}

// NewGroupWithTicketWithDefaults instantiates a new GroupWithTicket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithTicketWithDefaults() *GroupWithTicket {
	this := GroupWithTicket{}
	return &this
}

// GetTempTicket returns the TempTicket field value
func (o *GroupWithTicket) GetTempTicket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TempTicket
}

// GetTempTicketOk returns a tuple with the TempTicket field value
// and a boolean to check if the value has been set.
func (o *GroupWithTicket) GetTempTicketOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TempTicket, true
}

// SetTempTicket sets field value
func (o *GroupWithTicket) SetTempTicket(v string) {
	o.TempTicket = v
}

// GetGuid returns the Guid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupWithTicket) GetGuid() string {
	if o == nil || o.Guid.Get() == nil {
		var ret string
		return ret
	}

	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupWithTicket) GetGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// SetGuid sets field value
func (o *GroupWithTicket) SetGuid(v string) {
	o.Guid.Set(&v)
}

// GetName returns the Name field value
func (o *GroupWithTicket) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupWithTicket) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupWithTicket) SetName(v string) {
	o.Name = v
}

func (o GroupWithTicket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["temp_ticket"] = o.TempTicket
	}
	if true {
		toSerialize["guid"] = o.Guid.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGroupWithTicket struct {
	value *GroupWithTicket
	isSet bool
}

func (v NullableGroupWithTicket) Get() *GroupWithTicket {
	return v.value
}

func (v *NullableGroupWithTicket) Set(val *GroupWithTicket) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupWithTicket) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupWithTicket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupWithTicket(val *GroupWithTicket) *NullableGroupWithTicket {
	return &NullableGroupWithTicket{value: val, isSet: true}
}

func (v NullableGroupWithTicket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupWithTicket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


