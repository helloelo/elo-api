package app

import (
	"elo-api/internal/handler"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	router.HandleFunc("/", handler.Index)
}
