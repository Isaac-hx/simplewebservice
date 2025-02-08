// Package service for handling author object
package service

import (
	"encoding/json"
	"log"
	"mime"
	"net/http"
	"simplewebservice/config"
	"simplewebservice/internal/author"
	"simplewebservice/library"
	"simplewebservice/utils"
	"strconv"
)

type handlerAuthor struct {
	author *author.Author
}

func NewServeAuthor() *handlerAuthor {
	author := author.New()
	log.Println("berjalan")
	return &handlerAuthor{author: author}
}
func (handler *handlerAuthor) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	utils.LogServer(r)

	//Connect database
	db, err := config.Connect(library.Postgres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	//Set content type
	w.Header().Set("Content-type", "application/json")

	//Querying select data
	data, err := handler.author.GetAuthor(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&data)
}

func (handler *handlerAuthor) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	utils.LogServer(r)

	//request structure
	var authorReq struct {
		Name string `json:"name"`
	}

	//Connect database
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

	//parse data request
	reqBody := json.NewDecoder(r.Body)
	reqBody.DisallowUnknownFields()
	parseReqBody := reqBody.Decode(&authorReq)
	if parseReqBody != nil {
		http.Error(w, "Invalid json format!!", http.StatusBadRequest)
		return
	}

	//Querying create data
	err = handler.author.CreateAuthor(db, authorReq.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	//set content type
	w.Header().Set("Content-type", "application/json")
	var responseBody = map[string]interface{}{"Message": "Success,create data!!"}
	json.NewEncoder(w).Encode(&responseBody)
}

func (handler *handlerAuthor) GetAuthorById(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-type", "application/json")
	data, err := handler.author.GetAuthorById(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&data)

}
func (handler *handlerAuthor) DeleteAuthorById(w http.ResponseWriter, r *http.Request) {
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
	err = handler.author.DeleteAuthorById(db, id)
	if err != nil {
		if err.Error() == "0" {
			http.Error(w, "author not found!!", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-type", "application/json")
	message := map[string]interface{}{"Message": "Sucess deleted author!!"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&message)

}

func (handler *handlerAuthor) UpdateAuthorById(w http.ResponseWriter, r *http.Request) {
	utils.LogServer(r)

	var authorReq struct {
		Name string `json:"name"`
	}
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

	reqBody := json.NewDecoder(r.Body)
	reqBody.DisallowUnknownFields()
	parseReqBody := reqBody.Decode(&authorReq)
	if parseReqBody != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.author.UpdateAuthorById(db, id, authorReq.Name)
	if err != nil {
		if err.Error() == "0" {
			http.Error(w, "Book not updated!", http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	message := map[string]interface{}{"Message": "Succes updated author!!"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}
