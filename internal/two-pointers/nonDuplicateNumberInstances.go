package twopointers

func NonDuplicateNumbersInstance(nums []int) int {
	left := 1
	right := 1

	for right < len(nums) {
		// right pointer will scan the array
		// if right pointer finds a value that's unique (nums[right] != nums[right - 1]) then swap left and right then increment left
		// if the value is not unique, then continue scanning the array
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}

		right++
	}

	return left
}
