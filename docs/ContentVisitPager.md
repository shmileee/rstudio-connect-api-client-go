# ContentVisitPager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | **NullableString** | A full URL of the next page of results when the original request was made.  It will be &#x60;null&#x60; if the current response represents the last page of results.  | 
**Previous** | **NullableString** | A full URL of the previous page of results when the original request was made.  It will be &#x60;null&#x60; if the current response represents the first page of results.  | 
**Last** | **NullableString** | A full URL of the last page of results.  It will be &#x60;null&#x60; if the current response represents the last page of results.  | 
**Cursors** | [**AuditPagerCursors**](AuditPagerCursors.md) |  | 
**First** | **NullableString** | A full URL of the first page of results.  It will be &#x60;null&#x60; if the current response represents the first page of results.  | 

## Methods

### NewContentVisitPager

`func NewContentVisitPager(next NullableString, previous NullableString, last NullableString, cursors AuditPagerCursors, first NullableString, ) *ContentVisitPager`

NewContentVisitPager instantiates a new ContentVisitPager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentVisitPagerWithDefaults

`func NewContentVisitPagerWithDefaults() *ContentVisitPager`

NewContentVisitPagerWithDefaults instantiates a new ContentVisitPager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *ContentVisitPager) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ContentVisitPager) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ContentVisitPager) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *ContentVisitPager) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *ContentVisitPager) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *ContentVisitPager) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *ContentVisitPager) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *ContentVisitPager) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### SetPreviousNil

`func (o *ContentVisitPager) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *ContentVisitPager) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetLast

`func (o *ContentVisitPager) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *ContentVisitPager) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *ContentVisitPager) SetLast(v string)`

SetLast sets Last field to given value.


### SetLastNil

`func (o *ContentVisitPager) SetLastNil(b bool)`

 SetLastNil sets the value for Last to be an explicit nil

### UnsetLast
`func (o *ContentVisitPager) UnsetLast()`

UnsetLast ensures that no value is present for Last, not even an explicit nil
### GetCursors

`func (o *ContentVisitPager) GetCursors() AuditPagerCursors`

GetCursors returns the Cursors field if non-nil, zero value otherwise.

### GetCursorsOk

`func (o *ContentVisitPager) GetCursorsOk() (*AuditPagerCursors, bool)`

GetCursorsOk returns a tuple with the Cursors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursors

`func (o *ContentVisitPager) SetCursors(v AuditPagerCursors)`

SetCursors sets Cursors field to given value.


### GetFirst

`func (o *ContentVisitPager) GetFirst() string`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *ContentVisitPager) GetFirstOk() (*string, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *ContentVisitPager) SetFirst(v string)`

SetFirst sets First field to given value.


### SetFirstNil

`func (o *ContentVisitPager) SetFirstNil(b bool)`

 SetFirstNil sets the value for First to be an explicit nil

### UnsetFirst
`func (o *ContentVisitPager) UnsetFirst()`

UnsetFirst ensures that no value is present for First, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


