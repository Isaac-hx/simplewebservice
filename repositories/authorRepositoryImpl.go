package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"simplewebservice/author"
	"simplewebservice/database"
)

type authorPostgresRepository struct {
	db database.Database
}

func NewAuthorPostgresRepository(db database.Database) AuthorRepository {
	return &authorPostgresRepository{db: db}
}

func (a *authorPostgresRepository) InsertAuthorSQL(in *author.InsertAuthorDto) error {
	defer a.db.GetDb().Close()

	query := `INSERT INTO authors(name) VALUES($1) `
	_, err := a.db.GetDb().Exec(query, in.Name)
	if err != nil {
		return err
	}
	return nil
}

func (a *authorPostgresRepository) GetAuthorSQL(id int) (*author.GetBookDto, error) {
	defer a.db.GetDb().Close()

	var author author.GetBookDto
	query := `SELECT id,name FROM authors WHERE id = $1`
	err := a.db.GetDb().QueryRow(query, id).Scan(&author)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("0")
		}
		return nil, err
	}
	return &author, nil
}

func (a *authorPostgresRepository) GetListAuthorSQL(param string) (*[]author.GetBookDto, error) {
	defer a.db.GetDb().Close()

	var authors []author.GetBookDto
	query := fmt.Sprintf(`SELECT * FROM authors ORDER BY author_id %s`, param)
	row, err := a.db.GetDb().Query(query)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var eachAuthor author.GetBookDto
		err := row.Scan(&eachAuthor.Id, &eachAuthor.Name)
		if err != nil {
			return nil, err
		}
		authors = append(authors, eachAuthor)
	}
	row.Close()
	return &authors, nil
}

func (a *authorPostgresRepository) DeleteAuthorSQL(id int) error {
	defer a.db.GetDb().Close()

	query := `DELETE FROM authors WHERE author_id = $1`
	row, err := a.db.GetDb().Exec(query, id)
	if err != nil {
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("0")
	}
	return nil
}

func (a *authorPostgresRepository) UpdateAuthorSQL(id int, in *author.InsertAuthorDto) error {
	defer a.db.GetDb().Close()

	query := `UPDATE authors SET name=$1 WHERE author_id=$2`
	row, err := a.db.GetDb().Exec(query, in.Name, id)
	if err != nil {
		return err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("0")
	}
	return nil
}
