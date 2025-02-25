package handlers

import "net/http"

type Bookhandler interface {
	AddBook(w http.ResponseWriter, r *http.Request)
	SearchBookById(w http.ResponseWriter, r *http.Request)
}
