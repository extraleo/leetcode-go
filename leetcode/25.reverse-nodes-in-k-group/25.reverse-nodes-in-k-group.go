package leetcode

import "github.com/extraleo/alg/structures"

/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start

type ListNode=structures.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(first, last *ListNode) *ListNode {
	prev := last
	for first != last {
		next := first.Next
		first.Next = prev
		prev = first
		first = next
	}
	return prev
}

// @lc code=end

// 2个一组反转链表
func reverse2Group(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	third := reverse2Group(head.Next.Next)
	second := head.Next
	head.Next = third
	second.Next = head
	return second
}

