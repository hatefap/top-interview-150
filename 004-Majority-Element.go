package main

func majorityElement(nums []int) int {
	c, n := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if n == 0 {
			c, n = nums[i], 1
		} else if nums[i] == c {
			n++
		} else {
			n--
		}
	}
	return c
}
