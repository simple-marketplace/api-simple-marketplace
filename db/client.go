package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseClient(dsn string) *gorm.DB {
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return DB
}
