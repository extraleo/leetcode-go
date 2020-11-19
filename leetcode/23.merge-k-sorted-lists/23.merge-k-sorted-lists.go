/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
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

// @sort 
// todo 看起来只需要不断的merge, 但是, 我想是否两个两个排序的时候可以使用快排?
func mergeKLists(lists []*ListNode) *ListNode {
return nil
}

func mergeTwoLists(left, right *ListNode) *ListNode{
  if left == nil{
    return right
  }
  if right == nil{
    return left
  }

  if left.Val < right.Val{
    left.Next = mergeTwoLists(left.Next, right)
    return left
  }
  right.Next= mergeTwoLists(right.Next, left)
  return right
}


// @lc code=end

