package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string
	Username string `gorm:"uniqueIndex"`
	Password string
}
