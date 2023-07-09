package repositories

import (
	"absen/config"
	m "absen/middleware"
	"absen/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct{}

func InitUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (ur *UserRepositoryImpl) Register(userInput models.UserInput) (models.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	_, err = ur.GetByUsername(userInput.Username)

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

func (ur *UserRepositoryImpl) GetByUsername(username string) (models.User, error) {
	var user models.User

	err := config.DB.First(&user, "username = ?", username).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) Login(userInput models.UserAuth) (models.UserResponse, error) {
	var user models.User

	user, err := ur.GetByUsername(userInput.Username)

	if err != nil {
		return models.UserResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

	if err != nil {
		return models.UserResponse{}, err
	}

	token, err := m.CreateToken(user.ID, user.Name)
	if err != nil {
		return models.UserResponse{}, err
	}

	userResponse := models.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Role:     user.Role,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return userResponse, nil
}

func (ur *UserRepositoryImpl) Update(userInput models.UserInput, token string) (models.User, error) {
	user, err := m.VerifyToken(token)
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
