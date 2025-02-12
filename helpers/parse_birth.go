package helpers

import (
	"time"
)

func ParseBirthDate(birthDate string) (time.Time, error) {
	layout := "02/01/2006"
	date, err := time.Parse(layout, birthDate)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}
