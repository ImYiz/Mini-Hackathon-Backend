package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	connStr := "host=db user=postgres password=postgres dbname=kitchen port=5432 sslmode=disable"

	database, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	DB = database
}
