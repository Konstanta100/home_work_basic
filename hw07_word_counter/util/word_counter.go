package utilword

import (
	"strings"
	"unicode"
)

func CountWords(s string) map[string]int {
	dictionary := make(map[string]int)
	specSymbols := make(map[string]int)

	for _, sym := range s {
		if unicode.IsSpace(sym) || unicode.IsPunct(sym) {
			specSymbols[string(sym)]++
		}
	}

	for specSymbol := range specSymbols {
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
