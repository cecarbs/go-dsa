package warmup

import "unicode"

func ReverseVowels(str string) string {
	strRunes := []rune(str)
	i := 0
	j := len(str) - 1

	for i < j {
		if !isVowelMap(rune(strRunes[i])) {
			i++
			continue
		}
		if !isVowelMap(rune(strRunes[j])) {
			j--
			continue
		}

		strRunes[i], strRunes[j] = strRunes[j], strRunes[i]
		i++
		j--
	}
	return string(strRunes)
}

func isVowelSlice(char rune) bool {
	char = unicode.ToLower(char)
	vowels := "aeiouAEIOU"
	for _, vowel := range vowels {
		return char == vowel
	}
	return false
}

func isVowelMap(char rune) bool {
	char = unicode.ToLower(char)
	vowelMap := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	_, ok := vowelMap[char]
	return ok
}
