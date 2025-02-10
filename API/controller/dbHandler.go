package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connect() *sqlx.DB {
	// db_port := os.Getenv("DB_PORT")
	_ = godotenv.Load()

	db_user := os.Getenv("DB_USER")
	db_name := os.Getenv("DB_NAME")
	db_password := os.Getenv("PASSWORD")
	db_host := os.Getenv("HOST")
	dbaddress := fmt.Sprint("user=", db_user, " dbname=", db_name, " sslmode=disable password=", db_password, " host=", db_host, "")
	db, err := sqlx.Connect("postgres", dbaddress)
	if err != nil {
		log.Fatalln(err)
	}
	return db

}
