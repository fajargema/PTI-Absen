package admin

import (
	"absen/models"
	"time"
)

type AdminUserRepository interface {
	GetAll(token string) ([]models.User, error)
	GetByID(id, token string) (models.User, error)
	Create(userInput models.UserInput, token string) (models.User, error)
	Update(userInput models.UserInput, token, id string) (models.User, error)
	ChangePassword(UserInput models.UserChangePassword, token string) (models.User, error)
	Delete(id, token string) error
}

type AdminPresentRepository interface {
	GetAll(token, period string) ([]models.Present, error)
	GetByID(id, token string) (models.Present, error)
	Search(date time.Time, token string) ([]models.Present, error)
}
