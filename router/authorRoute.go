package router

import (
	"net/http"
	"simplewebservice/database"
	"simplewebservice/handlers"
	repositories "simplewebservice/repositories/author"
	"simplewebservice/usecases"
)

type AuthorRouter struct{}

func (r *AuthorRouter) ListAuthorRoute(db database.Database, app *http.ServeMux) {

	authorPostgresRepository := repositories.NewAuthorPostgresRepository(db)
	authorUsecaseImpl := usecases.NewAuthorUsecaseImpl(authorPostgresRepository)
	authorHandler := handlers.NewAuthorHttpHandler(authorUsecaseImpl)

	app.HandleFunc("POST /v1/author", authorHandler.AddAuthor)
	app.HandleFunc("GET /v1/author/{id}", authorHandler.SearchAuthorById)
	app.HandleFunc("DELETE /v1/author/{id}", authorHandler.DeleteAuthorById)
	app.HandleFunc("PUT /v1/author/{id}", authorHandler.UpdateAuthorById)
	app.HandleFunc("GET /v1/author", authorHandler.FindAllAuthors)
}

func NewAuthorRouter() *AuthorRouter {
	return &AuthorRouter{}
}
