package leetcode

func twoSum(nums []int, target int) []int {
	for i := range len(nums) {
		for _i := i + 1; _i < len(nums); _i++ {
			if nums[i]+nums[_i] == target {
				return []int{i, _i}
			}
		}
	}
	return nil
}
