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
		log.Printf("Error :%v %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&book)
}
