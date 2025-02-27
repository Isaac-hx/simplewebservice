// This file contain depedency of instance or object book
package repositories

import (
	"simplewebservice/internal/book"
)

type BookRepository interface {
	InsertBookSQL(in *book.InsertBookDto) error
	GetBookSQL(id int) (*book.GetBookDto, error)
	GetListBookSQL(param string) (*[]book.GetBookDto, error)
	DeleteBookSQL(id int) error
	UpdateBookSQL(id int, in *book.InsertBookDto) error
}
