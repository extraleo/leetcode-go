/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
 */

// @lc code=start

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package leetcode

import (
	"github.com/extraleo/alg/structures"
)

type TreeNode = structures.TreeNode

// Reference: https://leetcode.com/problems/count-complete-tree-nodes/discuss/61958/Concise-Java-solutions-O(log(n)2
// @again
func countNodes(root *TreeNode) int {
	h := height(root)
	if h < 0 {
		return 0
	}
	if height(root.Right) == h-1 { // 如果右节点的高度 == 整棵树的高度-1, 也就是说 右子树的高度 == 左子树的高度
		return (1 << h) + countNodes(root.Right) // 左侧是全数
	}
	return (1<< (h - 1)) + countNodes(root.Left)  // 右侧是全数,
}

func height(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return 1 + height(root.Left)
}

// Time O(N)
// That would be O(n). But... the actual solution has a gigantic optimization. It first walks all the way left and right to determine the height and whether it's a full tree, meaning the last row is full. If so, then the answer is just 2^height-1. And since always at least one of the two recursive calls is such a full tree, at least one of the two calls immediately stops. Again we have runtime O(log(n)^2).
func countNode1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNode1(root.Left) + countNode1(root.Right)
}

// @lc code=end
