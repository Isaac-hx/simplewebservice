// this file use to migration database (Will be using this file with ORM!!)
package migrations

import (
	"fmt"
	"log"
	"simplewebservice/book"
	"simplewebservice/database"
	"simplewebservice/utils"
	"strings"
)

// Migrate data book
func BookMigrate(db database.Database) {
	books := []book.InsertBookDto{
		{
			AuthorId:      1,
			Title:         "Sapiens: Riwayat Singkat Umat Manusia",
			Description:   "Narasi revolusioner tentang penciptaan dan evolusi umat manusia",
			TotalPage:     544,
			PublishedDate: utils.ParseDate("2014-09-04 00:00:00"),
			Price:         200000,
			CoverUrl:      "https://m.media-amazon.com/images/I/713jIoMO3UL.jpg",
		},
		{
			AuthorId:      1,
			Title:         "Homo Deus",
			Description:   "Eksplorasi masa depan manusia",
			TotalPage:     450,
			PublishedDate: utils.ParseDate("2017-02-21 00:00:00"),
			Price:         225000,
			CoverUrl:      "https://m.media-amazon.com/images/I/91HHqVTAJDL.jpg",
		},
		{
			AuthorId:      2,
			Title:         "Laskar Pelangi",
			Description:   "Kisah inspiratif tentang pendidikan di pelosok Indonesia",
			TotalPage:     529,
			PublishedDate: utils.ParseDate("2005-09-01 00:00:00"),
			Price:         150000,
			CoverUrl:      "https://example.com/laskar-pelangi.jpg",
		},
		{
			AuthorId:      3,
			Title:         "Bumi Manusia",
			Description:   "Roman sejarah tentang perjuangan di masa kolonial",
			TotalPage:     535,
			PublishedDate: utils.ParseDate("1980-08-01 00:00:00"),
			Price:         175000,
			CoverUrl:      "https://example.com/bumi-manusia.jpg",
		},
		{
			AuthorId:      4,
			Title:         "Filosofi Teras",
			Description:   "Pengenalan stoikisme dalam kehidupan modern",
			TotalPage:     346,
			PublishedDate: utils.ParseDate("2018-11-15 00:00:00"),
			Price:         130000,
			CoverUrl:      "https://example.com/filosofi-teras.jpg",
		},
		{
			AuthorId:      5,
			Title:         "Atomic Habits",
			Description:   "Cara membangun kebiasaan baik",
			TotalPage:     320,
			PublishedDate: utils.ParseDate("2018-10-16 00:00:00"),
			Price:         190000,
			CoverUrl:      "https://m.media-amazon.com/images/I/91PYouxG2dL.jpg",
		},
		{
			AuthorId:      6,
			Title:         "Dilan 1990",
			Description:   "Kisah cinta remaja di Bandung",
			TotalPage:     332,
			PublishedDate: utils.ParseDate("2014-04-01 00:00:00"),
			Price:         120000,
			CoverUrl:      "https://example.com/dilan-1990.jpg",
		},
		{
			AuthorId:      7,
			Title:         "The Psychology of Money",
			Description:   "Pemahaman tentang perilaku keuangan",
			TotalPage:     256,
			PublishedDate: utils.ParseDate("2020-09-08 00:00:00"),
			Price:         180000,
			CoverUrl:      "https://m.media-amazon.com/images/I/71gN4sWPHkL.jpg",
		},
		{
			AuthorId:      8,
			Title:         "Tenggelamnya Kapal Van Der Wijck",
			Description:   "Novel klasik tentang cinta dan tragedi",
			TotalPage:     255,
			PublishedDate: utils.ParseDate("1938-01-01 00:00:00"),
			Price:         110000,
			CoverUrl:      "https://example.com/van-der-wijck.jpg",
		},
		{
			AuthorId:      9,
			Title:         "Sebuah Seni untuk Bersikap Bodo Amat",
			Description:   "Panduan hidup minimalis dan praktis",
			TotalPage:     248,
			PublishedDate: utils.ParseDate("2016-09-13 00:00:00"),
			Price:         160000,
			CoverUrl:      "https://example.com/bodo-amat.jpg",
		},
	}

	// Query untuk batch insert (sama seperti sebelumnya)
	query := "INSERT INTO books (id, title, author_id, total_page, description, published_date, price, cover_url) VALUES "
	var values []interface{}
	valueStrings := make([]string, 0, len(books))

	for i, b := range books {
		startIdx := i*8 + 1
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
			startIdx, startIdx+1, startIdx+2, startIdx+3, startIdx+4, startIdx+5, startIdx+6, startIdx+7))
		values = append(values, i+1, b.Title, b.AuthorId, b.TotalPage, b.Description, b.PublishedDate, b.Price, b.CoverUrl)
	}

	query += strings.Join(valueStrings, ",")

	result, err := db.GetDb().Exec(query, values...)
	if err != nil {
		log.Fatalf("Error executing batch insert: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Error getting rows affected: %v", err)
	}

	log.Printf("Success migration data, %d rows affected!", rowsAffected)
}
