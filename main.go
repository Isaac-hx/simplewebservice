package main

import (
	"fmt"
	"net/http"
	"simplewebservice/controller"
	"simplewebservice/logger"
)

func main() {
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			controller.Getbook(w, r)
		case "POST":
			controller.CreateBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

	})
	http.HandleFunc("/card", controller.CreateCardIdentity)
	http.HandleFunc("/shape", controller.GetCalculateShape)
	http.HandleFunc("/shape/rotate", controller.GetRotateShape)

	logger.ListRoute("/book", "/shape", "/shape/rotate", "/card")
	fmt.Println("Service berjalan di localhost:8085/")

	http.ListenAndServe(":8085", nil)
}
