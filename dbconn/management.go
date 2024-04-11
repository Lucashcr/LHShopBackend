package dbconn

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// GetDB returns the database connection
func GetDB() *sql.DB {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("[ERROR]: Error opening database connection!")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("[ERROR]: Error pinging database!")
	}

	return db
}

func CloseDB() {
	db.Close()
	log.Println("[INFO]: Database connection closed!")
}
