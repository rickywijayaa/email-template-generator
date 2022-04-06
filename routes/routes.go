package routes

import (
	"email-template-generator/handler"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", handler.Get)

	return router
}
