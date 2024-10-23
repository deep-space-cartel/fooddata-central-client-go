/*
 * Food Data Central API
 *
 * The FoodData Central API provides REST access to FoodData Central (FDC). It is intended primarily to assist application developers wishing to incorporate nutrient data into their applications or websites.   To take full advantage of the API, developers should familiarize themselves with the database by reading the database documentation available via links on [Data Type Documentation](https://fdc.nal.usda.gov/data-documentation.html). This documentation provides the detailed definitions and descriptions needed to understand the data elements referenced in the API documentation.      Additional details about the API including rate limits, access, and licensing are available on the [FDC website](https://fdc.nal.usda.gov/api-guide.html)
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A copy of the criteria that were used in the search.
type FoodSearchCriteria struct {
	// Search terms to use in the search. The string may also include standard [search operators](https://fdc.nal.usda.gov/help.html#bkmk-2)
	Query string `json:"query,omitempty"`
	// Optional. Filter on a specific data type; specify one or more values in an array.
	DataType []string `json:"dataType,omitempty"`
	// Optional. Maximum number of results to return for the current page. Default is 50.
	PageSize int32 `json:"pageSize,omitempty"`
	// Optional. Page number to retrieve. The offset into the overall result set is expressed as (pageNumber * pageSize)
	PageNumber int32 `json:"pageNumber,omitempty"`
	// Optional. Specify one of the possible values to sort by that field. Note, dataType.keyword will be dataType and description.keyword will be description in future releases.
	SortBy string `json:"sortBy,omitempty"`
	// Optional. The sort direction for the results. Only applicable if sortBy is specified.
	SortOrder string `json:"sortOrder,omitempty"`
	// Optional. Filter results based on the brand owner of the food. Only applies to Branded Foods.
	BrandOwner string `json:"brandOwner,omitempty"`
	// Optional. Filter foods containing any of the specified trade channels.
	TradeChannel []string `json:"tradeChannel,omitempty"`
	// Filter foods published on or after this date. Format: YYYY-MM-DD
	StartDate string `json:"startDate,omitempty"`
	// Filter foods published on or before this date. Format: YYYY-MM-DD
	EndDate string `json:"endDate,omitempty"`
}