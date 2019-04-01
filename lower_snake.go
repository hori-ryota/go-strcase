package strcase

import (
	"strings"
)

func ToLowerSnake(s string) string {
	if s == "" {
		return s
	}
	ss := SplitIntoWords(s)
	for i, s := range ss {
		ss[i] = strings.ToLower(s)
	}
	return strings.Join(ss, "_")
}
