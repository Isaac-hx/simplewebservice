package models

type AuthorRequest struct {
	Name string `json:"name" binding:"required"`
}

type AuthorResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
