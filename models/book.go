package models

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	AuthorId  int    `json:"author_id"`
	TotalPage int    `json:"total_page"`
	Publisher string `json:"publisher"`
}
