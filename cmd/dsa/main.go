package main

import (
	"fmt"

	twopointers "github.com/cecarbs/dsa/internal/two-pointers"
)

func main() {
	// sentence := "TheQuickBrownFoxJumpsOverTheLazyDog"
	// sentence := "This is not a pangram"
	// result := warmup.CheckIfPangram(sentence)
	// fmt.Println(result)
	// testCase := "12345"
	// result := warmup.ValidPalindrome(testCase)
	// fmt.Println(result)
	// arr := []int{2, 3, 3, 3, 6, 9, 9}
	// result := twopointers.NonDuplicateNumbersInstance(arr)
	arr := []int{-2, -1, 0, 2, 3}
	result := twopointers.SquaringSortedArray(arr)
	fmt.Println(result)
}
