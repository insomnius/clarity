package noface

import (
	"fmt"

	"github.com/RadhiFadlillah/go-sastrawi"
)

func Test() {
	// Kalimat asal
	sentence := "Rakyat memenuhi halaman gedung untuk menyuarakan isi hatinya. Baca berita selengkapnya di http://www.kompas.com."

	// Pecah kalimat menjadi kata-kata menggunakan tokenizer
	tokenizer := sastrawi.NewTokenizer()
	words := tokenizer.Tokenize(sentence)

	// Ubah kata berimbuhan menjadi kata dasar
	stemmer := sastrawi.NewStemmer(sastrawi.DefaultDictionary)
	for _, word := range words {
		fmt.Printf("%s => %s\n", word, stemmer.Stem(word))
	}
}
