package main

// first try
func canJumpII(nums []int) int {
	n := len(nums)
	minJumps := make([]int, n)
	for i := 0; i < n; i++ {
		minJumps[i] = 100_000
	}
	minJumps[0] = 0
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if j+nums[j] >= i {
				minJumps[i] = min(minJumps[i], minJumps[j]+1)
			}
		}
	}
	return minJumps[n-1]
}
