package action

func FrequencyOfCharacterMap( data string ) map[string]int {
	
	freqMap := map[string]int {}

	for _, c := range data {
		freqMap[string(c)]++
	}

	return freqMap
}