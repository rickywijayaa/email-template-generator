package routes

import (
	"email-template-generator/app/user"
	"email-template-generator/auth"
	"email-template-generator/database"
	"email-template-generator/handler"
	"email-template-generator/middleware"

	"github.com/gin-gonic/gin"
)

func New(db *database.Connection) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	//Init Repo
	userRepository := user.NewRepository(db.DB)

	//Init Service
	userService := user.NewService(userRepository)
	authService := auth.NewJwtService()

	//Init Handler
	userHandler := handler.NewUserHandler(userService, authService)

	api := router.Group("/api/v1")
	api.POST("/login", userHandler.LoginUser)
	api.GET("/middleware", middleware.AuthMiddleware(userService, authService), userHandler.TestingMiddleware)

	return router
}
