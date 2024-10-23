# FoodSearchCriteria

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Search terms to use in the search. The string may also include standard [search operators](https://fdc.nal.usda.gov/help.html#bkmk-2) | [optional] [default to null]
**DataType** | **[]string** | Optional. Filter on a specific data type; specify one or more values in an array. | [optional] [default to null]
**PageSize** | **int32** | Optional. Maximum number of results to return for the current page. Default is 50. | [optional] [default to null]
**PageNumber** | **int32** | Optional. Page number to retrieve. The offset into the overall result set is expressed as (pageNumber * pageSize) | [optional] [default to null]
**SortBy** | **string** | Optional. Specify one of the possible values to sort by that field. Note, dataType.keyword will be dataType and description.keyword will be description in future releases. | [optional] [default to null]
**SortOrder** | **string** | Optional. The sort direction for the results. Only applicable if sortBy is specified. | [optional] [default to null]
**BrandOwner** | **string** | Optional. Filter results based on the brand owner of the food. Only applies to Branded Foods. | [optional] [default to null]
**TradeChannel** | **[]string** | Optional. Filter foods containing any of the specified trade channels. | [optional] [default to null]
**StartDate** | **string** | Filter foods published on or after this date. Format: YYYY-MM-DD | [optional] [default to null]
**EndDate** | **string** | Filter foods published on or before this date. Format: YYYY-MM-DD | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
