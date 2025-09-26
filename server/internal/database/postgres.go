package database

import (
	"github.com/Habeebamoo/Clivo/server/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	dsn, err := config.Get("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}