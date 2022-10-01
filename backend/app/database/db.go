package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test db error %v", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(d.Dsn), &gorm.Config{})
	return db, err
}
