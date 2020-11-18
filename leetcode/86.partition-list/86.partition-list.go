/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
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

import (
	"github.com/extraleo/alg/structures"
)

type ListNode = structures.ListNode

/**
*  空间复杂度是 n, 但是我感觉应该是有可能空间复杂是 1 的
*/
func partition(head *ListNode, x int) *ListNode {
	return partitionBuildNewOne(head, x)
}

func partitionBuildNewOne(head *ListNode, x int) *ListNode {
	mHead, lHead := &ListNode{0, nil}, &ListNode{0, nil}
	m, l := mHead, lHead
	for head != nil {
		if head.Val < x {
			l.Next = head
			l = l.Next
		} else {
			m.Next = head
			m = m.Next
		}
		head = head.Next
	}

	m.Next = nil
	l.Next = mHead.Next
	return lHead.Next

}

// @lc code=end
