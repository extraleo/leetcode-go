package sorting

import "math/rand"

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	var low, same, high []int
	pivot := nums[rand.Intn(len(nums))]

	for _, item := range nums {
		if item < pivot {
			low = append(low, item)
		} else if item == pivot {
			same = append(same, item)
		} else {
			high = append(high, item)
		}
	}
	return append(append(sortArray(low), same...), sortArray(high)...)
}
