package main

import (
	"fmt"

	"github.com/cecarbs/dsa/internal/warmup"
)

func main() {
	// sentence := "TheQuickBrownFoxJumpsOverTheLazyDog"
	// sentence := "This is not a pangram"
	// result := warmup.CheckIfPangram(sentence)
	// fmt.Println(result)
	// testCase := "12345"
	// result := warmup.ValidPalindrome(testCase)
	// fmt.Println(result)
	s := "paper"
	t := "repaap"
	fmt.Println(warmup.ValidAnagram(s, t))
}
