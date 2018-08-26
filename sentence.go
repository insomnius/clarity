package clarity

type Sentence struct {
	VagueSentence   string
	ClaritySentence string
	AlterHistory    []History
}

func (s *Sentence) Alter() {

}
