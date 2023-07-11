package services

import (
	"absen/models"
	"absen/repositories"
	"mime/multipart"
	"time"
)

type PresentService struct {
	repository repositories.PresentRepository
}

func InitPresentService() PresentService {
	return PresentService{
		repository: &repositories.PresentRepositoryImpl{},
	}
}

func (ps *PresentService) GetHomeWidget(token string) (models.HomeWidget, error) {
	return ps.repository.GetHomeWidget(token)
}

func (ps *PresentService) GetAll(token, period string) ([]models.Present, error) {
	return ps.repository.GetAll(token, period)
}

func (ps *PresentService) GetByID(id, token string) (models.Present, error) {
	return ps.repository.GetByID(id, token)
}

func (ps *PresentService) Search(date time.Time, token string) ([]models.Present, error) {
	return ps.repository.Search(date, token)
}

func (ps *PresentService) Create(financeInput models.PresentInput, token string, files []*multipart.FileHeader) (models.Present, error) {
	return ps.repository.Create(financeInput, token, files)
}
