package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const Path string = "/tmp/simplewebservice.log"

func LogServer(r *http.Request) {
	//Informasi client ip address
	clientIp := r.RemoteAddr
	//Informasi komputer client
	userAgent := r.Header.Get("User-Agent")
	//Informasi referer
	referer := r.Header.Get("Referer")
	//Informasi method yang digunakan
	method := r.Method
	//Infromasi path yang digunakan
	url := r.URL
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logPattern := fmt.Sprintf("[%s] %s %s %s %s %s ", timestamp, method, url, clientIp, userAgent, referer)
	//Ambil argumen dari package library
	WriteFileWithOS(Path, logPattern)
	log.Println(logPattern)
}

func WriteFileWithOS(path, text string) {
	//Membuka file atau membuat file jika tidak ditemukan
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal("Error : Tidak ditemukan file log.txt", err.Error())
	}
	_, err = file.WriteString(fmt.Sprintf("%s\n", text))
	if err != nil {
		log.Fatal("Error : Kesalahan dalam menulis log!", err.Error())
	}

	err = file.Sync()
	if err != nil {
		log.Fatal("Error : Kesalahan dalam menyimpan file!", err.Error())
	}

}

func ListRoute(dataRoute ...string) {
	fmt.Println("List route registered : ")
	for _, data := range dataRoute {
		fmt.Println(data)
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while reading .env file", err.Error())
	}
}

// parsing data from request
func ParseTimeDate(dateString string) (time.Time, error) {
	layoutFormat := "2006-01-02"
	date, err := time.Parse(layoutFormat, dateString)
	if err != nil {
		return date, err
	}
	return date, err
}
