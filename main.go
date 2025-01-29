package main

import (
	"fmt"
	"net/http"
	"simplewebservice/controller"
	"simplewebservice/logger"
)

func main() {
	http.HandleFunc("/books", controller.GetAllBooks)
	http.HandleFunc("/book", controller.GetBookById)
	logger.ListRoute("/books", "/book")
	fmt.Println("Service berjalan di localhost:8085/")

	http.ListenAndServe(":8085", nil)
}
