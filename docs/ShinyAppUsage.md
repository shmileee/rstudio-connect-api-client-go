# ShinyAppUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Started** | **time.Time** | The timestamp, in [RFC3339](https://tools.ietf.org/html/rfc3339) format, when the user opened the application.  | 
**ContentGuid** | **string** | The GUID, in [RFC4122](https://tools.ietf.org/html/rfc4122) format, of the Shiny application this information pertains to.  | 
**UserGuid** | **string** | The GUID, in [RFC4122](https://tools.ietf.org/html/rfc4122) format, of the user that visited the application.  | 
**DataVersion** | **int32** | The data version the record was recorded with.  The [Shiny Application Events](../admin/historical-information/#shiny-application-events) section of the RStudio Connect Admin Guide explains how to interpret &#x60;data_version&#x60; values.  | 
**Ended** | **time.Time** | The timestamp, in [RFC3339](https://tools.ietf.org/html/rfc3339) format, when the user left the application.  | 

## Methods

### NewShinyAppUsage

`func NewShinyAppUsage(started time.Time, contentGuid string, userGuid string, dataVersion int32, ended time.Time, ) *ShinyAppUsage`

NewShinyAppUsage instantiates a new ShinyAppUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShinyAppUsageWithDefaults

`func NewShinyAppUsageWithDefaults() *ShinyAppUsage`

NewShinyAppUsageWithDefaults instantiates a new ShinyAppUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStarted

`func (o *ShinyAppUsage) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *ShinyAppUsage) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *ShinyAppUsage) SetStarted(v time.Time)`

SetStarted sets Started field to given value.


### GetContentGuid

`func (o *ShinyAppUsage) GetContentGuid() string`

GetContentGuid returns the ContentGuid field if non-nil, zero value otherwise.

### GetContentGuidOk

`func (o *ShinyAppUsage) GetContentGuidOk() (*string, bool)`

GetContentGuidOk returns a tuple with the ContentGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuid

`func (o *ShinyAppUsage) SetContentGuid(v string)`

SetContentGuid sets ContentGuid field to given value.


### GetUserGuid

`func (o *ShinyAppUsage) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *ShinyAppUsage) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *ShinyAppUsage) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.


### GetDataVersion

`func (o *ShinyAppUsage) GetDataVersion() int32`

GetDataVersion returns the DataVersion field if non-nil, zero value otherwise.

### GetDataVersionOk

`func (o *ShinyAppUsage) GetDataVersionOk() (*int32, bool)`

GetDataVersionOk returns a tuple with the DataVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVersion

`func (o *ShinyAppUsage) SetDataVersion(v int32)`

SetDataVersion sets DataVersion field to given value.


### GetEnded

`func (o *ShinyAppUsage) GetEnded() time.Time`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *ShinyAppUsage) GetEndedOk() (*time.Time, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *ShinyAppUsage) SetEnded(v time.Time)`

SetEnded sets Ended field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


