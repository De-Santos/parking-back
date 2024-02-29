package initializers

import (
	"parking-back/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{}, &models.InvalidatedToken{})
	if err != nil {
		panic("Database syncing failed")
	}
}
