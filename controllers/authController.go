package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"parking-back/initializers"
	jwt2 "parking-back/jwt"
	"parking-back/models"
	"parking-back/obj"
	"parking-back/utils"
	"strconv"
	"time"
)

func Signup(c *gin.Context) {
	// Get the username and password off request body
	var body obj.SignupDto
	if c.Bind(&body) != nil {
		utils.ProcessBadResponse(c, "Failed to read body")
		return
	}

	err := initializers.V.Struct(body)
	if err != nil {
		utils.ProcessBadResponse(c, "Invalid request body: "+fmt.Sprint(err))
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
	var body obj.LoginDto

	if c.Bind(&body) != nil {
		utils.ProcessBadResponse(c, "Failed to read body")
		return
	}

	err := initializers.V.Struct(body)
	if err != nil {
		utils.ProcessBadResponse(c, "Invalid request body: "+fmt.Sprint(err))
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
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

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
	c.Header("Authorization", jwtToken)
	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

func Logout(c *gin.Context) {
	token, _ := c.Get("token")
	claims, _ := jwt2.ParseJwtClaims(token.(string))
	initializers.DB.Create(&models.InvalidatedToken{ID: utils.GetUint(claims.ID), Token: token.(string)})
	c.SetCookie("Authorization", "", -1, "", "", false, true)
}
