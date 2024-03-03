package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"parking-back/initializers"
	"parking-back/models"
)

func CheckUsernameExistence(username string) bool {
	var user models.User
	err := initializers.DB.
		Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)}).
		Where(&models.User{Username: username}).
		First(&user).
		Error
	if err != nil {
		return false
	}
	return user.ID != 0
}
