package warmup

func ReverseVowels(str string) string {
	i := 0
	j := len(str) - 1

	for i < j {
	}
}

func isVowelSlice(char rune) bool {
	vowels := "aeiouAEIOU"
	for _, vowel := range vowels {
		return char == vowel
	}
	return false
}

func isVowelMap(char rune) bool {
	vowelMap := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	_, ok := vowelMap[char]
	if ok {
		return true
	}
	return false
}
