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
	return reverseListHeadAndEnd(head)
}

func reverseListHeadAndEnd(head *ListNode) *ListNode{
	last :=  head
	for last.Next != nil{
		last= last.Next
	}

  return reverse(head,last)
}


func reverse(first, last *ListNode) *ListNode{
	var prev *ListNode
	curr := first
	for prev != last{
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// Recursion
func reverseListR(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	newHead:=reverseListR(head)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseListA1(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil{
		head.Next, pre, head = pre, head, head.Next
	}
	return pre
}

func reverseListA0(head *ListNode) *ListNode{
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

// @lc code=end

