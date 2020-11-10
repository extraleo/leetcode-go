/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start

package leetcode

import (
	"github.com/extraleo/alg/structures"
)

type ListNode = structures.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{0, &ListNode{}}
	dummy := newHead
	var1, var2, flag := 0, 0, 0
	for l1 != nil || l2 != nil || flag != 0 {
		if l1 == nil {
			var1 = 0
		} else {
			var1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			var2 = 0
		} else {
			var2 = l2.Val
			l2 = l2.Next
		}
		newHead.Next = &ListNode{Val: (var1 + var2 + flag) % 10}
		newHead = newHead.Next
		flag = (var1 + var2 + flag) / 10
	}
	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil || head.Next != nil {
		prev, head, head.Next = head, head.Next, prev
	}
	return prev
}

// @lc code=end
