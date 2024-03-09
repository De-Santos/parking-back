package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"parking-back/mapper"
	"parking-back/models"
	"parking-back/obj"
	"parking-back/repository"
	"parking-back/utils"
	"parking-back/utils/request"
)

func AddParking(c *gin.Context) {
	var body obj.ParkingDto
	if e := request.BindValidBody(c, &body); e != nil {
		utils.ProcessBadResponse(c, e.Message)
		return
	}

	user, _ := c.Get("user")
	entity := mapper.MapToParkingModelWithUser(body, user.(models.User).ID)
	createdParking, err := repository.CreateParking(entity)
	if err != nil {
		utils.ProcessBadResponse(c, "Create parking failed: "+fmt.Sprint(err))
		return
	}

	dto := mapper.MapToParkingDto(createdParking)

	c.JSON(http.StatusOK, dto)
}

func GetParkingList(c *gin.Context) {
	var query obj.SearchQuery
	if e := request.BindValidQuery(c, &query); e != nil {
		utils.ProcessBadResponse(c, e.Message)
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

func UpdateParking(c *gin.Context) {
	var body obj.ParkingDto
	if e := request.BindValidBody(c, &body); e != nil {
		utils.ProcessBadResponse(c, e.Message)
		return
	}

	parking := mapper.MapToParkingModel(body)
	updatedParking, err := repository.UpdateParking(parking)
	if err != nil {
		utils.ProcessBadResponse(c, "Update failed: "+fmt.Sprint(err))
		return
	}

	response := mapper.MapToParkingDto(updatedParking)
	c.JSON(http.StatusOK, response)
}

func DeleteParking(c *gin.Context) {
	var query obj.IdQuery
	if e := request.BindValidQuery(c, &query); e != nil {
		utils.ProcessBadResponse(c, e.Message)
		return
	}

	result := repository.DeleteParkingById(query.ID)
	if result == false {
		utils.ProcessBadResponse(c, "Delete failed")
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
