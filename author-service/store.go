package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	host = "localhost:5432"
	//user     = "author_user"
	user     = "postgres"
	password = ""
	database = "publication_manager"
)

func DBConnection() (*gorm.DB, error) {

	// host is in the form of: localhost:5432
	if os.Getenv("DB_HOST") != "" {
		host = os.Getenv("DB_HOST")
	}
	if os.Getenv("DB_NAME") != "" {
		database = os.Getenv("DB_NAME")
	}
	if os.Getenv("DB_USER") != "" {
		user = os.Getenv("DB_USER")
	}
	if os.Getenv("DB_PASS") != "" {
		password = os.Getenv("DB_PASS")
	}

	log.Printf("Connecting to database using: %s - %s - %s - %s\n", host, database, user, password)

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable",
			user, password, host, database,
		),
	)
}
