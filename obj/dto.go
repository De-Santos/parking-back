package obj

import (
	"parking-back/models"
	"time"
)

type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignupDto struct {
	FullName string `json:"full_name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ParkingDto struct {
	ID          uint               `json:"id"`
	Owner       string             `json:"owner" validate:"required"`
	Address     string             `json:"address" validate:"required"`
	Capacity    int                `json:"capacity" validate:"required,gt=0"`
	Coordinates models.Coordinates `json:"coordinates"`
	CreatedBy   string             `json:"created_by"`
	CreatedAt   time.Time          `json:"created_at"`
}

type PageableDtoWrapper struct {
	Limit int           `json:"limit"`
	Page  int           `json:"page"`
	Body  []interface{} `json:"body"`
}

func (pdw PageableDtoWrapper) New(pageable Pageable, body []interface{}) PageableDtoWrapper {
	if body == nil {
		body = []interface{}{}
	}
	return PageableDtoWrapper{Limit: pageable.GetLimit(), Page: pageable.GetPage(), Body: body}
}
