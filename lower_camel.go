package strcase

import (
	"strings"
)

func ToLowerCamel(s string) string {
	return ToLowerCamelWithInitialisms(s, CommonInitialisms)
}

func ToLowerCamelWithInitialisms(s string, initialisms map[string]bool) string {
	if s == "" {
		return s
	}
	ss := SplitIntoWordsWithInitialisms(s, initialisms)
	for i, s := range ss {
		if i == 0 {
			ss[i] = strings.ToLower(s)
			continue
		}
		if CommonInitialisms[strings.ToLower(s)] {
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
