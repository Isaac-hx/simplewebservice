package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"simplewebservice/internal/book"
	"simplewebservice/utils"
	"strconv"
)

type responseJson struct {
	Title         string  `json:"title"`
	AuthorId      int     `json:"author_id"`
	TotalPage     int     `json:"total_page"`
	Description   string  `json:"description"`
	PublishedDate string  `json:"published_date"`
	Price         float64 `json:"price"`
	CoverUrl      string  `json:"cover_url"`
}
type handlerBook struct {
	book *book.Book
	db   *sql.DB
}

func NewServeBook(db *sql.DB) *handlerBook {
	book := book.New()
	return &handlerBook{book: book, db: db}
}

func (handler *handlerBook) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	utils.LogServer(r)
	fmt.Println(handler.db)
	//Connect database
	defer handler.db.Close()

	filterTotalPage := r.URL.Query().Get("total_page")
	if filterTotalPage != "" {
		n, err := strconv.Atoi(filterTotalPage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, err := handler.book.SelectBookByFilter(handler.db, n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&data)
	} else {

		data, err := handler.book.GetAllBooks(handler.db)
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
	//Connect database
	defer handler.db.Close()

	reqId := r.PathValue("id")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := handler.book.GetBookById(handler.db, id)
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
	var bookReq responseJson

	//Connect database

	defer handler.db.Close()

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
	//verify the url is valid image

	if !utils.VerifyCoverUrl(bookReq.CoverUrl) {
		http.Error(w, "Invalid cover url!!", http.StatusBadRequest)
		return
	}
	//Call method create book from object handler
	err = handler.book.CreateBook(handler.db, bookReq.Title, bookReq.AuthorId, bookReq.TotalPage, bookReq.Description, date, bookReq.Price, bookReq.CoverUrl)
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

	var bookReq responseJson

	//Connect database

	defer handler.db.Close()
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
	if !utils.VerifyCoverUrl(bookReq.CoverUrl) {
		http.Error(w, "Invalid cover url!!", http.StatusBadRequest)
		return
	}

	err = handler.book.UpdateBookById(handler.db, id, bookReq.Title, bookReq.AuthorId, bookReq.TotalPage, bookReq.Description, date, bookReq.Price, bookReq.CoverUrl)
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

	//Connect database

	defer handler.db.Close()
	reqId := r.PathValue("id")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.book.DeleteBookById(handler.db, id)
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
