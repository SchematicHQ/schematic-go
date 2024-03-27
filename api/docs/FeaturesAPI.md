# \FeaturesAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckFlag**](FeaturesAPI.md#CheckFlag) | **Post** /flags/{key}/check | Check flag
[**CheckFlags**](FeaturesAPI.md#CheckFlags) | **Post** /flags/check | Check flags
[**CountCompaniesAudience**](FeaturesAPI.md#CountCompaniesAudience) | **Post** /audience/count-companies | Count Companies audience
[**CountFlagChecks**](FeaturesAPI.md#CountFlagChecks) | **Get** /flag-checks/count | Count flag checks
[**CountUsersAudience**](FeaturesAPI.md#CountUsersAudience) | **Post** /audience/count-users | Count Users audience
[**CreateFeature**](FeaturesAPI.md#CreateFeature) | **Post** /features | Create feature
[**CreateFlag**](FeaturesAPI.md#CreateFlag) | **Post** /flags | Create flag
[**DeleteFeature**](FeaturesAPI.md#DeleteFeature) | **Delete** /features/{feature_id} | Delete feature
[**DeleteFlag**](FeaturesAPI.md#DeleteFlag) | **Delete** /flags/{flag_id} | Delete flag
[**GetCompaniesAudience**](FeaturesAPI.md#GetCompaniesAudience) | **Post** /audience/get-companies | Get Companies audience
[**GetFeature**](FeaturesAPI.md#GetFeature) | **Get** /features/{feature_id} | Get feature
[**GetFlag**](FeaturesAPI.md#GetFlag) | **Get** /flags/{flag_id} | Get flag
[**GetFlagCheck**](FeaturesAPI.md#GetFlagCheck) | **Get** /flag-checks/{flag_check_id} | Get flag check
[**GetUsersAudience**](FeaturesAPI.md#GetUsersAudience) | **Post** /audience/get-users | Get Users audience
[**LatestFlagChecks**](FeaturesAPI.md#LatestFlagChecks) | **Get** /flag-checks/latest | Latest flag checks
[**ListFeatures**](FeaturesAPI.md#ListFeatures) | **Get** /features | List features
[**ListFlagChecks**](FeaturesAPI.md#ListFlagChecks) | **Get** /flag-checks | List flag checks
[**ListFlags**](FeaturesAPI.md#ListFlags) | **Get** /flags | List flags
[**RulesFlag**](FeaturesAPI.md#RulesFlag) | **Put** /flags/{flag_id}/rules | Rules flag
[**UpdateFeature**](FeaturesAPI.md#UpdateFeature) | **Put** /features/{feature_id} | Update feature
[**UpdateFlag**](FeaturesAPI.md#UpdateFlag) | **Put** /flags/{flag_id} | Update flag



## CheckFlag

> CheckFlagResponse CheckFlag(ctx, key).CheckFlagRequestBody(checkFlagRequestBody).Execute()

Check flag

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	key := "key_example" // string | key
	checkFlagRequestBody := *schematicapi.NewCheckFlagRequestBody() // CheckFlagRequestBody | 

	resp, r, err := client.API().FeaturesAPI.CheckFlag(context.Background(), key).CheckFlagRequestBody(checkFlagRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.CheckFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckFlag`: CheckFlagResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.CheckFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkFlagRequestBody** | [**CheckFlagRequestBody**](CheckFlagRequestBody.md) |  | 

### Return type

[**CheckFlagResponse**](CheckFlagResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckFlags

> CheckFlagsResponse CheckFlags(ctx).CheckFlagRequestBody(checkFlagRequestBody).Execute()

Check flags

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	checkFlagRequestBody := *schematicapi.NewCheckFlagRequestBody() // CheckFlagRequestBody | 

	resp, r, err := client.API().FeaturesAPI.CheckFlags(context.Background()).CheckFlagRequestBody(checkFlagRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.CheckFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckFlags`: CheckFlagsResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.CheckFlags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkFlagRequestBody** | [**CheckFlagRequestBody**](CheckFlagRequestBody.md) |  | 

### Return type

[**CheckFlagsResponse**](CheckFlagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountCompaniesAudience

> CountCompaniesAudienceResponse CountCompaniesAudience(ctx).AudienceRequestBody(audienceRequestBody).Execute()

Count Companies audience

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	audienceRequestBody := *schematicapi.NewAudienceRequestBody([]schematicapi.CreateOrUpdateConditionGroupRequestBody{*schematicapi.NewCreateOrUpdateConditionGroupRequestBody([]schematicapi.CreateOrUpdateConditionRequestBody{*schematicapi.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []schematicapi.CreateOrUpdateConditionRequestBody{*schematicapi.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}) // AudienceRequestBody | 

	resp, r, err := client.API().FeaturesAPI.CountCompaniesAudience(context.Background()).AudienceRequestBody(audienceRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.CountCompaniesAudience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountCompaniesAudience`: CountCompaniesAudienceResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.CountCompaniesAudience`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCompaniesAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audienceRequestBody** | [**AudienceRequestBody**](AudienceRequestBody.md) |  | 

### Return type

[**CountCompaniesAudienceResponse**](CountCompaniesAudienceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountFlagChecks

> CountFlagChecksResponse CountFlagChecks(ctx).FlagId(flagId).FlagIds(flagIds).Id(id).Limit(limit).Offset(offset).Execute()

Count flag checks

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	flagId := "flagId_example" // string |  (optional)
	flagIds := []string{"Inner_example"} // []string |  (optional)
	id := "id_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().FeaturesAPI.CountFlagChecks(context.Background()).FlagId(flagId).FlagIds(flagIds).Id(id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.CountFlagChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountFlagChecks`: CountFlagChecksResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.CountFlagChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountFlagChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flagId** | **string** |  | 
 **flagIds** | **[]string** |  | 
 **id** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountFlagChecksResponse**](CountFlagChecksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountUsersAudience

> CountUsersAudienceResponse CountUsersAudience(ctx).AudienceRequestBody(audienceRequestBody).Execute()

Count Users audience

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	audienceRequestBody := *schematicapi.NewAudienceRequestBody([]schematicapi.CreateOrUpdateConditionGroupRequestBody{*schematicapi.NewCreateOrUpdateConditionGroupRequestBody([]schematicapi.CreateOrUpdateConditionRequestBody{*schematicapi.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []schematicapi.CreateOrUpdateConditionRequestBody{*schematicapi.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}) // AudienceRequestBody | 

	resp, r, err := client.API().FeaturesAPI.CountUsersAudience(context.Background()).AudienceRequestBody(audienceRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.CountUsersAudience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountUsersAudience`: CountUsersAudienceResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.CountUsersAudience`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountUsersAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audienceRequestBody** | [**AudienceRequestBody**](AudienceRequestBody.md) |  | 

### Return type

[**CountUsersAudienceResponse**](CountUsersAudienceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFeature

> CreateFeatureResponse CreateFeature(ctx).CreateFeatureRequestBody(createFeatureRequestBody).Execute()

Create feature

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	createFeatureRequestBody := *schematicapi.NewCreateFeatureRequestBody("Description_example", "FeatureType_example", "Name_example") // CreateFeatureRequestBody | 

	resp, r, err := client.API().FeaturesAPI.CreateFeature(context.Background()).CreateFeatureRequestBody(createFeatureRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.CreateFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFeature`: CreateFeatureResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.CreateFeature`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFeatureRequestBody** | [**CreateFeatureRequestBody**](CreateFeatureRequestBody.md) |  | 

### Return type

[**CreateFeatureResponse**](CreateFeatureResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFlag

> CreateFlagResponse CreateFlag(ctx).CreateFlagRequestBody(createFlagRequestBody).Execute()

Create flag

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	createFlagRequestBody := *schematicapi.NewCreateFlagRequestBody(false, "Description_example", "FlagType_example", "Key_example", "Name_example") // CreateFlagRequestBody | 

	resp, r, err := client.API().FeaturesAPI.CreateFlag(context.Background()).CreateFlagRequestBody(createFlagRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.CreateFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFlag`: CreateFlagResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.CreateFlag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFlagRequestBody** | [**CreateFlagRequestBody**](CreateFlagRequestBody.md) |  | 

### Return type

[**CreateFlagResponse**](CreateFlagResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFeature

> DeleteFeatureResponse DeleteFeature(ctx, featureId).Execute()

Delete feature

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	featureId := "featureId_example" // string | feature_id

	resp, r, err := client.API().FeaturesAPI.DeleteFeature(context.Background(), featureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.DeleteFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFeature`: DeleteFeatureResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.DeleteFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | feature_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteFeatureResponse**](DeleteFeatureResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlag

> DeleteFlagResponse DeleteFlag(ctx, flagId).Execute()

Delete flag

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	flagId := "flagId_example" // string | flag_id

	resp, r, err := client.API().FeaturesAPI.DeleteFlag(context.Background(), flagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.DeleteFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFlag`: DeleteFlagResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.DeleteFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagId** | **string** | flag_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteFlagResponse**](DeleteFlagResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompaniesAudience

> GetCompaniesAudienceResponse GetCompaniesAudience(ctx).AudienceRequestBody(audienceRequestBody).Execute()

Get Companies audience

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	audienceRequestBody := *schematicapi.NewAudienceRequestBody([]schematicapi.CreateOrUpdateConditionGroupRequestBody{*schematicapi.NewCreateOrUpdateConditionGroupRequestBody([]schematicapi.CreateOrUpdateConditionRequestBody{*schematicapi.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []schematicapi.CreateOrUpdateConditionRequestBody{*schematicapi.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}) // AudienceRequestBody | 

	resp, r, err := client.API().FeaturesAPI.GetCompaniesAudience(context.Background()).AudienceRequestBody(audienceRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.GetCompaniesAudience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompaniesAudience`: GetCompaniesAudienceResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.GetCompaniesAudience`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audienceRequestBody** | [**AudienceRequestBody**](AudienceRequestBody.md) |  | 

### Return type

[**GetCompaniesAudienceResponse**](GetCompaniesAudienceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeature

> GetFeatureResponse GetFeature(ctx, featureId).Execute()

Get feature

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	featureId := "featureId_example" // string | feature_id

	resp, r, err := client.API().FeaturesAPI.GetFeature(context.Background(), featureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.GetFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeature`: GetFeatureResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.GetFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | feature_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFeatureResponse**](GetFeatureResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlag

> GetFlagResponse GetFlag(ctx, flagId).Execute()

Get flag

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	flagId := "flagId_example" // string | flag_id

	resp, r, err := client.API().FeaturesAPI.GetFlag(context.Background(), flagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.GetFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlag`: GetFlagResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.GetFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagId** | **string** | flag_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFlagResponse**](GetFlagResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagCheck

> GetFlagCheckResponse GetFlagCheck(ctx, flagCheckId).Execute()

Get flag check

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	flagCheckId := "flagCheckId_example" // string | flag_check_id

	resp, r, err := client.API().FeaturesAPI.GetFlagCheck(context.Background(), flagCheckId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.GetFlagCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlagCheck`: GetFlagCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.GetFlagCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagCheckId** | **string** | flag_check_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFlagCheckResponse**](GetFlagCheckResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersAudience

> GetUsersAudienceResponse GetUsersAudience(ctx).AudienceRequestBody(audienceRequestBody).Execute()

Get Users audience

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	audienceRequestBody := *schematicapi.NewAudienceRequestBody([]schematicapi.CreateOrUpdateConditionGroupRequestBody{*schematicapi.NewCreateOrUpdateConditionGroupRequestBody([]schematicapi.CreateOrUpdateConditionRequestBody{*schematicapi.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []schematicapi.CreateOrUpdateConditionRequestBody{*schematicapi.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}) // AudienceRequestBody | 

	resp, r, err := client.API().FeaturesAPI.GetUsersAudience(context.Background()).AudienceRequestBody(audienceRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.GetUsersAudience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersAudience`: GetUsersAudienceResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.GetUsersAudience`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audienceRequestBody** | [**AudienceRequestBody**](AudienceRequestBody.md) |  | 

### Return type

[**GetUsersAudienceResponse**](GetUsersAudienceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LatestFlagChecks

> LatestFlagChecksResponse LatestFlagChecks(ctx).FlagId(flagId).FlagIds(flagIds).Id(id).Limit(limit).Offset(offset).Execute()

Latest flag checks

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	flagId := "flagId_example" // string |  (optional)
	flagIds := []string{"Inner_example"} // []string |  (optional)
	id := "id_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().FeaturesAPI.LatestFlagChecks(context.Background()).FlagId(flagId).FlagIds(flagIds).Id(id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.LatestFlagChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LatestFlagChecks`: LatestFlagChecksResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.LatestFlagChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLatestFlagChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flagId** | **string** |  | 
 **flagIds** | **[]string** |  | 
 **id** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**LatestFlagChecksResponse**](LatestFlagChecksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatures

> ListFeaturesResponse ListFeatures(ctx).Limit(limit).Offset(offset).Execute()

List features

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().FeaturesAPI.ListFeatures(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.ListFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeatures`: ListFeaturesResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.ListFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListFeaturesResponse**](ListFeaturesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlagChecks

> ListFlagChecksResponse ListFlagChecks(ctx).FlagId(flagId).FlagIds(flagIds).Id(id).Limit(limit).Offset(offset).Execute()

List flag checks

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	flagId := "flagId_example" // string |  (optional)
	flagIds := []string{"Inner_example"} // []string |  (optional)
	id := "id_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().FeaturesAPI.ListFlagChecks(context.Background()).FlagId(flagId).FlagIds(flagIds).Id(id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.ListFlagChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFlagChecks`: ListFlagChecksResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.ListFlagChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFlagChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flagId** | **string** |  | 
 **flagIds** | **[]string** |  | 
 **id** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListFlagChecksResponse**](ListFlagChecksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlags

> ListFlagsResponse ListFlags(ctx).FeatureId(featureId).FlagIds(flagIds).Limit(limit).Offset(offset).Execute()

List flags

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	featureId := "featureId_example" // string |  (optional)
	flagIds := []string{"Inner_example"} // []string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().FeaturesAPI.ListFlags(context.Background()).FeatureId(featureId).FlagIds(flagIds).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.ListFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFlags`: ListFlagsResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.ListFlags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featureId** | **string** |  | 
 **flagIds** | **[]string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListFlagsResponse**](ListFlagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesFlag

> RulesFlagResponse RulesFlag(ctx, flagId).UpdateFlagRulesRequestBody(updateFlagRulesRequestBody).Execute()

Rules flag

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	flagId := "flagId_example" // string | flag_id
	updateFlagRulesRequestBody := *schematicapi.NewUpdateFlagRulesRequestBody([]schematicapi.CreateOrUpdateRuleRequestBody{*schematicapi.NewCreateOrUpdateRuleRequestBody([]schematicapi.CreateOrUpdateConditionGroupRequestBody{*schematicapi.NewCreateOrUpdateConditionGroupRequestBody([]schematicapi.CreateOrUpdateConditionRequestBody{*schematicapi.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []schematicapi.CreateOrUpdateConditionRequestBody{*schematicapi.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}, "Name_example", int32(123), false)}) // UpdateFlagRulesRequestBody | 

	resp, r, err := client.API().FeaturesAPI.RulesFlag(context.Background(), flagId).UpdateFlagRulesRequestBody(updateFlagRulesRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.RulesFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesFlag`: RulesFlagResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.RulesFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagId** | **string** | flag_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFlagRulesRequestBody** | [**UpdateFlagRulesRequestBody**](UpdateFlagRulesRequestBody.md) |  | 

### Return type

[**RulesFlagResponse**](RulesFlagResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeature

> UpdateFeatureResponse UpdateFeature(ctx, featureId).UpdateFeatureRequestBody(updateFeatureRequestBody).Execute()

Update feature

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	featureId := "featureId_example" // string | feature_id
	updateFeatureRequestBody := *schematicapi.NewUpdateFeatureRequestBody() // UpdateFeatureRequestBody | 

	resp, r, err := client.API().FeaturesAPI.UpdateFeature(context.Background(), featureId).UpdateFeatureRequestBody(updateFeatureRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.UpdateFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFeature`: UpdateFeatureResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.UpdateFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | feature_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFeatureRequestBody** | [**UpdateFeatureRequestBody**](UpdateFeatureRequestBody.md) |  | 

### Return type

[**UpdateFeatureResponse**](UpdateFeatureResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlag

> UpdateFlagResponse UpdateFlag(ctx, flagId).CreateFlagRequestBody(createFlagRequestBody).Execute()

Update flag

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	flagId := "flagId_example" // string | flag_id
	createFlagRequestBody := *schematicapi.NewCreateFlagRequestBody(false, "Description_example", "FlagType_example", "Key_example", "Name_example") // CreateFlagRequestBody | 

	resp, r, err := client.API().FeaturesAPI.UpdateFlag(context.Background(), flagId).CreateFlagRequestBody(createFlagRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.UpdateFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFlag`: UpdateFlagResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.UpdateFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagId** | **string** | flag_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFlagRequestBody** | [**CreateFlagRequestBody**](CreateFlagRequestBody.md) |  | 

### Return type

[**UpdateFlagResponse**](UpdateFlagResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

