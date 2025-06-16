package main

import (
	"log"
	"net/http"

	api "github.com/ajmacd/wordy/server"
)

func main() {
	mux := NewRouter()
	handler := WithCORS(mux)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/puzzle", api.PuzzleHandler)
	mux.HandleFunc("/api/found", api.FoundHandler)
	return mux
}

func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
