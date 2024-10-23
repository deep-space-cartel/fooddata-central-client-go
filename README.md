# Go API client for swagger

The FoodData Central API provides REST access to FoodData Central (FDC). It is intended primarily to assist application developers wishing to incorporate nutrient data into their applications or websites.   To take full advantage of the API, developers should familiarize themselves with the database by reading the database documentation available via links on [Data Type Documentation](https://fdc.nal.usda.gov/data-documentation.html). This documentation provides the detailed definitions and descriptions needed to understand the data elements referenced in the API documentation.      Additional details about the API including rate limits, access, and licensing are available on the [FDC website](https://fdc.nal.usda.gov/api-guide.html)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.1
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://nal.altarama.com/reft100.aspx?key=FoodData](https://nal.altarama.com/reft100.aspx?key=FoodData)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.nal.usda.gov/fdc*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FDCApi* | [**GetFood**](docs/FDCApi.md#getfood) | **Get** /v1/food/{fdcId} | Fetches details for one food item by FDC ID
*FDCApi* | [**GetFoods**](docs/FDCApi.md#getfoods) | **Get** /v1/foods | Fetches details for multiple food items using input FDC IDs
*FDCApi* | [**GetFoodsList**](docs/FDCApi.md#getfoodslist) | **Get** /v1/foods/list | Returns a paged list of foods, in the &#x27;abridged&#x27; format
*FDCApi* | [**GetFoodsSearch**](docs/FDCApi.md#getfoodssearch) | **Get** /v1/foods/search | Returns a list of foods that matched search (query) keywords
*FDCApi* | [**GetJsonSpec**](docs/FDCApi.md#getjsonspec) | **Get** /v1/json-spec | Returns this documentation in JSON format
*FDCApi* | [**GetYamlSpec**](docs/FDCApi.md#getyamlspec) | **Get** /v1/yaml-spec | Returns this documentation in JSON format
*FDCApi* | [**PostFoods**](docs/FDCApi.md#postfoods) | **Post** /v1/foods | Fetches details for multiple food items using input FDC IDs
*FDCApi* | [**PostFoodsList**](docs/FDCApi.md#postfoodslist) | **Post** /v1/foods/list | Returns a paged list of foods, in the &#x27;abridged&#x27; format
*FDCApi* | [**PostFoodsSearch**](docs/FDCApi.md#postfoodssearch) | **Post** /v1/foods/search | Returns a list of foods that matched search (query) keywords

## Documentation For Models

 - [AbridgedFoodItem](docs/AbridgedFoodItem.md)
 - [AbridgedFoodNutrient](docs/AbridgedFoodNutrient.md)
 - [BrandedFoodItem](docs/BrandedFoodItem.md)
 - [BrandedFoodItemLabelNutrients](docs/BrandedFoodItemLabelNutrients.md)
 - [BrandedFoodItemLabelNutrientsCalcium](docs/BrandedFoodItemLabelNutrientsCalcium.md)
 - [BrandedFoodItemLabelNutrientsCalories](docs/BrandedFoodItemLabelNutrientsCalories.md)
 - [BrandedFoodItemLabelNutrientsCarbohydrates](docs/BrandedFoodItemLabelNutrientsCarbohydrates.md)
 - [BrandedFoodItemLabelNutrientsFat](docs/BrandedFoodItemLabelNutrientsFat.md)
 - [BrandedFoodItemLabelNutrientsFiber](docs/BrandedFoodItemLabelNutrientsFiber.md)
 - [BrandedFoodItemLabelNutrientsIron](docs/BrandedFoodItemLabelNutrientsIron.md)
 - [BrandedFoodItemLabelNutrientsPotassium](docs/BrandedFoodItemLabelNutrientsPotassium.md)
 - [BrandedFoodItemLabelNutrientsProtein](docs/BrandedFoodItemLabelNutrientsProtein.md)
 - [BrandedFoodItemLabelNutrientsSaturatedFat](docs/BrandedFoodItemLabelNutrientsSaturatedFat.md)
 - [BrandedFoodItemLabelNutrientsSugars](docs/BrandedFoodItemLabelNutrientsSugars.md)
 - [BrandedFoodItemLabelNutrientsTransFat](docs/BrandedFoodItemLabelNutrientsTransFat.md)
 - [FoodAttribute](docs/FoodAttribute.md)
 - [FoodAttributeFoodAttributeType](docs/FoodAttributeFoodAttributeType.md)
 - [FoodCategory](docs/FoodCategory.md)
 - [FoodComponent](docs/FoodComponent.md)
 - [FoodListCriteria](docs/FoodListCriteria.md)
 - [FoodNutrient](docs/FoodNutrient.md)
 - [FoodNutrientDerivation](docs/FoodNutrientDerivation.md)
 - [FoodNutrientSource](docs/FoodNutrientSource.md)
 - [FoodPortion](docs/FoodPortion.md)
 - [FoodSearchCriteria](docs/FoodSearchCriteria.md)
 - [FoodUpdateLog](docs/FoodUpdateLog.md)
 - [FoodsCriteria](docs/FoodsCriteria.md)
 - [FoundationFoodItem](docs/FoundationFoodItem.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InputFoodFoundation](docs/InputFoodFoundation.md)
 - [InputFoodSurvey](docs/InputFoodSurvey.md)
 - [MeasureUnit](docs/MeasureUnit.md)
 - [Nutrient](docs/Nutrient.md)
 - [NutrientAcquisitionDetails](docs/NutrientAcquisitionDetails.md)
 - [NutrientAnalysisDetails](docs/NutrientAnalysisDetails.md)
 - [NutrientConversionFactors](docs/NutrientConversionFactors.md)
 - [RetentionFactor](docs/RetentionFactor.md)
 - [SampleFoodItem](docs/SampleFoodItem.md)
 - [SearchResult](docs/SearchResult.md)
 - [SearchResultFood](docs/SearchResultFood.md)
 - [SrLegacyFoodItem](docs/SrLegacyFoodItem.md)
 - [SurveyFoodItem](docs/SurveyFoodItem.md)
 - [WweiaFoodCategory](docs/WweiaFoodCategory.md)

## Documentation For Authorization

## ApiKeyAuth
- **Type**: API key

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

1) U.S. Department of Agriculture. Agricultural Research Service.

## Sources

* [FoodData Central API Guide](https://fdc.nal.usda.gov/api-guide.html)
* [Swagger Hub: FoodData Central API v1.0.1](https://app.swaggerhub.com/apis/fdcnal/food-data_central_api/1.0.1)


## Dependencies

```bash
go install -v github.com/go-critic/go-critic/cmd/gocritic@latest
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install golang.org/x/tools/cmd/goimports@latest
```
