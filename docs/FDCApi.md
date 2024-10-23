# {{classname}}

All URIs are relative to *https://api.nal.usda.gov/fdc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFood**](FDCApi.md#GetFood) | **Get** /v1/food/{fdcId} | Fetches details for one food item by FDC ID
[**GetFoods**](FDCApi.md#GetFoods) | **Get** /v1/foods | Fetches details for multiple food items using input FDC IDs
[**GetFoodsList**](FDCApi.md#GetFoodsList) | **Get** /v1/foods/list | Returns a paged list of foods, in the &#x27;abridged&#x27; format
[**GetFoodsSearch**](FDCApi.md#GetFoodsSearch) | **Get** /v1/foods/search | Returns a list of foods that matched search (query) keywords
[**GetJsonSpec**](FDCApi.md#GetJsonSpec) | **Get** /v1/json-spec | Returns this documentation in JSON format
[**GetYamlSpec**](FDCApi.md#GetYamlSpec) | **Get** /v1/yaml-spec | Returns this documentation in JSON format
[**PostFoods**](FDCApi.md#PostFoods) | **Post** /v1/foods | Fetches details for multiple food items using input FDC IDs
[**PostFoodsList**](FDCApi.md#PostFoodsList) | **Post** /v1/foods/list | Returns a paged list of foods, in the &#x27;abridged&#x27; format
[**PostFoodsSearch**](FDCApi.md#PostFoodsSearch) | **Post** /v1/foods/search | Returns a list of foods that matched search (query) keywords

# **GetFood**
> InlineResponse200 GetFood(ctx, fdcId, optional)
Fetches details for one food item by FDC ID

Retrieves a single food item by an FDC ID. Optional format and nutrients can be specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fdcId** | **string**| FDC id of the food to retrieve |
 **optional** | ***FDCApiGetFoodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FDCApiGetFoodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**| Optional. &#x27;abridged&#x27; for an abridged set of elements, &#x27;full&#x27; for all elements (default). |
 **nutrients** | [**optional.Interface of []int32**](int32.md)| Optional. List of up to 25 nutrient numbers. Only the nutrient information for the specified nutrients will be returned. Should be comma separated list (e.g. nutrients&#x3D;203,204) or repeating parameters (e.g. nutrients&#x3D;203&amp;nutrients&#x3D;204). If a food does not have any matching nutrients, the food will be returned with an empty foodNutrients element. |

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFoods**
> []Object GetFoods(ctx, fdcIds, optional)
Fetches details for multiple food items using input FDC IDs

Retrieves a list of food items by a list of up to 20 FDC IDs. Optional format and nutrients can be specified. Invalid FDC ID's or ones that are not found are omitted and an empty set is returned if there are no matches.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fdcIds** | [**[]string**](string.md)| List of multiple FDC ID&#x27;s. Should be comma separated list (e.g. fdcIds&#x3D;534358,373052) or repeating parameters (e.g. fdcIds&#x3D;534358&amp;fdcIds&#x3D;373052). |
 **optional** | ***FDCApiGetFoodsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FDCApiGetFoodsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**| Optional. &#x27;abridged&#x27; for an abridged set of elements, &#x27;full&#x27; for all elements (default). |
 **nutrients** | [**optional.Interface of []int32**](int32.md)| Optional. List of up to 25 nutrient numbers. Only the nutrient information for the specified nutrients will be returned. Should be comma separated list (e.g. nutrients&#x3D;203,204) or repeating parameters (e.g. nutrients&#x3D;203&amp;nutrients&#x3D;204). If a food does not have any matching nutrients, the food will be returned with an empty foodNutrients element. |

### Return type

**[]Object**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFoodsList**
> []AbridgedFoodItem GetFoodsList(ctx, optional)
Returns a paged list of foods, in the 'abridged' format

Retrieves a paged list of foods. Use the pageNumber parameter to page through the entire result set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FDCApiGetFoodsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FDCApiGetFoodsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataType** | [**optional.Interface of []string**](string.md)| Optional. Filter on a specific data type; specify one or more values in an array. |
 **pageSize** | **optional.Int32**| Optional. Maximum number of results to return for the current page. Default is 50. |
 **pageNumber** | **optional.Int32**| Optional. Page number to retrieve. The offset into the overall result set is expressed as (pageNumber * pageSize) |
 **sortBy** | **optional.String**| Optional. Specify one of the possible values to sort by that field. Note, dataType.keyword will be dataType and lowercaseDescription.keyword will be description in future releases. |
 **sortOrder** | **optional.String**| Optional. The sort direction for the results. Only applicable if sortBy is specified. |

### Return type

[**[]AbridgedFoodItem**](AbridgedFoodItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFoodsSearch**
> SearchResult GetFoodsSearch(ctx, query, optional)
Returns a list of foods that matched search (query) keywords

Search for foods using keywords. Results can be filtered by dataType and there are options for result page sizes or sorting.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| One or more search terms.  The string may include [search operators](https://fdc.nal.usda.gov/help.html#bkmk-2) |
 **optional** | ***FDCApiGetFoodsSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FDCApiGetFoodsSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataType** | [**optional.Interface of []string**](string.md)| Optional. Filter on a specific data type; specify one or more values in an array. |
 **pageSize** | **optional.Int32**| Optional. Maximum number of results to return for the current page. Default is 50. |
 **pageNumber** | **optional.Int32**| Optional. Page number to retrieve. The offset into the overall result set is expressed as (pageNumber * pageSize) |
 **sortBy** | **optional.String**| Optional. Specify one of the possible values to sort by that field. Note, dataType.keyword will be dataType and lowercaseDescription.keyword will be description in future releases. |
 **sortOrder** | **optional.String**| Optional. The sort direction for the results. Only applicable if sortBy is specified. |
 **brandOwner** | **optional.String**| Optional. Filter results based on the brand owner of the food. Only applies to Branded Foods |

### Return type

[**SearchResult**](SearchResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJsonSpec**
> GetJsonSpec(ctx, )
Returns this documentation in JSON format

The OpenAPI 3.0 specification for the FDC API rendered as JSON (JavaScript Object Notation)

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetYamlSpec**
> GetYamlSpec(ctx, )
Returns this documentation in JSON format

The OpenAPI 3.0 specification for the FDC API rendered as YAML (YAML Ain't Markup Language)

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFoods**
> []Object PostFoods(ctx, body)
Fetches details for multiple food items using input FDC IDs

Retrieves a list of food items by a list of up to 20 FDC IDs. Optional format and nutrients can be specified. Invalid FDC ID's or ones that are not found are omitted and an empty set is returned if there are no matches.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FoodsCriteria**](FoodsCriteria.md)|  |

### Return type

**[]Object**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFoodsList**
> []AbridgedFoodItem PostFoodsList(ctx, body)
Returns a paged list of foods, in the 'abridged' format

Retrieves a paged list of foods. Use the pageNumber parameter to page through the entire result set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FoodListCriteria**](FoodListCriteria.md)|  |

### Return type

[**[]AbridgedFoodItem**](AbridgedFoodItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFoodsSearch**
> SearchResult PostFoodsSearch(ctx, body)
Returns a list of foods that matched search (query) keywords

Search for foods using keywords. Results can be filtered by dataType and there are options for result page sizes or sorting.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FoodSearchCriteria**](FoodSearchCriteria.md)| The query string may also include standard [search operators](https://fdc.nal.usda.gov/help.html#bkmk-2) |

### Return type

[**SearchResult**](SearchResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
