package puzzle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// fetchRandomWords hits the Vercel API for `count` words,
// upper-cases them, and returns the slice or an error.
func fetchRandomWords(count int, length int) ([]string, error) {
	url := fmt.Sprintf("https://random-word-api.vercel.app/api?words=%d&length=%d", count, length)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetchRandomWords: HTTP GET failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetchRandomWords: non-200 status: %d", resp.StatusCode)
	}

	var words []string
	if err := json.NewDecoder(resp.Body).Decode(&words); err != nil {
		return nil, fmt.Errorf("fetchRandomWords: JSON decode failed: %w", err)
	}

	for i, w := range words {
		words[i] = strings.ToUpper(w)
	}
	return words, nil
}
