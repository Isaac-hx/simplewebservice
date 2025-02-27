// This file to define schema for data in the database
package book

import "time"

type InsertBookDto struct {
	Title         string
	AuthorId      int
	TotalPage     int
	Description   string
	PublishedDate time.Time
	Price         float64
	CoverUrl      string
}
type GetBookDto struct {
	Id            int
	Title         string
	AuthorName    string
	TotalPage     int
	Description   string
	PublishedDate time.Time
	Price         float64
	CoverUrl      string
}
