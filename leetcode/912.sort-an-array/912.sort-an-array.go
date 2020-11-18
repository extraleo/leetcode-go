/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 */

// @lc code=start

package leetcode

import (
	"math/rand"
)

func sortArray(nums []int) []int {
	return mergeSort(nums)
}

// todo
func hellSort(nums []int) []int{
	return nums
}

func mergeSort(nums []int) []int{
	if len(nums) < 2{
		return nums
	}
	mid := len(nums)/2 -1
	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func merge(left, right []int) []int{
	if left == nil || len(left) == 0{
		return right
	}
	if right == nil || len(right) == 0{
		return left
	}
	// todo
	// all := make([]int,0)

	return nil
}

func insertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		keyItem := nums[i]
		k := i - 1
		for k >= 0 && nums[k] > keyItem {
			nums[k+1] = nums[k]
			k--
		}
		nums[k+1] = keyItem
	}
	return nums
}

func selectSort(nums []int) []int {
	for index := range nums {
		min := nums[index]
		minIndex := index
		for k := index + 1; k < len(nums); k++ {
			if nums[k] < min {
				minIndex = k
				min = nums[k]
			}
		}
		nums[index], nums[minIndex] = nums[minIndex], nums[index]
	}
	return nums
}

func quickSort(nums []int) []int {
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

	return append(append(quickSort(low), same...), quickSort(high)...)
}

// @lc code=end
