package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load("/home/andrii/go/src/github.com/khoroshun/juliette/.env") //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")


	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string

	dbUri  = "postgres://rbhpkwwiiyxnqv:473d42b05449fbe855d4c47054635789df479e2feb4aafe9bb736e1fb441a118@ec2-54-225-129-101.compute-1.amazonaws.com:5432/dcbpq4tkr0pn6o"

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(
							&Account{},
							&BonusAccount{},
							&BonusTransaction{},
							&Order{}, &Client{},
							&DiscountAccount{},
							&DiscountChanges{},
							&DiscountDiapason{},
						   ) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}