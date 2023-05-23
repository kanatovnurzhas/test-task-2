package database

import "database/sql"

func ConnectionDB() (*sql.DB, error) {
	connectionStr := "user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
