package puzzle

import (
	"fmt"
	"log"
	"math/rand"
)

const (
	gridSize      = 10
	maxPlaceTries = 100
)

type Direction struct {
	dr int // row delta
	dc int // col delta
}

var directions = []Direction{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
	// {1, 1},   // down-right
	// {1, -1},  // down-left
	// {-1, 1},  // up-right
	// {-1, -1}, // up-left
}

func Generate() (PuzzleResponse, error) {
	// 1) Fetch words
	const wordCount = 5
	const wordLength = 4
	words, err := fetchRandomWords(wordCount, wordLength)
	if err != nil {
		return PuzzleResponse{}, fmt.Errorf("Generate: %w", err)
	}
	log.Printf("Generate: placing words %v\n", words)

	// 2) Create an empty grid (zero runes)
	grid := make([][]rune, gridSize)
	for i := range grid {
		grid[i] = make([]rune, gridSize)
	}

	// 3) Place each word on the empty grid
	for _, w := range words {
		placed := false
		log.Printf("  Word %q: starting placement\n", w)
		for attempt := 1; attempt <= maxPlaceTries; attempt++ {
			// pick dir, then a valid (r,c) so word stays in-bounds
			dir := directions[rand.Intn(len(directions))]
			dr, dc := dir.dr, dir.dc

			r := rand.Intn(gridSize)
			c := rand.Intn(gridSize)

			ok := canPlace(grid, w, r, c, dir)
			log.Printf("    Attempt %2d: r=%d c=%d dir=(%d,%d) canPlace=%v",
				attempt, r, c, dr, dc, ok)

			if !ok {
				continue
			}
			placeWord(grid, w, r, c, dir)
			log.Printf("    Placed %q at (r=%d,c=%d) dir=(%d,%d)",
				w, r, c, dr, dc)
			placed = true
			break
		}
		if !placed {
			return PuzzleResponse{}, fmt.Errorf(
				"Generate: failed to place word %q after %d tries",
				w, maxPlaceTries,
			)
		}
	}

	// 4) Fill every empty cell with random letters
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				grid[i][j] = rune(rand.Intn(26) + 'A')
			}
		}
	}

	// 5) Convert to string grid and return
	strGrid := make([][]string, gridSize)
	for i := range grid {
		strGrid[i] = make([]string, gridSize)
		for j, r := range grid[i] {
			strGrid[i][j] = string(r)
		}
	}
	return PuzzleResponse{Grid: strGrid, Words: words}, nil
}

func canPlace(grid [][]rune, w string, r, c int, dir Direction) bool {
	dr, dc := dir.dr, dir.dc
	for _, ch := range w {
		// bounds
		if r < 0 || r >= gridSize || c < 0 || c >= gridSize {
			return false
		}
		// collision: cell must be empty or match this letter
		if grid[r][c] != 0 && grid[r][c] != ch {
			return false
		}
		r += dr
		c += dc
	}
	return true
}

func placeWord(grid [][]rune, w string, r, c int, dir Direction) {
	dr, dc := dir.dr, dir.dc
	for _, ch := range w {
		grid[r][c] = ch
		r += dr
		c += dc
	}
}
