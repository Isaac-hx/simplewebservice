package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"simplewebservice/config"
	"simplewebservice/logger"
	"simplewebservice/models"
)

func Getbook(w http.ResponseWriter, r *http.Request) {
	//Menentukan tipe response apa yang akan dikembalikan client
	w.Header().Set("Content-type", "application/json")
	//connect database
	db, err := config.Connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	//Menutup koneksi database
	defer db.Close()

	queryString := "SELECT * FROM books"
	rows, err := db.Query(queryString)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//Menutup koneksi rows
	defer rows.Close()

	var listOfBooks []models.Book

	//Mengiterasi data yang didapatkan dari variabel rows pada object sql.Rows
	for rows.Next() {
		//Variabel penampung data pada setiap iterasi
		var eachBook models.Book
		//Menscan data yang diterima dan hasilnya nanti akan ditampung pada variabel eachbook
		err = rows.Scan(&eachBook.ID, &eachBook.Title, &eachBook.Author, &eachBook.TotalPage, &eachBook.Publisher)
		if err != nil {
			log.Println(err.Error())
			return
		}
		//Menambahkan data pada variabel eachBook ke slice listOfBook
		listOfBooks = append(listOfBooks, eachBook)

	}

	//Log server
	defer logger.LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listOfBooks)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	// Ambil ID dari query parameter
	id := r.URL.Query().Get("id")

	var book models.Book

	// Menyiapkan SQL statement berdasarkan ID
	//$1 merupakan placeholdar yang akan membinding parameter ke-1
	query := "SELECT id, title, author, total_page, publisher FROM books WHERE id = $1"
	err = db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.TotalPage, &book.Publisher)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		log.Println("Query Error: ", err.Error())
		return
	}

	// Log server
	logger.LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))

	// Menentukan tipe response dan mengirimkan data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	//Connect database
	db, err := config.Connect()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()

	//Membaca data dari body request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error from server !!", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	var parseData models.Book

	if err := json.Unmarshal(body, &parseData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		log.Println(err.Error())

		return
	}
	if parseData.TotalPage <= 0 {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)

		return
	}
	_, err = db.Exec("INSERT INTO books (title, author, total_page, publisher) VALUES ($1, $2, $3, $4)", parseData.Title, parseData.Author, int(parseData.TotalPage), parseData.Publisher)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	//Log server
	defer logger.LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))
	w.WriteHeader(http.StatusOK)
	responseBody := map[string]string{"Message": "Success create book!!"}
	json.NewEncoder(w).Encode(responseBody)

}

// update book controller
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	db, err := config.Connect()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	id := r.URL.Query().Get("id")
	var book models.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid format JSON!!!", http.StatusBadRequest)
		return
	}
	queryString := `
	UPDATE books
	SET title = $1, 
		author = $2, 
		total_page = $3, 
		publisher = $4
	WHERE id = $5`

	row, err := db.Exec(queryString, book.Title, book.Author, book.TotalPage, book.Publisher, id)
	if err != nil {
		http.Error(w, "Error deleting book!!!", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		http.Error(w, "Error retrieving affected rows", http.StatusInternalServerError)
		log.Println("RowsAffected error:", err.Error())
		return

	}
	if rowAffected != 1 {
		http.Error(w, "Book not found or not deleted", http.StatusNotFound)
		log.Printf("Delete operation affected %d rows\n", rowAffected)
		return
	}
	//Mencari data buku berdasarkan id
	//Log Server
	defer logger.LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))
	w.WriteHeader(http.StatusOK)
	message := map[string]string{"Message": "Book update successfully!"}
	json.NewEncoder(w).Encode(message)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db, err := config.Connect()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	defer db.Close()
	//Mengambil nilai id dari query parameter
	id := r.URL.Query().Get("id")
	//Mencari data buku berdasarkan id
	row, err := db.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Error deleting book!!!", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		http.Error(w, "Error retrieving affected rows", http.StatusInternalServerError)
		log.Println("RowsAffected error:", err.Error())
		return

	}
	if rowAffected != 1 {
		http.Error(w, "Book not found or not deleted", http.StatusNotFound)
		log.Printf("Delete operation affected %d rows\n", rowAffected)
		return
	}
	//Log server
	defer logger.LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))
	w.WriteHeader(http.StatusOK)
	message := map[string]string{"Message": "Delete book sucessfully!"}
	json.NewEncoder(w).Encode(message)
}
