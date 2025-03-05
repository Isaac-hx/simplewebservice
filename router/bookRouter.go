package router

import (
	"net/http"
	"simplewebservice/database"
	"simplewebservice/handlers"
	repositories "simplewebservice/repositories/book"
	"simplewebservice/usecases"
)

type BookRouter struct{}

func (r *BookRouter) ListBookRoute(db database.Database, app *http.ServeMux) {
	bookPostgresRepository := repositories.NewBookPostgresRepository(db)
	bookUsecaseImpl := usecases.NewBookUsecaseImpl(bookPostgresRepository)
	bookHandler := handlers.NewBookHttpHandler(bookUsecaseImpl)
	app.HandleFunc("POST /v1/book", bookHandler.AddBook)
	app.HandleFunc("GET /v1/book/{id}", bookHandler.SearchBookById)
	app.HandleFunc("DELETE /v1/book/{id}", bookHandler.DeleteBookById)
	app.HandleFunc("PUT /v1/book/{id}", bookHandler.UpdateBookById)
	app.HandleFunc("GET /v1/book", bookHandler.FindAllBooks)
}

func NewBookRoute() *BookRouter {
	return &BookRouter{}
}
