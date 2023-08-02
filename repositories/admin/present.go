package admin

import (
	"absen/config"
	"absen/helpers"
	m "absen/middleware"
	"absen/models"
	"time"
)

type AdminPresentRepositoryImpl struct{}

func InitAdminPresentRepository() AdminPresentRepository {
	return &AdminPresentRepositoryImpl{}
}

func (apr *AdminPresentRepositoryImpl) GetAll(token, period string) ([]models.Present, error) {
	var presents []models.Present

	_, err := m.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	if period != "" {
		startTime, err := helpers.GetPeriodTime(period)
		if err != nil {
			return nil, err
		}
		date := startTime.Format("2006-01-02")

		if err := config.DB.Where("date >= ?", date).Preload("User").Find(&presents).Error; err != nil {
			return nil, err
		}
	} else {
		if err := config.DB.Preload("User").Find(&presents).Error; err != nil {
			return nil, err
		}
	}

	return presents, nil
}

func (apr *AdminPresentRepositoryImpl) GetByID(id, token string) (models.Present, error) {
	var present models.Present

	_, err := m.VerifyToken(token)
	if err != nil {
		return models.Present{}, err
	}

	if err := config.DB.Preload("User").First(&present, "id = ?", id).Error; err != nil {
		return models.Present{}, err
	}

	return present, nil
}

func (apr *AdminPresentRepositoryImpl) Search(date time.Time, token string) ([]models.Present, error) {
	var presents []models.Present

	_, err := m.VerifyToken(token)
	if err != nil {
		return []models.Present{}, err
	}

	if err := config.DB.Where("date = ?", date).Preload("User").Find(&presents).Error; err != nil {
		return []models.Present{}, err
	}

	return presents, nil
}
