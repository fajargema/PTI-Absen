package repositories

import (
	"absen/models"
	"mime/multipart"
	"time"
)

type UserRepository interface {
	Register(UserInput models.UserInput) (models.User, error)
	GetByUsername(username string) (models.User, error)
	Login(UserInput models.UserAuth) (models.UserResponse, error)
	Update(UserInput models.UserInput, token string) (models.User, error)
}

type PresentRepository interface {
	GetAll(token string) ([]models.Present, error)
	GetByID(id, token string) (models.Present, error)
	GetHomeWidget(token string) (models.HomeWidget, error)
	Search(date time.Time, token string) ([]models.Present, error)
	Create(presentInput models.PresentInput, token string, files []*multipart.FileHeader) (models.Present, error)
}
