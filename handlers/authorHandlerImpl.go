package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"simplewebservice/models"
	"simplewebservice/usecases"
	"strconv"
)

type authorHttpHandler struct {
	authorUsecase usecases.AuthorUsecase
}

func NewAuthorHttpHandler(authorUsecase usecases.AuthorUsecase) *authorHttpHandler {
	return &authorHttpHandler{authorUsecase: authorUsecase}

}

func (a *authorHttpHandler) AddAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.AuthorRequest
	w.Header().Set("Content-type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		log.Printf("Error :%v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	err = a.authorUsecase.CreateAuthor(&author)
	if err != nil {
		log.Printf("Error : %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Succes created author!"))
}

func (a *authorHttpHandler) SearchAuthorById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Error : %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	authors, err := a.authorUsecase.FindAuthor(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			log.Printf("Error :%v", err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		default:
			log.Printf("Error :%v", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&authors)

}

func (a *authorHttpHandler) FindAllAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	param := r.URL.Query().Get("order")
	authors, err := a.authorUsecase.FindAllAuthor(param)
	if err != nil {
		switch err.Error() {
		case "0":
			log.Printf("Error : %v", err.Error())
			http.Error(w, "Invalid query params", http.StatusBadRequest)
			return
		default:
			log.Printf("Error : %v", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(authors)
}

func (a *authorHttpHandler) DeleteAuthorById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Erorr :%v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = a.authorUsecase.DeleteAuthor(id)
	if err != nil {
		switch err.Error() {
		case "0":
			log.Printf("Error : %v", "Id not found")
			http.Error(w, "ID not found", http.StatusNotFound)
			return
		default:
			log.Printf("Error : %v", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
	}

	w.WriteHeader(200)
	w.Write([]byte("Success deleted author!"))
}

func (a *authorHttpHandler) UpdateAuthorById(w http.ResponseWriter, r *http.Request) {
	var author models.AuthorRequest
	w.Header().Set("Content-type", "application/json")
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Error : %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		log.Printf("Error : %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = a.authorUsecase.UpdateAuthor(id, &author)
	if err != nil {
		switch err.Error() {
		case "0":
			log.Printf("Err :%v", "id not found")
			http.Error(w, "id not found", http.StatusNotFound)
			return
		default:
			log.Printf("Error : %v", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(200)
	w.Write([]byte("Succes updated author"))
}
