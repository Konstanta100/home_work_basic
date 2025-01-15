package utilword

import (
	"strings"
)

func CountWords(s string) map[string]int {
	dictionary := make(map[string]int)

	specSymbols := [...]string{"!", ",", "?", "\\", "/", ";", ".", "-", "\n", "\r", "\t", ":", "(", ")"}

	for _, specSymbol := range specSymbols {
		s = strings.ReplaceAll(s, specSymbol, " ")
	}

	slice := strings.Split(s, " ")

	for _, word := range slice {
		if len(word) == 0 {
			continue
		}

		word = strings.ToLower(word)
		dictionary[word]++
	}

	return dictionary
}
