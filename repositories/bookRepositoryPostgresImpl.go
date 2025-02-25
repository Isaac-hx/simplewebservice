package repositories

import (
	"database/sql"
	"simplewebservice/database"
	"simplewebservice/internal/book"
)

type bookPostgresRepository struct {
	db database.Database //depedency injection : save instance database which is implement interface database
}

// Method from object bookPostgresRepository
func (b *bookPostgresRepository) InsertBookSQL(in *book.InsertBookDto) error {
	//perform query here
	defer func() {
		b.db.GetDb().Close() //closed connection to database
	}()
	query := "INSERT INTO books(title,author_id,total_page,description,published_date,price,cover_url) VALUES($1,$2,$3,$4,$5,$6,$7)"
	_, err := b.db.GetDb().Exec(query, in.Title, in.AuthorId, in.TotalPage, in.Description, in.PublishedDate, in.Price, in.CoverUrl)
	if err != nil {
		return err
	}

	return nil
}

func (b *bookPostgresRepository) GetBookSQL(id int) *sql.Row {
	defer func() {
		b.db.GetDb().Close()
	}()

	query := `SELECT * FROM books WHERE id = $1`
	row := b.db.GetDb().QueryRow(query, id)
	return row

}

// Constructor of object bookPostgresRepository
// to instancing the object bookPostgresRepository
func NewBookPostgresRepository(db database.Database) BookRepository {
	return &bookPostgresRepository{db: db}
}
