package database

import (
	"fmt"

	"github.com/kenethrrizzo/golang-school/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(config *config.DatabaseConfig) (*gorm.DB, error) {
	// dsn: data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Name)
	// using mysql driver from gorm
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		logrus.Error("Error while connecting to database")
		return nil, err
	}
	logrus.Info("Successful database connection")
	return db, err
}
