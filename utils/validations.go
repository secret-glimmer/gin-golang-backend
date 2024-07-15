package utils

import "time"

func ValidateDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}
