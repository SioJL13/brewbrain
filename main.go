package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siojl13/brewbrain/db"
	"github.com/siojl13/brewbrain/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
