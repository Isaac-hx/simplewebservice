package author

import (
	"database/sql"
	"errors"
)

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// call constructor object author
func New() *Author {
	authorObject := &Author{}
	return authorObject
}

// Method yang digunakan untuk mengambil semua data dalam table authors
func (a *Author) GetAuthor(db *sql.DB) ([]Author, error) {
	var listAuthors []Author
	query := `SELECT * FROM authors`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var eachAuthor Author
		err = rows.Scan(&eachAuthor.Id, &eachAuthor.Name)
		if err != nil {
			return nil, err
		}
		listAuthors = append(listAuthors, eachAuthor)
	}

	return listAuthors, nil

}

func (a *Author) GetAuthorById(db *sql.DB, id int) (Author, error) {
	var author Author
	query := `SELECT * FROM authors WHERE author_id=$1`
	err := db.QueryRow(query, id).Scan(&author.Id, &author.Name)
	if err != nil {
		return author, err
	}

	return author, nil

}

func (a *Author) CreateAuthor(db *sql.DB, name string) error {
	query := `INSERT INTO authors(name) values($1)`
	_, err := db.Exec(query, name)
	if err != nil {
		return err
	}

	return nil
}

func (a *Author) DeleteAuthorById(db *sql.DB, id int) error {
	query := `DELETE FROM authors WHERE author_id=$1`
	row, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	rowDelete, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowDelete != 1 {
		return errors.New("0")
	}

	return nil
}

func (a *Author) UpdateAuthorById(db *sql.DB, id int, name string) error {
	query := `UPDATE authors SET name=$1 WHERE author_id=$2`
	row, err := db.Exec(query, name, id)
	if err != nil {
		return err
	}
	rowDelete, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowDelete != 1 {
		return errors.New("0")
	}
	return nil
}
