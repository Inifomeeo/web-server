package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}