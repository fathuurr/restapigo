package main

import (
	"log"
	"net/http"

	"restapigo/db"
	"restapigo/routes"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	router := routes.SetupRoutes(database)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
