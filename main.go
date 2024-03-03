package main

import (
	"github.com/gin-contrib/cors"
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

	r.Use(cors.Default())

	// Auth
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", middleware.RequireAuth, controllers.Logout)

	// Other
	r.GET("/username", controllers.CheckUsernameExistence)

	// Parking
	r.POST("/parking", middleware.RequireAuth, controllers.AddParking)
	r.GET("/parking", middleware.RequireAuth, controllers.GetParkingList)
	r.PUT("/parking", middleware.RequireAuth, controllers.UpdateParking)
	r.DELETE("/parking", middleware.RequireAuth, controllers.DeleteParking)

	// Car
	r.POST("/car", middleware.RequireAuth, controllers.AddCar)
	r.GET("/car", middleware.RequireAuth, controllers.GetCarList)
	r.PUT("/car", middleware.RequireAuth, controllers.UpdateCar)
	r.DELETE("/car", middleware.RequireAuth, controllers.DeleteCar)

	_ = r.Run()
}
