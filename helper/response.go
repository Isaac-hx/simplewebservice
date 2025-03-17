package helper

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func ResponseError(w http.ResponseWriter, err error) {
	log.Printf("Error :%v", err.Error())
	var errorData *CustomError
	if errors.As(err, &errorData) {
		http.Error(w, errorData.Error(), errorData.StatusCode)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func ResponseSucces(w http.ResponseWriter, args map[string]any) {
	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode(args)
	if err != nil {
		log.Fatalf("Error from : %v", err.Error())
	}
}
