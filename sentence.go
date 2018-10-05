package clarity

var with_history bool
var sentence_history SentenceHistory

type Sentence struct {
	VagueSentence   string `json:"vague_sentence"`
	ClaritySentence string `json:"clarity_sentence"`
}

type SentenceHistory struct {
	AlterHistory []History `json:"alter_history"`
}

func (s *Sentence) Alter() *Sentence {
	s.perform(false)

	return s
}

func (s *Sentence) AlterWithDebug() *Sentence {
	s.perform(true)

	return s
}

func (s *Sentence) Tokenize() []string {
	return Tokenizer(s.ClaritySentence)
}

func (s *Sentence) History() SentenceHistory {
	return sentence_history
}

func (s *Sentence) perform(debug bool) {
	with_history = debug

	s.ClaritySentence = s.VagueSentence

	s.casefold()
	s.stemming()
	s.removeStopword()
	s.normaliseBlankSpace()
	s.removeDuplicates()
}

func (s *Sentence) casefold() {
	process, word := Casefold(s.ClaritySentence)
	s.addHistory(s.ClaritySentence, word, process)
	s.ClaritySentence = word
}

func (s *Sentence) stemming() {
	process, word := Stem(s.ClaritySentence)
	s.addHistory(s.ClaritySentence, word, process)
	s.ClaritySentence = word
}

func (s *Sentence) removeStopword() {
	process, word := Stopwording(s.ClaritySentence)
	s.addHistory(s.ClaritySentence, word, process)
	s.ClaritySentence = word
}

func (s *Sentence) normaliseBlankSpace() {
	process, word := CleansingEarlyBlankSpace(s.ClaritySentence)
	s.addHistory(s.ClaritySentence, word, process)
	s.ClaritySentence = word

	process, word = CleansingExceesBlankSpace(s.ClaritySentence)
	s.addHistory(s.ClaritySentence, word, process)
	s.ClaritySentence = word
}

func (s *Sentence) removeDuplicates() {
	tokenized_sentence := s.Tokenize()

	process, word := CleansingDuplicates(tokenized_sentence)
	s.addHistory(s.ClaritySentence, word, process)
	s.ClaritySentence = word
}

func (s *Sentence) addHistory(vague string, clarified string, process string) {
	if with_history {
		sentence_history.AlterHistory = append(sentence_history.AlterHistory, History{Word: Word{VagueWord: vague, ClarifiedWord: clarified}, Process: process})
	}
}
