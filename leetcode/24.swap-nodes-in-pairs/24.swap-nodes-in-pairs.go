package leetcode

import "github.com/extraleo/alg/structures"

/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start

type ListNode = structures.ListNode

// 啊西, 写完 first-last 迭代之后不知道这个怎么写了
// 啊西, 直接迭代超出时间限制
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
        return head
    }
    newHead:=swapPairs(head.Next.Next)
    second:=head.Next
    second.Next =head
    head.Next=newHead
    return second
}

// @lc code=end
