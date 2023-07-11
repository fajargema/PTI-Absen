package admin

import (
	"absen/models"
)

type AdminUserRepository interface {
	GetAll(token string) ([]models.User, error)
	GetByID(id, token string) (models.User, error)
	Create(userInput models.UserInput, token string) (models.User, error)
	Update(userInput models.UserInput, token, id string) (models.User, error)
	Delete(id, token string) error
}
