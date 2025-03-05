// this file use to migration database (Will be using this file with ORM!!)
package migrations

import (
	"log"
	"simplewebservice/book"
	"simplewebservice/database"
	"time"
)

// Migrate data book
func BookMigrate(db database.Database) {
	var book1 book.InsertBookDto
	defer func() {
		db.GetDb().Close()
	}()

	publishedDate, err := time.Parse("2006-01-02 15:04:05", "2014-09-04 00:00:00")
	if err != nil {
		log.Fatal(err.Error())

	}
	book1 = book.InsertBookDto{
		AuthorId:      1,
		Title:         "Sapiens : Riwayat Singkat Umat Manusia",
		Description:   "Dari seorang sejarawan terkemuka hadir sebuah narasi revolusioner tentang penciptaan dan evolusi umat manusia",
		TotalPage:     544,
		PublishedDate: publishedDate,
		Price:         200.000,
		CoverUrl:      "https://m.media-amazon.com/images/I/713jIoMO3UL.jpg",
	}
	query := "INSERT INTO books VALUES($1,$2,$3,$4,$5,$6,$7,$8)"
	row, err := db.GetDb().Query(query, 1, book1.Title, book1.AuthorId, book1.TotalPage, book1.Description, book1.PublishedDate, book1.Price, book1.CoverUrl)
	if err != nil {
		log.Fatal(err.Error())
	}
	row.Close()

	log.Println("Sucess migration data,1 rows affected!")

}
