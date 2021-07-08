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

// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	// The user's desired username 
	Username string `json:"username"`
	// The user's first name 
	FirstName *string `json:"first_name,omitempty"`
	// The user's last name 
	LastName *string `json:"last_name,omitempty"`
	// The user's role. If `null` it will default to the role specified by the config setting `Authorization.DefaultUserRole`. 
	UserRole NullableString `json:"user_role,omitempty"`
	// Applies only to password authentication.  - When `true`, the created user will be asked to set their password on first login. The `password` request parameter cannot be set when this parameter is `true`. - When `false`, you must specify the `password`. 
	UserMustSetPassword *bool `json:"user_must_set_password,omitempty"`
	// Applies only to password authentication. Must be at least 6 characters long. Cannot be set when `user_must_set_password` is true. 
	Password *string `json:"password,omitempty"`
	// The user's email. 
	Email *string `json:"email,omitempty"`
	// Applies only to proxied authentication when `ProxyAuth.UniqueIdHeader` is being used.  It is required when SAML authentication is being used. 
	UniqueId *string `json:"unique_id,omitempty"`
}

// NewInlineObject1 instantiates a new InlineObject1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1(username string) *InlineObject1 {
	this := InlineObject1{}
	this.Username = username
	var userMustSetPassword bool = false
	this.UserMustSetPassword = &userMustSetPassword
	return &this
}

// NewInlineObject1WithDefaults instantiates a new InlineObject1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1WithDefaults() *InlineObject1 {
	this := InlineObject1{}
	var userMustSetPassword bool = false
	this.UserMustSetPassword = &userMustSetPassword
	return &this
}

// GetUsername returns the Username field value
func (o *InlineObject1) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *InlineObject1) SetUsername(v string) {
	o.Username = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *InlineObject1) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *InlineObject1) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *InlineObject1) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *InlineObject1) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *InlineObject1) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *InlineObject1) SetLastName(v string) {
	o.LastName = &v
}

// GetUserRole returns the UserRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1) GetUserRole() string {
	if o == nil || o.UserRole.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserRole.Get()
}

// GetUserRoleOk returns a tuple with the UserRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1) GetUserRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserRole.Get(), o.UserRole.IsSet()
}

// HasUserRole returns a boolean if a field has been set.
func (o *InlineObject1) HasUserRole() bool {
	if o != nil && o.UserRole.IsSet() {
		return true
	}

	return false
}

// SetUserRole gets a reference to the given NullableString and assigns it to the UserRole field.
func (o *InlineObject1) SetUserRole(v string) {
	o.UserRole.Set(&v)
}
// SetUserRoleNil sets the value for UserRole to be an explicit nil
func (o *InlineObject1) SetUserRoleNil() {
	o.UserRole.Set(nil)
}

// UnsetUserRole ensures that no value is present for UserRole, not even an explicit nil
func (o *InlineObject1) UnsetUserRole() {
	o.UserRole.Unset()
}

// GetUserMustSetPassword returns the UserMustSetPassword field value if set, zero value otherwise.
func (o *InlineObject1) GetUserMustSetPassword() bool {
	if o == nil || o.UserMustSetPassword == nil {
		var ret bool
		return ret
	}
	return *o.UserMustSetPassword
}

// GetUserMustSetPasswordOk returns a tuple with the UserMustSetPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetUserMustSetPasswordOk() (*bool, bool) {
	if o == nil || o.UserMustSetPassword == nil {
		return nil, false
	}
	return o.UserMustSetPassword, true
}

// HasUserMustSetPassword returns a boolean if a field has been set.
func (o *InlineObject1) HasUserMustSetPassword() bool {
	if o != nil && o.UserMustSetPassword != nil {
		return true
	}

	return false
}

// SetUserMustSetPassword gets a reference to the given bool and assigns it to the UserMustSetPassword field.
func (o *InlineObject1) SetUserMustSetPassword(v bool) {
	o.UserMustSetPassword = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *InlineObject1) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *InlineObject1) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *InlineObject1) SetPassword(v string) {
	o.Password = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineObject1) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineObject1) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InlineObject1) SetEmail(v string) {
	o.Email = &v
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise.
func (o *InlineObject1) GetUniqueId() string {
	if o == nil || o.UniqueId == nil {
		var ret string
		return ret
	}
	return *o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetUniqueIdOk() (*string, bool) {
	if o == nil || o.UniqueId == nil {
		return nil, false
	}
	return o.UniqueId, true
}

// HasUniqueId returns a boolean if a field has been set.
func (o *InlineObject1) HasUniqueId() bool {
	if o != nil && o.UniqueId != nil {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given string and assigns it to the UniqueId field.
func (o *InlineObject1) SetUniqueId(v string) {
	o.UniqueId = &v
}

func (o InlineObject1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.UserRole.IsSet() {
		toSerialize["user_role"] = o.UserRole.Get()
	}
	if o.UserMustSetPassword != nil {
		toSerialize["user_must_set_password"] = o.UserMustSetPassword
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.UniqueId != nil {
		toSerialize["unique_id"] = o.UniqueId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1 struct {
	value *InlineObject1
	isSet bool
}

func (v NullableInlineObject1) Get() *InlineObject1 {
	return v.value
}

func (v *NullableInlineObject1) Set(val *InlineObject1) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1(val *InlineObject1) *NullableInlineObject1 {
	return &NullableInlineObject1{value: val, isSet: true}
}

func (v NullableInlineObject1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


