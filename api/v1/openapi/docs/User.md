# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The user&#39;s username  | 
**ActiveTime** | **NullableTime** | The timestamp (in [RFC3339](https://tools.ietf.org/html/rfc3339) format) when the user was last active on the RStudio Connect server  | 
**FirstName** | **string** | The user&#39;s first name  | 
**LastName** | **string** | The user&#39;s last name  | 
**Locked** | **bool** | Whether or not the user is locked  | 
**UserRole** | **string** | The user&#39;s role  | 
**UpdatedTime** | **time.Time** | The timestamp (in [RFC3339](https://tools.ietf.org/html/rfc3339) format) when the user was last updated in the RStudio Connect server  | 
**Confirmed** | **bool** | When &#x60;false&#x60;, the created user must confirm their account through an email. This feature is unique to password authentication.  | 
**CreatedTime** | **time.Time** | The timestamp (in [RFC3339](https://tools.ietf.org/html/rfc3339) format) when the user was created in the RStudio Connect server  | 
**Guid** | **string** | The user&#39;s GUID, or unique identifier, in UUID [RFC4122](https://tools.ietf.org/html/rfc4122) format  | 
**Email** | **string** | The user&#39;s email  | 

## Methods

### NewUser

`func NewUser(username string, activeTime NullableTime, firstName string, lastName string, locked bool, userRole string, updatedTime time.Time, confirmed bool, createdTime time.Time, guid string, email string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetActiveTime

`func (o *User) GetActiveTime() time.Time`

GetActiveTime returns the ActiveTime field if non-nil, zero value otherwise.

### GetActiveTimeOk

`func (o *User) GetActiveTimeOk() (*time.Time, bool)`

GetActiveTimeOk returns a tuple with the ActiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTime

`func (o *User) SetActiveTime(v time.Time)`

SetActiveTime sets ActiveTime field to given value.


### SetActiveTimeNil

`func (o *User) SetActiveTimeNil(b bool)`

 SetActiveTimeNil sets the value for ActiveTime to be an explicit nil

### UnsetActiveTime
`func (o *User) UnsetActiveTime()`

UnsetActiveTime ensures that no value is present for ActiveTime, not even an explicit nil
### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetLocked

`func (o *User) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *User) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *User) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetUserRole

`func (o *User) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *User) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *User) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.


### GetUpdatedTime

`func (o *User) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *User) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *User) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.


### GetConfirmed

`func (o *User) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *User) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *User) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.


### GetCreatedTime

`func (o *User) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *User) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *User) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetGuid

`func (o *User) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *User) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *User) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


