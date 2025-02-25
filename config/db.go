package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Abstraksi interface pada setiap object database
type DatabaseConfig interface {
	ConnectDB() (*sql.DB, error)
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

func (p *PostgresDB) ConnectDB() (*sql.DB, error) {
	strConnection := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=%s`, p.Username, p.Password, p.Host, p.Port, p.DbName, p.SslMode)
	db, err := sql.Open("postgres", strConnection)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// object database mysql
type MysqlDB struct {
	host     string
	port     string
	username string
	password string
	dbname   string
}

func (m *MysqlDB) ConnectDB() (*sql.DB, error) {
	strConnection := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v`, m.username, m.password, m.host, m.port, m.dbname)
	//String koneksi menggunakan database postgres
	db, err := sql.Open("mysql", strConnection)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func NewConnection(conn DatabaseConfig) (*sql.DB, error) {
	return conn.ConnectDB()
}
