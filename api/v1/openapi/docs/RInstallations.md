# RInstallations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Installations** | Pointer to [**[]RInstallation**](RInstallation.md) | The array of information about RStudio Connect&#39;s R installations.  | [optional] 

## Methods

### NewRInstallations

`func NewRInstallations() *RInstallations`

NewRInstallations instantiates a new RInstallations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRInstallationsWithDefaults

`func NewRInstallationsWithDefaults() *RInstallations`

NewRInstallationsWithDefaults instantiates a new RInstallations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallations

`func (o *RInstallations) GetInstallations() []RInstallation`

GetInstallations returns the Installations field if non-nil, zero value otherwise.

### GetInstallationsOk

`func (o *RInstallations) GetInstallationsOk() (*[]RInstallation, bool)`

GetInstallationsOk returns a tuple with the Installations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallations

`func (o *RInstallations) SetInstallations(v []RInstallation)`

SetInstallations sets Installations field to given value.

### HasInstallations

`func (o *RInstallations) HasInstallations() bool`

HasInstallations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


