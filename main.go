package main

import (
	"fmt"
	"net/http"
	"simplewebservice/controller"
	"simplewebservice/utils"
)

func main() {

	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			//jika method GET terdeteksi memiliki query id
			if id := r.URL.Query().Get("id"); id != "" {
				controller.GetBookById(w, r)
				return
			}
			controller.Getbook(w, r)
		case "POST":
			controller.CreateBook(w, r)
		case "DELETE":
			controller.DeleteBook(w, r)
		case "PUT":
			controller.UpdateBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

	})
	http.HandleFunc("/card", controller.CreateCardIdentity)
	http.HandleFunc("/shape", controller.GetCalculateShape)
	http.HandleFunc("/shape/rotate", controller.GetRotateShape)
	utils.ListRoute("/book", "/shape", "/shape/rotate", "/card")
	fmt.Println("Service berjalan di localhost:8085/")

	http.ListenAndServe(":8085", nil)
}
