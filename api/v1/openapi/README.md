# Go API client for openapi

The RStudio Connect API can be used to perform certain
user actions remotely. You will need to install a tool or library
that can make HTTP requests, such as:

- [httr](http://httr.r-lib.org/) (R HTTP library)
- [cURL](https://curl.haxx.se/) (Linux tool for making HTTP calls)
- [requests](https://requests.readthedocs.io/en/master/) (Python HTTP library)

## Versioning of the API

The RStudio Connect Server API uses a simple, single number versioning scheme as noted
as the first part of each endpoint path.  This version number will only be incremented
in the event that non-backward compatible changes are made to an existing endpoint.
Note that this occurs on a per-endpoint basis; see the section on
[deprecation](#deprecation) below for more information.

Changes that are considered backward compatible are:

* New fields in responses.
* New non-required fields in requests.
* New endpoint behavior which does not violate the current functional intent of the
  endpoint.

Changes that are considered non-backward compatible are:

* Removal or rename of request or response fields.
* A change of the type or format of one or more request or response fields.
* Addition of new required request fields.
* A substantial deviation from the current functional intent of the endpoint.

The points relating to functional intent are assumed to be extremely rare as more
often such situations will result in a completely new endpoint, which makes the
change a backward compatible addition.

### Experimentation

RStudio Connect labels experimental endpoints in the API by including `/experimental`
in the endpoint path immediately after the version indicator.  If an endpoint is noted
as experimental, it should not be relied upon for any production work.  These are
endpoints that RStudio Connect is making available to our customers to solicit
feedback; they are subject to change without notice.  Such changes include anything
from altered request/response shapes, to complete abandonment of the endpoint.

This public review of an experimental endpoint will last as long as necessary to either
prove its viability or to determine that it’s not really needed.  The time for this
will vary based on the intricacies of each endpoint.  When the endpoint is finalized,
the next release of RStudio Connect will mark the experimental path as deprecated while
adding the endpoint without the `/experimental` prefix. The path with the experimental
prefix will be removed six months later.  The documentation for the endpoint will also
note, during that time, the original, experimental, path.

All experimental endpoints are clearly marked as such in this documentation.

### Deprecation and Removal of Old Versions

It is possible that RStudio Connect may decide to deprecate an endpoint.  This will
happen if either the endpoint serves no useful purpose because it’s functionality is
now handled by a different endpoint or because there is a newer version of the endpoint
that should be used.

If a deprecated endpoint is called, the response to it will include an extra HTTP
header called, `X-Deprecated-Endpoint` and will have as a value the path of the
endpoint that should be used instead.  If the functionality has no direct replacement,
the value will be set to `n/a`.

Deprecated versions of an endpoint will be supported for 1 year from the release date
of RStudio Connect in which the endpoint was marked as deprecated.  At that time, the
endpoint is subject to removal at the discretion of RStudio Connect.  The life cycle
of an endpoint will follow these steps.

1. The `/v1/endpoint` is public and in use by RStudio Connect customers.
1. RStudio Connect makes `/v2/experimental/endpoint` available for testing and feedback.
   Customers should still use `/v1/endpoint` for production work.
1. RStudio Connect moves version 2 of the endpoint out of experimentation so, all within
   the same release:
    1. `/v1/endpoint` is marked as deprecated.
    1. `/v2/experimental/endpoint` is marked as deprecated.
    1. `/v2/endpoint` is made public.
1. Six months later, `/v2/experimental/endpoint` is removed from the product.
1. Twelve months later, `/v1/endpoint` is removed from the product.

Note that it is possible that RStudio Connect may produce a new version of an existing
endpoint without making an experimental version of it first.  The same life cycle,
without those parts, will still be followed.

## Authentication

Most endpoints require you to identify yourself as a valid RStudio Connect
user. You do this by specifying an API Key when you make a call to the
server. The [API Keys](../user/api-keys/) chapter of the RStudio Connect
User Guide explains how to create an API Key.

### API Keys

API Keys are managed by each user in the RStudio Connect
dashboard. If you ever lose an API Key or otherwise feel it has
been compromised, use the dashboard to revoke the key and create
another one.

::alert danger
Keep your API Key safe.  If your RStudio Connect server's URL does not begin
with `https`, your API Key could be intercepted and used by a malicious actor.
::

Once you have an API Key, you can authenticate by passing the key with a prefix
of `\"Key \"` (the space is important) in the Authorization header.

Below are examples of invoking the \"Get R Information\" endpoint.

**cURL**
```bash
curl -H \"Authorization: Key XXXXXXXXXXX\" \\
     https://rstudioconnect.example.com/__api__/v1/server_settings/r
```

**R**
```r
library(httr)
apiKey <- \"XXXXXXXXXXX\"
result <- GET(\"https://rstudioconnect.example.com/__api__/v1/server_settings/r\",
  add_headers(Authorization = paste(\"Key\", apiKey)))
```

**Python**
```python
import requests
r = requests.get(
  'https://rstudioconnect.example.com/__api__/v1/server_settings/r',
  headers = { 'Authorization': 'Key XXXXXXXXXXX' }
)
```


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://support.rstudio.com/hc/en-us](https://support.rstudio.com/hc/en-us)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/__api__*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuditLogsApi* | [**GetAuditLogs**](docs/AuditLogsApi.md#getauditlogs) | **Get** /v1/audit_logs | Get audit logs
*BundlesApi* | [**DeleteBundle**](docs/BundlesApi.md#deletebundle) | **Delete** /v1/experimental/bundles/{id} | Delete bundle
*BundlesApi* | [**DownloadBundle**](docs/BundlesApi.md#downloadbundle) | **Get** /v1/experimental/bundles/{id}/download | Download the bundle archive
*BundlesApi* | [**GetBundle**](docs/BundlesApi.md#getbundle) | **Get** /v1/experimental/bundles/{id} | Get bundle details
*BundlesApi* | [**GetBundles**](docs/BundlesApi.md#getbundles) | **Get** /v1/experimental/content/{guid}/bundles | List bundles
*ContentApi* | [**CreateContent**](docs/ContentApi.md#createcontent) | **Post** /v1/experimental/content | Create content item
*ContentApi* | [**DeleteContent**](docs/ContentApi.md#deletecontent) | **Delete** /v1/experimental/content/{guid} | Delete content
*ContentApi* | [**DeployContentBundle**](docs/ContentApi.md#deploycontentbundle) | **Post** /v1/experimental/content/{guid}/deploy | Deploy deployment bundle
*ContentApi* | [**GetContent**](docs/ContentApi.md#getcontent) | **Get** /v1/experimental/content/{guid} | Get content details
*ContentApi* | [**UpdateContent**](docs/ContentApi.md#updatecontent) | **Post** /v1/experimental/content/{guid} | Update content
*ContentApi* | [**UploadContentBundle**](docs/ContentApi.md#uploadcontentbundle) | **Post** /v1/experimental/content/{guid}/upload | Upload deployment bundle
*GroupsApi* | [**AddGroupMember**](docs/GroupsApi.md#addgroupmember) | **Post** /v1/groups/{group_guid}/members | Add a group member
*GroupsApi* | [**CreateGroup**](docs/GroupsApi.md#creategroup) | **Post** /v1/groups | Create a group from caller-supplied details (Password, PAM, OAuth2, SAML, Proxied) 
*GroupsApi* | [**CreateRemoteGroup**](docs/GroupsApi.md#createremotegroup) | **Put** /v1/groups | Create a group using details from a remote authentication provider (LDAP) 
*GroupsApi* | [**DeleteGroup**](docs/GroupsApi.md#deletegroup) | **Delete** /v1/groups/{guid} | Delete a group
*GroupsApi* | [**GetGroup**](docs/GroupsApi.md#getgroup) | **Get** /v1/groups/{guid} | Get group details
*GroupsApi* | [**GetGroupMembers**](docs/GroupsApi.md#getgroupmembers) | **Get** /v1/groups/{group_guid}/members | Get group member details
*GroupsApi* | [**GetGroups**](docs/GroupsApi.md#getgroups) | **Get** /v1/groups | List or search for group details
*GroupsApi* | [**RemoveGroupMember**](docs/GroupsApi.md#removegroupmember) | **Delete** /v1/groups/{group_guid}/members/{user_guid} | Remove a group member
*GroupsApi* | [**SearchRemoteGroups**](docs/GroupsApi.md#searchremotegroups) | **Get** /v1/groups/remote | Search for group details from a remote provider
*InstrumentationApi* | [**GetContentVisits**](docs/InstrumentationApi.md#getcontentvisits) | **Get** /v1/instrumentation/content/visits | Get Content Visits
*InstrumentationApi* | [**GetShinyAppUsage**](docs/InstrumentationApi.md#getshinyappusage) | **Get** /v1/instrumentation/shiny/usage | Get Shiny App Usage
*RInformationApi* | [**GetRInformation**](docs/RInformationApi.md#getrinformation) | **Get** /v1/server_settings/r | Get R Information
*TasksApi* | [**GetTask**](docs/TasksApi.md#gettask) | **Get** /v1/experimental/tasks/{id} | Get task details
*UsersApi* | [**CreatePullUser**](docs/UsersApi.md#createpulluser) | **Put** /v1/users | Create a user using details from a remote authentication provider (LDAP, OAuth2) 
*UsersApi* | [**CreatePushUser**](docs/UsersApi.md#createpushuser) | **Post** /v1/users | Create a user from caller-supplied details (SAML, password, PAM, proxied) 
*UsersApi* | [**GetUser**](docs/UsersApi.md#getuser) | **Get** /v1/users/{guid} | Get user details
*UsersApi* | [**GetUsers**](docs/UsersApi.md#getusers) | **Get** /v1/users | List or search for user details
*UsersApi* | [**LockUser**](docs/UsersApi.md#lockuser) | **Post** /v1/users/{guid}/lock | Lock a user
*UsersApi* | [**SearchRemoteUsers**](docs/UsersApi.md#searchremoteusers) | **Get** /v1/users/remote | Search for user details from a remote provider
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Put** /v1/users/{guid} | Update a user


## Documentation For Models

 - [APIError](docs/APIError.md)
 - [AuditEntry](docs/AuditEntry.md)
 - [AuditLogs](docs/AuditLogs.md)
 - [AuditPager](docs/AuditPager.md)
 - [AuditPagerCursors](docs/AuditPagerCursors.md)
 - [Bundle](docs/Bundle.md)
 - [Bundles](docs/Bundles.md)
 - [Content](docs/Content.md)
 - [ContentDeploymentInstructions](docs/ContentDeploymentInstructions.md)
 - [ContentDeploymentTask](docs/ContentDeploymentTask.md)
 - [ContentUploadBundle](docs/ContentUploadBundle.md)
 - [ContentVisit](docs/ContentVisit.md)
 - [ContentVisitLogs](docs/ContentVisitLogs.md)
 - [ContentVisitPager](docs/ContentVisitPager.md)
 - [EditableUser](docs/EditableUser.md)
 - [Group](docs/Group.md)
 - [GroupMembers](docs/GroupMembers.md)
 - [GroupRemoteSearch](docs/GroupRemoteSearch.md)
 - [GroupWithTicket](docs/GroupWithTicket.md)
 - [Groups](docs/Groups.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject2](docs/InlineObject2.md)
 - [InlineObject3](docs/InlineObject3.md)
 - [InlineObject4](docs/InlineObject4.md)
 - [InlineObject5](docs/InlineObject5.md)
 - [InlineObject6](docs/InlineObject6.md)
 - [RInstallation](docs/RInstallation.md)
 - [RInstallations](docs/RInstallations.md)
 - [RemoteSearchResults](docs/RemoteSearchResults.md)
 - [ShinyAppUsage](docs/ShinyAppUsage.md)
 - [ShinyAppUsageLogs](docs/ShinyAppUsageLogs.md)
 - [ShinyAppUsagePager](docs/ShinyAppUsagePager.md)
 - [Task](docs/Task.md)
 - [User](docs/User.md)
 - [UserWithTicket](docs/UserWithTicket.md)
 - [Users](docs/Users.md)


## Documentation For Authorization



### apiKey

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@rstudio.com

