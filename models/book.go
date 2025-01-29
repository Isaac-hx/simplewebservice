package models

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	TotalPage int    `json:"total_page"`
	Publisher string `json:"publisher"`
}
