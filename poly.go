package polindrom

import (
	"unicode"
)

func IsPolindrom(str []rune) bool {
	i := 0
	j := len(str) - 1
	for {
		if j < i {
			return true
		}
		if !unicode.IsLetter(str[i]) {
			i++
			continue
		}
		if !unicode.IsLetter(str[j]) {
			j--
			continue
		}
		if unicode.ToLower(str[i]) != unicode.ToLower(str[j]) {
			return false
		}
		i++
		j--
	}
	return false
}
