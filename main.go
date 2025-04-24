package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siojl13/brewbrain/db"
	"github.com/siojl13/brewbrain/models"
	"github.com/siojl13/brewbrain/routes"
)

func main() {
	db.InitDB()

	if err := models.Migrate(db.DB); err != nil {
		panic(err)
	}

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
