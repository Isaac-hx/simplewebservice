package config

import (
	"database/sql"
)

func Connect() (*sql.DB, error) {
	//String koneksi menggunakan database postgres
	db, err := sql.Open("postgres", "host=localhost port=5432 user=isaachx password=saydimas78 dbname=database-golang sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
