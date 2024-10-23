# SearchResultFood

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FdcId** | **int32** | Unique ID of the food. | [default to null]
**DataType** | **string** | The type of the food data. | [optional] [default to null]
**Description** | **string** | The description of the food. | [default to null]
**FoodCode** | **string** | Any A unique ID identifying the food within FNDDS. | [optional] [default to null]
**FoodNutrients** | [**[]AbridgedFoodNutrient**](AbridgedFoodNutrient.md) |  | [optional] [default to null]
**PublicationDate** | **string** | Date the item was published to FDC. | [optional] [default to null]
**ScientificName** | **string** | The scientific name of the food. | [optional] [default to null]
**BrandOwner** | **string** | Brand owner for the food. Only applies to Branded Foods. | [optional] [default to null]
**GtinUpc** | **string** | GTIN or UPC code identifying the food. Only applies to Branded Foods. | [optional] [default to null]
**Ingredients** | **string** | The list of ingredients (as it appears on the product label). Only applies to Branded Foods. | [optional] [default to null]
**NdbNumber** | **int32** | Unique number assigned for foundation foods. Only applies to Foundation and SRLegacy Foods. | [optional] [default to null]
**AdditionalDescriptions** | **string** | Any additional descriptions of the food. | [optional] [default to null]
**AllHighlightFields** | **string** | allHighlightFields | [optional] [default to null]
**Score** | **float32** | Relative score indicating how well the food matches the search criteria. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
