package config

import (
	"log"
	"testing"

	_ "github.com/lib/pq"
)

// Tes function s
var (
	postgres = &PostgresDB{
		Host:     "localhost",
		Username: "isaachx",
		Password: "saydimas78",
		Port:     "5432",
		SslMode:  "disable",
		DbName:   "database-golang"}
	postgresMustEqual string = "postgres://isaachx:saydimas78@localhost:5432/database-golang?sslmode=disable"
)

func TestConnectionString(t *testing.T) {
	_, stringPattern := postgres.stringConnection()
	//method yang digunakan untuk mencetak nilai parameter menjadi log error
	t.Logf("Connection string pattern : %s", stringPattern)
	if stringPattern != postgresMustEqual {
		//Metthod yang digunakan untuk memunculkan log dan diikuti keterangan fail pada testing
		t.Errorf("Failed! string must equal! %s", postgresMustEqual)
	}
}

func TestConnectionToPostgres(t *testing.T) {
	db, err := Connect(postgres)
	if err != nil {
		t.Fatalf("Failed to connect to PostgreSQL: %v", err) // Hentikan pengujian jika koneksi gagal
	}
	defer db.Close() // Pastikan koneksi ditutup setelah selesai

	// Pastikan koneksi berhasil dengan memeriksa statistik
	stats := db.Ping()
	t.Logf("Connection successful! Open connections")
	if stats != nil {
		t.Errorf("Connection database failed")
	}
}

//Benchmark database

func BenchmarkDatabase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		db, err := Connect(postgres)

		if err != nil {
			log.Fatal(err.Error()) // Hentikan pengujian jika koneksi gagal
			return
		}
		defer db.Close() // Pastikan koneksi ditutup setelah selesai
	}

}
