// Generate returns a deterministic 10×10 puzzle and word list.
//
//	func Generate() PuzzleResponse {
//		return PuzzleResponse{
//			Grid: [][]string{
//				{"C", "A", "T", "U", "D", "A", "X", "I", "H", "H"},
//				{"E", "X", "D", "D", "O", "G", "V", "X", "R", "C"},
//				{"S", "B", "I", "R", "D", "N", "B", "A", "C", "G"},
//				{"H", "Q", "T", "A", "R", "G", "F", "W", "U", "W"},
//				{"R", "N", "H", "O", "S", "I", "I", "Z", "A", "Y"},
//				{"Z", "F", "W", "N", "K", "I", "S", "E", "G", "Y"},
//				{"K", "D", "C", "M", "D", "L", "H", "L", "T", "I"},
//				{"Z", "B", "X", "O", "R", "D", "M", "C", "R", "J"},
//				{"U", "T", "L", "S", "G", "W", "C", "B", "V", "H"},
//				{"Y", "J", "C", "H", "D", "M", "I", "O", "U", "L"},
//			},
//			Words: []string{"CAT", "DOG", "BIRD", "FISH"},
//		}
//	}
package puzzle

import (
	"fmt"
	"math/rand"
)

const (
	gridSize      = 10
	maxPlaceTries = 100
)

// Direction vectors for 8 possible orientations
var directions = [8][2]int{
	{0, 1},   // right
	{1, 0},   // down
	{0, -1},  // left
	{-1, 0},  // up
	{1, 1},   // down-right
	{1, -1},  // down-left
	{-1, 1},  // up-right
	{-1, -1}, // up-left
}

// Generate constructs a grid with the given words placed randomly.
func Generate() (PuzzleResponse, error) {
	// Fetch random words
	const wordCount = 4
	words, err := fetchRandomWords(wordCount)
	if err != nil {
		return PuzzleResponse{}, fmt.Errorf("Generate: %w", err)
	}

	// Initialize empty grid with placeholders
	grid := make([][]rune, gridSize)
	for i := range grid {
		grid[i] = make([]rune, gridSize)
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}

	// Place each word
	for _, w := range words {
		placed := false
		for try := 0; try < maxPlaceTries && !placed; try++ {
			dir := directions[rand.Intn(len(directions))]
			r := rand.Intn(gridSize)
			c := rand.Intn(gridSize)
			if canPlace(grid, w, r, c, dir) {
				placeWord(grid, w, r, c, dir)
				placed = true
			}
		}
		if !placed {
			return PuzzleResponse{}, fmt.Errorf("could not place word %q after %d tries", w, maxPlaceTries)
		}
	}

	// Fill empty cells with random letters
	for i := range gridSize {
		for j := range gridSize {
			if grid[i][j] == 0 {
				grid[i][j] = rune(rand.Intn(26) + 'A')
			}
		}
	}

	// Convert rune grid to string grid
	strGrid := make([][]string, gridSize)
	for i := range grid {
		strGrid[i] = make([]string, gridSize)
		for j, r := range grid[i] {
			strGrid[i][j] = string(r)
		}
	}

	return PuzzleResponse{Grid: strGrid, Words: words}, nil
}

// canPlace checks if word w can fit into grid starting at (r,c) moving in dir
func canPlace(grid [][]rune, w string, r, c int, dir [2]int) bool {
	dr, dc := dir[0], dir[1]
	for _, ch := range w {
		// Check bounds
		if r < 0 || r >= gridSize || c < 0 || c >= gridSize {
			return false
		}
		// Check conflict: either empty or same letter
		if grid[r][c] != 0 && grid[r][c] != ch {
			return false
		}
		r += dr
		c += dc
	}
	return true
}

// placeWord writes word w into grid at (r,c) moving in dir
func placeWord(grid [][]rune, w string, r, c int, dir [2]int) {
	dr, dc := dir[0], dir[1]
	for _, ch := range w {
		grid[r][c] = ch
		r += dr
		c += dc
	}
}
