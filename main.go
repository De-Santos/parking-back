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
	initializers.InitializeVariables()
}

func main() {
	r := gin.Default()

	// Auth
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", middleware.RequireAuth, controllers.Logout)

	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/parking", middleware.RequireAuth, controllers.AddParking)

	_ = r.Run()
}
