# GroupWithTicket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TempTicket** | **string** | This value is for actions that require a &#x60;temp_ticket&#x60;, such as adding an LDAP group to the Connect server.  | 
**Guid** | **NullableString** | The group&#39;s unique identifier in [RFC4122](https://tools.ietf.org/html/rfc4122) format. When a group does not exist in the RStudio Connect server, this property will be &#x60;null&#x60;.  | 
**Name** | **string** | The group name  | 

## Methods

### NewGroupWithTicket

`func NewGroupWithTicket(tempTicket string, guid NullableString, name string, ) *GroupWithTicket`

NewGroupWithTicket instantiates a new GroupWithTicket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithTicketWithDefaults

`func NewGroupWithTicketWithDefaults() *GroupWithTicket`

NewGroupWithTicketWithDefaults instantiates a new GroupWithTicket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTempTicket

`func (o *GroupWithTicket) GetTempTicket() string`

GetTempTicket returns the TempTicket field if non-nil, zero value otherwise.

### GetTempTicketOk

`func (o *GroupWithTicket) GetTempTicketOk() (*string, bool)`

GetTempTicketOk returns a tuple with the TempTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempTicket

`func (o *GroupWithTicket) SetTempTicket(v string)`

SetTempTicket sets TempTicket field to given value.


### GetGuid

`func (o *GroupWithTicket) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *GroupWithTicket) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *GroupWithTicket) SetGuid(v string)`

SetGuid sets Guid field to given value.


### SetGuidNil

`func (o *GroupWithTicket) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *GroupWithTicket) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetName

`func (o *GroupWithTicket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupWithTicket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupWithTicket) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


