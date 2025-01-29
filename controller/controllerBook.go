package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simplewebservice/logger"
	"simplewebservice/models"
	"strconv"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	//Log server
	defer logger.LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))

	//Menentukan tipe response apa yang akan dikembalikan client
	w.Header().Set("Content-type", "application/json")

	//jika method yang digunakan adalah get
	if r.Method == "GET" {
		res, err := json.Marshal(models.ListOfBooks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		w.Write(res)
		return
	}
	http.Error(w, "Books not found!", http.StatusBadRequest)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	//Log server

	defer logger.LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))

	w.Header().Set("Content-type", "application/json")
	if r.Method == "GET" {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		for _, data := range models.ListOfBooks {
			//jika id ditemukan
			if data.ID == id {
				//proses encode data menjadi json
				response, err := json.Marshal(data)
				//jika terjadi error pada encode data
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				//response balik ke client dari json
				w.Write(response)
				return

			}
		}
		http.Error(w, "Book not found!", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
