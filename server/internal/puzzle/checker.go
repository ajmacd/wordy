package puzzle

// CheckWord returns true if `word` is in the provided list.
func CheckWord(word string, list []string) bool {
	for _, w := range list {
		if w == word {
			return true
		}
	}
	return false
}
