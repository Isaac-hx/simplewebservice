package service

import (
	"encoding/json"
	"mime"
	"net/http"
	"simplewebservice/config"
	"simplewebservice/internal/book"
	"simplewebservice/library"
	"simplewebservice/utils"
	"strconv"
)

type reqBook struct {
	Title         string `json:"title"`
	AuthorId      int    `json:"author_id"`
	TotalPage     int    `json:"total_page"`
	Description   string `json:"description"`
	PublishedDate string `json:"published_date"`
}
type handlerBook struct {
	book *book.Book
}

func NewServeBook() *handlerBook {
	book := book.New()
	return &handlerBook{book: book}
}

func (handler *handlerBook) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	utils.LogServer(r)

	db, err := config.Connect(library.Postgres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	filterTotalPage := r.URL.Query().Get("total_page")
	if filterTotalPage != "" {
		n, err := strconv.Atoi(filterTotalPage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, err := handler.book.SelectBookByFilter(db, n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&data)
	} else {

		data, err := handler.book.GetAllBooks(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&data)
	}

}

func (handler *handlerBook) GetBookById(w http.ResponseWriter, r *http.Request) {
	utils.LogServer(r)
	db, err := config.Connect(library.Postgres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	reqId := r.PathValue("id")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.book.GetBookById(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&data)
}

func (handler *handlerBook) CreateBook(w http.ResponseWriter, r *http.Request) {
	utils.LogServer(r)
	var bookReq reqBook

	db, err := config.Connect(library.Postgres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	//enforce json type
	contentType := r.Header.Get("Content-type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	//jika parsing gagal
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//jika data yang dikirimkan bukan type json
	if mediaType != "application/json" {
		http.Error(w, "Error : Expect application/json data", http.StatusUnsupportedMediaType)
		return
	}
	reqBody := json.NewDecoder(r.Body)
	reqBody.DisallowUnknownFields()
	//parsing json to object bookReq
	parseBody := reqBody.Decode(&bookReq)
	if parseBody != nil {
		http.Error(w, "Invalid format json!!", http.StatusBadRequest)
		return
	}
	//parsing published string to time.Time
	date, err := utils.ParseTimeDate(bookReq.PublishedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Call method create book from object handler
	err = handler.book.CreateBook(db, bookReq.Title, bookReq.AuthorId, bookReq.TotalPage, bookReq.Description, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := map[string]interface{}{"Message": "Data created sucessfully!!"}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&message)
}

func (handler *handlerBook) UpdateBookById(w http.ResponseWriter, r *http.Request) {
	utils.LogServer(r)

	var bookReq reqBook

	db, err := config.Connect(library.Postgres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	//enforce json type
	contentType := r.Header.Get("Content-type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	//jika parsing gagal
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//jika data yang dikirimkan bukan type json
	if mediaType != "application/json" {
		http.Error(w, "Error : Expect application/json data", http.StatusUnsupportedMediaType)
		return
	}
	reqId := r.PathValue("id")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reqBody := json.NewDecoder(r.Body)
	reqBody.DisallowUnknownFields()

	parseBody := reqBody.Decode(&bookReq)
	if parseBody != nil {
		http.Error(w, "Invalid format json!!", http.StatusBadRequest)
		return
	}
	date, err := utils.ParseTimeDate(bookReq.PublishedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.book.UpdateBookById(db, id, bookReq.Title, bookReq.AuthorId, bookReq.TotalPage, bookReq.Description, date)
	if err != nil {
		if err.Error() == "0" {
			http.Error(w, "Book not updated!", http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	message := map[string]interface{}{"Message": "Succes updated book!!"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)

}

func (handler *handlerBook) DeleteBookById(w http.ResponseWriter, r *http.Request) {
	utils.LogServer(r)

	db, err := config.Connect(library.Postgres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	reqId := r.PathValue("id")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.book.DeleteBookById(db, id)
	if err != nil {
		if err.Error() == "0" {
			http.Error(w, "Book not deleted!", http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	message := map[string]interface{}{"Message": "Succes deleted book!!"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}
