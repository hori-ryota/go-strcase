package strcase

import (
	"unicode"
)

func SplitIntoWords(s string) []string {
	if s == "" {
		return []string{}
	}
	if len(s) == 1 {
		return []string{s}
	}
	tt := []string{}
	d := 0
	runes := []rune(s)
	for i := 1; i < len(runes); i++ {
		switch {
		case !unicode.IsLetter(runes[i]):
			tt = append(tt, string(runes[d:i]))
			i++
			d = i
		case unicode.IsUpper(runes[i]) && unicode.IsLower(runes[i-1]):
			tt = append(tt, string(runes[d:i]))
			d = i
		case unicode.IsUpper(runes[i]) && unicode.IsUpper(runes[i-1]):
			if i != len(runes)-1 && unicode.IsLower(runes[i+1]) {
				tt = append(tt, string(runes[d:i]))
				d = i
			}
		}
	}
	tt = append(tt, string(runes[d:]))
	return tt
}
