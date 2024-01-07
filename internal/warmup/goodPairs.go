package warmup

func GoodPairs(nums []int) int {
	pairCount := 0
	pairCountMap := make(map[int]int)
	for _, num := range nums {
		if count, ok := pairCountMap[num]; ok {
			if count > 1 {
				pairCount = updateCount(pairCount, count)
			}
		}
	}
	return pairCount
}

func updateCount(pairCount int, numCount int) int {
	pairCount += numCount - 1
	return pairCount
}
