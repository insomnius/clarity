package clarity

import (
	"strings"
)

func Casefold(sentence string) string {
	return strings.ToLower(sentence)
}
