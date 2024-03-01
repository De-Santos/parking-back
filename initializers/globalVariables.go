package initializers

import "github.com/go-playground/validator/v10"

var V *validator.Validate

func InitializeVariables() {
	V = validator.New(validator.WithRequiredStructEnabled())
}
