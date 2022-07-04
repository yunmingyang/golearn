package word

import (
	"unicode"
)


func IsPalindrom(s string) bool {
	// var letters []rune
	letters := make([]rune, 0, len(s))

	for _, s := range s {
		if unicode.IsLetter(s) {
			letters = append(letters, unicode.ToLower(s))
		}
	}

	// for i := range letters {
	// 	if letters[i] != letters[len(letters) - i - 1] {
	// 		return false
	// 	}
	// }
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters) - i - 1] {
			return false
		}
	}

	return true
}