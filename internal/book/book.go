package book

import (
	"database/sql"
	"errors"
	"time"
)

type Book struct {
	Id            int       `json:"id"`
	Title         string    `json:"title"`
	AuthorId      int       `json:"author_id"` //reference key from table authorss
	TotalPage     string    `json:"total_page"`
	Description   string    `json:"description"`
	PublishedDate time.Time `json:"published_date"`
	Price         float64   `json:"price"`
	CoverUrl      string    `json:"cover_url"`
}

// call construct object book
func New() *Book {
	bookObject := &Book{}
	return bookObject
}

type BookWithAuthor struct {
	Book
	AuthorName string `json:"name"`
}

func (b *Book) GetAllBooks(db *sql.DB) ([]BookWithAuthor, error) {
	var listOfBooks []BookWithAuthor
	query := `SELECT 
	books.id,
	books.title, 
	books.author_id,
	books.total_page,
	books.description,
	books.published_date,
	books.price,
	books.cover_url,
	authors.name FROM books
	INNER JOIN authors ON books.author_id=authors.author_id `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var eachBook BookWithAuthor
		err := rows.Scan(&eachBook.Id, &eachBook.Title, &eachBook.AuthorId, &eachBook.TotalPage, &eachBook.Description, &eachBook.PublishedDate, &eachBook.Price, &eachBook.CoverUrl, &eachBook.AuthorName)

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
	err := db.QueryRow(query, id).Scan(&book.Id, &book.Title, &book.AuthorId, &book.TotalPage, &book.Description, &book.PublishedDate, &book.Price, &book.CoverUrl)
	if err != nil {
		return book, err

	}

	return book, nil
}

func (b *Book) CreateBook(db *sql.DB, title string, author_id int, total_page int, description string, published_date time.Time, price float64, cover_url string) error {
	query := `INSERT INTO books(title,author_id,total_page,description,published_date,price,cover_url) VALUES($1,$2,$3,$4,$5,$6,$7)`
	_, err := db.Exec(query, title, author_id, total_page, description, published_date, price, cover_url)
	if err != nil {
		return err
	}
	return nil
}

func (b *Book) UpdateBookById(db *sql.DB, id int, title string, author_id int, total_page int, description string, published_date time.Time, price float64, cover_url string) error {
	query := `UPDATE books SET title=$1, author_id=$2, total_page=$3, description=$4,published_date=$5,price=$6,cover_url=$7 WHERE id=$8`
	row, err := db.Exec(query, title, author_id, total_page, description, published_date, price, cover_url, id)
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
		err = rows.Scan(&eachBook.Id, &eachBook.Title, &eachBook.AuthorId, &eachBook.TotalPage, &eachBook.Description, &eachBook.PublishedDate, &eachBook.CoverUrl, &eachBook.Price)
		if err != nil {
			return nil, err
		}

		filteredBooks = append(filteredBooks, eachBook)

	}
	return filteredBooks, nil

}
