package db

import (
	"log"
	"main/internal/config"
	"main/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgreSQL(cfg config.DBConfig) (*gorm.DB, error) {
	dsn := cfg.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		log.Fatalf("PostgreSQL connection failed : %v", err)
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.User{},
	)

	if err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}
}