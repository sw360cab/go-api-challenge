package db

import (
	"c3lx/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_STR")
	if Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatal(fmt.Sprintf("Database is unreachable: %v\n", err))
	}
	// Migrate the schema
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Challenge{})

	// Delete permanently
	// cmp. https://gorm.io/docs/delete.html#Delete-permanently
	Db.Unscoped().Where("1 = 1").Delete(&models.User{})
	Db.Unscoped().Where("1 = 1").Delete(&models.Challenge{})

	var users = []models.User{{Username: "Bob"}, {Username: "Alice"}}
	Db.Create(&users)

	var challenges = []models.Challenge{{
		Name:        "Get to Steppin'",
		Description: "Get at least 10,000 steps a day for a month",
	}, {
		Name:        "You are what you eat",
		Description: "Eat two healthy meals a day for two weeks",
	}, {
		Name:        "Rip Van Winkle",
		Description: "Average seven hours of sleep for five nights in a row",
	}}
	Db.Create(challenges)
	// unavailable challenge
	Db.Model(models.Challenge{}).Create(
		map[string]interface{}{
			"Name":        "Not available challenge",
			"Description": "A challenge that is not available",
			"Available":   false})
}
