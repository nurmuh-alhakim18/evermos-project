package helpers

import (
	"time"
)

func ParseBirthDate(birthDate string) (string, error) {
	layout := "02/01/2006"
	date, err := time.Parse(layout, birthDate)
	if err != nil {
		return "", err
	}

	return date.Format("2006-01-02"), err
}

func BirthDateToIndoFormat(birthDate string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, birthDate)
	if err != nil {
		return "", err
	}

	return parsedTime.Format("02/01/2006"), err
}
