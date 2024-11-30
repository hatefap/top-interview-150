package main

func removeDuplicatesII(nums []int) int {
	m := make(map[int]int)
	last := 0
	for i := 0; i < len(nums); i++ {
		if n, ok := m[nums[i]]; ok && n > 1 {
			continue
		}
		m[nums[i]] = m[nums[i]] + 1
		nums[last] = nums[i]
		last++
	}
	return last
}
