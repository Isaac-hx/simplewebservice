// This file to define schema for data in the database
package book

import "time"

type InsertBookDto struct {
	Title         string    `json:"title"`
	AuthorId      int       `json:"author_id"` //reference key from table authorss
	TotalPage     int       `json:"total_page"`
	Description   string    `json:"description"`
	PublishedDate time.Time `json:"published_date"`
	Price         float64   `json:"price"`
	CoverUrl      string    `json:"cover_url"`
}
