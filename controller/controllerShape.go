package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"simplewebservice/logger"
	"simplewebservice/models"
)

func GetCalculateShape(w http.ResponseWriter, r *http.Request) {
	// Log Server
	defer logger.LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))
	w.Header().Set("Content-type", "application/json")
	body, err := io.ReadAll(r.Body)
	log.Println(body)
	if err != nil {
		http.Error(w, "Error from server !!", http.StatusInternalServerError)
		return
	}

	var data models.Cube
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	//Memanggil fungsi calculate shape dari models
	result := models.CalculateShape(data)
	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func GetRotateShape(w http.ResponseWriter, r *http.Request) {
	// Log Server
	defer logger.LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))
	w.Header().Set("Content-type", "application/json")
	body, err := io.ReadAll(r.Body)
	log.Println(body)
	if err != nil {
		http.Error(w, "Error from server !!", http.StatusInternalServerError)
		return
	}

	var data models.Cube
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	//Memanggil fungsi calculate shape dari models
	result := models.CalculateRotateShape(&data, 45)
	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
