package action

import (
	"regexp"
	"strings"
)

var re string

func init() {
	latinClasses := []string {
		"a-z",
		"\u00E0-\u00E5", // à-å
		"\u00E7-\u00EF", // ç-ï
		"\u00F1-\u00F6", // ñ-ö
		"\u00F9-\u00FD", // ù-ý
		"\u00FF", // ÿ
	}

	joinedLatinClasses := strings.Join(latinClasses, "")
	re = "^[" + joinedLatinClasses + "]*$"
}

func IsStringAPalindromeCandidate(s string) bool {
	matched, _ := regexp.MatchString(re, s) 
	return matched
}
