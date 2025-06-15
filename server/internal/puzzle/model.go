package puzzle

// PuzzleResponse is the JSON returned by GET /puzzle.
type PuzzleResponse struct {
	Grid  [][]string `json:"grid"`
	Words []string   `json:"words"`
}

// FoundRequest is the JSON expected by POST /found.
type FoundRequest struct {
	Word string `json:"word"`
}
