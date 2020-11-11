/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
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
	"reflect"
	"testing"

	"github.com/extraleo/alg/structures"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
	}
	tests := []struct {
		put args
		want []int
	}{
		{args{structures.Ints2List([]int{1,2,3,4,5}),2,4}, []int{1,4,3,2,5}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := structures.List2Ints(reverseBetween(tt.put.head, tt.put.m, tt.put.n)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
