package main

import "math"

func main() {

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	lenSum := len(nums1) + len(nums2)

	mn_2 := int(math.Floor(float64(lenSum / 2)))

	tmp := 1
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] <= nums2[j] {
			i++
			tmp++
		} else {
			j++
			tmp++
		}
		if tmp == mn_2 {
			if lenSum%2 == 0 {

			}
		}
	}

	return 0
}
