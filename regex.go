package clarity

import (
	"regexp"
)

var regexEarlyBlankSpace = regexp.MustCompile(`^ +`)
var regexExceesBlankSpace = regexp.MustCompile(` {2,}`)
