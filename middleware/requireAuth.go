package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"parking-back/initializers"
	jwt2 "parking-back/jwt"
	"parking-back/models"
	"parking-back/repository"
	"parking-back/utils"
	"time"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie off request
	cookieToken, err := c.Cookie("Authorization")
	var headerToken jwt2.HeaderToken
	err2 := c.BindHeader(&headerToken)
	var token string

	if err == nil {
		token = cookieToken
	} else if err2 == nil {
		token, err = headerToken.GetToken()
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decode and validate jwt
	claims, err := jwt2.ParseJwtClaims(token)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Check the exp
	if time.Now().After(claims.ExpiresAt.Time) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if exist, err := repository.IsTokenInvalidated(utils.GetUint64(claims.ID)); exist || err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Find the user with token sub
	var user models.User
	initializers.DB.First(&user, claims.Subject)

	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Attach to request
	c.Set("user", user)
	c.Set("token", token)

	// Continue
	c.Next()
}
