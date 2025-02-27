package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"simplewebservice/models"
	"simplewebservice/usecases"
	"strconv"
)

type bookHttpHandler struct {
	bookUsecase usecases.BookUsecase
}

func NewBookHttpHandler(bookUsecase usecases.BookUsecase) *bookHttpHandler {
	return &bookHttpHandler{
		bookUsecase: bookUsecase,
	}
}

func (h *bookHttpHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var reqBody models.BookRequest
	w.Header().Set("Content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		log.Printf("Error :%v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.bookUsecase.CreateBook(&reqBody)
	if err != nil {
		log.Printf("Error :%v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Book created!"))
}

func (h *bookHttpHandler) SearchBookById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	w.Header().Set("Content-type", "application/json")

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Error :%v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book, err := h.bookUsecase.FindBook(id)
	if err != nil {
		switch err.Error() {
		case "0":
			log.Printf("Error :%v", "ID Not Found")
			http.Error(w, "ID Not Found", http.StatusNotFound)
			return
		default:
			log.Printf("Error :%v", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&book)
}

func (h *bookHttpHandler) DeleteBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Error :%v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.bookUsecase.DeleteBook(id)
	if err != nil {
		switch err.Error() {
		case "0":
			log.Printf("Error :%v", "ID Not Found")
			http.Error(w, "ID Not Found", http.StatusNotFound)
			return
		default:
			log.Printf("Error :%v", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(200)
	w.Write([]byte("Sucess deleted book!"))
}

func (h *bookHttpHandler) UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var book models.BookRequest
	w.Header().Set("Content-type", "application/json")
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Error :%v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Printf("Error :%v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.bookUsecase.EditBook(id, &book)
	if err != nil {
		switch err.Error() {
		case "0":
			log.Printf("Error :%v", "ID not found")
			http.Error(w, "ID Not found", http.StatusNotFound)
			return
		default:
			log.Printf("Error :%v", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(200)
	w.Write([]byte("Sucess updated book!"))
}

func (h *bookHttpHandler) FindAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	param := r.URL.Query().Get("order")
	book, err := h.bookUsecase.FindAllBook(param)
	if err != nil {
		log.Printf("Erorr : %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&book)
}
