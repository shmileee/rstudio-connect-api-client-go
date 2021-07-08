# ContentVisit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentGuid** | **string** | The GUID, in [RFC4122](https://tools.ietf.org/html/rfc4122) format, of the content this information pertains to.  | 
**RenderingId** | Pointer to **int32** | The ID of the rendering the user visited.  This will be &#x60;null&#x60; for static content.  | [optional] 
**VariantKey** | Pointer to **string** | The key of the variant the user visited.  This will be &#x60;null&#x60; for static content.  | [optional] 
**DataVersion** | **int32** | The data version the record was recorded with.  The [Rendered and Static Content Visit Events](../admin/historical-information/#rendered-and-static-content-visit-events) section of the RStudio Connect Admin Guide explains how to interpret &#x60;data_version&#x60; values.  | 
**Time** | **time.Time** | The timestamp, in [RFC3339](https://tools.ietf.org/html/rfc3339) format, when the user visited the content.  | 
**BundleId** | **int32** | The ID of the particular bundle used.  | 
**UserGuid** | **string** | The GUID, in [RFC4122](https://tools.ietf.org/html/rfc4122) format, of the user that visited the content.  | 

## Methods

### NewContentVisit

`func NewContentVisit(contentGuid string, dataVersion int32, time time.Time, bundleId int32, userGuid string, ) *ContentVisit`

NewContentVisit instantiates a new ContentVisit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentVisitWithDefaults

`func NewContentVisitWithDefaults() *ContentVisit`

NewContentVisitWithDefaults instantiates a new ContentVisit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentGuid

`func (o *ContentVisit) GetContentGuid() string`

GetContentGuid returns the ContentGuid field if non-nil, zero value otherwise.

### GetContentGuidOk

`func (o *ContentVisit) GetContentGuidOk() (*string, bool)`

GetContentGuidOk returns a tuple with the ContentGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuid

`func (o *ContentVisit) SetContentGuid(v string)`

SetContentGuid sets ContentGuid field to given value.


### GetRenderingId

`func (o *ContentVisit) GetRenderingId() int32`

GetRenderingId returns the RenderingId field if non-nil, zero value otherwise.

### GetRenderingIdOk

`func (o *ContentVisit) GetRenderingIdOk() (*int32, bool)`

GetRenderingIdOk returns a tuple with the RenderingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderingId

`func (o *ContentVisit) SetRenderingId(v int32)`

SetRenderingId sets RenderingId field to given value.

### HasRenderingId

`func (o *ContentVisit) HasRenderingId() bool`

HasRenderingId returns a boolean if a field has been set.

### GetVariantKey

`func (o *ContentVisit) GetVariantKey() string`

GetVariantKey returns the VariantKey field if non-nil, zero value otherwise.

### GetVariantKeyOk

`func (o *ContentVisit) GetVariantKeyOk() (*string, bool)`

GetVariantKeyOk returns a tuple with the VariantKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantKey

`func (o *ContentVisit) SetVariantKey(v string)`

SetVariantKey sets VariantKey field to given value.

### HasVariantKey

`func (o *ContentVisit) HasVariantKey() bool`

HasVariantKey returns a boolean if a field has been set.

### GetDataVersion

`func (o *ContentVisit) GetDataVersion() int32`

GetDataVersion returns the DataVersion field if non-nil, zero value otherwise.

### GetDataVersionOk

`func (o *ContentVisit) GetDataVersionOk() (*int32, bool)`

GetDataVersionOk returns a tuple with the DataVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVersion

`func (o *ContentVisit) SetDataVersion(v int32)`

SetDataVersion sets DataVersion field to given value.


### GetTime

`func (o *ContentVisit) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ContentVisit) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ContentVisit) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetBundleId

`func (o *ContentVisit) GetBundleId() int32`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *ContentVisit) GetBundleIdOk() (*int32, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *ContentVisit) SetBundleId(v int32)`

SetBundleId sets BundleId field to given value.


### GetUserGuid

`func (o *ContentVisit) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *ContentVisit) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *ContentVisit) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


