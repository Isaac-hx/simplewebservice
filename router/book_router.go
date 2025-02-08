// this package contains route based on object book
package router

import (
	"net/http"
	"simplewebservice/service"
)

func BookRoute(mux *http.ServeMux) {
	bookServer := service.NewServeBook()

	mux.HandleFunc("GET /book", bookServer.GetAllBooks)
	mux.HandleFunc("GET /book/{id}", bookServer.GetBookById)
	mux.HandleFunc("POST /book", bookServer.CreateBook)
	mux.HandleFunc("DELETE /book/{id}", bookServer.DeleteBookById)
	mux.HandleFunc("PUT /book/{id}", bookServer.UpdateBookById)
}
