# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | **string** | The unique identifier  | 
**Name** | **string** | The group name  | 
**OwnerGuid** | **NullableString** | The group owner&#39;s unique identifier. When using LDAP, or Proxied authentication with group provisioning enabled this property will always be &#x60;null&#x60;.  | 

## Methods

### NewGroup

`func NewGroup(guid string, name string, ownerGuid NullableString, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *Group) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Group) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Group) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerGuid

`func (o *Group) GetOwnerGuid() string`

GetOwnerGuid returns the OwnerGuid field if non-nil, zero value otherwise.

### GetOwnerGuidOk

`func (o *Group) GetOwnerGuidOk() (*string, bool)`

GetOwnerGuidOk returns a tuple with the OwnerGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGuid

`func (o *Group) SetOwnerGuid(v string)`

SetOwnerGuid sets OwnerGuid field to given value.


### SetOwnerGuidNil

`func (o *Group) SetOwnerGuidNil(b bool)`

 SetOwnerGuidNil sets the value for OwnerGuid to be an explicit nil

### UnsetOwnerGuid
`func (o *Group) UnsetOwnerGuid()`

UnsetOwnerGuid ensures that no value is present for OwnerGuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


