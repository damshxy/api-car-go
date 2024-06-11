package database

import (
	"fmt"
	"os"

	"github.com/damshxy/api-car-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgUser := os.Getenv("PG_USERNAME")
	pgPass := os.Getenv("PG_PASSWORD")
	pgDB := os.Getenv("PG_DB")

	dsn := fmt.Sprintf( "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", pgHost, pgUser, pgPass, pgDB, pgPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err = DB.AutoMigrate(
		&models.User{},
		&models.Car{},
	); err != nil {
		panic("failed to migrate database")
	}

	fmt.Println("Connection Opened to Database")
}