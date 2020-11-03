package leetcode

import "github.com/extraleo/alg/structures"

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start

type ListNode = structures.ListNode

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head.Next
	for{
		if fast == nil || fast.Next == nil{
			return false
		}
		if slow == fast{
			return true
		}
		slow=slow.Next
		fast=fast.Next.Next
	}
}

func hasCycleMap(head *ListNode) bool{
	nodeMap := make(map[*ListNode]bool)
	for head != nil{
		if nodeMap[head] {
			return true
		}
		nodeMap[head]=true
		head=head.Next
	} 
  return false
}

// @lc code=end