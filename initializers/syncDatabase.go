package initializers

import (
	"parking-back/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Database syncing failed")
	}
}
