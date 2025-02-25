// This file contain depedency of instance or object book
package repositories

import (
	"database/sql"
	"simplewebservice/internal/book"
)

type BookRepository interface {
	InsertBookSQL(in *book.InsertBookDto) error
	GetBookSQL(id int) *sql.Row
}
