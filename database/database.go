// This file to depedency structure for switch database or mocking database
package database

import "database/sql"

type Database interface {
	GetDb() *sql.DB
	TestPing()
	GetStat()
}
