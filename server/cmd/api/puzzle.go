package handler

import (
	"net/http"

	api "github.com/ajmacd/wordy/server"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	api.PuzzleHandler(w, r)
}
