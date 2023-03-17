package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {

	log.Println("hello")

	app := application{}

	mux := app.routes()

	log.Println("Starting Server on port 8080...")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}
