package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbURL := os.Getenv("db_connection")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&WeatherEntity{})
	if err != nil {
		panic(err)
	}

	return db
}

type WeatherEntity struct {
	ID        int64     `gorm:"primary_key;auto_increment"`
	CreatedAt time.Time `json:"createdat"`
	Query     string    `json:"query"`
	Data      string    `json:"data"`
}
