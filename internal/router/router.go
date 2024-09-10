package router

import (
	"github.com/gin-gonic/gin"
	"myproject/internal/app"
	"myproject/internal/handlers"
	"myproject/internal/middleware"
)

func SetupRouter(a *app.App) *gin.Engine {
	r := gin.Default()

	authHandler := handlers.NewAuthHandler(a.Services.Auth)
	userHandler := handlers.NewUserHandler(a.Services.User)
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
		user := api.Group("/user").Use(middleware.AuthMiddleware(a.Config))
		{
			user.GET("/profile", userHandler.Profile)
		}
	}

	return r
}
