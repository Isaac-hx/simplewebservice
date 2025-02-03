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

// update book controller
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//Log Server
	defer logger.LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))

	w.Header().Set("Content-type", "application/json")
	if r.Method != "PUT" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if id := r.URL.Query().Get("id"); id != "" {
		param, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		//Mencari data buku berdasarkan id
		for i, data := range models.ListOfBooks {
			//jika id ditemukan
			if data.ID == param {
				body, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "Error from server !!", http.StatusInternalServerError)
					return
				}

				//Memuat data buku yang akan diupdate
				var book models.Book
				if err := json.Unmarshal(body, &book); err != nil {
					http.Error(w, "Invalid JSON format", http.StatusBadRequest)
					return
				}
				//Mengupdate id data buku
				book.ID = param

				//Mengupdate data buku ke list buku
				models.ListOfBooks[i] = book
				w.WriteHeader(http.StatusOK)
				res := map[string]string{"message": "Book updated sucessfully!!"}
				json.NewEncoder(w).Encode(res)
				return
			}

		}
	}
	http.Error(w, "Book not found!", http.StatusNotFound)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//Log server
	defer logger.LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))

	w.Header().Set("Content-type", "application/json")
	if r.Method != "DELETE" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	//Mengambil nilai id dari query parameter
	if id := r.URL.Query().Get("id"); id != "" {
		param, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		//Mencari data buku berdasarkan id
		for i, data := range models.ListOfBooks {
			if data.ID == param {
				//Menghapus buku dari data list menggunakan metode slicing array
				models.ListOfBooks = append(models.ListOfBooks[:i], models.ListOfBooks[i+1:]...)
				w.WriteHeader(http.StatusOK)
				res := map[string]string{"message": "Book deleted sucessfully!!"}
				json.NewEncoder(w).Encode(res)
				return
			}
		}
	}
	http.Error(w, "Book not found!", http.StatusNotFound)

}
