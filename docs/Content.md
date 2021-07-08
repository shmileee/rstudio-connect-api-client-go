# Content

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxConnsPerProcess** | Pointer to **NullableInt32** | Specifies the maximum number of client connections allowed to an individual process. Incoming connections which will exceed this limit are routed to a new process or rejected. When &#x60;null&#x60;, the default &#x60;Scheduler.MaxConnsPerProcess&#x60; is used. Applies only to content types that are executed on demand.  | [optional] 
**RunAs** | Pointer to **NullableString** | The UNIX user that executes this content. When &#x60;null&#x60;, the default &#x60;Applications.RunAs&#x60; is used. Applies only to executable content types - not &#x60;static&#x60;.  | [optional] [readonly] 
**MaxProcesses** | Pointer to **NullableInt32** | Specifies the total number of concurrent processes allowed for a single interactive application. When &#x60;null&#x60;, the default &#x60;Scheduler.MaxProcesses&#x60; is used. Applies only to content types that are executed on demand.  | [optional] 
**AccessType** | Pointer to **string** | Access type describes how this content manages its viewers. The value &#x60;all&#x60; is the most permissive; any visitor to RStudio Connect will be able to view this content. The value &#x60;logged_in&#x60; indicates that all RStudio Connect accounts may view the content. The &#x60;acl&#x60; value lets specifically enumerated users and groups view the content. Users configured as collaborators may always view content.  | [optional] [default to "acl"]
**RunAsCurrentUser** | Pointer to **bool** | Indicates if this content is allowed to execute as the logged-in user when using PAM authentication. Applies only to executable content types - not &#x60;static&#x60;.  | [optional] [readonly] [default to false]
**CreatedTime** | Pointer to **time.Time** | The timestamp (RFC3339) indicating when this content was created.  | [optional] [readonly] 
**InitTimeout** | Pointer to **NullableInt32** | The maximum number of seconds allowed for an interactive application to start. RStudio Connect must be able to connect to a newly launched Shiny application, for example, before this threshold has elapsed. When &#x60;null&#x60;, the default &#x60;Scheduler.InitTimeout&#x60; is used. Applies only to content types that are executed on demand.  | [optional] 
**MinProcesses** | Pointer to **NullableInt32** | Specifies the minimum number of concurrent processes allowed for a single interactive application. When &#x60;null&#x60;, the default &#x60;Scheduler.MinProcesses&#x60; is used. Applies only to content types that are executed on demand.  | [optional] 
**ReadTimeout** | Pointer to **NullableInt32** | Maximum number of seconds allowed without data received from a client connection. A value of &#x60;0&#x60; means a lack of client (browser) interaction never causes the connection to close. When &#x60;null&#x60;, the default &#x60;Scheduler.ReadTimeout&#x60; is used. Applies only to content types that are executed on demand.  | [optional] 
**PyVersion** | Pointer to **NullableString** | The version of the Python interpreter associated with this content. The value &#x60;null&#x60; represents that a Python interpreter is not used by this content or that the Python package environment has not been successfully restored. Automatically assigned upon the successful deployment of a bundle.  | [optional] [readonly] 
**Guid** | Pointer to **string** | The unique identifier of this content item.  | [optional] [readonly] 
**Title** | Pointer to **NullableString** | The title of this content.  | [optional] 
**AppMode** | Pointer to **string** | The runtime model for this content. Has a value of &#x60;unknown&#x60; before data is deployed to this item. Automatically assigned upon the first successful bundle deployment.  See the Content description for additional details.  | [optional] [readonly] 
**Role** | Pointer to **string** | The relationship of the accessing user to this content. A value of &#x60;owner&#x60; is returned for the content owner. &#x60;editor&#x60; indicates a collaborator. The &#x60;viewer&#x60; value is given to users who are permitted to view the content. A &#x60;none&#x60; role is returned for administrators who cannot view the content but are permitted to view its configuration. Computed at the time of the request.  | [optional] [readonly] 
**Description** | Pointer to **string** | A rich description of this content.  | [optional] 
**HasParameters** | Pointer to **bool** | True when R Markdown rendered content allows parameter configuration. Automatically assigned upon the first successful bundle deployment. Applies only to content with an &#x60;app_mode&#x60; of &#x60;rmd-static&#x60;.  | [optional] [readonly] 
**RVersion** | Pointer to **NullableString** | The version of the R interpreter associated with this content. The value &#x60;null&#x60; represents that an R interpreter is not used by this content or that the R package environment has not been successfully restored. Automatically assigned upon the successful deployment of a bundle.  | [optional] [readonly] 
**LoadFactor** | Pointer to **NullableFloat32** | Controls how aggressively new processes are spawned. When &#x60;null&#x60;, the default &#x60;Scheduler.LoadFactor&#x60; is used. Applies only to content types that are executed on demand.  | [optional] 
**IdleTimeout** | Pointer to **NullableInt32** | The maximum number of seconds a worker process for an interactive application to remain alive after it goes idle (no active connections). When &#x60;null&#x60;, the default &#x60;Scheduler.IdleTimeout&#x60; is used. Applies only to content types that are executed on demand.  | [optional] 
**BundleId** | Pointer to **NullableString** | The identifier for the active deployment bundle. Automatically assigned upon the successful deployment of that bundle.  | [optional] [readonly] 
**ConnectionTimeout** | Pointer to **NullableInt32** | Maximum number of seconds allowed without data sent or received across a client connection. A value of &#x60;0&#x60; means connections will never time-out (not recommended). When &#x60;null&#x60;, the default &#x60;Scheduler.ConnectionTimeout&#x60; is used. Applies only to content types that are executed on demand.  | [optional] 
**Name** | **string** | A simple, URL-friendly identifier. Allows alpha-numeric characters, hyphens (&#x60;\&quot;-\&quot;&#x60;), and underscores (&#x60;\&quot;_\&quot;&#x60;).  | 
**Url** | Pointer to **string** | The URL associated with this content. Computed from the associated vanity URL or the identifiers for this content.  | [optional] [readonly] 
**ContentCategory** | Pointer to **string** | Describes the specialization of the content runtime model. Automatically assigned upon the first successful bundle deployment.  | [optional] [readonly] 
**LastDeployedTime** | Pointer to **time.Time** | The timestamp (RFC3339) indicating when this content last had a successful bundle deployment performed.  | [optional] [readonly] 
**OwnerGuid** | Pointer to **string** | The unique identifier of the user who created this content item. Automatically assigned when the content is created.  | [optional] [readonly] 

## Methods

### NewContent

`func NewContent(name string, ) *Content`

NewContent instantiates a new Content object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentWithDefaults

`func NewContentWithDefaults() *Content`

NewContentWithDefaults instantiates a new Content object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxConnsPerProcess

`func (o *Content) GetMaxConnsPerProcess() int32`

GetMaxConnsPerProcess returns the MaxConnsPerProcess field if non-nil, zero value otherwise.

### GetMaxConnsPerProcessOk

`func (o *Content) GetMaxConnsPerProcessOk() (*int32, bool)`

GetMaxConnsPerProcessOk returns a tuple with the MaxConnsPerProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnsPerProcess

`func (o *Content) SetMaxConnsPerProcess(v int32)`

SetMaxConnsPerProcess sets MaxConnsPerProcess field to given value.

### HasMaxConnsPerProcess

`func (o *Content) HasMaxConnsPerProcess() bool`

HasMaxConnsPerProcess returns a boolean if a field has been set.

### SetMaxConnsPerProcessNil

`func (o *Content) SetMaxConnsPerProcessNil(b bool)`

 SetMaxConnsPerProcessNil sets the value for MaxConnsPerProcess to be an explicit nil

### UnsetMaxConnsPerProcess
`func (o *Content) UnsetMaxConnsPerProcess()`

UnsetMaxConnsPerProcess ensures that no value is present for MaxConnsPerProcess, not even an explicit nil
### GetRunAs

`func (o *Content) GetRunAs() string`

GetRunAs returns the RunAs field if non-nil, zero value otherwise.

### GetRunAsOk

`func (o *Content) GetRunAsOk() (*string, bool)`

GetRunAsOk returns a tuple with the RunAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAs

`func (o *Content) SetRunAs(v string)`

SetRunAs sets RunAs field to given value.

### HasRunAs

`func (o *Content) HasRunAs() bool`

HasRunAs returns a boolean if a field has been set.

### SetRunAsNil

`func (o *Content) SetRunAsNil(b bool)`

 SetRunAsNil sets the value for RunAs to be an explicit nil

### UnsetRunAs
`func (o *Content) UnsetRunAs()`

UnsetRunAs ensures that no value is present for RunAs, not even an explicit nil
### GetMaxProcesses

`func (o *Content) GetMaxProcesses() int32`

GetMaxProcesses returns the MaxProcesses field if non-nil, zero value otherwise.

### GetMaxProcessesOk

`func (o *Content) GetMaxProcessesOk() (*int32, bool)`

GetMaxProcessesOk returns a tuple with the MaxProcesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProcesses

`func (o *Content) SetMaxProcesses(v int32)`

SetMaxProcesses sets MaxProcesses field to given value.

### HasMaxProcesses

`func (o *Content) HasMaxProcesses() bool`

HasMaxProcesses returns a boolean if a field has been set.

### SetMaxProcessesNil

`func (o *Content) SetMaxProcessesNil(b bool)`

 SetMaxProcessesNil sets the value for MaxProcesses to be an explicit nil

### UnsetMaxProcesses
`func (o *Content) UnsetMaxProcesses()`

UnsetMaxProcesses ensures that no value is present for MaxProcesses, not even an explicit nil
### GetAccessType

`func (o *Content) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *Content) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *Content) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *Content) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetRunAsCurrentUser

`func (o *Content) GetRunAsCurrentUser() bool`

GetRunAsCurrentUser returns the RunAsCurrentUser field if non-nil, zero value otherwise.

### GetRunAsCurrentUserOk

`func (o *Content) GetRunAsCurrentUserOk() (*bool, bool)`

GetRunAsCurrentUserOk returns a tuple with the RunAsCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsCurrentUser

`func (o *Content) SetRunAsCurrentUser(v bool)`

SetRunAsCurrentUser sets RunAsCurrentUser field to given value.

### HasRunAsCurrentUser

`func (o *Content) HasRunAsCurrentUser() bool`

HasRunAsCurrentUser returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Content) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Content) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Content) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Content) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetInitTimeout

`func (o *Content) GetInitTimeout() int32`

GetInitTimeout returns the InitTimeout field if non-nil, zero value otherwise.

### GetInitTimeoutOk

`func (o *Content) GetInitTimeoutOk() (*int32, bool)`

GetInitTimeoutOk returns a tuple with the InitTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitTimeout

`func (o *Content) SetInitTimeout(v int32)`

SetInitTimeout sets InitTimeout field to given value.

### HasInitTimeout

`func (o *Content) HasInitTimeout() bool`

HasInitTimeout returns a boolean if a field has been set.

### SetInitTimeoutNil

`func (o *Content) SetInitTimeoutNil(b bool)`

 SetInitTimeoutNil sets the value for InitTimeout to be an explicit nil

### UnsetInitTimeout
`func (o *Content) UnsetInitTimeout()`

UnsetInitTimeout ensures that no value is present for InitTimeout, not even an explicit nil
### GetMinProcesses

`func (o *Content) GetMinProcesses() int32`

GetMinProcesses returns the MinProcesses field if non-nil, zero value otherwise.

### GetMinProcessesOk

`func (o *Content) GetMinProcessesOk() (*int32, bool)`

GetMinProcessesOk returns a tuple with the MinProcesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinProcesses

`func (o *Content) SetMinProcesses(v int32)`

SetMinProcesses sets MinProcesses field to given value.

### HasMinProcesses

`func (o *Content) HasMinProcesses() bool`

HasMinProcesses returns a boolean if a field has been set.

### SetMinProcessesNil

`func (o *Content) SetMinProcessesNil(b bool)`

 SetMinProcessesNil sets the value for MinProcesses to be an explicit nil

### UnsetMinProcesses
`func (o *Content) UnsetMinProcesses()`

UnsetMinProcesses ensures that no value is present for MinProcesses, not even an explicit nil
### GetReadTimeout

`func (o *Content) GetReadTimeout() int32`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *Content) GetReadTimeoutOk() (*int32, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *Content) SetReadTimeout(v int32)`

SetReadTimeout sets ReadTimeout field to given value.

### HasReadTimeout

`func (o *Content) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.

### SetReadTimeoutNil

`func (o *Content) SetReadTimeoutNil(b bool)`

 SetReadTimeoutNil sets the value for ReadTimeout to be an explicit nil

### UnsetReadTimeout
`func (o *Content) UnsetReadTimeout()`

UnsetReadTimeout ensures that no value is present for ReadTimeout, not even an explicit nil
### GetPyVersion

`func (o *Content) GetPyVersion() string`

GetPyVersion returns the PyVersion field if non-nil, zero value otherwise.

### GetPyVersionOk

`func (o *Content) GetPyVersionOk() (*string, bool)`

GetPyVersionOk returns a tuple with the PyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPyVersion

`func (o *Content) SetPyVersion(v string)`

SetPyVersion sets PyVersion field to given value.

### HasPyVersion

`func (o *Content) HasPyVersion() bool`

HasPyVersion returns a boolean if a field has been set.

### SetPyVersionNil

`func (o *Content) SetPyVersionNil(b bool)`

 SetPyVersionNil sets the value for PyVersion to be an explicit nil

### UnsetPyVersion
`func (o *Content) UnsetPyVersion()`

UnsetPyVersion ensures that no value is present for PyVersion, not even an explicit nil
### GetGuid

`func (o *Content) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Content) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Content) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Content) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetTitle

`func (o *Content) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Content) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Content) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Content) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Content) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Content) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAppMode

`func (o *Content) GetAppMode() string`

GetAppMode returns the AppMode field if non-nil, zero value otherwise.

### GetAppModeOk

`func (o *Content) GetAppModeOk() (*string, bool)`

GetAppModeOk returns a tuple with the AppMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppMode

`func (o *Content) SetAppMode(v string)`

SetAppMode sets AppMode field to given value.

### HasAppMode

`func (o *Content) HasAppMode() bool`

HasAppMode returns a boolean if a field has been set.

### GetRole

`func (o *Content) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Content) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Content) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Content) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetDescription

`func (o *Content) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Content) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Content) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Content) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHasParameters

`func (o *Content) GetHasParameters() bool`

GetHasParameters returns the HasParameters field if non-nil, zero value otherwise.

### GetHasParametersOk

`func (o *Content) GetHasParametersOk() (*bool, bool)`

GetHasParametersOk returns a tuple with the HasParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasParameters

`func (o *Content) SetHasParameters(v bool)`

SetHasParameters sets HasParameters field to given value.

### HasHasParameters

`func (o *Content) HasHasParameters() bool`

HasHasParameters returns a boolean if a field has been set.

### GetRVersion

`func (o *Content) GetRVersion() string`

GetRVersion returns the RVersion field if non-nil, zero value otherwise.

### GetRVersionOk

`func (o *Content) GetRVersionOk() (*string, bool)`

GetRVersionOk returns a tuple with the RVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRVersion

`func (o *Content) SetRVersion(v string)`

SetRVersion sets RVersion field to given value.

### HasRVersion

`func (o *Content) HasRVersion() bool`

HasRVersion returns a boolean if a field has been set.

### SetRVersionNil

`func (o *Content) SetRVersionNil(b bool)`

 SetRVersionNil sets the value for RVersion to be an explicit nil

### UnsetRVersion
`func (o *Content) UnsetRVersion()`

UnsetRVersion ensures that no value is present for RVersion, not even an explicit nil
### GetLoadFactor

`func (o *Content) GetLoadFactor() float32`

GetLoadFactor returns the LoadFactor field if non-nil, zero value otherwise.

### GetLoadFactorOk

`func (o *Content) GetLoadFactorOk() (*float32, bool)`

GetLoadFactorOk returns a tuple with the LoadFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadFactor

`func (o *Content) SetLoadFactor(v float32)`

SetLoadFactor sets LoadFactor field to given value.

### HasLoadFactor

`func (o *Content) HasLoadFactor() bool`

HasLoadFactor returns a boolean if a field has been set.

### SetLoadFactorNil

`func (o *Content) SetLoadFactorNil(b bool)`

 SetLoadFactorNil sets the value for LoadFactor to be an explicit nil

### UnsetLoadFactor
`func (o *Content) UnsetLoadFactor()`

UnsetLoadFactor ensures that no value is present for LoadFactor, not even an explicit nil
### GetIdleTimeout

`func (o *Content) GetIdleTimeout() int32`

GetIdleTimeout returns the IdleTimeout field if non-nil, zero value otherwise.

### GetIdleTimeoutOk

`func (o *Content) GetIdleTimeoutOk() (*int32, bool)`

GetIdleTimeoutOk returns a tuple with the IdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeout

`func (o *Content) SetIdleTimeout(v int32)`

SetIdleTimeout sets IdleTimeout field to given value.

### HasIdleTimeout

`func (o *Content) HasIdleTimeout() bool`

HasIdleTimeout returns a boolean if a field has been set.

### SetIdleTimeoutNil

`func (o *Content) SetIdleTimeoutNil(b bool)`

 SetIdleTimeoutNil sets the value for IdleTimeout to be an explicit nil

### UnsetIdleTimeout
`func (o *Content) UnsetIdleTimeout()`

UnsetIdleTimeout ensures that no value is present for IdleTimeout, not even an explicit nil
### GetBundleId

`func (o *Content) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *Content) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *Content) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *Content) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### SetBundleIdNil

`func (o *Content) SetBundleIdNil(b bool)`

 SetBundleIdNil sets the value for BundleId to be an explicit nil

### UnsetBundleId
`func (o *Content) UnsetBundleId()`

UnsetBundleId ensures that no value is present for BundleId, not even an explicit nil
### GetConnectionTimeout

`func (o *Content) GetConnectionTimeout() int32`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *Content) GetConnectionTimeoutOk() (*int32, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *Content) SetConnectionTimeout(v int32)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *Content) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### SetConnectionTimeoutNil

`func (o *Content) SetConnectionTimeoutNil(b bool)`

 SetConnectionTimeoutNil sets the value for ConnectionTimeout to be an explicit nil

### UnsetConnectionTimeout
`func (o *Content) UnsetConnectionTimeout()`

UnsetConnectionTimeout ensures that no value is present for ConnectionTimeout, not even an explicit nil
### GetName

`func (o *Content) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Content) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Content) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *Content) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Content) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Content) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Content) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetContentCategory

`func (o *Content) GetContentCategory() string`

GetContentCategory returns the ContentCategory field if non-nil, zero value otherwise.

### GetContentCategoryOk

`func (o *Content) GetContentCategoryOk() (*string, bool)`

GetContentCategoryOk returns a tuple with the ContentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCategory

`func (o *Content) SetContentCategory(v string)`

SetContentCategory sets ContentCategory field to given value.

### HasContentCategory

`func (o *Content) HasContentCategory() bool`

HasContentCategory returns a boolean if a field has been set.

### GetLastDeployedTime

`func (o *Content) GetLastDeployedTime() time.Time`

GetLastDeployedTime returns the LastDeployedTime field if non-nil, zero value otherwise.

### GetLastDeployedTimeOk

`func (o *Content) GetLastDeployedTimeOk() (*time.Time, bool)`

GetLastDeployedTimeOk returns a tuple with the LastDeployedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeployedTime

`func (o *Content) SetLastDeployedTime(v time.Time)`

SetLastDeployedTime sets LastDeployedTime field to given value.

### HasLastDeployedTime

`func (o *Content) HasLastDeployedTime() bool`

HasLastDeployedTime returns a boolean if a field has been set.

### GetOwnerGuid

`func (o *Content) GetOwnerGuid() string`

GetOwnerGuid returns the OwnerGuid field if non-nil, zero value otherwise.

### GetOwnerGuidOk

`func (o *Content) GetOwnerGuidOk() (*string, bool)`

GetOwnerGuidOk returns a tuple with the OwnerGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGuid

`func (o *Content) SetOwnerGuid(v string)`

SetOwnerGuid sets OwnerGuid field to given value.

### HasOwnerGuid

`func (o *Content) HasOwnerGuid() bool`

HasOwnerGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


