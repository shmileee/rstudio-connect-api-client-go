# AuditPagerCursors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | **NullableString** | A cursor ID that can be used with the &#x60;next&#x60; query param to get the next page of results.  | 
**Previous** | **NullableString** | A cursor ID that can be used with the &#x60;previous&#x60; query param to get the previous page of results.  | 

## Methods

### NewAuditPagerCursors

`func NewAuditPagerCursors(next NullableString, previous NullableString, ) *AuditPagerCursors`

NewAuditPagerCursors instantiates a new AuditPagerCursors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditPagerCursorsWithDefaults

`func NewAuditPagerCursorsWithDefaults() *AuditPagerCursors`

NewAuditPagerCursorsWithDefaults instantiates a new AuditPagerCursors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *AuditPagerCursors) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AuditPagerCursors) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AuditPagerCursors) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *AuditPagerCursors) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *AuditPagerCursors) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *AuditPagerCursors) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *AuditPagerCursors) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *AuditPagerCursors) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### SetPreviousNil

`func (o *AuditPagerCursors) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *AuditPagerCursors) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


