package clarity

func FromSentence(sentence string) Sentence {

	return Sentence{VagueSentence: sentence}
}

func FromParagraph(paragraph string) Paragraph {

	return Paragraph{VagueParagraph: paragraph}
}

func FromArticle(article string) Article {

	return Article{VagueArticle: article}
}
