package twopointers

import "sort"

func TripletSumToZero(nums []int) [][]int {
	var triplets [][]int

	// TODO:
	// 1. sort the array first
	// 2. since x + y + z = 0, it's also fair to say that x + y = -z
	// 3. you cannot have duplicates so you can skip a number if x, y, or z are the same
	// 4. put a pointer at the start of the array and one at the end
	// 5. if the sum of the sum equals the the target (current iteration), then add to the list
	// 6. if sum is less, then we need to increment, the left; if it is greater, decrement the right
	//
	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			triplets = append(triplets, []int{nums[0], nums[1], nums[2]})
			return triplets
		}
	}
	sort.Ints(nums)

	curr := 0

	for curr < len(nums)-2 {
		left := curr + 1
		right := len(nums) - 1
		// if the current number is the same as the previous we just go to the next number
		if curr-1 >= 0 && nums[curr] == nums[curr-1] {
			curr++
			continue
		}
		result := findPair(nums, left, right, curr)
		if len(result) > 0 {
			triplets = append(triplets, result...)
		}
		curr++
	}

	return triplets
}

func findPair(arr []int, left int, right int, curr int) [][]int {
	var result [][]int
	for left < right {
		if arr[left]+arr[right] == -arr[curr] {
			result = append(result, []int{arr[curr], arr[left], arr[right]})
			left++
			right--
		}
		if arr[left]+arr[right] > -arr[curr] {
			right--
			continue
		}
		if arr[left]+arr[right] < -arr[curr] {
			left++
			continue
		}
	}
	return result
}

// [-3, -2, -1, 0, 1, 1, 2]
// [-4, -1, -1, 0, 1, 2]
