package admin

import (
	"absen/models"
	repo "absen/repositories/admin"
)

type AdminUserService struct {
	repository repo.AdminUserRepository
}

func InitAdminUserService() AdminUserService {
	return AdminUserService{
		repository: &repo.AdminUserRepositoryImpl{},
	}
}

func (aus *AdminUserService) GetAll(token string) ([]models.User, error) {
	return aus.repository.GetAll(token)
}

func (aus *AdminUserService) GetByID(id, token string) (models.User, error) {
	return aus.repository.GetByID(id, token)
}

func (aus *AdminUserService) Create(userInput models.UserInput, token string) (models.User, error) {
	return aus.repository.Create(userInput, token)
}

func (aus *AdminUserService) Update(userInput models.UserInput, token, id string) (models.User, error) {
	return aus.repository.Update(userInput, token, id)
}

func (aus *AdminUserService) Delete(id, token string) error {
	return aus.repository.Delete(id, token)
}
