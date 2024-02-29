package models

type InvalidatedToken struct {
	ID    uint   `gorm:"primarykey"`
	Token string `gorm:"unique"`
}
