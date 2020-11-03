package leetcode

import (
	"github.com/extraleo/alg/structures"
)

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start


type ListNode = structures.ListNode

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil{
		head.Next, pre, head = pre, head, head.Next
	}
	return pre
}

func reverseListA(head *ListNode) *ListNode{
	var pre *ListNode
 	curr:=head
	
	for curr != nil{
		next := curr.Next
		curr.Next =pre
		pre = curr
		curr = next
	}
	return pre
}

func reverseListR(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	newHead:=reverseListR(head)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
// @lc code=end

