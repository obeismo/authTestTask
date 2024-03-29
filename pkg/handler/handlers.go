package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", homePage)

	auth := router.Group("/auth")
	{
		auth.POST("/issuetokens", signUp)
		auth.POST("/refresh", signIn)
	}

	return router
}
