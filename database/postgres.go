// This file instance postgres and implementation depedency Database
package database

import (
	"database/sql"
	"fmt"
	"log"
	"simplewebservice/config"
	"sync"
)

type postgresDatabase struct {
	Db *sql.DB
}

// Initializing variable
var (
	once       sync.Once
	dbInstance *postgresDatabase
)

// method get dbInstance
func (p *postgresDatabase) GetDb() *sql.DB {
	return dbInstance.Db
}

func (p *postgresDatabase) TestPing() {
	if err := dbInstance.Db.Ping(); err != nil {
		log.Fatal("Test connection failed %v", err.Error())

	}
	log.Println("Sucess connected to database!")
}

// constructor for struct PostgresDatabase
// return type interface database
func NewDatabasePostgres(conf *config.Config) Database {
	once.Do(func() {

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			conf.Db.Host,
			conf.Db.User,
			conf.Db.Password,
			conf.Db.DBName,
			conf.Db.Port,
			conf.Db.SSLMode,
			conf.Db.TimeZone)
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			panic("Failed to connect database!")
		}

		dbInstance = &postgresDatabase{Db: db}
	})
	return dbInstance
}
