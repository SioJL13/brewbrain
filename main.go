package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/siojl13/brewbrain/db"
	"github.com/siojl13/brewbrain/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/brews", getBrews)
	server.GET("/brews/:id", getBrew)
	server.POST("/brews", createBrews)

	server.Run(":8080")
}

func getBrews(context *gin.Context) {
	brews, err := models.GetAllBrews()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, brews)
}

func getBrew(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	brew, err := models.GetBrewByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	context.JSON(http.StatusOK, brew)
}

func createBrews(context *gin.Context) {
	var brew models.Brew

	err := context.ShouldBindJSON(&brew)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	brew.ID = 1

	err = brew.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "brew created", "brew": brew})
}
