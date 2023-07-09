package middleware

import (
	"absen/config"
	"absen/models"
	"absen/utils"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId uint, name string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"name":    name,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(utils.GetConfig("JWT_SECRET_KEY")))
}

func VerifyToken(tokenString string) (models.User, error) {
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}
	var user models.User

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetConfig("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return user, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.ID = uint(claims["user_id"].(float64))
	} else {
		return user, errors.New("Invalid token")
	}

	return user, nil
}

func IsAllowedRole(tokenString string, allowedRoles []string) (bool, error) {
	user, err := VerifyToken(tokenString)
	if err != nil {
		return false, err
	}

	var User models.User
	if err := config.DB.Model(&models.User{}).Where("id = ?", user.ID).First(&User).Error; err != nil {
		return false, err
	}

	for _, role := range allowedRoles {
		if User.Role == role {
			return true, nil
		}
	}

	return false, nil
}
