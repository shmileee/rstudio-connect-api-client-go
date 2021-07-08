# InlineObject1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The user&#39;s desired username  | 
**FirstName** | Pointer to **string** | The user&#39;s first name  | [optional] 
**LastName** | Pointer to **string** | The user&#39;s last name  | [optional] 
**UserRole** | Pointer to **NullableString** | The user&#39;s role. If &#x60;null&#x60; it will default to the role specified by the config setting &#x60;Authorization.DefaultUserRole&#x60;.  | [optional] 
**UserMustSetPassword** | Pointer to **bool** | Applies only to password authentication.  - When &#x60;true&#x60;, the created user will be asked to set their password on first login. The &#x60;password&#x60; request parameter cannot be set when this parameter is &#x60;true&#x60;. - When &#x60;false&#x60;, you must specify the &#x60;password&#x60;.  | [optional] [default to false]
**Password** | Pointer to **string** | Applies only to password authentication. Must be at least 6 characters long. Cannot be set when &#x60;user_must_set_password&#x60; is true.  | [optional] 
**Email** | Pointer to **string** | The user&#39;s email.  | [optional] 
**UniqueId** | Pointer to **string** | Applies only to proxied authentication when &#x60;ProxyAuth.UniqueIdHeader&#x60; is being used.  It is required when SAML authentication is being used.  | [optional] 

## Methods

### NewInlineObject1

`func NewInlineObject1(username string, ) *InlineObject1`

NewInlineObject1 instantiates a new InlineObject1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject1WithDefaults

`func NewInlineObject1WithDefaults() *InlineObject1`

NewInlineObject1WithDefaults instantiates a new InlineObject1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *InlineObject1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InlineObject1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InlineObject1) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetFirstName

`func (o *InlineObject1) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InlineObject1) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InlineObject1) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *InlineObject1) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *InlineObject1) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InlineObject1) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InlineObject1) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InlineObject1) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUserRole

`func (o *InlineObject1) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *InlineObject1) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *InlineObject1) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *InlineObject1) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### SetUserRoleNil

`func (o *InlineObject1) SetUserRoleNil(b bool)`

 SetUserRoleNil sets the value for UserRole to be an explicit nil

### UnsetUserRole
`func (o *InlineObject1) UnsetUserRole()`

UnsetUserRole ensures that no value is present for UserRole, not even an explicit nil
### GetUserMustSetPassword

`func (o *InlineObject1) GetUserMustSetPassword() bool`

GetUserMustSetPassword returns the UserMustSetPassword field if non-nil, zero value otherwise.

### GetUserMustSetPasswordOk

`func (o *InlineObject1) GetUserMustSetPasswordOk() (*bool, bool)`

GetUserMustSetPasswordOk returns a tuple with the UserMustSetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMustSetPassword

`func (o *InlineObject1) SetUserMustSetPassword(v bool)`

SetUserMustSetPassword sets UserMustSetPassword field to given value.

### HasUserMustSetPassword

`func (o *InlineObject1) HasUserMustSetPassword() bool`

HasUserMustSetPassword returns a boolean if a field has been set.

### GetPassword

`func (o *InlineObject1) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineObject1) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineObject1) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlineObject1) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEmail

`func (o *InlineObject1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineObject1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineObject1) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineObject1) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUniqueId

`func (o *InlineObject1) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *InlineObject1) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *InlineObject1) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *InlineObject1) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


