package warmup

func ValidAnagram(s, t string) bool {
	stringMap := make(map[rune]int)
	for _, letter := range s {
		stringMap[letter]++
	}

	for _, letter := range t {
		stringMap[letter]--
	}

	for _, value := range stringMap {
		if value != 0 {
			return false
		}
	}
	return true
}
