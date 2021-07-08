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

// EditableUser struct for EditableUser
type EditableUser struct {
	// The user's username 
	Username string `json:"username"`
	// The user's first name 
	FirstName string `json:"first_name"`
	// The user's last name 
	LastName string `json:"last_name"`
	// The user's role 
	UserRole string `json:"user_role"`
	// The timestamp (in [RFC3339](https://tools.ietf.org/html/rfc3339) format) when the user was last updated in the RStudio Connect server 
	UpdatedTime time.Time `json:"updated_time"`
	// The user's email 
	Email string `json:"email"`
}

// NewEditableUser instantiates a new EditableUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditableUser(username string, firstName string, lastName string, userRole string, updatedTime time.Time, email string) *EditableUser {
	this := EditableUser{}
	this.Username = username
	this.FirstName = firstName
	this.LastName = lastName
	this.UserRole = userRole
	this.UpdatedTime = updatedTime
	this.Email = email
	return &this
}

// NewEditableUserWithDefaults instantiates a new EditableUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditableUserWithDefaults() *EditableUser {
	this := EditableUser{}
	return &this
}

// GetUsername returns the Username field value
func (o *EditableUser) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *EditableUser) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *EditableUser) SetUsername(v string) {
	o.Username = v
}

// GetFirstName returns the FirstName field value
func (o *EditableUser) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *EditableUser) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *EditableUser) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *EditableUser) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *EditableUser) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *EditableUser) SetLastName(v string) {
	o.LastName = v
}

// GetUserRole returns the UserRole field value
func (o *EditableUser) GetUserRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserRole
}

// GetUserRoleOk returns a tuple with the UserRole field value
// and a boolean to check if the value has been set.
func (o *EditableUser) GetUserRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserRole, true
}

// SetUserRole sets field value
func (o *EditableUser) SetUserRole(v string) {
	o.UserRole = v
}

// GetUpdatedTime returns the UpdatedTime field value
func (o *EditableUser) GetUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value
// and a boolean to check if the value has been set.
func (o *EditableUser) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedTime, true
}

// SetUpdatedTime sets field value
func (o *EditableUser) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = v
}

// GetEmail returns the Email field value
func (o *EditableUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *EditableUser) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *EditableUser) SetEmail(v string) {
	o.Email = v
}

func (o EditableUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["first_name"] = o.FirstName
	}
	if true {
		toSerialize["last_name"] = o.LastName
	}
	if true {
		toSerialize["user_role"] = o.UserRole
	}
	if true {
		toSerialize["updated_time"] = o.UpdatedTime
	}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableEditableUser struct {
	value *EditableUser
	isSet bool
}

func (v NullableEditableUser) Get() *EditableUser {
	return v.value
}

func (v *NullableEditableUser) Set(val *EditableUser) {
	v.value = val
	v.isSet = true
}

func (v NullableEditableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableEditableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditableUser(val *EditableUser) *NullableEditableUser {
	return &NullableEditableUser{value: val, isSet: true}
}

func (v NullableEditableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


