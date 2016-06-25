package handler

import (
	"fmt"
	"net/http"
)

// PostMatch is a handler to get index
func PostMatch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PostMatch")
}

// PutResult is a handler to get index
func PutResult(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PutResult")
}
