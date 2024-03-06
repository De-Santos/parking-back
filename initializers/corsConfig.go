package initializers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization")
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	return cors.New(config)
}
