# EditableUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The user&#39;s username  | 
**FirstName** | **string** | The user&#39;s first name  | 
**LastName** | **string** | The user&#39;s last name  | 
**UserRole** | **string** | The user&#39;s role  | 
**UpdatedTime** | **time.Time** | The timestamp (in [RFC3339](https://tools.ietf.org/html/rfc3339) format) when the user was last updated in the RStudio Connect server  | 
**Email** | **string** | The user&#39;s email  | 

## Methods

### NewEditableUser

`func NewEditableUser(username string, firstName string, lastName string, userRole string, updatedTime time.Time, email string, ) *EditableUser`

NewEditableUser instantiates a new EditableUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditableUserWithDefaults

`func NewEditableUserWithDefaults() *EditableUser`

NewEditableUserWithDefaults instantiates a new EditableUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *EditableUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EditableUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EditableUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetFirstName

`func (o *EditableUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EditableUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EditableUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *EditableUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EditableUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EditableUser) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetUserRole

`func (o *EditableUser) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *EditableUser) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *EditableUser) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.


### GetUpdatedTime

`func (o *EditableUser) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *EditableUser) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *EditableUser) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.


### GetEmail

`func (o *EditableUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EditableUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EditableUser) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


