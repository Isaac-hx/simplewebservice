// Package service for handling author object
package service

import (
	"database/sql"
	"encoding/json"
	"mime"
	"net/http"
	"simplewebservice/internal/author"
	"simplewebservice/utils"
	"strconv"
)

type handlerAuthor struct {
	author *author.Author
	db     *sql.DB
}

func NewServeAuthor(db *sql.DB) *handlerAuthor {
	author := author.New()
	return &handlerAuthor{author: author, db: db}
}
func (handler *handlerAuthor) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	utils.LogServer(r)

	//Connect database
	defer handler.db.Close()
	//Set content type
	w.Header().Set("Content-type", "application/json")

	//Querying select data
	data, err := handler.author.GetAuthor(handler.db)
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

	//parse data request
	reqBody := json.NewDecoder(r.Body)
	reqBody.DisallowUnknownFields()
	parseReqBody := reqBody.Decode(&authorReq)
	if parseReqBody != nil {
		http.Error(w, "Invalid json format!!", http.StatusBadRequest)
		return
	}

	//Querying create data
	err = handler.author.CreateAuthor(handler.db, authorReq.Name)
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

	//Connect database
	defer handler.db.Close()
	reqId := r.PathValue("id")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	data, err := handler.author.GetAuthorById(handler.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&data)

}
func (handler *handlerAuthor) DeleteAuthorById(w http.ResponseWriter, r *http.Request) {
	utils.LogServer(r)

	//Connect database

	defer handler.db.Close()
	reqId := r.PathValue("id")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = handler.author.DeleteAuthorById(handler.db, id)
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
	//Connect database

	defer handler.db.Close()
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
		http.Error(w, "Error while parsing body!", http.StatusBadRequest)
		return
	}

	err = handler.author.UpdateAuthorById(handler.db, id, authorReq.Name)
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
