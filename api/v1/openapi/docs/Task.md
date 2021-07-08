# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Numeric indication as to the cause of an error. Non-zero when an error has occured.  | [optional] [readonly] 
**Last** | Pointer to **int32** | The total number of output lines produced so far. Use as the value to &#x60;first&#x60; in the next request to only fetch output lines beyond what you have already received.  | [optional] [readonly] 
**Finished** | Pointer to **bool** | Indicates that a task has completed.  | [optional] [readonly] 
**Error** | Pointer to **string** | Description of the error. An empty string when no error has occurred.  | [optional] [readonly] 
**Output** | Pointer to **[]string** | An array containing lines of output produced by the task. The initial line of output is dictated by the &#x60;first&#x60; input parameter. The offset of the last output line is indicated by the &#x60;last&#x60; response field.  | [optional] [readonly] 
**Id** | Pointer to **string** | The identifier for this task.  | [optional] [readonly] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Task) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Task) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Task) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *Task) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLast

`func (o *Task) GetLast() int32`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *Task) GetLastOk() (*int32, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *Task) SetLast(v int32)`

SetLast sets Last field to given value.

### HasLast

`func (o *Task) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetFinished

`func (o *Task) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *Task) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *Task) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *Task) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetError

`func (o *Task) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Task) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Task) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Task) HasError() bool`

HasError returns a boolean if a field has been set.

### GetOutput

`func (o *Task) GetOutput() []string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *Task) GetOutputOk() (*[]string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *Task) SetOutput(v []string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *Task) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetId

`func (o *Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Task) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


