package db

import (
	"c3lx/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	//err = godotenv.Load()
	dsn := os.Getenv("DB_STR")
	// env variable may be provided even if .env file is not present. e.g. GitHub Actions
	if err != nil && dsn == "" {
		log.Fatal("Error loading .env file")
	}
	if Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatal(fmt.Sprintf("Database is unreachable: %v\n", err))
	}
	// Migrate the schema
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Challenge{})
}
