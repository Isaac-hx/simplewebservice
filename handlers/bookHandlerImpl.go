package handlers

import (
	"encoding/json"
	"net/http"
	"simplewebservice/helper"
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
	//request logic
	var reqBody models.BookRequest
	w.Header().Set("Content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		helper.ResponseError(w, helper.NewCustomError(err.Error(), 400, err))
		return
	}

	err = h.bookUsecase.CreateBook(&reqBody)

	//response logic
	if err != nil {
		helper.ResponseError(w, err)
		return
	}

	helper.ResponseSucces(w, map[string]any{"Message": "Success created book!"})
}

func (h *bookHttpHandler) SearchBookById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	w.Header().Set("Content-type", "application/json")

	id, err := strconv.Atoi(idString)
	if err != nil {
		helper.ResponseError(w, helper.ErrInvalidPathValue)
		return
	}
	book, err := h.bookUsecase.FindBook(id)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}

	helper.ResponseSucces(w, map[string]any{"Message": "Success get book", "book": book})
}

func (h *bookHttpHandler) DeleteBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		helper.ResponseError(w, helper.ErrInvalidPathValue)
		return
	}
	err = h.bookUsecase.DeleteBook(id)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}

	helper.ResponseSucces(w, map[string]any{"Message": "Success deleted book!"})
}

func (h *bookHttpHandler) UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var book models.BookRequest
	w.Header().Set("Content-type", "application/json")
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		helper.ResponseError(w, helper.ErrInvalidPathValue)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		helper.ResponseError(w, helper.NewCustomError(err.Error(), 400, err))
		return
	}
	err = h.bookUsecase.EditBook(id, &book)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}
	helper.ResponseSucces(w, map[string]any{"Message": "Success updated book!"})
}

func (h *bookHttpHandler) FindAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	param := r.URL.Query().Get("order")
	book, err := h.bookUsecase.FindAllBook(param)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}
	helper.ResponseSucces(w, map[string]any{
		"Message": "Sucess get book!",
		"Book":    &book,
	})
}
