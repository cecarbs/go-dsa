package warmup

import "math"

func shortestDistance(words []string, word1 string, word2 string) int {
	if len(words) < 1 {
		return 0
	}
	shortestDistance := math.MaxInt32
	pointerOne, pointerTwo := -1, -1

	for idx, word := range words {
		// if the first occurence is either word 1 or 2, then we know we need to look for the other
		switch word {
		case word1:
			pointerOne = idx
		case word2:
			pointerTwo = idx
		}

		if pointerOne != -1 && pointerTwo != -1 {
			shortestDistance = int(math.Min(float64(shortestDistance), math.Abs(float64(pointerOne-pointerTwo))))
		}
	}
	return shortestDistance
}
