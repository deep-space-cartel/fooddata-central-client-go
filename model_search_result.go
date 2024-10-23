/*
 * Food Data Central API
 *
 * The FoodData Central API provides REST access to FoodData Central (FDC). It is intended primarily to assist application developers wishing to incorporate nutrient data into their applications or websites.   To take full advantage of the API, developers should familiarize themselves with the database by reading the database documentation available via links on [Data Type Documentation](https://fdc.nal.usda.gov/data-documentation.html). This documentation provides the detailed definitions and descriptions needed to understand the data elements referenced in the API documentation.      Additional details about the API including rate limits, access, and licensing are available on the [FDC website](https://fdc.nal.usda.gov/api-guide.html)
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SearchResult struct {
	FoodSearchCriteria *FoodSearchCriteria `json:"foodSearchCriteria,omitempty"`
	// The total number of foods found matching the search criteria.
	TotalHits int32 `json:"totalHits,omitempty"`
	// The current page of results being returned.
	CurrentPage int32 `json:"currentPage,omitempty"`
	// The total number of pages found matching the search criteria.
	TotalPages int32 `json:"totalPages,omitempty"`
	// The list of foods found matching the search criteria. See Food Fields below.
	Foods []SearchResultFood `json:"foods,omitempty"`
}