/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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

import "github.com/extraleo/alg/structures"

// TreeNode Definition
type TreeNode = structures.TreeNode

func postorderTraversal(root *TreeNode) []int {
	return postorderTraversalRecursion1(root)
}

// 递归
func postorderTraversalRecursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var output []int
	postorderHandler(root, &output)
	return output
}

func postorderHandler(root *TreeNode, output *[]int) {
	if root != nil {
		postorderHandler(root.Left, output)
		postorderHandler(root.Right, output)
		*output = append(*output, root.Val)
	}
}

func postorderTraversalRecursion1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	output := []int{root.Val}
	return append(append(postorderTraversalRecursion1(root.Left), postorderTraversalRecursion1(root.Right)...), output...)
}

// @lc code=end
