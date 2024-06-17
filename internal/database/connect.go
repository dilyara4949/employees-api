package database

import (
	"database/sql"
	"fmt"
	"github.com/dilyara4949/employees-api/internal/config"
	_ "github.com/lib/pq"
)

func ConnectPostgres(cfg config.DB) (*sql.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	fmt.Println(url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
