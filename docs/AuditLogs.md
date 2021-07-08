# AuditLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**AuditPager**](AuditPager.md) |  | [optional] 
**Results** | Pointer to [**[]AuditEntry**](AuditEntry.md) | The audit logs  | [optional] 

## Methods

### NewAuditLogs

`func NewAuditLogs() *AuditLogs`

NewAuditLogs instantiates a new AuditLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogsWithDefaults

`func NewAuditLogsWithDefaults() *AuditLogs`

NewAuditLogsWithDefaults instantiates a new AuditLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *AuditLogs) GetPaging() AuditPager`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *AuditLogs) GetPagingOk() (*AuditPager, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *AuditLogs) SetPaging(v AuditPager)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *AuditLogs) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *AuditLogs) GetResults() []AuditEntry`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AuditLogs) GetResultsOk() (*[]AuditEntry, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AuditLogs) SetResults(v []AuditEntry)`

SetResults sets Results field to given value.

### HasResults

`func (o *AuditLogs) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


