package routes

import (
	"email-template-generator/database"
	"email-template-generator/handler"

	"github.com/gin-gonic/gin"
)

func New(db *database.Connection) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", handler.Get)

	return router
}
