// This file contain instance muxServer and implementation of interface server
package server

import (
	"fmt"
	"log"
	"net/http"
	"simplewebservice/config"
	"simplewebservice/database"
	"simplewebservice/handlers"
	"simplewebservice/repositories"
	"simplewebservice/usecases"
	"simplewebservice/utils"
)

type muxServer struct {
	app  *http.ServeMux
	db   database.Database
	conf *config.Config
}

func (ms *muxServer) Start() {
	ms.app.HandleFunc("GET /v1/health", func(w http.ResponseWriter, r *http.Request) {
		utils.LogServer(r)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK!"))
	})

	//Initialize object NewBookPostgresRepository
	bookPostgresRepository := repositories.NewBookPostgresRepository(ms.db)
	bookUsecaseImpl := usecases.NewBookUsecaseImpl(bookPostgresRepository)
	bookHandler := handlers.NewBookHttpHandler(bookUsecaseImpl)

	ms.app.HandleFunc("POST /v1/book", bookHandler.AddBook)
	ms.app.HandleFunc("GET /v1/book/{id}", bookHandler.SearchBookById)
	ms.app.HandleFunc("DELETE /v1/book/{id}", bookHandler.DeleteBookById)
	ms.app.HandleFunc("PUT /v1/book/{id}", bookHandler.UpdateBookById)
	ms.app.HandleFunc("GET /v1/book", bookHandler.FindAllBooks)

	authorPostgresRepository := repositories.NewAuthorPostgresRepository(ms.db)
	authorUsecaseImpl := usecases.NewAuthorUsecaseImpl(authorPostgresRepository)
	authorHandler := handlers.NewAuthorHttpHandler(authorUsecaseImpl)

	ms.app.HandleFunc("POST /v1/author", authorHandler.AddAuthor)
	ms.app.HandleFunc("GET /v1/author/{id}", authorHandler.SearchAuthorById)
	ms.app.HandleFunc("DELETE /v1/author/{id}", authorHandler.DeleteAuthorById)
	ms.app.HandleFunc("PUT /v1/author/{id}", authorHandler.UpdateAuthorById)
	ms.app.HandleFunc("GET /v1/author", authorHandler.FindAllAuthors)

	serverPort := fmt.Sprintf(":%d", ms.conf.Server.Port)
	log.Printf("Server running in addr %s", serverPort)

	//Running server
	err := http.ListenAndServe(serverPort, ms.app)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

// Construct to object muxServer
func NewServerMux(conf *config.Config, db database.Database) Server {
	muxApp := http.NewServeMux()
	muxServer := muxServer{app: muxApp, db: db, conf: conf}
	return &muxServer
}
