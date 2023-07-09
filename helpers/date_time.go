package helpers

import (
	"absen/models"
	"errors"
	"time"
)

func FormatDate(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, errors.New("Tanggal kosong")
	}
	var formatedDate time.Time

	layoutFormat := "2006-01-02"
	formatedDate, err := time.Parse(layoutFormat, date)
	if err != nil {
		return formatedDate, err
	}

	return formatedDate, nil
}

func StringToHour(s string) (models.ChekInHours, error) {
	t, err := time.Parse("15:04:05", s)
	if err != nil {
		return models.ChekInHours{}, err
	}
	return models.ChekInHours{Time: t}, nil
}
