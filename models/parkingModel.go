package models

import (
	"gorm.io/gorm"
	"time"
)

type Parking struct {
	gorm.Model
	ParkingName string
	Owner       string
	Address     string
	Capacity    int
	Coordinates Coordinates `gorm:"embedded"`
	CreatedByID uint
	CreatedBy   User `gorm:"foreignKey:CreatedByID"`
}

func (p *Parking) GetUpdatedColumns() map[string]interface{} {
	return map[string]interface{}{
		"parking_name": p.ParkingName,
		"owner":        p.Owner,
		"address":      p.Address,
		"capacity":     p.Capacity,
		"lat":          p.Coordinates.Lat,
		"lng":          p.Coordinates.Lng,
	}
}

type Coordinates struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type Car struct {
	gorm.Model
	Vpr        string
	Arrived    time.Time
	Expiration time.Time
	ParkingId  uint
	Parking    Parking `gorm:"foreignKey:ParkingId"`
}

func (c *Car) GetUpdatedColumns() map[string]interface{} {
	return map[string]interface{}{
		"vpr":        c.Vpr,
		"arrived":    c.Arrived,
		"expiration": c.Expiration,
	}
}
