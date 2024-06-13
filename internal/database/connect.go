package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "postgres"
	dbUrl    = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
)

func ConnectPostgres() (*sql.DB, error) {
	url := fmt.Sprintf(dbUrl, host, port, user, password, dbname)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to database")
	return db, nil
}
