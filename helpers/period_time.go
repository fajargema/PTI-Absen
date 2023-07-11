package helpers

import (
	"errors"
	"time"
)

func GetPeriodTime(period string) (time.Time, error) {
	now := time.Now()
	weekday := int(now.Weekday())

	if period == "today" {
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()), nil
	} else if period == "weekly" {
		weekStart := now.AddDate(0, 0, -weekday+1)
		return time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, now.Location()), nil
	} else if period == "monthly" {
		monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		return monthStart, nil
	}

	return time.Time{}, errors.New("Invalid period")
}
