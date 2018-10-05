package clarity

import (
	"strings"
)

func CleansingEarlyBlankSpace(sentence string) (string, string) {
	sentence = regexEarlyBlankSpace.ReplaceAllString(sentence, "")

	return "cleansing_early_blank_space", sentence
}

func CleansingExceesBlankSpace(sentence string) (string, string) {
	sentence = regexExceesBlankSpace.ReplaceAllString(sentence, " ")

	return "cleansing_excees_blank_space", sentence
}

func CleansingDuplicates(sentence []string) (string, string) {
	entry := make(map[string]bool)
	unique_word := []string{}

	for _, word := range sentence {
		if _, unique := entry[word]; !unique {
			entry[word] = true
			unique_word = append(unique_word, word)
		}
	}

	cleansed_sentences := strings.Join(unique_word, " ")

	return "cleansing_duplicates", cleansed_sentences
}
