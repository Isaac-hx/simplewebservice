package handlers

import "net/http"

type AuthorHandler interface {
	AddAuthor(w http.ResponseWriter, r *http.Request)
	SearchAuthorById(w http.ResponseWriter, r *http.Request)
	DeleteAuthorById(w http.ResponseWriter, r *http.Request)
	UpdateAuthorById(w http.ResponseWriter, r *http.Request)
	FindAllAuthors(w http.ResponseWriter, r *http.Request)
}
