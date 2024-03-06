package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"parking-back/obj"
	"parking-back/repository"
	"parking-back/utils"
	"parking-back/utils/request"
)

func CheckUsernameExistence(c *gin.Context) {
	var query obj.StringQuery
	if e := request.BindValidQuery(c, &query); e != nil {
		utils.ProcessBadResponse(c, e.Message)
		return
	}

	result := repository.CheckUsernameExistence(query.String)
	c.JSON(http.StatusOK, result)
}
