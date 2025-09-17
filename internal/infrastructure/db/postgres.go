package db

import (
	"fintrack-backend/internal/infrastructure/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(config config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DBUrl()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	return db

}
