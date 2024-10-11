package action

import "regexp"

func IsStringAPalindromeCandidate(s string) bool {

	matched, _ := regexp.Match(`^[a-z]*$`, []byte(s))

	return matched
}
