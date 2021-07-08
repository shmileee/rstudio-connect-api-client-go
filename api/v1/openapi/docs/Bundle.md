# Bundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentGuid** | Pointer to **string** | The identifier of the owning content.  | [optional] [readonly] 
**RVersion** | Pointer to **NullableString** | The version of the R interpreter used when last restoring this bundle. The value &#x60;null&#x60; represents that an R interpreter is not used by this bundle or that the R package environment has not been successfully restored.  R version is not disclosed to users with a \&quot;viewer\&quot; role and returned with the value &#x60;null&#x60;.  | [optional] [readonly] 
**Active** | Pointer to **bool** | Indicates if this bundle is active for the owning content.  | [optional] [readonly] 
**CreatedTime** | Pointer to **time.Time** | The timestamp (RFC3339) of when this bundle was created.  | [optional] [readonly] 
**PyVersion** | Pointer to **NullableString** | The version of the Python interpreter used when last restoring this bundle.  The value &#x60;null&#x60; represents that a Python interpreter is not used by this bundle or that the Python package environment has not been successfully restored.  Python version is not disclosed to users with a \&quot;viewer\&quot; role and returned with the value &#x60;null&#x60;.  | [optional] [readonly] 
**Id** | Pointer to **string** | The identifier for this bundle.  | [optional] [readonly] 
**Size** | Pointer to **float32** | On-disk size in bytes of the tar.gz file associated with this bundle. Zero when there is no on-disk file.  | [optional] [readonly] 

## Methods

### NewBundle

`func NewBundle() *Bundle`

NewBundle instantiates a new Bundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleWithDefaults

`func NewBundleWithDefaults() *Bundle`

NewBundleWithDefaults instantiates a new Bundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentGuid

`func (o *Bundle) GetContentGuid() string`

GetContentGuid returns the ContentGuid field if non-nil, zero value otherwise.

### GetContentGuidOk

`func (o *Bundle) GetContentGuidOk() (*string, bool)`

GetContentGuidOk returns a tuple with the ContentGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuid

`func (o *Bundle) SetContentGuid(v string)`

SetContentGuid sets ContentGuid field to given value.

### HasContentGuid

`func (o *Bundle) HasContentGuid() bool`

HasContentGuid returns a boolean if a field has been set.

### GetRVersion

`func (o *Bundle) GetRVersion() string`

GetRVersion returns the RVersion field if non-nil, zero value otherwise.

### GetRVersionOk

`func (o *Bundle) GetRVersionOk() (*string, bool)`

GetRVersionOk returns a tuple with the RVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRVersion

`func (o *Bundle) SetRVersion(v string)`

SetRVersion sets RVersion field to given value.

### HasRVersion

`func (o *Bundle) HasRVersion() bool`

HasRVersion returns a boolean if a field has been set.

### SetRVersionNil

`func (o *Bundle) SetRVersionNil(b bool)`

 SetRVersionNil sets the value for RVersion to be an explicit nil

### UnsetRVersion
`func (o *Bundle) UnsetRVersion()`

UnsetRVersion ensures that no value is present for RVersion, not even an explicit nil
### GetActive

`func (o *Bundle) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Bundle) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Bundle) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Bundle) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Bundle) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Bundle) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Bundle) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Bundle) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetPyVersion

`func (o *Bundle) GetPyVersion() string`

GetPyVersion returns the PyVersion field if non-nil, zero value otherwise.

### GetPyVersionOk

`func (o *Bundle) GetPyVersionOk() (*string, bool)`

GetPyVersionOk returns a tuple with the PyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPyVersion

`func (o *Bundle) SetPyVersion(v string)`

SetPyVersion sets PyVersion field to given value.

### HasPyVersion

`func (o *Bundle) HasPyVersion() bool`

HasPyVersion returns a boolean if a field has been set.

### SetPyVersionNil

`func (o *Bundle) SetPyVersionNil(b bool)`

 SetPyVersionNil sets the value for PyVersion to be an explicit nil

### UnsetPyVersion
`func (o *Bundle) UnsetPyVersion()`

UnsetPyVersion ensures that no value is present for PyVersion, not even an explicit nil
### GetId

`func (o *Bundle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bundle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bundle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Bundle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSize

`func (o *Bundle) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Bundle) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Bundle) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Bundle) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


