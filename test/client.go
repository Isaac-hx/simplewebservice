package test

import (
	"encoding/json"
	"net/http"
	"simplewebservice/models"

	_ "github.com/lib/pq"
)

var baseUrl = "http://localhost:8085"

func fetchBooks() ([]models.Book, error) {
	var err error
	//Mendeklarasikan nilai dereference instance object http.client
	var client = &http.Client{}
	var data []models.Book

	//Memanggil fungsi http.NewRequest yang mengembalikan nilai pointer reference http.Request
	//Parameter pertama,berisikan tipe method yang akan digunakan seperti `POST` atau `GET`
	//Parameter kedua,adalah tujuan url yang ingin di request
	//Parameter ketiga,adalah form data request jika terdapat body
	req, err := http.NewRequest("GET", baseUrl+"/books", nil)
	if err != nil {
		return nil, err
	}

	//Memanggil method yang ada pada instance `client` dan mengeksekusi method Do
	//,dengan menyisipkan argumen req yang telah dibuat object sebelumnya pada variabel `req`
	//Method ini akan mengembalikan object http.Response
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//response data yang telah diambil perlu diclose setelah tidak dipakai.
	defer res.Body.Close()

	//data pada variabel `res` yang terdapat pada property body
	//Mendecode data menjadi data bertipe pointer `&data`
	//Perbedaan antara json.marshal dan newdecoder
	//json.Marshal digunakan untuk mengubah nilai Go menjadi representasi []byte,cocok untuk menghasilkan json dan menyimpannya untuk kemudian di olah lebih lanjut..
	//json.Newencoder digunakan untuk membuat object encoder yang menulis JSON ke io.writer,fungsi ini tidak menghasilkan nilai []byte,tetapi langsung menulis ke tujuan yang ditentukan.Cocok digunakan untuk menulis data json langsung ke io.Writer
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil

}
