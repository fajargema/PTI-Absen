package admin

import (
	"absen/models"
	repo "absen/repositories/admin"
	"time"
)

type AdminPresentService struct {
	repository repo.AdminPresentRepository
}

func InitAdminPresentService() AdminPresentService {
	return AdminPresentService{
		repository: &repo.AdminPresentRepositoryImpl{},
	}
}

func (aps *AdminPresentService) GetAll(token, period string) ([]models.Present, error) {
	return aps.repository.GetAll(token, period)
}

func (aps *AdminPresentService) GetByID(id, token string) (models.Present, error) {
	return aps.repository.GetByID(id, token)
}

func (aps *AdminPresentService) Search(date time.Time, token string) ([]models.Present, error) {
	return aps.repository.Search(date, token)
}
