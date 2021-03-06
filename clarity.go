package clarity

import (
	"regexp"
	"strings"

	sastrawi "github.com/RadhiFadlillah/go-sastrawi"
)

func init() {
	newStopwords(
		"ada", "lah", "adanya", "lain", "adalah", "lainnya", "adapun", "melainkan", "agak",
		"selaku", "laku", "agaknya", "lalu", "agar", "melalui", "akan", "terlalu", "akankah",
		"lama", "akhirnya", "akhir", "lamanya", "aku", "pronomia", "selama", "akulah", "amat",
		"selamanya", "amatlah", "lebih", "anda", "terlebih", "andalah", "bermacam", "macam",
		"antar", "diantaranya", "antara", "semacam", "antaranya", "maka", "diantara", "makanya",
		"apa", "makin", "apaan", "malah", "mengapa", "malahan", "apabila", "mampu", "apakah",
		"mampukah", "apalagi", "mana", "pr", "apatah", "manakala", "atau", "manalagi", "ataukah",
		"masih", "ataupun", "masihkah", "bagai", "semasih", "bagaikan", "masing", "sebagai",
		"sebagainya", "mau", "bagaimana", "maupun", "bagaimanapun", "semaunya", "sebagaimana",
		"memang", "bagaimanakah", "bagainamakah", "mereka", "bagi", "merekalah", "bahkan", "meski",
		"bahwa", "meskipun", "bahwasanya", "bahwasannya", "semula", "mula", "sebaliknya", "balik",
		"mungkin", "banyak", "mungkinkah", "sebanyak", "numeralia", "nah", "beberapa", "namun",
		"seberapa", "nanti", "begini", "nantinya", "beginian", "nyaris", "beginikah", "oleh", "beginilah",
		"olehnya", "sebegini", "seorang", "orang", "begitu", "seseorang", "begitukah", "pada", "begitulah",
		"padanya", "begitupun", "padahal", "sebegitu", "paling", "belum", "sepanjang", "panjang", "belumlah",
		"pantas", "sebelum", "sepantasnya", "sebelumnya", "sepantasnyalah", "sebenarnya", "benar", "para",
		"berapa", "pasti", "berapakah", "pastilah", "berapalah", "per", "paticle", "berapapun", "pernah",
		"betulkah", "betul", "pula", "sebetulnya", "pun", "biasa", "merupakan", "rupa", "biasanya", "rupanya",
		"bila", "serupa", "bilakah", "saat", "bisa", "saatnya", "bisakah", "sesaat", "sebisanya", "saja", "boleh",
		"sajalah", "bolehkah", "saling", "bolehlah", "bersama", "sama", "buat", "bukan", "bukankah", "bukanlah",
		"sesama", "bukannya", "sambil", "cuma", "sampai", "percuma", "sana", "dahulu", "sangat", "dalam",
		"sangatlah", "dan", "saya", "dapat", "sayalah", "dari", "se", "daripada", "sebab", "dekat",
		"sebabnya", "demi", "sebuah", "demikian", "tersebut", "sebut", "demikianlah", "tersebutlah",
		"sedemikian", "sedang", "dengan", "sedangkan", "depan", "sedikit", "di", "sedikitnya", "dia",
		"segala", "dialah", "segalanya", "dini", "segera", "diri", "sesegera", "dirinya", "sejak", "terdiri",
		"sejenak", "dong", "sekali", "dulu", "sekalian", "enggak", "sekalipun", "enggaknya", "sesekali", "entah",
		"sekaligus", "entahlah", "sekarang", "terhadap", "hadap", "sekaranglah", "terhadapnya", "sekitar", "hal",
		"sekitarnya", "hampir", "sela", "hanya", "selain", "hanyalah", "selalu", "harus", "seluruh", "numeral",
		"haruslah", "seluruhnya", "harusnya", "semakin", "seharusnya", "sementara", "hendak", "sempat", "hendaklah",
		"semua", "hendaknya", "semuanya", "hingga", "sendiri", "sehingga", "sendirinya", "ia", "seolah",
		"ialah", "ibarat", "seperti", "ingin", "sepertinya", "inginkah", "sering", "inginkan", "seringnya",
		"ini", "serta", "inikah", "siapa", "inilah", "siapakah", "itu", "siapapun", "itukah", "disini", "sini",
		"itulah", "disinilah", "jangan", "jangankan", "sinilah", "janganlah", "sesuatu", "suatu", "jika",
		"sesuatunya", "jikalau", "juga", "sesudah", "sudah", "justru", "sesudahnya", "kala", "kalau", "sudahkah",
		"kalaulah", "sudahlah", "kalaupun", "supaya", "berkali", "kali", "tadi", "tadinya", "kalian", "tak",
		"kami", "tanpa", "kamilah", "setelah", "telah", "kamu", "kamulah", "tentang", "kan", "tentu", "kapan",
		"tentulah", "kapankah", "tentunya", "kapanpun", "tertentu", "dikarenakan", "karena", "seterusnya", "terus",
		"tapi", "tetapi", "karenanya", "ke", "setiap", "tiap", "kecil", "kemudian", "setidak", "tidak", "kenapa",
		"setidaknya", "kepada", "kepadanya", "tidakkah", "ketika", "tidaklah", "seketika", "toh", "khususnya",
		"khusus", "waduh", "kini", "wah", "kinilah", "wahai", "kiranya", "kira", "sewaktu", "waktu", "sekiranya",
		"walau", "kita", "walaupun", "kitalah", "wong", "kok", "yaitu", "lagi", "yakni", "lagian", "yang",
		"berada", "masa", "keadaan", "semasa", "masalah", "akhiri", "masalahnya", "berakhir", "termasuk",
		"masuk", "berakhirlah", "semata", "mata", "berakhirnya", "diakhiri", "diminta", "minta",
		"diakhirinya", "dimintai", "mengakhiri", "meminta", "terakhir", "memintakan", "artinya",
		"arti", "berarti", "mirip", "asal", "dimisalkan", "misal", "asalkan", "memisalkan",
		"atas", "awal", "misalkan", "awalnya", "misalnya", "berawal", "semisal", "berbagai",
		"semisalnya", "bagian", "bermula", "sebagian", "baik", "mulanya", "sebaik", "dimulai",
		"mulai", "dimulailah", "sebaiknya", "dimulainya", "bakal", "memulai", "bakalan", "mulailah",
		"terbanyak", "dimungkinkan", "bapak", "kemungkinan", "baru", "kemungkinannya", "bawah",
		"memungkinkan", "belakang", "menaiki", "naik", "belakangan", "menanti", "benarkah", "benarlah",
		"menantikan", "beri", "menyatakan", "nyata", "berikan", "nyatanya", "diberi", "ternyata",
		"diberikan", "pak", "diberikannya", "memberi", "dipastikan", "memberikan", "memastikan", "besar",
		"penting", "sebesar", "pentingnya", "diperlukan", "perlu", "kebetulan", "diperlukannya", "dibuat",
		"memerlukan", "dibuatnya", "diperbuat", "perlukah", "diperbuatnya", "perlunya", "membuat", "seperlunya",
		"memperbuat", "pertama", "bulan", "bung", "memihak", "pihak", "cara", "caranya", "pihaknya", "secara",
		"sepihak", "cukup", "pukul", "cukupkah", "dipunyai", "punya", "cukuplah", "mempunyai", "secukupnya",
		"terdahulu", "merasa", "rasa", "didapat", "mendapat", "rasanya", "mendapatkan", "terasa", "terdapat",
		"rata", "berdatangan", "datang", "berupa", "disampaikan", "didatangkan", "kesampaian", "mendatang",
		"menyampaikan", "mendatangi", "mendatangkan", "sampaikan", "dua", "sesampai", "kedua", "tersampaikan",
		"keduanya", "menyangkut", "sangkut", "empat", "satu", "seenaknya", "enak", "disebut", "digunakan",
		"guna", "disebutkan", "dipergunakan", "disebutkannya", "menyebutkan", "gunakan", "mempergunakan",
		"sebutlah", "menggunakan", "sebutnya", "hari", "keseluruhan", "berkehendak", "keseluruhannya",
		"menghendaki", "menyeluruh", "diibaratkan", "sendirian", "diibaratkannya", "bersiap", "siap",
		"ibaratkan", "ibaratnya", "mempersiapkan", "mengibaratkan", "menyiapkan", "mengibaratkannya", "ibu",
		"dipersoalkan", "soal", "berikut", "ikut", "mempersoalkan", "berikutnya", "persoalan", "diingat",
		"ingat", "soalnya", "diingatkan", "diketahui", "tahu", "diketahuinya", "mengetahui", "mengingat",
		"mengingatkan", "tahun", "seingat", "ditambahkan", "tambah", "teringat", "menambahkan", "berkeinginan",
		"tambahnya", "diinginkan", "tampak", "keinginan", "tampaknya", "menginginkan", "ditandaskan", "tandas",
		"jadi", "menandaskan", "jadilah", "adjectice", "jadinya", "tandasnya", "menjadi", "bertanya", "tanya",
		"terjadi", "terjadilah", "dipertanyakan", "terjadinya", "ditanya", "jauh", "ditanyai", "sejauh",
		"ditanyakan", "dijawab", "jawab", "mempertanyakan", "menanya", "jawaban", "menanyai", "jawabnya",
		"menanyakan", "menjawab", "pertanyaan", "dijelaskan", "jelas", "pertanyakan", "dijelaskannya",
		"tanyakan", "jelaskan", "tanyanya", "jelaslah", "ditegaskan", "tegas", "jelasnya", "menegaskan",
		"menjelaskan", "berjumlah", "jumlah", "tegasnya", "setempat", "tempat", "jumlahnya", "sejumlah",
		"setengah", "tengah", "sekadar", "kadar", "sekadarnya", "tepat", "kasus", "berkata", "kata", "tetap",
		"dikatakan", "setiba", "tiba", "dikatakannya", "setibanya", "katakan", "katakanlah", "tiga", "katanya",
		"setinggi", "tinggi", "mengatakan", "mengatakannya", "ditujukan", "tuju", "sekecil", "menuju", "keluar",
		"tertuju", "kembali", "ditunjuk", "tunjuk", "berkenaan", "kena", "ditunjuki", "mengenai", "ditunjukkan",
		"bekerja", "kerja", "ditunjukkannya", "dikerjakan", "ditunjuknya", "mengerjakan", "menunjuk", "dikira",
		"menunjuki", "diperkirakan", "menunjukkan", "menunjuknya", "memperkirakan", "berturut", "turut", "mengira",
		"terkira", "menurut", "kurang", "sekurang", "bertutur", "tutur", "sekurangnya", "dituturkan", "berlainan",
		"dituturkannya", "dilakukan", "menuturkan", "melakukan", "berlalu", "tuturnya", "dilalui", "diucapkan",
		"ucap", "keterlaluan", "diucapkannya", "kelamaan", "mengucapkan", "berlangsung", "langsung",
		"mengucapkannya", "lanjut", "lanjutnya", "ucapnya", "selanjutnya", "berujar", "ujar", "berlebihan",
		"lewat", "ujarnya", "dilihat", "lihat", "umum", "diperlihatkan", "umumnya", "kelihatan", "diungkapkan",
		"ungkap", "kelihatannya", "mengungkapkan", "melihat", "melihatnya", "ungkapnya", "memperlihatkan", "untuk",
		"terlihat", "usah", "kelima", "lima", "seusai", "usai", "luar", "terutama", "utama", "bermaksud", "maksud",
		"dimaksud", "waktunya", "dimaksudkan", "meyakini", "yakin", "dimaksudkannya", "meyakinkan", "dimaksudnya",
		"semampu", "semampunya",
	)
}

// FromSentece create clarifier specific for sentence.
func FromSentence(sentence string) Sentence {
	s := createSentence(sentence)
	return s
}

// // FromParagraph create clarifier specific for paragraph.
// func FromParagraph(paragraph string) Paragraph {

// 	return Paragraph{VagueParagraph: paragraph}
// }

// // FromArticle create clarifier specific for article
// func FromArticle(article string) Article {

// 	return Article{VagueArticle: article}
// }

var (
	// regexEarlyBlankSpace is regex to remove space at the beginning of the sentence, phrase or article
	regexEarlyBlankSpace = regexp.MustCompile(`^ +`)

	// regexExceesBlankSpace is regex to remove excess blank space
	regexExceesBlankSpace = regexp.MustCompile(` {2,}`)

	// stopWords is variable use stopword value
	stopWords StopWord
)

// history use to store history of an information retrieval process
type history struct {
	word struct {
		before string
		after  string
	}
	process string
}

// Stopword is a list of stopword
type StopWord []string

// Tokenized is bag of words
type Tokenized []string

// Json use to convert all value to json value
type Json map[string]interface{}

// casefold return lowercase string
func casefold(sentence string) (string, string) {
	return "casefolding", strings.ToLower(sentence)
}

// stem return basic word
func stem(sentence string) (string, string) {
	words := tokenize(sentence)

	stemmer := sastrawi.NewStemmer(sastrawi.DefaultDictionary)
	for key, word := range words {
		words[key] = stemmer.Stem(word)
	}

	return "stemming", strings.Join(words, " ")
}

// newStopwords create a list of stopwords
func newStopwords(stopword ...string) {
	stopWords = stopword
}

// stopWording remove all stopword from tokenize string
func stopwording(sentence string) (string, string) {

	for _, stopword := range stopWords {
		regexStopword := regexp.MustCompile(`\b` + stopword + `\b`)
		sentence = regexStopword.ReplaceAllString(sentence, "")
	}

	return "stopwording", sentence
}

// cleansingEarlyBlankSpace remove all blank space at the beginning of sentence, phrase or article.
func cleansingEarlyBlankSpace(sentence string) (string, string) {
	sentence = regexEarlyBlankSpace.ReplaceAllString(sentence, "")

	return "cleansing_early_blank_space", sentence
}

// cleansingExceesBlankSpace remove all excess blanks space and replace it with one blank space
func cleansingExceesBlankSpace(sentence string) (string, string) {
	sentence = regexExceesBlankSpace.ReplaceAllString(sentence, " ")

	return "cleansing_excees_blank_space", sentence
}

// cleansingDuplicates remove all of duplicated word from tokenized string
func cleansingDuplicates(sentence []string) (string, string) {
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

// tokenize create bag of words from string
func tokenize(sentence string) []string {
	tokenizer := sastrawi.NewTokenizer()
	return tokenizer.Tokenize(sentence)
}
