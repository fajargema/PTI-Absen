package services

import (
	"absen/models"
	"absen/repositories"
)

type UserService struct {
	repository repositories.UserRepository
}

func InitUserService() UserService {
	return UserService{
		repository: &repositories.UserRepositoryImpl{},
	}
}

func (us *UserService) GetByUsername(username string) (models.User, error) {
	return us.repository.GetByUsername(username)
}

func (us *UserService) Register(userInput models.UserInput) (models.User, error) {
	return us.repository.Register(userInput)
}

func (us *UserService) Login(userInput models.UserAuth) (models.UserResponse, error) {
	return us.repository.Login(userInput)
}

func (us *UserService) Update(userInput models.UserInput, token string) (models.User, error) {
	return us.repository.Update(userInput, token)
}

func (us *UserService) ChangePassword(userInput models.UserChangePassword, token string) (models.User, error) {
	return us.repository.ChangePassword(userInput, token)
}
