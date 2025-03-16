// This package is for load config file data from .env file
package config

import (
	"log"
	"os"
	"simplewebservice/utils"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

// Instance for object server
type Server struct {
	Port int
}

// Instance for object database
type Db struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}
type Pool struct {
	MaxOpenConns         int
	MaxIdleConns         int
	MaxOpenConnsLifetime int
	MaxIdleConnsLifetime int
}

// Instance for object Config
type Config struct {
	Server *Server
	Db     *Db
	Pool   *Pool
}

var (
	once           sync.Once
	configInstance *Config
)

// This function is created to build a construction
func GetConfig() *Config {
	//This method only run once
	once.Do(func() {
		//to test this function you must insert argument `../.env` in godotenv.Load
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatalf("Error from read .env file !,%v", err.Error())
		}
		//database config
		host := os.Getenv("DB_HOST")
		dbPortString := os.Getenv("DB_PORT")
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_DBNAME")
		sslmode := os.Getenv("DB_SSLMODE")
		timezone := os.Getenv("DB_TIMEZONE")
		dbPort, err := strconv.Atoi(dbPortString)
		if err != nil {
			log.Fatalf("Error from parsing data .env!,%v", err.Error())
		}

		//server config
		serverPortString := os.Getenv("SERVER_PORT")
		serverPort, err := strconv.Atoi(serverPortString)
		if err != nil {
			log.Fatalf("Error from parsing data .env!,%v", err.Error())

		}

		//pool config
		maxOpenCons := os.Getenv("MAX_OPEN_CONS")
		maxIdleCons := os.Getenv("MAX_IDLE_CONS")
		maxOpenLifetime := os.Getenv("MAX_OPENS_LIFETIME")
		maxIdleLifetime := os.Getenv("MAX_IDLE_LIFETIME")
		arrPool, err := utils.ConvertInt(maxOpenCons, maxIdleCons, maxOpenLifetime, maxIdleLifetime)
		if err != nil {
			log.Fatalf("Error from parsing data .env!,%v", err.Error())
		}

		configInstance = &Config{
			Db:     &Db{Host: host, User: username, Password: password, DBName: dbname, Port: dbPort, SSLMode: sslmode, TimeZone: timezone},
			Server: &Server{Port: serverPort},
			Pool:   &Pool{MaxOpenConns: arrPool[0], MaxIdleConns: arrPool[1], MaxOpenConnsLifetime: arrPool[2], MaxIdleConnsLifetime: arrPool[3]},
		}

	})
	return configInstance
}
