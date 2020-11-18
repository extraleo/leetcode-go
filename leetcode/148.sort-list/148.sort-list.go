/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
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

// @sort
func sortList(head *ListNode) *ListNode {
	return mergeSortIt(head)
}


// Your runtime beats 47.16 % of golang submissions
// Your memory usage beats 12.63 % of golang submissions (8 MB)
func mergeSortIt(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// find mid
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tail := slow
	slow = slow.Next
	tail.Next = nil
	return mergeTwoListsIt(mergeSortIt(head),mergeSortIt(slow))
}

func mergeTwoListsIt(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	}

	if l2 != nil {
		head.Next = l2
	}

	return dummy.Next
}

// @sort merge sort
// Your runtime beats 26.29 % of golang submissions
// Your memory usage beats 7.22 % of golang submissions (8.8 MB)
// time: O(nlogn) time and O(1) memory
func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// find mid
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tail := slow
	slow = slow.Next
	tail.Next = nil

	return mergeTwoListsRecursion(mergeSort(head), mergeSort(slow))
}

func mergeTwoListsRecursion(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRecursion(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoListsRecursion(l2.Next, l1)
	return l2
}

// @lc code=end
