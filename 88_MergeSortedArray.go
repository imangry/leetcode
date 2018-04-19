package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	if n == 0 {
		return
	}
	length := m + n
	tmp := length - 1
	i, j := m-1, n-1
	for ; i >= 0 && j >= 0; {
		if nums1[i] >= nums2[j] {
			nums1[tmp] = nums1[i]
			tmp --
			i--
		} else {
			nums1[tmp] = nums2[j]
			tmp--
			j--
		}
	}
	for j >= 0 {
		nums1[tmp] = nums2[j]
		tmp --
		j--
	}
}
