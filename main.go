package main

import (
	"fmt"
	"log"
	"net/http"

	"elo-api/internal/app"
)

func main() {

	router := app.InitRouter()

	fmt.Println("Listenning on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
