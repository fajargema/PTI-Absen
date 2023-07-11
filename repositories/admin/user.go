package admin

import (
	"absen/config"
	m "absen/middleware"
	"absen/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AdminUserRepositoryImpl struct{}

func InitAdminUserRepository() AdminUserRepository {
	return &AdminUserRepositoryImpl{}
}

func (aur *AdminUserRepositoryImpl) GetAll(token string) ([]models.User, error) {
	var users []models.User

	_, err := m.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	if err := config.DB.Where("role = ?", "isUser").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (aur *AdminUserRepositoryImpl) GetByID(id, token string) (models.User, error) {
	_, err := m.VerifyToken(token)
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	if err := config.DB.Where("role = ?", "isUser").First(&user, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (aur *AdminUserRepositoryImpl) Create(userInput models.UserInput, token string) (models.User, error) {
	_, err := m.VerifyToken(token)
	if err != nil {
		return models.User{}, err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	_, err = getByUsername(userInput.Username)

	if err == nil {
		return models.User{}, errors.New("Username sudah ada")
	}

	var createdUser models.User = models.User{
		Name:     userInput.Name,
		Role:     "isUser",
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: string(password),
	}

	result := config.DB.Create(&createdUser)

	if err := result.Error; err != nil {
		return models.User{}, err
	}

	if err := config.DB.Last(&createdUser).Error; err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (aur *AdminUserRepositoryImpl) Update(userInput models.UserInput, id, token string) (models.User, error) {
	_, err := m.VerifyToken(token)
	if err != nil {
		return models.User{}, err
	}
	user, err := aur.GetByID(id, token)
	if err != nil {
		return models.User{}, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user.Name = userInput.Name
	user.Username = userInput.Username
	user.Email = userInput.Email
	user.Password = string(password)

	err = config.DB.Save(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (aur *AdminUserRepositoryImpl) Delete(id, token string) error {
	_, err := m.VerifyToken(token)
	if err != nil {
		return err
	}

	user, err := aur.GetByID(id, token)
	if err != nil {
		return err
	}

	err = config.DB.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func getByUsername(username string) (models.User, error) {
	var user models.User

	err := config.DB.First(&user, "username = ?", username).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
