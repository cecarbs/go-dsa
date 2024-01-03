package warmup

import (
	"fmt"
	"unicode"
)

func ValidPalindrome(str string) bool {
	i := 0
	j := len(str) - 1

	for i < j {

		if !unicode.IsLetter(rune(str[i])) && !unicode.IsDigit(rune(str[i])) {
			fmt.Println(str[i])
			i++
			continue
		}
		if !unicode.IsLetter(rune(str[j])) && !unicode.IsDigit(rune(str[i])) {
			fmt.Println(str[i])
			j--
			continue
		}

		if unicode.ToLower(rune(str[i])) == unicode.ToLower(rune(str[j])) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
