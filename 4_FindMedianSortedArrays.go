package main

func main() {

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	lenSum := len(nums1) + len(nums2)

	mn_2 := (lenSum + 1) / 2

	tmp := 1
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		var flag bool
		if nums1[i] <= nums2[j] {
			i++
			tmp++
			flag = false
		} else {
			j++
			tmp++
			flag = true
		}

		if tmp == mn_2 {

		}


	}

	return 0
}
