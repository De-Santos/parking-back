package dto

type LoginBody struct {
	Username string
	Password string
}

type SignupBody struct {
	FullName string
	Username string
	Password string
}
