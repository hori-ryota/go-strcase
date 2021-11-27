package strcase

import (
	"strings"
)

// ToUpperSnake convert string to UPPER_SNAKE.
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
