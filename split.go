package strcase

import (
	"strings"
	"unicode"
)

// SplitIntoWords splits string to words.
func SplitIntoWords(s string) []string {
	return SplitIntoWordsWithInitialisms(s, CommonInitialisms)
}

// SplitIntoWords splits string to words with initialisms.
func SplitIntoWordsWithInitialisms(s string, initialisms map[string]bool) []string {
	if s == "" {
		return []string{}
	}
	if len(s) == 1 {
		return []string{s}
	}
	if initialisms == nil {
		initialisms = map[string]bool{}
	}
	words := []string{}
	ss := strings.FieldsFunc(s, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})
	for _, s := range ss {
		if s == "" {
			continue
		}
		runes := []rune(s)
		d := 0
		for i := 0; i < len(runes); i++ {
			switch {
			case i+1 < len(runes) && !unicode.IsUpper(runes[i]) && unicode.IsUpper(runes[i+1]):
				words = append(words, string(runes[d:i+1]))
				d = i + 1
				continue
			case i+2 < len(runes) && unicode.IsUpper(runes[i+1]) && unicode.IsLower(runes[i+2]):
				if initialisms[strings.ToLower(string(runes[d:i+2]))] && runes[i+2] == 's' {
					continue
				}
				words = append(words, string(runes[d:i+1]))
				d = i + 1
				continue
			}
		}
		words = append(words, string(runes[d:]))
	}
	return words
}
