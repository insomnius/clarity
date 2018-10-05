package clarity

import (
	"strings"
)

func Casefold(sentence string) (string, string) {
	return "casefolding", strings.ToLower(sentence)
}
