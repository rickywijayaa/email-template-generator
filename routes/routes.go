package routes

import (
	"email-template-generator/app/user"
	"email-template-generator/database"
	"email-template-generator/handler"

	"github.com/gin-gonic/gin"
)

func New(db *database.Connection) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	//Init Repo
	userRepository := user.NewRepository(db.DB)

	//Init Service
	userService := user.NewService(userRepository)

	//Init Handler
	userHandler := handler.NewUserHandler(userService)

	api := router.Group("/api/v1")
	api.POST("/login", userHandler.LoginUser)

	return router
}
