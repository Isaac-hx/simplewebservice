package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

const path string = "/home/isaaachx/Documents/log.txt"

func LogServer(textLog string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	timeLog := fmt.Sprintf("[%s] %s", timestamp, textLog)
	WriteFileWithOS(path, timeLog)
	log.Println(textLog)
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
