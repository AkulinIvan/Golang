package config

import (
	"fmt"
	"os"

	"github.com/AkulinIvan/Golang/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
//подключение к базе данных
func PostgresConnection() {
	godotenv.Load()
	dbname := os.Getenv("POSTGRES_DBNAME")
	dbuser := os.Getenv("POSTGRES_USER")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbhost := os.Getenv("POSTGRES_HOST")

	connection := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", dbhost, dbuser, dbname, dbpassword)

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic("db connection failed")
	}
	DB = db
	fmt.Println("db connection successfully")
	AutoMigrate(db)
}
//автомиграция при каждом запуске приложения
func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&model.Tasks{},
	)
}
