package main

import (
	"simplewebservice/config"
	"simplewebservice/database"
	"simplewebservice/server"

	_ "github.com/lib/pq"
)

func main() {
	// data, err := utils.LoadEnv()
	// if err != nil {
	// 	log.Fatalf("Error from load .env %v", err.Error())

	// }
	// //initialising object database

	// conn, err := config.NewConnection(&data)
	// if err != nil {
	// 	log.Fatalf("Error from initializing database %v", err.Error())
	// }

	// //initialising object servermux
	// mux := http.NewServeMux()
	// //registered route
	// router.AuthorRoute(mux, conn)
	// router.BookRoute(mux, conn)

	// utils.ListRoute("/book", "/author")
	// log.Println("Service berjalan di localhost:8085/")

	// http.ListenAndServe(":8085", mux)

	// Calling constructor getConfig
	conf := config.GetConfig()
	// Inializing database
	db := database.NewDatabasePostgres(conf)
	db.GetStat()
	//Run application
	server.NewServerMux(conf, db).Start()

	// //Migrate book data
	// migrations.BookMigrate(db)
}
