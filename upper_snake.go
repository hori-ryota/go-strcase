package strcase

import (
	"strings"
)

func ToUpperSnake(s string) string {
	if s == "" {
		return s
	}
	ss := SplitIntoWords(s)
	for i, s := range ss {
		ss[i] = strings.ToUpper(s)
	}
	return strings.Join(ss, "_")
}
