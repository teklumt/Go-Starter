package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router  = gin.Default()

func IntializeRoutes() {
	publicRoutes := Router.Group("v1/")

	publicRoutes.GET("health", HealthCheck)
	
}

func HealthCheck(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{
		"message": "API is Up and Running",
	})



}