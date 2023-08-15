package repositories

import (
	"absen/config"
	"absen/helpers"
	m "absen/middleware"
	"absen/models"
	"encoding/json"
	"errors"
	"mime/multipart"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type PresentRepositoryImpl struct{}

func InitPresentRepository() PresentRepository {
	return &PresentRepositoryImpl{}
}

func (pr *PresentRepositoryImpl) GetHomeWidget(token string) (models.HomeWidget, error) {
	homeWidget := models.HomeWidget{}

	user, err := m.VerifyToken(token)
	if err != nil {
		return models.HomeWidget{}, err
	}

	layoutFormat := "2006-01-02"
	now := time.Now()
	nowFormat := now.Format(layoutFormat)

	// Get Check-In Time
	var presentInTime models.Present
	if err := config.DB.Preload("User").
		Select("time").
		Where("date = ? AND user_id = ? AND status = ?", nowFormat, user.ID, 0).
		Order("time DESC").
		First(&presentInTime).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			homeWidget.CheckIn.Time = "belum absen"
		} else {
			return models.HomeWidget{}, err
		}
	} else {
		homeWidget.CheckIn.Time = presentInTime.Time.Format("15:04:05")
	}

	// Get Check-In Latitude
	var presentInLat models.Present
	if err := config.DB.Preload("User").
		Select("latitude").
		Where("date = ? AND user_id = ? AND status = ?", nowFormat, user.ID, 0).
		Order("time DESC").
		First(&presentInLat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			homeWidget.CheckIn.Latitude = "belum absen"
		} else {
			return models.HomeWidget{}, err
		}
	} else {
		homeWidget.CheckIn.Latitude = presentInLat.Latitude
	}

	// Get Check-In Longitude
	var presentInLong models.Present
	if err := config.DB.Preload("User").
		Select("longitude").
		Where("date = ? AND user_id = ? AND status = ?", nowFormat, user.ID, 0).
		Order("time DESC").
		First(&presentInLong).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			homeWidget.CheckIn.Longitude = "belum absen"
		} else {
			return models.HomeWidget{}, err
		}
	} else {
		homeWidget.CheckIn.Longitude = presentInLong.Longitude
	}

	// Get Check-In Location
	var presentInLoc models.Present
	if err := config.DB.Preload("User").
		Select("location").
		Where("date = ? AND user_id = ? AND status = ?", nowFormat, user.ID, 0).
		Order("time DESC").
		First(&presentInLoc).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			homeWidget.CheckIn.Location = "belum absen"
		} else {
			return models.HomeWidget{}, err
		}
	} else {
		homeWidget.CheckIn.Location = presentInLoc.Location
	}

	// Get Check-Out Last
	var presentOutTime models.Present
	if err := config.DB.Preload("User").
		Select("time").
		Where("date = ? AND user_id = ? AND status = ?", nowFormat, user.ID, 1).
		Order("time DESC").
		First(&presentOutTime).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			homeWidget.CheckOut.Time = "belum absen"
		} else {
			return models.HomeWidget{}, err
		}
	} else {
		homeWidget.CheckOut.Time = presentOutTime.Time.Format("15:04:05")
	}

	// Get Check-In Latitude
	var presentOutLat models.Present
	if err := config.DB.Preload("User").
		Select("latitude").
		Where("date = ? AND user_id = ? AND status = ?", nowFormat, user.ID, 1).
		Order("time DESC").
		First(&presentOutLat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			homeWidget.CheckOut.Latitude = "belum absen"
		} else {
			return models.HomeWidget{}, err
		}
	} else {
		homeWidget.CheckOut.Latitude = presentOutLat.Latitude
	}

	// Get Check-In Longitude
	var presentOutLong models.Present
	if err := config.DB.Preload("User").
		Select("longitude").
		Where("date = ? AND user_id = ? AND status = ?", nowFormat, user.ID, 1).
		Order("time DESC").
		First(&presentOutLong).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			homeWidget.CheckOut.Longitude = "belum absen"
		} else {
			return models.HomeWidget{}, err
		}
	} else {
		homeWidget.CheckOut.Longitude = presentOutLong.Longitude
	}

	// Get Check-Out Location
	var presentOutLoc models.Present
	if err := config.DB.Preload("User").
		Select("location").
		Where("date = ? AND user_id = ? AND status = ?", nowFormat, user.ID, 1).
		Order("time DESC").
		First(&presentOutLoc).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			homeWidget.CheckOut.Location = "belum absen"
		} else {
			return models.HomeWidget{}, err
		}
	} else {
		homeWidget.CheckOut.Location = presentOutLoc.Location
	}

	return homeWidget, nil
}

func (pr *PresentRepositoryImpl) GetAll(token, period string) ([]models.Present, error) {
	var presents []models.Present

	user, err := m.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	if period != "" {
		startTime, err := helpers.GetPeriodTime(period)
		if err != nil {
			return nil, err
		}
		date := startTime.Format("2006-01-02")

		if err := config.DB.Where("user_id = ? AND date >= ?", user.ID, date).Preload("User").Find(&presents).Error; err != nil {
			return nil, err
		}
	} else {
		if err := config.DB.Where("user_id = ?", user.ID).Preload("User").Find(&presents).Error; err != nil {
			return nil, err
		}
	}

	return presents, nil
}

func (pr *PresentRepositoryImpl) GetByID(id, token string) (models.Present, error) {
	var present models.Present

	user, err := m.VerifyToken(token)
	if err != nil {
		return models.Present{}, err
	}

	if err := config.DB.Preload("User").First(&present, "id = ? AND user_id = ?", id, user.ID).Error; err != nil {
		return models.Present{}, err
	}

	return present, nil
}

func (pr *PresentRepositoryImpl) Search(date time.Time, token string) ([]models.Present, error) {
	var presents []models.Present

	user, err := m.VerifyToken(token)
	if err != nil {
		return []models.Present{}, err
	}

	if err := config.DB.Where("date = ? AND user_id = ?", date, user.ID).Preload("User").Find(&presents).Error; err != nil {
		return []models.Present{}, err
	}

	return presents, nil
}

func (pr *PresentRepositoryImpl) Create(presentInput models.PresentInput, token string, files []*multipart.FileHeader) (models.Present, error) {
	user, err := m.VerifyToken(token)
	if err != nil {
		return models.Present{}, err
	}

	var User models.User
	if err := config.DB.Where("id = ?", user.ID).First(&User).Error; err != nil {
		return models.Present{}, err
	}

	lat, _ := strconv.ParseFloat(presentInput.Latitude, 64)
	long, _ := strconv.ParseFloat(presentInput.Longitude, 64)
	dis := helpers.HaversineDistance(lat, long)
	distance := dis * 1000
	distanceFormatted := strconv.FormatFloat(distance, 'f', -1, 64)

	formatedDate, err := helpers.FormatDate(presentInput.DateInput)
	if err != nil {
		return models.Present{}, err
	}
	presentInput.Date = formatedDate

	checkInTime, err := helpers.StringToHour(presentInput.Time)
	if err != nil {
		return models.Present{}, err
	}

	var status int8
	if checkInTime.After(time.Date(checkInTime.Year(), checkInTime.Month(), checkInTime.Day(), 17, 0, 0, 0, checkInTime.Location())) && checkInTime.Before(time.Date(checkInTime.Year(), checkInTime.Month(), checkInTime.Day(), 20, 0, 0, 0, checkInTime.Location())) {
		status = 1
	} else if checkInTime.After(time.Date(checkInTime.Year(), checkInTime.Month(), checkInTime.Day(), 8, 0, 0, 0, checkInTime.Location())) && checkInTime.Before(time.Date(checkInTime.Year(), checkInTime.Month(), checkInTime.Day(), 12, 0, 0, 0, checkInTime.Location())) {
		status = 0
	} else {
		return models.Present{}, errors.New("Melebihi jam absen")
	}

	var count int64
	if err := config.DB.Model(&models.Present{}).Where("date = ? AND status = ? AND user_id = ?", presentInput.DateInput, status, user.ID).Count(&count).Error; err != nil {
		return models.Present{}, err
	}

	if count > 0 {
		if status == 0 {
			return models.Present{}, errors.New("Kamu sudah absen masuk")
		} else {
			return models.Present{}, errors.New("Kamu sudah absen keluar")
		}
	}

	getLocation, err := helpers.GetLocation(presentInput.Latitude, presentInput.Longitude)
	if err != nil {
		return models.Present{}, err
	}
	var displayLoc models.GeocodingResponse
	err = json.Unmarshal([]byte(getLocation), &displayLoc)
	if err != nil {
		return models.Present{}, err
	}

	var urls []string
	svc, err := config.CreateS3Client()
	if err != nil {
		return models.Present{}, err
	}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return models.Present{}, err
		}
		defer src.Close()

		url, err := helpers.UploadToS3(svc, src, file.Filename)
		if err != nil {
			return models.Present{}, err
		}

		urls = append(urls, url)
	}

	jsonURLs, err := json.Marshal(urls)
	if err != nil {
		return models.Present{}, err
	}

	createdPresent := models.Present{
		Date:      &presentInput.Date,
		Time:      checkInTime,
		URL:       string(jsonURLs),
		Longitude: presentInput.Longitude,
		Latitude:  presentInput.Latitude,
		Location:  displayLoc.Results[4].FormattedAddress,
		Distance:  distanceFormatted,
		Status:    status,
		UserID:    user.ID,
		User:      User,
	}

	result := config.DB.Create(&createdPresent)
	if err := result.Error; err != nil {
		return models.Present{}, err
	}

	if err := config.DB.Last(&createdPresent).Error; err != nil {
		return models.Present{}, err
	}

	return createdPresent, nil
}
