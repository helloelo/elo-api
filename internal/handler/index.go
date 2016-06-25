package handler

import (
	"fmt"
	"net/http"
)

// Index is a handler to get index
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Elo")
}
