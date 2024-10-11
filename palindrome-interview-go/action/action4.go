package action

import (
	"sort"
	"strings"
)

func MakePalindromeFromFreqCharacterMap( freqCharacterMap map[string]int ) string {
	
	if len( freqCharacterMap ) == 0 {
		return ""
	}

	var palindromeLength int
	var keyList []string

	for k,v := range freqCharacterMap {
		palindromeLength += v
		keyList = append(keyList, k)
	}

	sort.Strings(keyList)

	palindrome := make([]string, palindromeLength)

	left := 0
	right := palindromeLength - 1

	for _, k := range keyList {

		count := freqCharacterMap[k]

		if count % 2 == 1 {
			palindrome[palindromeLength / 2] = k
			count--
		}

		for ; count > 0; {

			palindrome[left] = k
			palindrome[right] = k

			left++
			right--

			count -= 2
		}
	}

	return strings.Join(palindrome, "")
}