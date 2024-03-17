package database

import (
	"fmt"
	"log"
	"os"
	"sesi11/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	host     = "localhost"
// 	port     = "5432"
// 	user     = "postgres"
// 	password = "1234"
// 	dbname   = "postgres"
// )

const (
	host     = os.Getenv("PGHOST")
	port     = os.Getenv("PGPORT")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbname   = os.Getenv("PGDATABASE")
)

var (
	db  *gorm.DB
	err error
)

func GetDB() *gorm.DB {
	return db
}

func StartDB() {
	psqInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(psqInfo), &gorm.Config{})
	fmt.Println(psqInfo)
	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}
	db.Debug().AutoMigrate(model.User{}, model.Product{})
}
