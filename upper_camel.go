package strcase

import (
	"strings"
)

// ToUpperCamel convert string to UpperCamel.
func ToUpperCamel(s string) string {
	return ToUpperCamelWithInitialisms(s, CommonInitialisms)
}

// ToUpperCamelWithInitialisms convert string to UpperCamel with initialisms.
func ToUpperCamelWithInitialisms(s string, initialisms map[string]bool) string {
	if s == "" {
		return s
	}
	if initialisms == nil {
		initialisms = map[string]bool{}
	}
	ss := SplitIntoWordsWithInitialisms(s, initialisms)
	for i, s := range ss {
		if initialisms[strings.ToLower(s)] {
			ss[i] = strings.ToUpper(s)
			continue
		}
		if strings.ToLower(s[len(s)-1:]) == "s" && initialisms[strings.ToLower(s[:len(s)-1])] {
			ss[i] = strings.ToUpper(s[:len(s)-1]) + "s"
			continue
		}
		ss[i] = strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
	}
	return strings.Join(ss, "")
}
