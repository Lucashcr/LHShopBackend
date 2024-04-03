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
	if db == nil {
		log.Println("[INFO]: Creating database connection...")

		var err error
		db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal(err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("[INFO]: Database connection created!")
	}

	return db
}
