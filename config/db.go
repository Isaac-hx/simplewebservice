package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Abstraksi interface pada setiap object database
type DatabaseConfig interface {
	stringConnection() (string, string)
}

// object database postgres
type PostgresDB struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SslMode  string
}

func (p *PostgresDB) stringConnection() (string, string) {
	return "postgres", fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=%s`, p.Username, p.Password, p.Host, p.Port, p.DbName, p.SslMode)
}

// object database mysql
type MysqlDB struct {
	host     string
	port     string
	username string
	password string
	dbname   string
}

func (m *MysqlDB) stringConnection() (string, string) {
	return "mysql", fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v`, m.username, m.password, m.host, m.port, m.dbname)
}

func Connect(DC DatabaseConfig) (*sql.DB, error) {
	//String koneksi menggunakan database postgres
	db, err := sql.Open(DC.stringConnection())
	if err != nil {
		return nil, err
	}
	return db, nil
}
