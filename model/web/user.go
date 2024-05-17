package web

type UserUpdateServiceRequest struct {
	Name  string `json:"name"`
	Email string `validate:"email" json:"email"`
}

type UserServiceRequest struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UserLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}
