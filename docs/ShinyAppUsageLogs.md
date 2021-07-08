# ShinyAppUsageLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**ShinyAppUsagePager**](ShinyAppUsagePager.md) |  | [optional] 
**Results** | Pointer to [**[]ShinyAppUsage**](ShinyAppUsage.md) | The Shiny application usage logs  | [optional] 

## Methods

### NewShinyAppUsageLogs

`func NewShinyAppUsageLogs() *ShinyAppUsageLogs`

NewShinyAppUsageLogs instantiates a new ShinyAppUsageLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShinyAppUsageLogsWithDefaults

`func NewShinyAppUsageLogsWithDefaults() *ShinyAppUsageLogs`

NewShinyAppUsageLogsWithDefaults instantiates a new ShinyAppUsageLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *ShinyAppUsageLogs) GetPaging() ShinyAppUsagePager`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ShinyAppUsageLogs) GetPagingOk() (*ShinyAppUsagePager, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ShinyAppUsageLogs) SetPaging(v ShinyAppUsagePager)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ShinyAppUsageLogs) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *ShinyAppUsageLogs) GetResults() []ShinyAppUsage`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ShinyAppUsageLogs) GetResultsOk() (*[]ShinyAppUsage, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ShinyAppUsageLogs) SetResults(v []ShinyAppUsage)`

SetResults sets Results field to given value.

### HasResults

`func (o *ShinyAppUsageLogs) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


