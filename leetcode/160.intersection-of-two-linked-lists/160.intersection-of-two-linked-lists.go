/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := lenListNode(headA)
	lenB := lenListNode(headB)
	stepA, stepB := headA, headB
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			stepA = stepA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			stepB = stepB.Next
		}
	}

	for stepA != nil {
		if stepA == stepB {
			return stepA
		}
		stepA = stepA.Next
		stepB = stepB.Next
	}
	return nil
}

func lenListNode(head *ListNode) int {
	if head == nil {
		return 0
	}
	count := 0
	for head != nil {
		head = head.Next
		count++
	}
	return count
}

// @lc code=end
