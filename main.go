package main

import (
	"log"
	"net/http"
	"simplewebservice/router"
	"simplewebservice/utils"
)

func main() {
	//initialising object servermux
	mux := http.NewServeMux()
	//registered route
	router.AuthorRoute(mux)
	router.BookRoute(mux)

	utils.ListRoute("/book", "/author")
	log.Println("Service berjalan di localhost:8085/")

	http.ListenAndServe(":8085", mux)
}
