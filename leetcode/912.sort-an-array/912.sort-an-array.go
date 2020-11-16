/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 */

// @lc code=start

package leetcode

import "math/rand"

func sortArray(nums []int) []int {
	return quick(nums)
}

func quick(nums []int) []int {
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

	return append(append(quick(low), same...), quick(high)...)
}

// @lc code=end
