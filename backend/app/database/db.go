package database

import (
	domain "cupcake/app/models"
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
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

func NewTest() (*Database, error) {
	db, err := gorm.Open(sqlite.Open("gorm"), &gorm.Config{})

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	err = db.AutoMigrate(&domain.Bet{}, &domain.Bracket{}, &domain.Match{},
		&domain.NationalTeam{}, &domain.User{}, &domain.UserPoints{},
		&domain.Joker{}, &domain.Groups{})

	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

func New(config *DatabaseConfig) (*Database, error) {
	var db *gorm.DB
	var err error
	dsn := "user=" + config.Username + " password=" + config.Password + " dbname=" + config.Database + " host=" + config.Host + " port=" + strconv.Itoa(config.Port) + " TimeZone=UTC"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	err = db.AutoMigrate(&domain.Bet{}, &domain.Bracket{}, &domain.Match{},
		&domain.NationalTeam{}, &domain.User{}, &domain.UserPoints{},
		&domain.Joker{}, &domain.Groups{})

	if err != nil {
		return nil, err
	}

	return &Database{db}, err
}
