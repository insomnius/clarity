package noface

import (
	sastrawi "github.com/RadhiFadlillah/go-sastrawi"
)

func Tokenizer(sentence string) []string {
	tokenizer := sastrawi.NewTokenizer()
	return tokenizer.Tokenize(sentence)
}
