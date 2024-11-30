package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	total := m + n
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[total-1] = nums1[m-1]
			m--
		} else {
			nums1[total-1] = nums2[n-1]
			n--
		}
		total--
	}

	if n > 0 && m == 0 {
		for total > 0 {
			nums1[total-1] = nums2[n-1]
			n--
			total--
		}
	}
}
