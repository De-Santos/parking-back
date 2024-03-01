package models

import (
	"gorm.io/gorm"
	"time"
)

type Parking struct {
	gorm.Model
	Owner       string
	Address     string
	Capacity    int
	Coordinates Coordinates `gorm:"embedded"`
	CreatedByID uint
	CreatedBy   User `gorm:"foreignKey:CreatedByID"`
}

type Coordinates struct {
	Lat float32
	Lng float32
}

type Car struct {
	gorm.Model
	Vpr        string
	Arrived    time.Time
	Expiration time.Time
	ParkingId  uint
	Parking    Parking `gorm:"foreignKey:ParkingId"`
}
