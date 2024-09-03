package database

import (
	"fmt"
	"log"

	"github.com/damshxy/api-car-go/config"
	"github.com/damshxy/api-car-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.PGHOST,
		c.PGPORT,
		c.PGUSER,
		c.PGPASSWORD,
		c.PGDATABASE,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Car{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	DB = db
	log.Println("Successfully connected to database")

	return DB, nil
}