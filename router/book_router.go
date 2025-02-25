// this package contains route based on object book
package router

import (
	"database/sql"
	"net/http"
	"simplewebservice/service"
)

func BookRoute(mux *http.ServeMux, db *sql.DB) {
	bookServer := service.NewServeBook(db)

	mux.HandleFunc("GET /book", bookServer.GetAllBooks)
	mux.HandleFunc("GET /book/{id}", bookServer.GetBookById)
	mux.HandleFunc("POST /book", bookServer.CreateBook)
	mux.HandleFunc("DELETE /book/{id}", bookServer.DeleteBookById)
	mux.HandleFunc("PUT /book/{id}", bookServer.UpdateBookById)
}
