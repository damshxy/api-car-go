package database

import (
	"fmt"

	"github.com/damshxy/api-car-go/config"
	"github.com/damshxy/api-car-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDatabasePG(c *config.Config) {
	var err error

	dsn := fmt.Sprintf( "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.PGHost,
		c.PGUser,
		c.PGPass,
		c.PGDBName,
		c.PGPort,
	)

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