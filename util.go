package tfidf

import (
	"strings"
	"unicode"
	"regexp"
)

// init initializes the package by pre-compiling the regex for the stripping method
func init() {
	r, err := regexp.Compile(`[\s\p{Zs}]{2,}`)

	if nil != err {
		panic(err)
	}

	stripSpacesRegex = r
}

// stripSpacesRegex in ram cache for the compiled regex
var stripSpacesRegex *regexp.Regexp

// StripSpacesRegex removes all the duplicates spaces and all newlines through regex
// it produces one space between words text
func StripSpacesRegex(t string) string {
	r := stripSpacesRegex.ReplaceAllString(t, " ")

	return r
}

// StripSpacesLoop removes all the duplicates spaces and all newlines through a loop
// it produces one space between words text
func StripSpacesLoop(t string) string {
	var lsp bool = false

	r := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			if ('\t' == r || '\n' == r) && !lsp {
				lsp = true
				return ' '
			}
			if lsp || '\n' == r {
				return -1
			} else {
				lsp = true
				return r
			}
		} else {
			lsp = false
			return r
		}
	}, t)

	return r
}
