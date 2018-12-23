package clarity

import (
	"encoding/json"
)

type Sentence interface {
	// Alter perform information extraction and retrieval process
	Alter() Sentence

	// VagueSentence return vague sentence of Sentence
	VagueSentence() (vagueSentence string)
	// ClaritySentence return the final result of information and retrieval process
	ClaritySentence() (claritySentence string)

	// AsJson convert sentence into json format
	AsJson() []byte

	// WithHistory will notify the Sentence to store all of transformation history
	WithHistory() Sentence
	// WithHistory will notify the Sentence to not store all of transformation history
	WithoutHistory() Sentence

	// History list of all process with before after of transformed word
	History() []history
}

func createSentence(s string) Sentence {
	return &sentence{vagueSentence: s}
}

type sentence struct {
	vagueSentence     string
	claritySentence   string
	sentenceHistory   []history
	tokenizedSentence Tokenized
	withHistory       bool
}

func (s *sentence) Alter() Sentence {
	s.perform()
	return s
}

func (s *sentence) VagueSentence() string {
	return s.vagueSentence
}

func (s *sentence) ClaritySentence() string {
	return s.claritySentence
}

func (s *sentence) WithHistory() Sentence {
	s.withHistory = true
	return s
}

func (s *sentence) WithoutHistory() Sentence {
	s.withHistory = false
	return s
}

func (s *sentence) AsJson() []byte {
	to_be_encoded := Json{
		"vague_sentence":   s.vagueSentence,
		"clarity_sentence": s.claritySentence,
	}

	if s.withHistory {
		var sentenceHistory []Json
		for _, history := range s.sentenceHistory {
			sentenceHistory = append(sentenceHistory, Json{
				"word":    Json{"before": history.word.before, "after": history.word.after},
				"process": history.process,
			})
		}

		to_be_encoded["history"] = sentenceHistory
	}

	encoded_json, err := json.Marshal(to_be_encoded)
	if err != nil {
		return []byte(``)
	}

	return encoded_json
}

func (s *sentence) perform() {
	s.claritySentence = s.vagueSentence

	s.casefold()
	s.stem()
	s.removeStopword()
	s.normaliseBlankSpace()

	s.tokenizedSentence = tokenize(s.claritySentence)

	s.removeDup()
}

func (s *sentence) History() []history {
	return s.sentenceHistory
}

func (s *sentence) casefold() {
	process, word := casefold(s.claritySentence)
	s.addHistory(s.claritySentence, word, process)
	s.claritySentence = word
}

func (s *sentence) stem() {
	process, word := stem(s.claritySentence)
	s.addHistory(s.claritySentence, word, process)
	s.claritySentence = word
}

func (s *sentence) removeStopword() {
	process, word := stopwording(s.claritySentence)
	s.addHistory(s.claritySentence, word, process)
	s.claritySentence = word
}

func (s *sentence) normaliseBlankSpace() {
	process, word := cleansingEarlyBlankSpace(s.claritySentence)
	s.addHistory(s.claritySentence, word, process)
	s.claritySentence = word

	process, word = cleansingExceesBlankSpace(s.claritySentence)
	s.addHistory(s.claritySentence, word, process)
	s.claritySentence = word
}

func (s *sentence) removeDup() {
	process, word := cleansingDuplicates(s.tokenizedSentence)
	s.addHistory(s.claritySentence, word, process)
	s.claritySentence = word
}

func (s *sentence) addHistory(before string, after string, process string) {
	if s.withHistory == false {
		return
	}

	h := history{process: process}
	h.word.before = before
	h.word.after = after

	s.sentenceHistory = append(s.sentenceHistory, h)
}
