# ContentVisitLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**ContentVisitPager**](ContentVisitPager.md) |  | [optional] 
**Results** | Pointer to [**[]ContentVisit**](ContentVisit.md) | The content visit logs  | [optional] 

## Methods

### NewContentVisitLogs

`func NewContentVisitLogs() *ContentVisitLogs`

NewContentVisitLogs instantiates a new ContentVisitLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentVisitLogsWithDefaults

`func NewContentVisitLogsWithDefaults() *ContentVisitLogs`

NewContentVisitLogsWithDefaults instantiates a new ContentVisitLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *ContentVisitLogs) GetPaging() ContentVisitPager`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ContentVisitLogs) GetPagingOk() (*ContentVisitPager, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ContentVisitLogs) SetPaging(v ContentVisitPager)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ContentVisitLogs) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *ContentVisitLogs) GetResults() []ContentVisit`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ContentVisitLogs) GetResultsOk() (*[]ContentVisit, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ContentVisitLogs) SetResults(v []ContentVisit)`

SetResults sets Results field to given value.

### HasResults

`func (o *ContentVisitLogs) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


