package main

import (
	"github.com/gin-gonic/gin"
	"parking-back/controllers"
	"parking-back/initializers"
	"parking-back/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	// Auth
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", middleware.RequireAuth, controllers.Logout)

	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	_ = r.Run()
}
