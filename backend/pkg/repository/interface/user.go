package interfaces

import (
	"github.com/chungvan2301/shoeshop/backend/pkg/ultis/models"
)

type UserRepository interface {
	GetUserDetail(userID string) (models.UserDetail, error)
	RegisterUser(user models.UserInput) error
	EditUser(user models.UserUpdate) error
	GetUserPassword(userID string) (string, error)
	DeleteUser(userID string) error
	GetUserByEmail(email string) (models.UserLoginResponse, error)
}
