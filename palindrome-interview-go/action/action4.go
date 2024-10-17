package action

import (
	"slices"
	"maps"
)

func MakePalindromeFromFreqCharacterMap( freqCharacterMap map[string]int ) string {

	if len( freqCharacterMap ) == 0 {
		return ""
	}

	var palindromeLength int

	for _,v := range freqCharacterMap {
		palindromeLength += v
	}

	keyList := slices.Sorted(maps.Keys(freqCharacterMap))

	palindrome := make([]rune, palindromeLength)

	left := 0
	right := palindromeLength - 1

	for _, k := range keyList {

		count := freqCharacterMap[k]
		r := []rune(k)[0]

		if count % 2 == 1 {
			palindrome[palindromeLength / 2] = r
			count--
		}

		for ; count > 0; {

			palindrome[left] = r
			palindrome[right] = r

			left++
			right--

			count -= 2
		}
	}

	return string(palindrome)
}