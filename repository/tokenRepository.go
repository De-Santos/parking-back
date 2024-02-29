package repository

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"parking-back/initializers"
	"parking-back/models"
)

func IsTokenInvalidated(id uint64) (bool, error) {
	var expToken models.InvalidatedToken
	db := initializers.DB.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})
	if err := db.First(&expToken, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
