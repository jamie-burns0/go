package action

func IsFreqCharacterMapAPalindromeCandidate( freqCharacterMap map[string]int) bool {
	
	if len( freqCharacterMap ) == 0 {
		return false
	}

	foundOddNumber := false

	for _,v := range freqCharacterMap {
		if v % 2 != 0 {
			if foundOddNumber {
				return false
			}
			foundOddNumber = true
		}
	}
	return true
}