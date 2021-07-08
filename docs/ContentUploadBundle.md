# ContentUploadBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | Pointer to **string** | Identifier for the newly created bundle.  | [optional] [readonly] 
**Size** | Pointer to **float32** | The number of bytes received.  | [optional] [readonly] 

## Methods

### NewContentUploadBundle

`func NewContentUploadBundle() *ContentUploadBundle`

NewContentUploadBundle instantiates a new ContentUploadBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentUploadBundleWithDefaults

`func NewContentUploadBundleWithDefaults() *ContentUploadBundle`

NewContentUploadBundleWithDefaults instantiates a new ContentUploadBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *ContentUploadBundle) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *ContentUploadBundle) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *ContentUploadBundle) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *ContentUploadBundle) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetSize

`func (o *ContentUploadBundle) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentUploadBundle) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentUploadBundle) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ContentUploadBundle) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


