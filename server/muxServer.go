// This file contain instance muxServer and implementation of interface server
package server

import (
	"fmt"
	"log"
	"net/http"
	"simplewebservice/config"
	"simplewebservice/database"
	"simplewebservice/router"

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

	//Initialize object route
	bookRoute := router.NewBookRoute()
	bookRoute.ListBookRoute(ms.db, ms.app)
	authorRoute := router.NewAuthorRouter()
	authorRoute.ListAuthorRoute(ms.db, ms.app)

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
