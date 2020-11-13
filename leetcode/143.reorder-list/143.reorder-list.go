/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
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


// @note
// 123654 => 162534
func reorderList(head *ListNode) *ListNode {
  if head == nil|| head.Next == nil{
    return head
  }
  fast, slow := head, head
  for fast.Next != nil && fast.Next.Next != nil{
    fast = fast.Next.Next
    slow = slow.Next
  }
	mid := slow // mid = 3
  mid.Next = reverse(slow.Next)  // 123654 

  p1 := head // 1
  p2 := mid.Next // 6


  // 123654 => 12354 
  //             654 
  // 62354
  // 162354
  // p1 = 2
  // p2 = 5
  for p1 != mid{ 
    mid.Next = p2.Next
    p2.Next = p1.Next
    p1.Next = p2
    p1 = p2.Next 
    p2 = mid.Next 
  }

  return head
}


func reverse(head *ListNode) *ListNode{
  var prev *ListNode
  for head != nil{
    next:=head.Next
    head.Next = prev
    prev = head
    head = next
  }
  return prev
}
// @lc code=end

