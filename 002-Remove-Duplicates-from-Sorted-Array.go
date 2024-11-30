package main

func removeDuplicates(nums []int) int {
	set := make(map[int]struct{})
	last := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			continue
		}
		set[nums[i]] = struct{}{}
		nums[last] = nums[i]
		last++
	}
	return last
}
