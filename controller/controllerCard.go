package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"simplewebservice/models"
	"simplewebservice/utils"
)

func CreateCardIdentity(w http.ResponseWriter, r *http.Request) {
	//Log Server
	defer utils.LogServer(r)
	w.Header().Set("Content-type", "application/json")
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	response, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error from server !!", http.StatusInternalServerError)
		return
	}
	var data map[string]interface{}
	if err := json.Unmarshal(response, &data); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}
	switch data["card"].(string) {
	case "KTP":
		var ktp models.KTP
		ktp.FirstName = data["first_name"].(string)
		ktp.LastName = data["last_name"].(string)
		ktp.Gender = data["gender"].(bool)
		ktp.NumberCivilization = int(data["nik"].(float64))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(models.GetCardInformation(ktp)))
		return
	case "SIM":
		var sim models.SIM
		sim.FirstName = data["first_name"].(string)
		sim.LastName = data["last_name"].(string)
		sim.Gender = data["gender"].(bool)
		sim.NumberDriver = int(data["nik"].(float64))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(models.GetCardInformation(sim)))
		return

	default:
		http.Error(w, "Card not found", http.StatusNotFound)
		return
	}
}
