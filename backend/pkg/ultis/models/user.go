package models

type UserInput struct {
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
	Email     string `form:"email" validate:"required,email"`
	Password  string `form:"password" validate:"required,min=8,max=20"`
}

type UserUpdate struct {
	ID          string `form:"_id"`
	FirstName   string `form:"first_name"`
	LastName    string `form:"last_name"`
	Email       string `form:"email" validate:"required,email"`
	OldPassword string `form:"old_password" validate:"required,min=8,max=20"`
	NewPassword string `form:"new_password" validate:"required,min=8,max=20"`
}

type UserDetail struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserLogin struct {
	Email    string `form:"email,omitempty" validate:"required,email"`
	Password string `form:"password" validate:"required,min=8,max=20"`
}

type UserLoginResponse struct {
	ID       string `json:"_id"`
	Password string `json:"password"`
}
