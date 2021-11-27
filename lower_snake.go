package strcase

import (
	"strings"
)

// ToLowerSnake convert string to lower_snake.
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
