/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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

// package leetcode

// import "github.com/extraleo/alg/structures"

// type TreeNode = structures.TreeNode

func preorderTraversal(root *TreeNode) []int {
	return preorderTraversalRecursion1(root)
}

func preorderTraversalInteration(root *TreeNode) []int {
	return nil
}

// 递归
func preorderTraversalRecursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	o := []int{root.Val}
	return append(append(o, preorderTraversalRecursion(root.Left)...), preorderTraversalRecursion(root.Right)...)
}

// 递归
func preorderTraversalRecursion1(root *TreeNode) []int{
	var output []int
	preorderHandler(root, &output)
	return output
}

func preorderHandler(root *TreeNode, output *[]int){
	if root != nil{
		*output = append(*output, root.Val)
		preorderHandler(root.Left, output)
		preorderHandler(root.Right, output)
	}
}

// @lc code=end

