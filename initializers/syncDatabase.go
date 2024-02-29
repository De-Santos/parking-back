package initializers

import (
	"parking-back/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.InvalidatedToken{},
		&models.Parking{},
		&models.Car{})
	if err != nil {
		panic("Database syncing failed")
	}
}
