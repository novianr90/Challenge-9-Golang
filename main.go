package main

import (
	"challenge-9/models"
	"challenge-9/routers"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		username = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		dbName   = os.Getenv("DB_NAME")
	)

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbName)

	db, err = gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func main() {
	routers.StartServer(db).Run(":3000")
}
