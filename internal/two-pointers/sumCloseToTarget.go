package twopointers

import (
	"math"
	"sort"
)

func SumCloseToTarget(nums []int, targetSum int) int {
	sort.Ints(nums)

	curr := 0
	left := curr + 1
	right := len(nums) - 1

	smallestDifference := math.MaxInt

	for curr < len(nums)-2 {
		result := findPairs(nums, curr, left, right, targetSum)
		if result == targetSum {
			return targetSum
		}

		smallestDifference = int(math.Min(float64(smallestDifference), float64(result)))
	}
	return smallestDifference
}

// if this function returns 0 then we can immediately return
func findPairs(nums []int, curr int, left int, right int, targetSum int) int {
	smallestDifference := math.MaxInt

	for left < right {
		// need to find the smallest difference
		// calculate the difference b/t target sum and sum of all numbers in curr position
		if nums[left]+nums[right]+nums[curr] == 0 {
			return targetSum
		}
		currentSmallestDifference := targetSum - nums[left] - nums[right] - nums[curr]
		smallestDifference = int(math.Min(float64(smallestDifference), float64(currentSmallestDifference)))
	}
	return smallestDifference
}

func float(currentSmallestDifference int) {
	panic("unimplemented")
}
