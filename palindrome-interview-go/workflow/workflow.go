package workflow

import "jamieburns.me/palidrome-interview/action"

func Execute(s string) string {

	if ! action.IsStringAPalindromeCandidate(s) {
		return "Not a palindrome candidate. Contains characters other than lowercase alphabetic characters"
	}

	freqCharacterMap := action.FrequencyOfCharacterMap(s)

	if ! action.IsFreqCharacterMapAPalindromeCandidate(freqCharacterMap) {
		return "Not a palindrome candidate. Contains an unusable frequency of characters"
	}

	return action.MakePalindromeFromFreqCharacterMap( freqCharacterMap )
}
