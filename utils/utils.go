package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ProcessBadResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
}

func GetUint64(str string) uint64 {
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		fmt.Println("Error parsing uint str:", err)
		panic(err)
	}
	return id
}

func GetUint(str string) uint {
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		fmt.Println("Error parsing uint str:", err)
		panic(err)
	}
	return uint(id)
}

func GetOffset(page int, limit int) int {
	return (page - 1) * limit
}
