package main

import (
	"log"
	"net/http"

	"github.com/ajmacd/wordy/internal/api"
)

func main() {
	mux := api.NewRouter()
	handler := api.WithCORS(mux)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
