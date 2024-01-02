package main

import (
	"fmt"

	"github.com/cecarbs/dsa/internal/warmup"
)

func main() {
	// sentence := "TheQuickBrownFoxJumpsOverTheLazyDog"
	sentence := "This is not a pangram"
	result := warmup.CheckIfPangram(sentence)
	fmt.Println(result)
}
