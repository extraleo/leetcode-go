/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */

// @lc code=start

package leetcode

import "github.com/extraleo/alg/structures"

type ListNode = structures.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil{
		return head
	}
	dummy := &ListNode{
		Val: val+1,
		Next: head,
	}
	curr := dummy
	for curr != nil && curr.Next != nil{
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		}else{
			curr = curr.Next
		}
	}
	return dummy.Next
}
// @lc code=end

