package twopointers

func SquaringSortedArray(nums []int) []int {
	var startingPoint int
	var squares []int

	for idx, num := range nums {
		if num >= 0 {
			startingPoint = idx
			break
		}
	}

	left := startingPoint - 1
	right := startingPoint

	for left >= 0 || right < len(nums) {
		if right >= len(nums) {
			leftSquare := nums[left] * nums[left]
			squares = append(squares, leftSquare)
			left--
		} else if left <= 0 {
			rightSquare := nums[right] * nums[right]
			squares = append(squares, rightSquare)
			right++
		} else {
			rightSquare := nums[right] * nums[right]
			leftSquare := nums[left] * nums[left]
			if rightSquare < leftSquare {
				squares = append(squares, rightSquare, leftSquare)
			} else {
				squares = append(squares, leftSquare, rightSquare)
			}
			left--
			right++
		}
	}
	return squares
}

// alternative solution is to set a pointer at each end of the array and compare the values since the array is sorted
