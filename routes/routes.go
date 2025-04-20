package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/brews", getBrews)
	server.GET("/brews/:id", getBrew)
	server.POST("/brews", createBrew)
	server.PUT("/brews/:id", updateBrew)
	server.DELETE("/brews/:id", deleteBrew)

	server.POST("/brews/calculate", calculateRecipe)
}
