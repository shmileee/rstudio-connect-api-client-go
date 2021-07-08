# Bundles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | The total number of bundles that belong to the given content.  | [optional] 
**Results** | Pointer to [**[]Bundle**](Bundle.md) | The bundles list  | [optional] 
**CurrentPage** | Pointer to **int32** | The current page of results.  | [optional] 

## Methods

### NewBundles

`func NewBundles() *Bundles`

NewBundles instantiates a new Bundles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundlesWithDefaults

`func NewBundlesWithDefaults() *Bundles`

NewBundlesWithDefaults instantiates a new Bundles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *Bundles) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Bundles) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Bundles) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Bundles) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetResults

`func (o *Bundles) GetResults() []Bundle`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Bundles) GetResultsOk() (*[]Bundle, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Bundles) SetResults(v []Bundle)`

SetResults sets Results field to given value.

### HasResults

`func (o *Bundles) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetCurrentPage

`func (o *Bundles) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *Bundles) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *Bundles) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *Bundles) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


