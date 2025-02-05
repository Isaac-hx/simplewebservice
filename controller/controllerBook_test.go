package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

var url = "http://localhost:8085/book"
var client = &http.Client{}

// Body JSON yang valid
var bodyReqPost = []byte(`{
		"title":"Dunia asd",
		"author":"Jostein Gaarder",
		"total_page":500,
		"publisher":"Gramdedia"
	}`)
var bodyReqPut = []byte(`{
		"title":"3475 MDPL",
		"author":"Stein Hawk",
		"total_page":200,
		"publisher":"Kompas Koran"
	}`)

type output map[string]interface{}

func TestPostBook(t *testing.T) {

	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyReqPost))
	if err != nil {
		t.Fatal(err)
	}

	// Kirim permintaan
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer res.Body.Close() // Pastikan respons body ditutup

	// Log respons
	t.Logf("Response returned %v", res)

	var dataResult output
	// Decode body respons
	err = json.NewDecoder(res.Body).Decode(&dataResult)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err.Error())
	}

	// Log sukses decode JSON
	t.Logf("Success JSON decode: %v", dataResult)

	// Validasi status kode dan pesan respons
	if res.StatusCode != 200 {
		t.Fatalf("Failed! Status code: %v, message: %v", res.StatusCode, dataResult["message"])
	}
}

func TestGetAllBooks(t *testing.T) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer res.Body.Close() // Pastikan respons body ditutup

	var dataResult []output
	err = json.NewDecoder(res.Body).Decode(&dataResult)
	if err != nil {
		t.Fatal("Failed to decode json format!!", err.Error())
	}

	t.Logf("Success JSON decode: %v", dataResult)
	if res.StatusCode != 200 {
		t.Fatalf("Failed! Status code: %v, message: %v", res.StatusCode, dataResult)
	}

}

func TestUpdateBookById(t *testing.T) {

	req, err := http.NewRequest("PUT", url+"?id=12", bytes.NewReader(bodyReqPut))
	if err != nil {
		t.Fatal(err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer res.Body.Close() // Pastikan respons body ditutup

	// Log respons
	t.Logf("Response returned %v", res)

	// Decode body respons
	var dataResult output
	err = json.NewDecoder(res.Body).Decode(&dataResult)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err.Error())
	}

	// Log sukses decode JSON
	t.Logf("Success JSON decode: %v", dataResult)

	// Validasi status kode dan pesan respons
	if res.StatusCode != 200 {
		t.Fatalf("Failed! Status code: %v, message: %v", res.StatusCode, dataResult["message"])
	}

}

func TestDeleteBookById(t *testing.T) {

	req, err := http.NewRequest("DELETE", url+"?id=6", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer res.Body.Close() // Pastikan respons body ditutup

	// Log respons
	t.Logf("Response returned %v", res)

	// Decode body respons
	var dataResult output
	err = json.NewDecoder(res.Body).Decode(&dataResult)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err.Error())
	}

	// Log sukses decode JSON
	t.Logf("Success JSON decode: %v", dataResult)

	// Validasi status kode dan pesan respons
	if res.StatusCode != 200 {
		t.Fatalf("Failed! Status code: %v, message: %v", res.StatusCode, dataResult["message"])
	}

}
