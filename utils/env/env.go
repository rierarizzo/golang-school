package env

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const (
	LOCAL      = "local"
	DEVELOP    = "develop"
	STAGING    = "staging"
	PRODUCTION = "production"
)

func MustGet(key string) string {
	val := os.Getenv(key)
	if val == "" && key != "PORT" {
		logrus.Error("Error getting env key")
		panic("Env key missing " + key)
	}
	return val
}

func CheckDotEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil && os.Getenv("ENVIRONMENT") == LOCAL {
		logrus.Error("Error loading .env file")
	}
}
