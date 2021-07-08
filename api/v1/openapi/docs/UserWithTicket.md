# UserWithTicket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The user&#39;s username  | 
**ActiveTime** | **NullableTime** | The timestamp (in [RFC3339](https://tools.ietf.org/html/rfc3339) format) when the user was last active on the RStudio Connect server  | 
**FirstName** | **string** | The user&#39;s first name  | 
**LastName** | **string** | The user&#39;s last name  | 
**Locked** | **bool** | Whether or not the user is locked  | 
**TempTicket** | **string** | This value is for actions that require a &#x60;temp_ticket&#x60;, such as adding an LDAP or OAuth2 user to the Connect server.  | 
**UserRole** | **string** | The user&#39;s role  | 
**UpdatedTime** | **time.Time** | The timestamp (in [RFC3339](https://tools.ietf.org/html/rfc3339) format) when the user was last updated in the RStudio Connect server  | 
**Confirmed** | **bool** | Whether or not the user is confirmed. This property will always be &#x60;true&#x60; for all authentication providers except password.  | 
**CreatedTime** | **time.Time** | The timestamp (in [RFC3339](https://tools.ietf.org/html/rfc3339) format) when the user was created in the RStudio Connect server  | 
**Guid** | **NullableString** | The user&#39;s GUID, or unique identifier in [RFC4122](https://tools.ietf.org/html/rfc4122) format. When a user does not exist in the RStudio Connect server, this property will be &#x60;null&#x60;.  | 
**Email** | **string** | The user&#39;s email  | 

## Methods

### NewUserWithTicket

`func NewUserWithTicket(username string, activeTime NullableTime, firstName string, lastName string, locked bool, tempTicket string, userRole string, updatedTime time.Time, confirmed bool, createdTime time.Time, guid NullableString, email string, ) *UserWithTicket`

NewUserWithTicket instantiates a new UserWithTicket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithTicketWithDefaults

`func NewUserWithTicketWithDefaults() *UserWithTicket`

NewUserWithTicketWithDefaults instantiates a new UserWithTicket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserWithTicket) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserWithTicket) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserWithTicket) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetActiveTime

`func (o *UserWithTicket) GetActiveTime() time.Time`

GetActiveTime returns the ActiveTime field if non-nil, zero value otherwise.

### GetActiveTimeOk

`func (o *UserWithTicket) GetActiveTimeOk() (*time.Time, bool)`

GetActiveTimeOk returns a tuple with the ActiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTime

`func (o *UserWithTicket) SetActiveTime(v time.Time)`

SetActiveTime sets ActiveTime field to given value.


### SetActiveTimeNil

`func (o *UserWithTicket) SetActiveTimeNil(b bool)`

 SetActiveTimeNil sets the value for ActiveTime to be an explicit nil

### UnsetActiveTime
`func (o *UserWithTicket) UnsetActiveTime()`

UnsetActiveTime ensures that no value is present for ActiveTime, not even an explicit nil
### GetFirstName

`func (o *UserWithTicket) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserWithTicket) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserWithTicket) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UserWithTicket) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserWithTicket) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserWithTicket) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetLocked

`func (o *UserWithTicket) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *UserWithTicket) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *UserWithTicket) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetTempTicket

`func (o *UserWithTicket) GetTempTicket() string`

GetTempTicket returns the TempTicket field if non-nil, zero value otherwise.

### GetTempTicketOk

`func (o *UserWithTicket) GetTempTicketOk() (*string, bool)`

GetTempTicketOk returns a tuple with the TempTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempTicket

`func (o *UserWithTicket) SetTempTicket(v string)`

SetTempTicket sets TempTicket field to given value.


### GetUserRole

`func (o *UserWithTicket) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *UserWithTicket) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *UserWithTicket) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.


### GetUpdatedTime

`func (o *UserWithTicket) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *UserWithTicket) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *UserWithTicket) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.


### GetConfirmed

`func (o *UserWithTicket) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *UserWithTicket) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *UserWithTicket) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.


### GetCreatedTime

`func (o *UserWithTicket) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *UserWithTicket) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *UserWithTicket) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetGuid

`func (o *UserWithTicket) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *UserWithTicket) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *UserWithTicket) SetGuid(v string)`

SetGuid sets Guid field to given value.


### SetGuidNil

`func (o *UserWithTicket) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *UserWithTicket) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetEmail

`func (o *UserWithTicket) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserWithTicket) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserWithTicket) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


