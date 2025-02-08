//this package contains route based on object author

package router

import (
	"net/http"
	"simplewebservice/service"
)

func AuthorRoute(mux *http.ServeMux) {
	authorServer := service.NewServeAuthor()

	mux.HandleFunc("GET /author", authorServer.GetAllAuthors)
	mux.HandleFunc("POST /author", authorServer.CreateAuthor)
	mux.HandleFunc("GET /author/{id}", authorServer.GetAuthorById)
	mux.HandleFunc("DELETE /author/{id}", authorServer.DeleteAuthorById)
	mux.HandleFunc("PUT /author/{id}", authorServer.UpdateAuthorById)
}
