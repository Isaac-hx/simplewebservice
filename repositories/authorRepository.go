package repositories

import "simplewebservice/internal/author"

type AuthorRepository interface {
	InsertAuthorSQL(in *author.InsertAuthorDto) error
	GetAuthorSQL(id int) (*author.GetBookDto, error)
	DeleteAuthorSQL(id int) error
	GetListAuthorSQL(param string) (*[]author.GetBookDto, error)
	UpdateAuthorSQL(id int, in *author.InsertAuthorDto) error
}
