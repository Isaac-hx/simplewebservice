// This package contain object from request user
package models

import "time"

type BookRequest struct {
	Title         string  `json:"title" binding:"required"`
	AuthorId      int     `json:"author_id" binding:"required"` //reference key from table authorss
	TotalPage     int     `json:"total_page" binding:"required"`
	Description   string  `json:"description" binding:"required"`
	PublishedDate string  `json:"published_date" binding:"required"`
	Price         float64 `json:"price" binding:"required"`
	CoverUrl      string  `json:"cover_url" binding:"required"`
}

type BookResponse struct {
	Id            int       `json:"id"`
	Title         string    `json:"title"`
	AuthorName    string    `json:"author_name"` //reference key from table authorss
	TotalPage     int       `json:"total_page"`
	Description   string    `json:"description"`
	PublishedDate time.Time `json:"published_date"`
	Price         float64   `json:"price"`
	CoverUrl      string    `json:"cover_url"`
}
