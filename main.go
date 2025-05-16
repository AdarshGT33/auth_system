package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/handlers"
	"main.go/middlewares"
)

func init() {
	handlers.LoadEnv()
	handlers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "UPDATE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authRoute := r.Group("/auth")
	protectedRoute := r.Group("/verified")

	r.GET("/")
	authRoute.POST("/signup", controllers.RegisterUser)
	authRoute.POST("/login", controllers.LoginUser)
	protectedRoute.GET("/dashboard", middlewares.AuthMiddleware())

	r.Run(":8000")
}
