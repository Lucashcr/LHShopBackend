package dbconn

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// GetDB returns the database connection
func GetDB() *sql.DB {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("[ERROR]: DATABASE_URL not set!")
		os.Exit(1)
	}

	var err error
	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("[ERROR]: Error opening database connection!")
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		log.Fatal("[ERROR]: Error pinging database!")
		os.Exit(1)
	}

	return db
}

func CloseDB() {
	db.Close()
	log.Println("[INFO]: Database connection closed!")
}
