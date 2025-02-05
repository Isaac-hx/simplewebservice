package library

import (
	"os"
	"simplewebservice/config"
	"simplewebservice/utils"
)

var Postgres config.DatabaseConfig

func init() {
	//Load environment variable
	utils.LoadEnv()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DBNAME")
	sslmode := os.Getenv("DB_SSLMODE")
	Postgres = &config.PostgresDB{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DbName:   dbname,
		SslMode:  sslmode}
}
