# AuditEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventDescription** | **string** | Description of action  | 
**UserId** | **string** | User ID of the actor who made the audit action  | 
**UserDescription** | **string** | Description of the actor  | 
**Time** | **time.Time** | Timestamp in [RFC3339](https://tools.ietf.org/html/rfc3339) format when action was taken  | 
**Action** | **string** | Audit action taken  | 
**Id** | **string** | ID of the audit action  | 

## Methods

### NewAuditEntry

`func NewAuditEntry(eventDescription string, userId string, userDescription string, time time.Time, action string, id string, ) *AuditEntry`

NewAuditEntry instantiates a new AuditEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEntryWithDefaults

`func NewAuditEntryWithDefaults() *AuditEntry`

NewAuditEntryWithDefaults instantiates a new AuditEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventDescription

`func (o *AuditEntry) GetEventDescription() string`

GetEventDescription returns the EventDescription field if non-nil, zero value otherwise.

### GetEventDescriptionOk

`func (o *AuditEntry) GetEventDescriptionOk() (*string, bool)`

GetEventDescriptionOk returns a tuple with the EventDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDescription

`func (o *AuditEntry) SetEventDescription(v string)`

SetEventDescription sets EventDescription field to given value.


### GetUserId

`func (o *AuditEntry) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuditEntry) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuditEntry) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserDescription

`func (o *AuditEntry) GetUserDescription() string`

GetUserDescription returns the UserDescription field if non-nil, zero value otherwise.

### GetUserDescriptionOk

`func (o *AuditEntry) GetUserDescriptionOk() (*string, bool)`

GetUserDescriptionOk returns a tuple with the UserDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDescription

`func (o *AuditEntry) SetUserDescription(v string)`

SetUserDescription sets UserDescription field to given value.


### GetTime

`func (o *AuditEntry) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AuditEntry) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AuditEntry) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetAction

`func (o *AuditEntry) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditEntry) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditEntry) SetAction(v string)`

SetAction sets Action field to given value.


### GetId

`func (o *AuditEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditEntry) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


