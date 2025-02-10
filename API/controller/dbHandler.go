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
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		_ = godotenv.Load()
	}

	db_user := os.Getenv("DB_USER")
	db_name := os.Getenv("DB_NAME")
	db_password := os.Getenv("PASSWORD")
	db_host := os.Getenv("HOST")
	db_port := os.Getenv("DB_PORT")

	// Default PostgreSQL port if not set
	if db_port == "" {
		db_port = "5432"
	}

	// Database connection string
	dbaddress := fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable password=%s host=%s port=%s",
		db_user, db_name, db_password, db_host, db_port,
	)

	// Connect to PostgreSQL
	db, err := sqlx.Connect("postgres", dbaddress)
	if err != nil {
		log.Println("Database connection failed:", err)
		return nil
	}

	log.Println("Connected to PostgreSQL database!")
	return db
}
