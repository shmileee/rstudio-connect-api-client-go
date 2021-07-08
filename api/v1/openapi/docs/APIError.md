# APIError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | The specific code for the type of error returned.  | 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | **string** | Some text which describes the problem that was encountered.  | 

## Methods

### NewAPIError

`func NewAPIError(code int32, error_ string, ) *APIError`

NewAPIError instantiates a new APIError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIErrorWithDefaults

`func NewAPIErrorWithDefaults() *APIError`

NewAPIErrorWithDefaults instantiates a new APIError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *APIError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *APIError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *APIError) SetCode(v int32)`

SetCode sets Code field to given value.


### GetPayload

`func (o *APIError) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *APIError) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *APIError) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *APIError) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *APIError) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *APIError) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetError

`func (o *APIError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *APIError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *APIError) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


