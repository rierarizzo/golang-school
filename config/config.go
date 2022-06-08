package config

import (
	"github.com/kenethrrizzo/golang-school/utils/env"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Environment string
	Port        string
	Database    *DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func NewConfig() (*Config) {
	env.CheckDotEnv()
	port := env.MustGet("PORT")
	if port == "" {
		port = "8080"
	}

	databaseConfig := DatabaseConfig{
		Host:     env.MustGet("DATABASE_HOST"),
		Port:     env.MustGet("DATABASE_PORT"),
		User:     env.MustGet("DATABASE_USER"),
		Password: env.MustGet("DATABASE_PASSWORD"),
		Name:     env.MustGet("DATABASE_NAME"),
	}

	logrus.Info("Successful configuration loading")
	return &Config{
		Environment: env.MustGet("ENVIRONMENT"),
		Port:        port,
		Database:    &databaseConfig,
	}
}
