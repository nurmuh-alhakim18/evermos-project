package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

var Env map[string]string

func LoadConfig() {
	var err error
	Env, err = godotenv.Read(".env")
	if err != nil {
		log.Fatalf("failed to read env file: %v", err)
	}
}

func GetEnv(key, value string) string {
	val, ok := Env[key]
	if !ok {
		val = value
	}

	return val
}
