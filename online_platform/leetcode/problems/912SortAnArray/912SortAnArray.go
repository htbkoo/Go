package _12SortAnArray

import "math"

func sort(nums []int, lo int, hi int) []int {
	if lo >= hi {
		return nums[lo : lo+1]
	}

	mid := lo + (hi-lo)/2
	left := sort(nums, lo, mid)
	right := sort(nums, mid+1, hi)

	result := make([]int, len(left)+len(right))
	l := 0
	r := 0
	for l < len(left) || r < len(right) {
		leftVal := math.MaxInt
		if l < len(left) {
			leftVal = left[l]
		}

		rightVal := math.MaxInt
		if r < len(right) {
			rightVal = right[r]
		}

		if rightVal < leftVal {
			result[l+r] = right[r]
			r++
		} else {
			result[l+r] = left[l]
			l++
		}
	}

	return result
}

func sortArray(nums []int) []int {
	return sort(nums, 0, len(nums)-1)
}
