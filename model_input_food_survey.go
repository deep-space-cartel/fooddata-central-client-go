/*
 * Food Data Central API
 *
 * The FoodData Central API provides REST access to FoodData Central (FDC). It is intended primarily to assist application developers wishing to incorporate nutrient data into their applications or websites.   To take full advantage of the API, developers should familiarize themselves with the database by reading the database documentation available via links on [Data Type Documentation](https://fdc.nal.usda.gov/data-documentation.html). This documentation provides the detailed definitions and descriptions needed to understand the data elements referenced in the API documentation.      Additional details about the API including rate limits, access, and licensing are available on the [FDC website](https://fdc.nal.usda.gov/api-guide.html)
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package fdc

// applies to Survey (FNDDS). Not all inputFoods will have all fields.
type InputFoodSurvey struct {
	Id                    int32            `json:"id,omitempty"`
	Amount                float32          `json:"amount,omitempty"`
	FoodDescription       string           `json:"foodDescription,omitempty"`
	IngredientCode        int32            `json:"ingredientCode,omitempty"`
	IngredientDescription string           `json:"ingredientDescription,omitempty"`
	IngredientWeight      float32          `json:"ingredientWeight,omitempty"`
	PortionCode           string           `json:"portionCode,omitempty"`
	PortionDescription    string           `json:"portionDescription,omitempty"`
	SequenceNumber        int32            `json:"sequenceNumber,omitempty"`
	SurveyFlag            int32            `json:"surveyFlag,omitempty"`
	Unit                  string           `json:"unit,omitempty"`
	InputFood             *SurveyFoodItem  `json:"inputFood,omitempty"`
	RetentionFactor       *RetentionFactor `json:"retentionFactor,omitempty"`
}
