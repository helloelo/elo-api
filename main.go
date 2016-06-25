package main

import (
	"fmt"
	"log"
	"net/http"

	"elo-api/internal/app"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	app.InitRoutes(router)

	fmt.Println("Listenning on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
