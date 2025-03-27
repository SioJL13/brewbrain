package models

import "time"

type Brew struct {
	ID             uint
	CoffeeName     string
	CoffeeType     string  // honey, natural, lavado
	CoffeeGrams    float64 `binding:"required"`
	GrindSize      int     `binding:"required"` // numeric grind size
	WaterGrams     float64 `binding:"required"`
	BrewingMethod  string  `binding:"required"` // kalita, breville, aeropress, etc
	BrewTime       int     `binding:"required"` //in seconds
	ExtractionTime int     // in seconds
	WaterTemp      float64 // Celcius
	GrinderType    string  // breville, timemore
	CreatedAt      time.Time
}

var brews = []Brew{}

func (b Brew) Save() {
	//TODO: later add database
	brews = append(brews, b)
}

func GetAllBrews() []Brew {
	return brews
}
