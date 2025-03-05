package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"simplewebservice/book"
	"simplewebservice/database"
)

type bookPostgresRepository struct {
	db database.Database //depedency injection : save instance database which is implement interface database
}

// Method from object bookPostgresRepository
func (b *bookPostgresRepository) InsertBookSQL(in *book.InsertBookDto) error {
	defer b.db.GetDb().Close()

	query := "INSERT INTO books(title,author_id,total_page,description,published_date,price,cover_url) VALUES($1,$2,$3,$4,$5,$6,$7)"
	_, err := b.db.GetDb().Exec(query, in.Title, in.AuthorId, in.TotalPage, in.Description, in.PublishedDate, in.Price, in.CoverUrl)
	if err != nil {
		return err
	}

	return nil
}

func (b *bookPostgresRepository) GetBookSQL(id int) (*book.GetBookDto, error) {
	defer b.db.GetDb().Close()

	var book book.GetBookDto
	query := `SELECT 
	books.id,books.title,
	authors.name,books.total_page,
	books.description,books.published_date,
	books.price,books.cover_url 
	FROM books
	INNER JOIN authors
	ON books.author_id = authors.author_id
	WHERE id = $1`
	err := b.db.GetDb().QueryRow(query, id).Scan(&book.Id, &book.Title, &book.AuthorName, &book.TotalPage, &book.Description, &book.PublishedDate, &book.Price, &book.CoverUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("0")
		}
		return nil, err
	}
	return &book, nil

}

func (b *bookPostgresRepository) DeleteBookSQL(id int) error {
	defer b.db.GetDb().Close()

	query := "DELETE FROM books WHERE id = $1"
	row, err := b.db.GetDb().Exec(query, id)
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

func (b *bookPostgresRepository) UpdateBookSQL(id int, in *book.InsertBookDto) error {
	defer b.db.GetDb().Close()

	query := `UPDATE books 
	SET title=$1,
	author_id=$2,
	description=$3,
	total_page=$4,
	price=$5,
	cover_url=$6,
	published_date=$7
	WHERE id = $8`

	row, err := b.db.GetDb().Exec(query, in.Title, in.AuthorId, in.Description, in.TotalPage, in.Price, in.CoverUrl, in.PublishedDate, id)
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

func (b *bookPostgresRepository) GetListBookSQL(param string) (*[]book.GetBookDto, error) {
	defer b.db.GetDb().Close()

	var books []book.GetBookDto
	query := fmt.Sprintf(`SELECT 
    books.id, books.title,
    books.price, books.cover_url, authors.name 
    FROM books 
    INNER JOIN authors 
    ON books.author_id = authors.author_id
	ORDER BY books.id %s`, param)
	rows, err := b.db.GetDb().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var eachBook book.GetBookDto
		err := rows.Scan(&eachBook.Id, &eachBook.Title, &eachBook.Price, &eachBook.CoverUrl, &eachBook.AuthorName)
		if err != nil {
			return nil, err
		}
		books = append(books, eachBook)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &books, nil
}

// Constructor of object bookPostgresRepository
// to instancing the object bookPostgresRepository
// to return out the interface bookrepository
func NewBookPostgresRepository(db database.Database) BookRepository {
	return &bookPostgresRepository{db: db}
}
