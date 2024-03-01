package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"parking-back/initializers"
	"parking-back/mapper"
	"parking-back/models"
	"parking-back/obj"
	"parking-back/repository"
	"parking-back/utils"
)

func AddParking(c *gin.Context) {
	var body obj.ParkingDto
	if c.Bind(&body) != nil {
		utils.ProcessBadResponse(c, "Failed to read dto")
		return
	}

	err := initializers.V.Struct(body)
	if err != nil {
		utils.ProcessBadResponse(c, "Invalid request body: "+fmt.Sprint(err))
		return
	}

	user, _ := c.Get("user")
	entity := mapper.MapToParkingModel(body, user.(models.User).ID)

	initializers.DB.Create(&entity)
}

func GetParkingList(c *gin.Context) {
	var query obj.ParkingSearchQuery
	if c.BindQuery(&query) != nil {
		utils.ProcessBadResponse(c, "Invalid query params")
		return
	}

	err := initializers.V.Struct(query)
	if err != nil {
		utils.ProcessBadResponse(c, "Invalid request query: "+fmt.Sprint(err))
		return
	}

	parkingPage := repository.GetParkingPage(&query)
	parkingDtoPage := mapper.MapToParkingDtoList(parkingPage)

	var interfaceSlice []interface{}
	for _, parkingDto := range parkingDtoPage {
		interfaceSlice = append(interfaceSlice, parkingDto)
	}

	c.JSON(http.StatusOK, obj.PageableDtoWrapper{}.New(&query, interfaceSlice))
}
