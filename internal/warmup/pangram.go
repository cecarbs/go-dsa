package warmup

import (
	"fmt"
	"strings"
	"unicode"
)

func CheckIfPangram(sentence string) bool {
	sentence = strings.ToLower(sentence)
	fmt.Println(sentence)
	hashmap := make(map[rune]int)

	for _, char := range sentence {
		if !unicode.IsLetter(char) {
			continue
		}
		_, ok := hashmap[char]

		if ok {
			hashmap[char] = hashmap[char] + 1
		} else {
			hashmap[char] = 1
		}
	}

	return len(hashmap) == 26
}
