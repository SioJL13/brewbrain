package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siojl13/brewbrain/models"
)

func main() {
	server := gin.Default()

	server.GET("/brews", getBrews)
	server.POST("/brews", createBrews)

	server.Run(":8080")
}

func getBrews(context *gin.Context) {
	brews := models.GetAllBrews()
	context.JSON(http.StatusOK, brews)
}

func createBrews(context *gin.Context) {
	var brew models.Brew

	err := context.ShouldBindJSON(&brew)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	brew.ID = 1

	brew.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "brew created", "brew": brew})
}
