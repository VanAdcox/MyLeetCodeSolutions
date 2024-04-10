package leetcode

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return returnMedian(nums2)
	}
	if len(nums2) == 0 {
		return returnMedian(nums1)
	}
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	return returnMedian(nums1)
}

func returnMedian(array []int) float64 {
	if len(array)%2 != 0 {
		return float64(array[len(array)/2])
	} else {
		return float64(array[len(array)/2-1]+array[len(array)/2]) / 2
	}
}
