/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 */

// @lc code=start

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package leetcode

import "github.com/extraleo/alg/structures"

type ListNode = structures.ListNode

// @again
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: 0, Next: nil}
	cur, pre := head, dummy
	for cur != nil {
		next := cur.Next
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur
		pre = dummy
		cur = next
	}
	return dummy.Next
}

// @lc code=end

func insertionSortArray(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		k := i - 1
		for k >= 0 && nums[k] > nums[i] {
			nums[k+1] = nums[k]
			k--
		}
		nums[k+1] = nums[i]
	}
	return nums
}
