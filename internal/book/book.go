package book

import (
	"database/sql"
	"errors"
)

type Book struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	AuthorId    int    `json:"author_id"` //reference key from table authorss
	TotalPage   string `json:"total_page"`
	Description string `json:"description"`
}

func New() *Book {
	bookObject := &Book{}
	return bookObject
}

func (b *Book) GetAllBooks(db *sql.DB) ([]Book, error) {
	var listOfBooks []Book
	query := `SELECT * FROM books`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var eachBook Book
		err := rows.Scan(&eachBook.Id, &eachBook.Title, &eachBook.AuthorId, &eachBook.TotalPage, &eachBook.Description)
		if err != nil {
			return nil, err
		}
		listOfBooks = append(listOfBooks, eachBook)
	}
	return listOfBooks, nil
}

func (b *Book) GetBookById(db *sql.DB, id int) (Book, error) {
	var book Book
	query := `SELECT * FROM books WHERE id=$1`
	err := db.QueryRow(query, id).Scan(&book.Id, &book.Title, &book.AuthorId, &book.TotalPage, &book.Description)
	if err != nil {
		return book, err

	}

	return book, nil
}

func (b *Book) CreateBook(db *sql.DB, title string, author_id int, total_page int, description string) error {
	query := `INSERT INTO books(title,author_id,total_page,description) VALUES($1,$2,$3,$4)`
	_, err := db.Exec(query, title, author_id, total_page, description)
	if err != nil {
		return err
	}
	return nil
}

func (b *Book) UpdateBookById(db *sql.DB, id int, title string, author_id int, total_page int, description string) error {
	query := `UPDATE books SET title=$1, author_id=$2, total_page=$3, description=$4 WHERE id=$5`
	row, err := db.Exec(query, title, author_id, total_page, description, id)
	if err != nil {
		return err
	}

	//jika rows effected tidak didapatkan
	rowAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected != 1 {
		return errors.New("0")
	}

	return nil

}

func (b *Book) DeleteBookById(db *sql.DB, id int) error {
	query := `DELETE from books WHERE id=$1`
	row, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	rowAffected, err := row.RowsAffected()
	if rowAffected != 1 {
		return errors.New("0")
	}
	if err != nil {
		return err
	}
	return nil

}

//create a method to filter that book have a total_page > n

func (b *Book) SelectBookByFilter(db *sql.DB, n int) ([]Book, error) {
	var filteredBooks []Book
	query := `SELECT * FROM books WHERE total_page > $1`
	rows, err := db.Query(query, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var eachBook Book
		err = rows.Scan(&eachBook.Id, &eachBook.Title, &eachBook.AuthorId, &eachBook.TotalPage, &eachBook.Description)
		if err != nil {
			return nil, err
		}

		filteredBooks = append(filteredBooks, eachBook)

	}
	return filteredBooks, nil

}
