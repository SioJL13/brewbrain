package models

type CalculateRecipeRequest struct {
	CoffeeGrams   *float64 `json:"coffeeGrams,omitempty"`
	WaterGrams    *float64 `json:"waterGrams,omitempty"`
	BrewingMethod string   `json:"brewingMethod,omitempty"`
	Strength      string   `json:"strength,omitempty"`
	GrinderType   string   `json:"grinderType,omitempty"`
	CoffeeName    string   `json:"coffeeName,omitempty"`
}

type CalculateRecipeResponse struct {
	CoffeeGrams float64 `json:"coffeeGrams"`
	WaterGrams  float64 `json:"waterGrams"`
	GrindSize   int     `json:"grindSize"`
	BrewTime    int     `json:"brewTime"`
}
