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

	listenAddr := fmt.Sprintf("%s:%d", config.Config.Host, config.Config.Port)

	fmt.Println("Application started on " + listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
