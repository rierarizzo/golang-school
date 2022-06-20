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

func NewConfig() *Config {
	env.CheckDotEnv()
	port := env.MustGet("PORT")
	if port == "" {
		port = "8080"
	}

	dBConfig := DatabaseConfig{
		Host:     env.MustGet("DB_HOST"),
		Port:     env.MustGet("DB_PORT"),
		Name:     env.MustGet("DB_NAME"),
		User:     env.MustGet("DB_USER"),
		Password: env.MustGet("DB_PASSWORD"),
	}

	logrus.Info("Successful configuration loading")
	return &Config{
		Environment: env.MustGet("ENVIRONMENT"),
		Port:        port,
		Database:    &dBConfig,
	}
}
