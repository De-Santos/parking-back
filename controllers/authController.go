package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"parking-back/dtos"
	"parking-back/initializers"
	jwt2 "parking-back/jwt"
	"parking-back/models"
	"parking-back/utils"
	"strconv"
	"time"
)

func Signup(c *gin.Context) {
	// Get the username and password off request body
	var body dtos.SignupDto
	if c.Bind(&body) != nil {
		utils.ProcessBadResponse(c, "Failed to read body")
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		utils.ProcessBadResponse(c, "Failed to hash password")
		return
	}

	// Create the user
	user := models.User{FullName: body.FullName, Username: body.Username, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		utils.ProcessBadResponse(c, "Failed to create user")
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	// Get the username and password off request body
	var body dtos.LoginDto

	if c.Bind(&body) != nil {
		utils.ProcessBadResponse(c, "Failed to read body")
		return
	}

	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		utils.ProcessBadResponse(c, "Invalid username or password")
		return
	}

	// Compare sent in password with saved user password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		utils.ProcessBadResponse(c, "Invalid username or password")
		return
	}

	// Generate a jwt token
	jwtToken, err := jwt2.BuildJwt(strconv.Itoa(int(user.ID)))
	if err != nil {
		utils.ProcessBadResponse(c, "Failed to create token")
		return
	}

	// Sent it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", jwtToken, int(time.Now().Add(time.Hour*24).Unix()), "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Logout(c *gin.Context) {
	token, _ := c.Cookie("Authorization")
	claims, _ := jwt2.ParseJwtClaims(token)
	initializers.DB.Create(&models.InvalidatedToken{ID: utils.GetUint(claims.ID), Token: token})
	c.SetCookie("Authorization", "", -1, "", "", false, true)
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in",
	})
}
