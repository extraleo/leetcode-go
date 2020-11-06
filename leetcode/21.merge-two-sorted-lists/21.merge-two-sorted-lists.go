/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start

package leetcode

import "github.com/extraleo/alg/structures"

type ListNode = structures.ListNode


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

}


// using dummy node to create list
// seems not very fast:
// Your runtime beats 58.55% of golang submissions
// Your memory usage beats 62.46% of golang submissions
func mergeTwoListIteration(l1,l2 *ListNode) *ListNode{
		if l1 == nil{
		return l2
	}

	if l2 == nil{
		return l1
	}
	dummy := &ListNode{0, &ListNode{}}
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		}else{
			head.Next = l2
			l2 = l2.Next
		}
		head=head.Next
	}
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil{
		head.Next =l2
	}
	return dummy.Next
}

// @lc code=end
