/*
 * RStudio Connect API Reference
 *
 * The RStudio Connect API can be used to perform certain user actions remotely. You will need to install a tool or library that can make HTTP requests, such as:  - [httr](http://httr.r-lib.org/) (R HTTP library) - [cURL](https://curl.haxx.se/) (Linux tool for making HTTP calls) - [requests](https://requests.readthedocs.io/en/master/) (Python HTTP library)  ## <a name=\"versioning-policy\">Versioning of the API</a>  The RStudio Connect Server API uses a simple, single number versioning scheme as noted as the first part of each endpoint path.  This version number will only be incremented in the event that non-backward compatible changes are made to an existing endpoint. Note that this occurs on a per-endpoint basis; see the section on [deprecation](#deprecation) below for more information.  Changes that are considered backward compatible are:  * New fields in responses. * New non-required fields in requests. * New endpoint behavior which does not violate the current functional intent of the   endpoint.  Changes that are considered non-backward compatible are:  * Removal or rename of request or response fields. * A change of the type or format of one or more request or response fields. * Addition of new required request fields. * A substantial deviation from the current functional intent of the endpoint.  The points relating to functional intent are assumed to be extremely rare as more often such situations will result in a completely new endpoint, which makes the change a backward compatible addition.  ### Experimentation  RStudio Connect labels experimental endpoints in the API by including `/experimental` in the endpoint path immediately after the version indicator.  If an endpoint is noted as experimental, it should not be relied upon for any production work.  These are endpoints that RStudio Connect is making available to our customers to solicit feedback; they are subject to change without notice.  Such changes include anything from altered request/response shapes, to complete abandonment of the endpoint.  This public review of an experimental endpoint will last as long as necessary to either prove its viability or to determine that it’s not really needed.  The time for this will vary based on the intricacies of each endpoint.  When the endpoint is finalized, the next release of RStudio Connect will mark the experimental path as deprecated while adding the endpoint without the `/experimental` prefix. The path with the experimental prefix will be removed six months later.  The documentation for the endpoint will also note, during that time, the original, experimental, path.  All experimental endpoints are clearly marked as such in this documentation.  ### <a name=\"deprecation\">Deprecation and Removal of Old Versions</a>  It is possible that RStudio Connect may decide to deprecate an endpoint.  This will happen if either the endpoint serves no useful purpose because it’s functionality is now handled by a different endpoint or because there is a newer version of the endpoint that should be used.  If a deprecated endpoint is called, the response to it will include an extra HTTP header called, `X-Deprecated-Endpoint` and will have as a value the path of the endpoint that should be used instead.  If the functionality has no direct replacement, the value will be set to `n/a`.  Deprecated versions of an endpoint will be supported for 1 year from the release date of RStudio Connect in which the endpoint was marked as deprecated.  At that time, the endpoint is subject to removal at the discretion of RStudio Connect.  The life cycle of an endpoint will follow these steps.  1. The `/v1/endpoint` is public and in use by RStudio Connect customers. 1. RStudio Connect makes `/v2/experimental/endpoint` available for testing and feedback.    Customers should still use `/v1/endpoint` for production work. 1. RStudio Connect moves version 2 of the endpoint out of experimentation so, all within    the same release:     1. `/v1/endpoint` is marked as deprecated.     1. `/v2/experimental/endpoint` is marked as deprecated.     1. `/v2/endpoint` is made public. 1. Six months later, `/v2/experimental/endpoint` is removed from the product. 1. Twelve months later, `/v1/endpoint` is removed from the product.  Note that it is possible that RStudio Connect may produce a new version of an existing endpoint without making an experimental version of it first.  The same life cycle, without those parts, will still be followed.  ## <a name=\"authentication\">Authentication</a>  Most endpoints require you to identify yourself as a valid RStudio Connect user. You do this by specifying an API Key when you make a call to the server. The [API Keys](../user/api-keys/) chapter of the RStudio Connect User Guide explains how to create an API Key.  ### <a name=\"api-keys\">API Keys</a>  API Keys are managed by each user in the RStudio Connect dashboard. If you ever lose an API Key or otherwise feel it has been compromised, use the dashboard to revoke the key and create another one.  ::alert danger Keep your API Key safe.  If your RStudio Connect server's URL does not begin with `https`, your API Key could be intercepted and used by a malicious actor. ::  Once you have an API Key, you can authenticate by passing the key with a prefix of `\"Key \"` (the space is important) in the Authorization header.  Below are examples of invoking the \"Get R Information\" endpoint.  **cURL** ```bash curl -H \"Authorization: Key XXXXXXXXXXX\" \\      https://rstudioconnect.example.com/__api__/v1/server_settings/r ```  **R** ```r library(httr) apiKey <- \"XXXXXXXXXXX\" result <- GET(\"https://rstudioconnect.example.com/__api__/v1/server_settings/r\",   add_headers(Authorization = paste(\"Key\", apiKey))) ```  **Python** ```python import requests r = requests.get(   'https://rstudioconnect.example.com/__api__/v1/server_settings/r',   headers = { 'Authorization': 'Key XXXXXXXXXXX' } ) ``` 
 *
 * API version: 1.0.1
 * Contact: support@rstudio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Content The content object models all the \"things\" you may deploy to RStudio Connect. This includes Shiny applications, R Markdown documents, Plumber APIs, TensorFlow Model APIs, and plots.  Content has an `app_mode` field that represents the type of content represented by this item and its runtime model. This field initially has a value of `unknown` and receives its final value upon its first successful bundle deploy.  * `rmd-static` - An [R Markdown](https://rmarkdown.rstudio.com) document or site. * `shiny` - R code defining a [Shiny application](https://shiny.rstudio.com). * `rmd-shiny` - An [R Markdown](https://rmarkdown.rstudio.com) document with a Shiny runtime. * `static` - Content deployed without source; often HTML and plots. * `api` - R code defining a [Plumber API](https://www.rplumber.io). * `jupyter-static` - A Jupyter Notebook. * `tensorflow-saved-model` - A TensorFlow Model API. * `unknown` - No known runtime model.  Content marked as `shiny`, `rmd-shiny`, `api`, or `tensorflow-saved-model` is executed on demand as requests arrive.  Content marked as `rmd-static` or `jupyter-static` renders from source to output HTML. This rendering can occur based on some schedule or when explicitly triggered. It is *not* on each visit. Viewers of this type of content see a previously rendered result.  Content marked `static` is presented to viewers in its deployed form.  The `content_category` field refines the type of content specified by `app_mode`. It is empty by default. The `rsconnect` R package may assign a value when analyzing your content and building its manifest and bundle. Plots (images created in the RStudio IDE and presented in the \"Plots\" pane) have an `app_mode` of `static` and receive a `content_category` of `plot` to distinguish them from other HTML documents. Pinned static data sets have an `app_mode` of `static` and a `content_category` of `pin`. Multi-document R Markdown sites have an `app_mode` of `rmd-static` and a `content_category` of `site`.  Fields marked \"read-only\" cannot be written by [`POST /experimental/content`](#createContent) or [`POST /experimental/content/{guid}`](#updateContent).  The fields `bundle_id`, `app_mode`, `content_category`, `has_parameters`, `r_version`, and `py_version` are computed as a consequence of a [`POST /experimental/content/{guid}/deploy`](#deployContentBundle) deployment operation.  The `run_as` and `run_as_current_user` fields are read-only as fields of Content objects. A future API will allow adjustment of these properties. Use the RStudio Connect dashboard to adjust what Unix user executes your content. 
type Content struct {
	// Specifies the maximum number of client connections allowed to an individual process. Incoming connections which will exceed this limit are routed to a new process or rejected. When `null`, the default `Scheduler.MaxConnsPerProcess` is used. Applies only to content types that are executed on demand. 
	MaxConnsPerProcess NullableInt32 `json:"max_conns_per_process,omitempty"`
	// The UNIX user that executes this content. When `null`, the default `Applications.RunAs` is used. Applies only to executable content types - not `static`. 
	RunAs NullableString `json:"run_as,omitempty"`
	// Specifies the total number of concurrent processes allowed for a single interactive application. When `null`, the default `Scheduler.MaxProcesses` is used. Applies only to content types that are executed on demand. 
	MaxProcesses NullableInt32 `json:"max_processes,omitempty"`
	// Access type describes how this content manages its viewers. The value `all` is the most permissive; any visitor to RStudio Connect will be able to view this content. The value `logged_in` indicates that all RStudio Connect accounts may view the content. The `acl` value lets specifically enumerated users and groups view the content. Users configured as collaborators may always view content. 
	AccessType *string `json:"access_type,omitempty"`
	// Indicates if this content is allowed to execute as the logged-in user when using PAM authentication. Applies only to executable content types - not `static`. 
	RunAsCurrentUser *bool `json:"run_as_current_user,omitempty"`
	// The timestamp (RFC3339) indicating when this content was created. 
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// The maximum number of seconds allowed for an interactive application to start. RStudio Connect must be able to connect to a newly launched Shiny application, for example, before this threshold has elapsed. When `null`, the default `Scheduler.InitTimeout` is used. Applies only to content types that are executed on demand. 
	InitTimeout NullableInt32 `json:"init_timeout,omitempty"`
	// Specifies the minimum number of concurrent processes allowed for a single interactive application. When `null`, the default `Scheduler.MinProcesses` is used. Applies only to content types that are executed on demand. 
	MinProcesses NullableInt32 `json:"min_processes,omitempty"`
	// Maximum number of seconds allowed without data received from a client connection. A value of `0` means a lack of client (browser) interaction never causes the connection to close. When `null`, the default `Scheduler.ReadTimeout` is used. Applies only to content types that are executed on demand. 
	ReadTimeout NullableInt32 `json:"read_timeout,omitempty"`
	// The version of the Python interpreter associated with this content. The value `null` represents that a Python interpreter is not used by this content or that the Python package environment has not been successfully restored. Automatically assigned upon the successful deployment of a bundle. 
	PyVersion NullableString `json:"py_version,omitempty"`
	// The unique identifier of this content item. 
	Guid *string `json:"guid,omitempty"`
	// The title of this content. 
	Title NullableString `json:"title,omitempty"`
	// The runtime model for this content. Has a value of `unknown` before data is deployed to this item. Automatically assigned upon the first successful bundle deployment.  See the Content description for additional details. 
	AppMode *string `json:"app_mode,omitempty"`
	// The relationship of the accessing user to this content. A value of `owner` is returned for the content owner. `editor` indicates a collaborator. The `viewer` value is given to users who are permitted to view the content. A `none` role is returned for administrators who cannot view the content but are permitted to view its configuration. Computed at the time of the request. 
	Role *string `json:"role,omitempty"`
	// A rich description of this content. 
	Description *string `json:"description,omitempty"`
	// True when R Markdown rendered content allows parameter configuration. Automatically assigned upon the first successful bundle deployment. Applies only to content with an `app_mode` of `rmd-static`. 
	HasParameters *bool `json:"has_parameters,omitempty"`
	// The version of the R interpreter associated with this content. The value `null` represents that an R interpreter is not used by this content or that the R package environment has not been successfully restored. Automatically assigned upon the successful deployment of a bundle. 
	RVersion NullableString `json:"r_version,omitempty"`
	// Controls how aggressively new processes are spawned. When `null`, the default `Scheduler.LoadFactor` is used. Applies only to content types that are executed on demand. 
	LoadFactor NullableFloat32 `json:"load_factor,omitempty"`
	// The maximum number of seconds a worker process for an interactive application to remain alive after it goes idle (no active connections). When `null`, the default `Scheduler.IdleTimeout` is used. Applies only to content types that are executed on demand. 
	IdleTimeout NullableInt32 `json:"idle_timeout,omitempty"`
	// The identifier for the active deployment bundle. Automatically assigned upon the successful deployment of that bundle. 
	BundleId NullableString `json:"bundle_id,omitempty"`
	// Maximum number of seconds allowed without data sent or received across a client connection. A value of `0` means connections will never time-out (not recommended). When `null`, the default `Scheduler.ConnectionTimeout` is used. Applies only to content types that are executed on demand. 
	ConnectionTimeout NullableInt32 `json:"connection_timeout,omitempty"`
	// A simple, URL-friendly identifier. Allows alpha-numeric characters, hyphens (`\"-\"`), and underscores (`\"_\"`). 
	Name string `json:"name"`
	// The URL associated with this content. Computed from the associated vanity URL or the identifiers for this content. 
	Url *string `json:"url,omitempty"`
	// Describes the specialization of the content runtime model. Automatically assigned upon the first successful bundle deployment. 
	ContentCategory *string `json:"content_category,omitempty"`
	// The timestamp (RFC3339) indicating when this content last had a successful bundle deployment performed. 
	LastDeployedTime *time.Time `json:"last_deployed_time,omitempty"`
	// The unique identifier of the user who created this content item. Automatically assigned when the content is created. 
	OwnerGuid *string `json:"owner_guid,omitempty"`
}

// NewContent instantiates a new Content object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContent(name string) *Content {
	this := Content{}
	var accessType string = "acl"
	this.AccessType = &accessType
	var runAsCurrentUser bool = false
	this.RunAsCurrentUser = &runAsCurrentUser
	this.Name = name
	return &this
}

// NewContentWithDefaults instantiates a new Content object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentWithDefaults() *Content {
	this := Content{}
	var accessType string = "acl"
	this.AccessType = &accessType
	var runAsCurrentUser bool = false
	this.RunAsCurrentUser = &runAsCurrentUser
	return &this
}

// GetMaxConnsPerProcess returns the MaxConnsPerProcess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Content) GetMaxConnsPerProcess() int32 {
	if o == nil || o.MaxConnsPerProcess.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxConnsPerProcess.Get()
}

// GetMaxConnsPerProcessOk returns a tuple with the MaxConnsPerProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Content) GetMaxConnsPerProcessOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxConnsPerProcess.Get(), o.MaxConnsPerProcess.IsSet()
}

// HasMaxConnsPerProcess returns a boolean if a field has been set.
func (o *Content) HasMaxConnsPerProcess() bool {
	if o != nil && o.MaxConnsPerProcess.IsSet() {
		return true
	}

	return false
}

// SetMaxConnsPerProcess gets a reference to the given NullableInt32 and assigns it to the MaxConnsPerProcess field.
func (o *Content) SetMaxConnsPerProcess(v int32) {
	o.MaxConnsPerProcess.Set(&v)
}
// SetMaxConnsPerProcessNil sets the value for MaxConnsPerProcess to be an explicit nil
func (o *Content) SetMaxConnsPerProcessNil() {
	o.MaxConnsPerProcess.Set(nil)
}

// UnsetMaxConnsPerProcess ensures that no value is present for MaxConnsPerProcess, not even an explicit nil
func (o *Content) UnsetMaxConnsPerProcess() {
	o.MaxConnsPerProcess.Unset()
}

// GetRunAs returns the RunAs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Content) GetRunAs() string {
	if o == nil || o.RunAs.Get() == nil {
		var ret string
		return ret
	}
	return *o.RunAs.Get()
}

// GetRunAsOk returns a tuple with the RunAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Content) GetRunAsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RunAs.Get(), o.RunAs.IsSet()
}

// HasRunAs returns a boolean if a field has been set.
func (o *Content) HasRunAs() bool {
	if o != nil && o.RunAs.IsSet() {
		return true
	}

	return false
}

// SetRunAs gets a reference to the given NullableString and assigns it to the RunAs field.
func (o *Content) SetRunAs(v string) {
	o.RunAs.Set(&v)
}
// SetRunAsNil sets the value for RunAs to be an explicit nil
func (o *Content) SetRunAsNil() {
	o.RunAs.Set(nil)
}

// UnsetRunAs ensures that no value is present for RunAs, not even an explicit nil
func (o *Content) UnsetRunAs() {
	o.RunAs.Unset()
}

// GetMaxProcesses returns the MaxProcesses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Content) GetMaxProcesses() int32 {
	if o == nil || o.MaxProcesses.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxProcesses.Get()
}

// GetMaxProcessesOk returns a tuple with the MaxProcesses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Content) GetMaxProcessesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxProcesses.Get(), o.MaxProcesses.IsSet()
}

// HasMaxProcesses returns a boolean if a field has been set.
func (o *Content) HasMaxProcesses() bool {
	if o != nil && o.MaxProcesses.IsSet() {
		return true
	}

	return false
}

// SetMaxProcesses gets a reference to the given NullableInt32 and assigns it to the MaxProcesses field.
func (o *Content) SetMaxProcesses(v int32) {
	o.MaxProcesses.Set(&v)
}
// SetMaxProcessesNil sets the value for MaxProcesses to be an explicit nil
func (o *Content) SetMaxProcessesNil() {
	o.MaxProcesses.Set(nil)
}

// UnsetMaxProcesses ensures that no value is present for MaxProcesses, not even an explicit nil
func (o *Content) UnsetMaxProcesses() {
	o.MaxProcesses.Unset()
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *Content) GetAccessType() string {
	if o == nil || o.AccessType == nil {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetAccessTypeOk() (*string, bool) {
	if o == nil || o.AccessType == nil {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *Content) HasAccessType() bool {
	if o != nil && o.AccessType != nil {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *Content) SetAccessType(v string) {
	o.AccessType = &v
}

// GetRunAsCurrentUser returns the RunAsCurrentUser field value if set, zero value otherwise.
func (o *Content) GetRunAsCurrentUser() bool {
	if o == nil || o.RunAsCurrentUser == nil {
		var ret bool
		return ret
	}
	return *o.RunAsCurrentUser
}

// GetRunAsCurrentUserOk returns a tuple with the RunAsCurrentUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetRunAsCurrentUserOk() (*bool, bool) {
	if o == nil || o.RunAsCurrentUser == nil {
		return nil, false
	}
	return o.RunAsCurrentUser, true
}

// HasRunAsCurrentUser returns a boolean if a field has been set.
func (o *Content) HasRunAsCurrentUser() bool {
	if o != nil && o.RunAsCurrentUser != nil {
		return true
	}

	return false
}

// SetRunAsCurrentUser gets a reference to the given bool and assigns it to the RunAsCurrentUser field.
func (o *Content) SetRunAsCurrentUser(v bool) {
	o.RunAsCurrentUser = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Content) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Content) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *Content) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetInitTimeout returns the InitTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Content) GetInitTimeout() int32 {
	if o == nil || o.InitTimeout.Get() == nil {
		var ret int32
		return ret
	}
	return *o.InitTimeout.Get()
}

// GetInitTimeoutOk returns a tuple with the InitTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Content) GetInitTimeoutOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InitTimeout.Get(), o.InitTimeout.IsSet()
}

// HasInitTimeout returns a boolean if a field has been set.
func (o *Content) HasInitTimeout() bool {
	if o != nil && o.InitTimeout.IsSet() {
		return true
	}

	return false
}

// SetInitTimeout gets a reference to the given NullableInt32 and assigns it to the InitTimeout field.
func (o *Content) SetInitTimeout(v int32) {
	o.InitTimeout.Set(&v)
}
// SetInitTimeoutNil sets the value for InitTimeout to be an explicit nil
func (o *Content) SetInitTimeoutNil() {
	o.InitTimeout.Set(nil)
}

// UnsetInitTimeout ensures that no value is present for InitTimeout, not even an explicit nil
func (o *Content) UnsetInitTimeout() {
	o.InitTimeout.Unset()
}

// GetMinProcesses returns the MinProcesses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Content) GetMinProcesses() int32 {
	if o == nil || o.MinProcesses.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MinProcesses.Get()
}

// GetMinProcessesOk returns a tuple with the MinProcesses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Content) GetMinProcessesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MinProcesses.Get(), o.MinProcesses.IsSet()
}

// HasMinProcesses returns a boolean if a field has been set.
func (o *Content) HasMinProcesses() bool {
	if o != nil && o.MinProcesses.IsSet() {
		return true
	}

	return false
}

// SetMinProcesses gets a reference to the given NullableInt32 and assigns it to the MinProcesses field.
func (o *Content) SetMinProcesses(v int32) {
	o.MinProcesses.Set(&v)
}
// SetMinProcessesNil sets the value for MinProcesses to be an explicit nil
func (o *Content) SetMinProcessesNil() {
	o.MinProcesses.Set(nil)
}

// UnsetMinProcesses ensures that no value is present for MinProcesses, not even an explicit nil
func (o *Content) UnsetMinProcesses() {
	o.MinProcesses.Unset()
}

// GetReadTimeout returns the ReadTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Content) GetReadTimeout() int32 {
	if o == nil || o.ReadTimeout.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ReadTimeout.Get()
}

// GetReadTimeoutOk returns a tuple with the ReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Content) GetReadTimeoutOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReadTimeout.Get(), o.ReadTimeout.IsSet()
}

// HasReadTimeout returns a boolean if a field has been set.
func (o *Content) HasReadTimeout() bool {
	if o != nil && o.ReadTimeout.IsSet() {
		return true
	}

	return false
}

// SetReadTimeout gets a reference to the given NullableInt32 and assigns it to the ReadTimeout field.
func (o *Content) SetReadTimeout(v int32) {
	o.ReadTimeout.Set(&v)
}
// SetReadTimeoutNil sets the value for ReadTimeout to be an explicit nil
func (o *Content) SetReadTimeoutNil() {
	o.ReadTimeout.Set(nil)
}

// UnsetReadTimeout ensures that no value is present for ReadTimeout, not even an explicit nil
func (o *Content) UnsetReadTimeout() {
	o.ReadTimeout.Unset()
}

// GetPyVersion returns the PyVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Content) GetPyVersion() string {
	if o == nil || o.PyVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.PyVersion.Get()
}

// GetPyVersionOk returns a tuple with the PyVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Content) GetPyVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PyVersion.Get(), o.PyVersion.IsSet()
}

// HasPyVersion returns a boolean if a field has been set.
func (o *Content) HasPyVersion() bool {
	if o != nil && o.PyVersion.IsSet() {
		return true
	}

	return false
}

// SetPyVersion gets a reference to the given NullableString and assigns it to the PyVersion field.
func (o *Content) SetPyVersion(v string) {
	o.PyVersion.Set(&v)
}
// SetPyVersionNil sets the value for PyVersion to be an explicit nil
func (o *Content) SetPyVersionNil() {
	o.PyVersion.Set(nil)
}

// UnsetPyVersion ensures that no value is present for PyVersion, not even an explicit nil
func (o *Content) UnsetPyVersion() {
	o.PyVersion.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *Content) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *Content) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *Content) SetGuid(v string) {
	o.Guid = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Content) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Content) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Content) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Content) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Content) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Content) UnsetTitle() {
	o.Title.Unset()
}

// GetAppMode returns the AppMode field value if set, zero value otherwise.
func (o *Content) GetAppMode() string {
	if o == nil || o.AppMode == nil {
		var ret string
		return ret
	}
	return *o.AppMode
}

// GetAppModeOk returns a tuple with the AppMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetAppModeOk() (*string, bool) {
	if o == nil || o.AppMode == nil {
		return nil, false
	}
	return o.AppMode, true
}

// HasAppMode returns a boolean if a field has been set.
func (o *Content) HasAppMode() bool {
	if o != nil && o.AppMode != nil {
		return true
	}

	return false
}

// SetAppMode gets a reference to the given string and assigns it to the AppMode field.
func (o *Content) SetAppMode(v string) {
	o.AppMode = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *Content) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *Content) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *Content) SetRole(v string) {
	o.Role = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Content) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Content) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Content) SetDescription(v string) {
	o.Description = &v
}

// GetHasParameters returns the HasParameters field value if set, zero value otherwise.
func (o *Content) GetHasParameters() bool {
	if o == nil || o.HasParameters == nil {
		var ret bool
		return ret
	}
	return *o.HasParameters
}

// GetHasParametersOk returns a tuple with the HasParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetHasParametersOk() (*bool, bool) {
	if o == nil || o.HasParameters == nil {
		return nil, false
	}
	return o.HasParameters, true
}

// HasHasParameters returns a boolean if a field has been set.
func (o *Content) HasHasParameters() bool {
	if o != nil && o.HasParameters != nil {
		return true
	}

	return false
}

// SetHasParameters gets a reference to the given bool and assigns it to the HasParameters field.
func (o *Content) SetHasParameters(v bool) {
	o.HasParameters = &v
}

// GetRVersion returns the RVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Content) GetRVersion() string {
	if o == nil || o.RVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.RVersion.Get()
}

// GetRVersionOk returns a tuple with the RVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Content) GetRVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RVersion.Get(), o.RVersion.IsSet()
}

// HasRVersion returns a boolean if a field has been set.
func (o *Content) HasRVersion() bool {
	if o != nil && o.RVersion.IsSet() {
		return true
	}

	return false
}

// SetRVersion gets a reference to the given NullableString and assigns it to the RVersion field.
func (o *Content) SetRVersion(v string) {
	o.RVersion.Set(&v)
}
// SetRVersionNil sets the value for RVersion to be an explicit nil
func (o *Content) SetRVersionNil() {
	o.RVersion.Set(nil)
}

// UnsetRVersion ensures that no value is present for RVersion, not even an explicit nil
func (o *Content) UnsetRVersion() {
	o.RVersion.Unset()
}

// GetLoadFactor returns the LoadFactor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Content) GetLoadFactor() float32 {
	if o == nil || o.LoadFactor.Get() == nil {
		var ret float32
		return ret
	}
	return *o.LoadFactor.Get()
}

// GetLoadFactorOk returns a tuple with the LoadFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Content) GetLoadFactorOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoadFactor.Get(), o.LoadFactor.IsSet()
}

// HasLoadFactor returns a boolean if a field has been set.
func (o *Content) HasLoadFactor() bool {
	if o != nil && o.LoadFactor.IsSet() {
		return true
	}

	return false
}

// SetLoadFactor gets a reference to the given NullableFloat32 and assigns it to the LoadFactor field.
func (o *Content) SetLoadFactor(v float32) {
	o.LoadFactor.Set(&v)
}
// SetLoadFactorNil sets the value for LoadFactor to be an explicit nil
func (o *Content) SetLoadFactorNil() {
	o.LoadFactor.Set(nil)
}

// UnsetLoadFactor ensures that no value is present for LoadFactor, not even an explicit nil
func (o *Content) UnsetLoadFactor() {
	o.LoadFactor.Unset()
}

// GetIdleTimeout returns the IdleTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Content) GetIdleTimeout() int32 {
	if o == nil || o.IdleTimeout.Get() == nil {
		var ret int32
		return ret
	}
	return *o.IdleTimeout.Get()
}

// GetIdleTimeoutOk returns a tuple with the IdleTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Content) GetIdleTimeoutOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IdleTimeout.Get(), o.IdleTimeout.IsSet()
}

// HasIdleTimeout returns a boolean if a field has been set.
func (o *Content) HasIdleTimeout() bool {
	if o != nil && o.IdleTimeout.IsSet() {
		return true
	}

	return false
}

// SetIdleTimeout gets a reference to the given NullableInt32 and assigns it to the IdleTimeout field.
func (o *Content) SetIdleTimeout(v int32) {
	o.IdleTimeout.Set(&v)
}
// SetIdleTimeoutNil sets the value for IdleTimeout to be an explicit nil
func (o *Content) SetIdleTimeoutNil() {
	o.IdleTimeout.Set(nil)
}

// UnsetIdleTimeout ensures that no value is present for IdleTimeout, not even an explicit nil
func (o *Content) UnsetIdleTimeout() {
	o.IdleTimeout.Unset()
}

// GetBundleId returns the BundleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Content) GetBundleId() string {
	if o == nil || o.BundleId.Get() == nil {
		var ret string
		return ret
	}
	return *o.BundleId.Get()
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Content) GetBundleIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BundleId.Get(), o.BundleId.IsSet()
}

// HasBundleId returns a boolean if a field has been set.
func (o *Content) HasBundleId() bool {
	if o != nil && o.BundleId.IsSet() {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given NullableString and assigns it to the BundleId field.
func (o *Content) SetBundleId(v string) {
	o.BundleId.Set(&v)
}
// SetBundleIdNil sets the value for BundleId to be an explicit nil
func (o *Content) SetBundleIdNil() {
	o.BundleId.Set(nil)
}

// UnsetBundleId ensures that no value is present for BundleId, not even an explicit nil
func (o *Content) UnsetBundleId() {
	o.BundleId.Unset()
}

// GetConnectionTimeout returns the ConnectionTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Content) GetConnectionTimeout() int32 {
	if o == nil || o.ConnectionTimeout.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ConnectionTimeout.Get()
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Content) GetConnectionTimeoutOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConnectionTimeout.Get(), o.ConnectionTimeout.IsSet()
}

// HasConnectionTimeout returns a boolean if a field has been set.
func (o *Content) HasConnectionTimeout() bool {
	if o != nil && o.ConnectionTimeout.IsSet() {
		return true
	}

	return false
}

// SetConnectionTimeout gets a reference to the given NullableInt32 and assigns it to the ConnectionTimeout field.
func (o *Content) SetConnectionTimeout(v int32) {
	o.ConnectionTimeout.Set(&v)
}
// SetConnectionTimeoutNil sets the value for ConnectionTimeout to be an explicit nil
func (o *Content) SetConnectionTimeoutNil() {
	o.ConnectionTimeout.Set(nil)
}

// UnsetConnectionTimeout ensures that no value is present for ConnectionTimeout, not even an explicit nil
func (o *Content) UnsetConnectionTimeout() {
	o.ConnectionTimeout.Unset()
}

// GetName returns the Name field value
func (o *Content) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Content) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Content) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Content) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Content) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Content) SetUrl(v string) {
	o.Url = &v
}

// GetContentCategory returns the ContentCategory field value if set, zero value otherwise.
func (o *Content) GetContentCategory() string {
	if o == nil || o.ContentCategory == nil {
		var ret string
		return ret
	}
	return *o.ContentCategory
}

// GetContentCategoryOk returns a tuple with the ContentCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetContentCategoryOk() (*string, bool) {
	if o == nil || o.ContentCategory == nil {
		return nil, false
	}
	return o.ContentCategory, true
}

// HasContentCategory returns a boolean if a field has been set.
func (o *Content) HasContentCategory() bool {
	if o != nil && o.ContentCategory != nil {
		return true
	}

	return false
}

// SetContentCategory gets a reference to the given string and assigns it to the ContentCategory field.
func (o *Content) SetContentCategory(v string) {
	o.ContentCategory = &v
}

// GetLastDeployedTime returns the LastDeployedTime field value if set, zero value otherwise.
func (o *Content) GetLastDeployedTime() time.Time {
	if o == nil || o.LastDeployedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastDeployedTime
}

// GetLastDeployedTimeOk returns a tuple with the LastDeployedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetLastDeployedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastDeployedTime == nil {
		return nil, false
	}
	return o.LastDeployedTime, true
}

// HasLastDeployedTime returns a boolean if a field has been set.
func (o *Content) HasLastDeployedTime() bool {
	if o != nil && o.LastDeployedTime != nil {
		return true
	}

	return false
}

// SetLastDeployedTime gets a reference to the given time.Time and assigns it to the LastDeployedTime field.
func (o *Content) SetLastDeployedTime(v time.Time) {
	o.LastDeployedTime = &v
}

// GetOwnerGuid returns the OwnerGuid field value if set, zero value otherwise.
func (o *Content) GetOwnerGuid() string {
	if o == nil || o.OwnerGuid == nil {
		var ret string
		return ret
	}
	return *o.OwnerGuid
}

// GetOwnerGuidOk returns a tuple with the OwnerGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Content) GetOwnerGuidOk() (*string, bool) {
	if o == nil || o.OwnerGuid == nil {
		return nil, false
	}
	return o.OwnerGuid, true
}

// HasOwnerGuid returns a boolean if a field has been set.
func (o *Content) HasOwnerGuid() bool {
	if o != nil && o.OwnerGuid != nil {
		return true
	}

	return false
}

// SetOwnerGuid gets a reference to the given string and assigns it to the OwnerGuid field.
func (o *Content) SetOwnerGuid(v string) {
	o.OwnerGuid = &v
}

func (o Content) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxConnsPerProcess.IsSet() {
		toSerialize["max_conns_per_process"] = o.MaxConnsPerProcess.Get()
	}
	if o.RunAs.IsSet() {
		toSerialize["run_as"] = o.RunAs.Get()
	}
	if o.MaxProcesses.IsSet() {
		toSerialize["max_processes"] = o.MaxProcesses.Get()
	}
	if o.AccessType != nil {
		toSerialize["access_type"] = o.AccessType
	}
	if o.RunAsCurrentUser != nil {
		toSerialize["run_as_current_user"] = o.RunAsCurrentUser
	}
	if o.CreatedTime != nil {
		toSerialize["created_time"] = o.CreatedTime
	}
	if o.InitTimeout.IsSet() {
		toSerialize["init_timeout"] = o.InitTimeout.Get()
	}
	if o.MinProcesses.IsSet() {
		toSerialize["min_processes"] = o.MinProcesses.Get()
	}
	if o.ReadTimeout.IsSet() {
		toSerialize["read_timeout"] = o.ReadTimeout.Get()
	}
	if o.PyVersion.IsSet() {
		toSerialize["py_version"] = o.PyVersion.Get()
	}
	if o.Guid != nil {
		toSerialize["guid"] = o.Guid
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.AppMode != nil {
		toSerialize["app_mode"] = o.AppMode
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.HasParameters != nil {
		toSerialize["has_parameters"] = o.HasParameters
	}
	if o.RVersion.IsSet() {
		toSerialize["r_version"] = o.RVersion.Get()
	}
	if o.LoadFactor.IsSet() {
		toSerialize["load_factor"] = o.LoadFactor.Get()
	}
	if o.IdleTimeout.IsSet() {
		toSerialize["idle_timeout"] = o.IdleTimeout.Get()
	}
	if o.BundleId.IsSet() {
		toSerialize["bundle_id"] = o.BundleId.Get()
	}
	if o.ConnectionTimeout.IsSet() {
		toSerialize["connection_timeout"] = o.ConnectionTimeout.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.ContentCategory != nil {
		toSerialize["content_category"] = o.ContentCategory
	}
	if o.LastDeployedTime != nil {
		toSerialize["last_deployed_time"] = o.LastDeployedTime
	}
	if o.OwnerGuid != nil {
		toSerialize["owner_guid"] = o.OwnerGuid
	}
	return json.Marshal(toSerialize)
}

type NullableContent struct {
	value *Content
	isSet bool
}

func (v NullableContent) Get() *Content {
	return v.value
}

func (v *NullableContent) Set(val *Content) {
	v.value = val
	v.isSet = true
}

func (v NullableContent) IsSet() bool {
	return v.isSet
}

func (v *NullableContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContent(val *Content) *NullableContent {
	return &NullableContent{value: val, isSet: true}
}

func (v NullableContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


