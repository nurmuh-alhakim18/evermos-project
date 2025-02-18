package helpers

import (
	"crypto/rand"
	"fmt"
)

func GenerateShortID() (string, error) {
	numBytes := 3
	b := make([]byte, numBytes)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", b), nil
}
