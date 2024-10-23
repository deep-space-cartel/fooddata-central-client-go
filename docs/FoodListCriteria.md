# FoodListCriteria

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **[]string** | Optional. Filter on a specific data type; specify one or more values in an array. | [optional] [default to null]
**PageSize** | **int32** | Optional. Maximum number of results to return for the current page. Default is 50. | [optional] [default to null]
**PageNumber** | **int32** | Optional. Page number to retrieve. The offset into the overall result set is expressed as (pageNumber * pageSize) | [optional] [default to null]
**SortBy** | **string** | Optional. Specify one of the possible values to sort by that field. Note, dataType.keyword will be dataType and lowercaseDescription.keyword will be description in future releases. | [optional] [default to null]
**SortOrder** | **string** | Optional. The sort direction for the results. Only applicable if sortBy is specified. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
