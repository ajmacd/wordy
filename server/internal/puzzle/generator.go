package puzzle

// Generate returns a deterministic 10×10 puzzle and word list.
func Generate() PuzzleResponse {
	return PuzzleResponse{
		Grid: [][]string{
			{"C", "A", "T", "U", "D", "A", "X", "I", "H", "H"},
			{"E", "X", "D", "D", "O", "G", "V", "X", "R", "C"},
			{"S", "B", "I", "R", "D", "N", "B", "A", "C", "G"},
			{"H", "Q", "T", "A", "R", "G", "F", "W", "U", "W"},
			{"R", "N", "H", "O", "S", "I", "I", "Z", "A", "Y"},
			{"Z", "F", "W", "N", "K", "I", "S", "E", "G", "Y"},
			{"K", "D", "C", "M", "D", "L", "H", "L", "T", "I"},
			{"Z", "B", "X", "O", "R", "D", "M", "C", "R", "J"},
			{"U", "T", "L", "S", "G", "W", "C", "B", "V", "H"},
			{"Y", "J", "C", "H", "D", "M", "I", "O", "U", "L"},
		},
		Words: []string{"CAT", "DOG", "BIRD", "FISH"},
	}
}
