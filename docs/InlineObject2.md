# InlineObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The user&#39;s new username  | [optional] 
**FirstName** | Pointer to **string** | The user&#39;s new first name  | [optional] 
**LastName** | Pointer to **string** | The user&#39;s new last name  | [optional] 
**Email** | Pointer to **string** | The user&#39;s new email.  | [optional] 
**UserRole** | Pointer to **string** | The user&#39;s new role. Note that you can only downgrade yourself. Administrators can change other users&#39; roles to any valid role.  | [optional] 

## Methods

### NewInlineObject2

`func NewInlineObject2() *InlineObject2`

NewInlineObject2 instantiates a new InlineObject2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject2WithDefaults

`func NewInlineObject2WithDefaults() *InlineObject2`

NewInlineObject2WithDefaults instantiates a new InlineObject2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *InlineObject2) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InlineObject2) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InlineObject2) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InlineObject2) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *InlineObject2) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InlineObject2) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InlineObject2) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *InlineObject2) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *InlineObject2) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InlineObject2) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InlineObject2) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InlineObject2) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *InlineObject2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineObject2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineObject2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineObject2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUserRole

`func (o *InlineObject2) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *InlineObject2) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *InlineObject2) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *InlineObject2) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


