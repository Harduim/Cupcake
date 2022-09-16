package database

import (
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

type Database struct {
	*gorm.DB
}

func New(config *DatabaseConfig) (*Database, error) {
	var db *gorm.DB
	var err error
	dsn := "user=" + config.Username + " password=" + config.Password + " dbname=" + config.Database + " host=" + config.Host + " port=" + strconv.Itoa(config.Port) + " TimeZone=UTC"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return &Database{db}, err
}
