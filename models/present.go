package models

import (
	"database/sql/driver"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Present struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	Time      ChekInHours
	Date      *time.Time     `json:"date" form:"date" gorm:"type:date"`
	URL       string         `json:"url" form:"url"`
	Longitude string         `json:"longitude" form:"longitude"`
	Latitude  string         `json:"latitude" form:"latitude"`
	Location  string         `json:"location" form:"location"`
	Distance  string         `json:"distance" form:"distance"`
	Status    int8           `json:"status" form:"status"`
	UserID    uint           `json:"user_id" form:"user_id"`
	User      User           `gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type ChekInHours struct {
	time.Time
}

func (t *ChekInHours) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		timeValue, err := time.Parse("15:04:05", string(v))
		if err != nil {
			return err
		}
		t.Time = timeValue
		return nil
	case string:
		timeValue, err := time.Parse("15:04:05", v)
		if err != nil {
			return err
		}
		t.Time = timeValue
		return nil
	default:
		return errors.New("type not supported")
	}

}

func (t ChekInHours) Value() (driver.Value, error) {
	return t.Format("15:04:05"), nil
}

func (t ChekInHours) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.Format("15:04:05") + "\""), nil
}

type PresentInput struct {
	Date      time.Time `json:"-"`
	DateInput string    `gorm:"-" json:"date" form:"date"`
	Time      string    `json:"time" form:"time" validate:"required"`
	URL       string    `json:"url" form:"url"`
	Longitude string    `json:"longitude" form:"longitude" validate:"required"`
	Latitude  string    `json:"latitude" form:"latitude" validate:"required"`
	Location  string    `json:"location" form:"location"`
	Status    int8      `json:"status" form:"status"`
	UserID    uint      `json:"user_id" form:"user_id"`
}
