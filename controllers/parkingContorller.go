package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"parking-back/dtos"
	"parking-back/initializers"
	"parking-back/mapper"
	"parking-back/models"
	"parking-back/utils"
)

func AddParking(c *gin.Context) {
	// Get data from request
	var body dtos.ParkingDto
	if c.Bind(&body) != nil {
		utils.ProcessBadResponse(c, "Failed to read dto")
		return
	}

	// Validate data
	err := initializers.V.Struct(body)
	if err != nil {
		utils.ProcessBadResponse(c, "Invalid request body: "+fmt.Sprint(err))
		return
	}

	// Map to model
	user, _ := c.Get("user")
	entity := mapper.MapToParkingModel(body, user.(models.User).ID)

	fmt.Println(user)

	// Save into database
	initializers.DB.Create(&entity)
}
