package model

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func openDatabase() (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN: os.Getenv("DATABASE_URL"),
	}), &gorm.Config{})
}
