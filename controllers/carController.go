package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"parking-back/mapper"
	"parking-back/obj"
	"parking-back/repository"
	"parking-back/utils"
	"parking-back/utils/request"
)

func AddCar(c *gin.Context) {
	var body obj.CarDto
	if e := request.BindValidBody(c, &body); e != nil {
		utils.ProcessBadResponse(c, e.Message)
		return
	}

	entity := mapper.MapToCarModel(body)
	repository.SaveCar(entity)

	c.JSON(http.StatusOK, gin.H{})
}

func GetCarList(c *gin.Context) {
	var query obj.SearchQuery
	if e := request.BindValidQuery(c, &query); e != nil {
		utils.ProcessBadResponse(c, e.Message)
		fmt.Println(e.Err)
		return
	}

	carPage := repository.GetCarPage(&query, uint(query.Context))
	parkingDtoPage := mapper.MapToCarDtoList(carPage)

	var interfaceSlice []interface{}
	for _, carDto := range parkingDtoPage {
		interfaceSlice = append(interfaceSlice, carDto)
	}

	c.JSON(http.StatusOK, obj.PageableDtoWrapper{}.New(&query, interfaceSlice))
}

func UpdateCar(c *gin.Context) {
	var body obj.CarDto
	if e := request.BindValidBody(c, &body); e != nil {
		utils.ProcessBadResponse(c, e.Message)
		return
	}

	car := mapper.MapToCarModel(body)
	updatedCar, err := repository.UpdateCar(car)
	if err != nil {
		utils.ProcessBadResponse(c, "Update failed: "+fmt.Sprint(err))
		return
	}

	response := mapper.MapToCarDto(updatedCar)
	c.JSON(http.StatusOK, response)
}

func DeleteCar(c *gin.Context) {
	var query obj.IdQuery
	if e := request.BindValidQuery(c, &query); e != nil {
		utils.ProcessBadResponse(c, e.Message)
		return
	}

	result := repository.DeleteCarById(query.ID)
	if result == false {
		utils.ProcessBadResponse(c, "Delete failed")
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
