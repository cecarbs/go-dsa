package twopointers

func PairWithTargetSum(nums []int, targetSum int) []int {
	i := 0
	j := len(nums) - 1

	for i < j {
		if nums[i]+nums[j] == targetSum {
			return []int{i, j}
		}
		if nums[i]+nums[j] > targetSum {
			j--
			continue
		}
		if nums[i]+nums[j] < targetSum {
			i++
			continue
		}
	}
	return []int{-1, -1}
}
