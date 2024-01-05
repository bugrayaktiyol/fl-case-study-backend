// utils.go
package main

import (
	"os"

	"github.com/bugrayaktiyol/fl-case-study-backend/api/users"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupEnvironment() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_PATH")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrateModels(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
}
