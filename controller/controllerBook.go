package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"simplewebservice/logger"
	"simplewebservice/models"
	"strconv"
)

func Getbook(w http.ResponseWriter, r *http.Request) {
	//Log server
	defer logger.LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))
	//Menentukan tipe response apa yang akan dikembalikan client
	w.Header().Set("Content-type", "application/json")

	//jika dia memiliki query parameter
	if id := r.URL.Query().Get("id"); id != "" {
		//konvert nilai string ke integer
		param, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		GetBookById(w, r, param)
		return

	}

	res, err := json.Marshal(models.ListOfBooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request, id int) {
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
			//response balik ke client melalui format data
			w.Write(response)
			return

		}
	}
	http.Error(w, "Book not found!", http.StatusNotFound)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	//Log server
	defer logger.LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))

	w.Header().Set("Content-type", "application/json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error from server !!", http.StatusInternalServerError)
		return
	}

	var data models.Book
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	models.ListOfBooks = append(models.ListOfBooks, data)
	w.WriteHeader(http.StatusCreated)
	res := map[string]string{"message": "Book created sucessfully!!"}
	json.NewEncoder(w).Encode(res)

}
