package library

import "simplewebservice/config"

var Postgres config.DatabaseConfig

func init() {
	Postgres = &config.PostgresDB{
		Host:     "localhost",
		Port:     "5432",
		Username: "isaachx",
		Password: "saydimas78",
		DbName:   "database-golang",
		SslMode:  "disable"}
}
