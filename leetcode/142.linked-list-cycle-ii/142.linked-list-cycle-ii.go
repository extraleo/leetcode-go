/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start

package leetcode

import "github.com/extraleo/alg/structures"

type ListNode = structures.ListNode

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow1, fast1 := head, fast
			for {
				if slow1 == fast1 {
					return slow1
				}
				slow1 = slow1.Next
				fast1 = fast1.Next
			}
		}
	}
}

// seems map also resovle the cycle ii but:
// 16/16 cases passed (8 ms)
// Your runtime beats 36.67 % of golang submissions
// Your memory usage beats 7.34 % of golang submissions (5.2 MB)
func hasCycleMapII(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool)
	for head != nil {
		if nodeMap[head] {
			return head
		}
		nodeMap[head] = true
		head = head.Next
	}
	return nil
}

// @lc code=end
