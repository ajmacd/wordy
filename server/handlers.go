package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ajmacd/wordy/server/internal/puzzle"
)

// PuzzleHandler returns a static puzzle.
func PuzzleHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := puzzle.Generate()
	if err != nil {
		log.Printf("Error generating puzzle: %v\n", err)
		http.Error(w, "Puzzle generation failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// FoundHandler logs when a client reports a found word.
func FoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var req puzzle.FoundRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	log.Printf("Word found: %s\n", req.Word)
	w.WriteHeader(http.StatusNoContent)
}
