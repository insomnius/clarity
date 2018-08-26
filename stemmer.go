package noface

import (
	"strings"

	sastrawi "github.com/RadhiFadlillah/go-sastrawi"
)

func Stem(sentence string) string {
	words := Tokenizer(sentence)

	stemmer := sastrawi.NewStemmer(sastrawi.DefaultDictionary)
	for key, word := range words {
		words[key] = stemmer.Stem(word)
	}

	return strings.Join(words, " ")
}
