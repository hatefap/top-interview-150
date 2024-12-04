package main

func canJump(nums []int) bool {
	reach := nums[0]
	for i := 1; i < len(nums); i++ {
		if reach < i {
			return false
		}
		reach = max(reach, nums[i]+i)
	}
	return reach >= len(nums)-1
}
