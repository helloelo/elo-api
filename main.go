package main

import (
	"fmt"
	"log"
	"net/http"

	"elo-api/internal/app"
	"elo-api/internal/config"
)

func main() {

	router := app.InitRouter()
	config.InitConfig()

	fmt.Println("Listenning on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
