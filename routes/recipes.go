package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siojl13/brewbrain/models"
)

const DEFAULT_RATIO = 16.0    // Std 1:16 coffee to water ratio
const DEFAULT_BREW_TIME = 180 // Default brew time in sec

func calculateRecipe(context *gin.Context) {
	var req models.CalculateRecipeRequest

	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var coffeeGrams, waterGrams float64

	if req.CoffeeGrams != nil {
		coffeeGrams = *req.CoffeeGrams
		waterGrams = coffeeGrams * DEFAULT_RATIO
	} else if req.WaterGrams != nil {
		waterGrams = *req.WaterGrams
		coffeeGrams = waterGrams / DEFAULT_RATIO
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Must provide either coffeeGrams or waterGrams"})
		return
	}

	//TODO: Check grindSize (?)
	//TODO: fetch historical brew data for certain coffee name, only if provided

	res := models.CalculateRecipeResponse{
		CoffeeGrams: coffeeGrams,
		WaterGrams:  waterGrams,
		GrindSize:   0,
		BrewTime:    DEFAULT_BREW_TIME,
	}

	context.JSON(http.StatusOK, res)

}
