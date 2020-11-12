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
	"reflect"
	"testing"

	"github.com/extraleo/alg/structures"
)

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		put args
		want []int
	}{
		{args{structures.Ints2List([]int {1,2,3,4,5,6})}, []int{1,6,2,5,3,4}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := reorderList(tt.put.head); !reflect.DeepEqual(structures.List2Ints(got), (tt.want)) {
				t.Errorf("reorderList() = %v, want %v", got, tt.want)
			}
		})
	}
}
