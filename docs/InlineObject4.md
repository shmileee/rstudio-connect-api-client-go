# InlineObject4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TempTicket** | **string** | The temporary ticket used for creating a group on the RStudio Connect server. It is obtained from the [&#x60;GET /groups/remote&#x60;](#searchRemoteGroups) endpoint.  | 

## Methods

### NewInlineObject4

`func NewInlineObject4(tempTicket string, ) *InlineObject4`

NewInlineObject4 instantiates a new InlineObject4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject4WithDefaults

`func NewInlineObject4WithDefaults() *InlineObject4`

NewInlineObject4WithDefaults instantiates a new InlineObject4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTempTicket

`func (o *InlineObject4) GetTempTicket() string`

GetTempTicket returns the TempTicket field if non-nil, zero value otherwise.

### GetTempTicketOk

`func (o *InlineObject4) GetTempTicketOk() (*string, bool)`

GetTempTicketOk returns a tuple with the TempTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempTicket

`func (o *InlineObject4) SetTempTicket(v string)`

SetTempTicket sets TempTicket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


