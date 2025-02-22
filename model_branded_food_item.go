/*
 * Food Data Central API
 *
 * The FoodData Central API provides REST access to FoodData Central (FDC). It is intended primarily to assist application developers wishing to incorporate nutrient data into their applications or websites.   To take full advantage of the API, developers should familiarize themselves with the database by reading the database documentation available via links on [Data Type Documentation](https://fdc.nal.usda.gov/data-documentation.html). This documentation provides the detailed definitions and descriptions needed to understand the data elements referenced in the API documentation.      Additional details about the API including rate limits, access, and licensing are available on the [FDC website](https://fdc.nal.usda.gov/api-guide.html)
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package fdc

type BrandedFoodItem struct {
	FdcId                    int32                          `json:"fdcId"`
	AvailableDate            string                         `json:"availableDate,omitempty"`
	BrandOwner               string                         `json:"brandOwner,omitempty"`
	DataSource               string                         `json:"dataSource,omitempty"`
	DataType                 string                         `json:"dataType"`
	Description              string                         `json:"description"`
	FoodClass                string                         `json:"foodClass,omitempty"`
	GtinUpc                  string                         `json:"gtinUpc,omitempty"`
	HouseholdServingFullText string                         `json:"householdServingFullText,omitempty"`
	Ingredients              string                         `json:"ingredients,omitempty"`
	ModifiedDate             string                         `json:"modifiedDate,omitempty"`
	PublicationDate          string                         `json:"publicationDate,omitempty"`
	ServingSize              interface{}                    `json:"servingSize,omitempty"`
	ServingSizeUnit          string                         `json:"servingSizeUnit,omitempty"`
	PreparationStateCode     string                         `json:"preparationStateCode,omitempty"`
	BrandedFoodCategory      string                         `json:"brandedFoodCategory,omitempty"`
	TradeChannel             []string                       `json:"tradeChannel,omitempty"`
	GpcClassCode             int32                          `json:"gpcClassCode,omitempty"`
	FoodNutrients            []FoodNutrient                 `json:"foodNutrients,omitempty"`
	FoodUpdateLog            []FoodUpdateLog                `json:"foodUpdateLog,omitempty"`
	LabelNutrients           *BrandedFoodItemLabelNutrients `json:"labelNutrients,omitempty"`
	PackageWeight            string                         `json:"packageWeight,omitempty"`
	ShortDescription         string                         `json:"shortDescription,omitempty"`
	DiscontinuedDate         string                         `json:"discontinuedDate,omitempty"`
	BrandName                string                         `json:"brandName,omitempty"`
	SubbrandName             string                         `json:"subbrandName,omitempty"`
	MarketCountry            string                         `json:"marketCountry,omitempty"`
	// FoodComponents []
	// FoodPortions []

}
