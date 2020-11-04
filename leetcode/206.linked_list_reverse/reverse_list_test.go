package leetcode

import (
	"reflect"
	"testing"

	"github.com/extraleo/alg/structures"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "first",
			args: args{
				head: structures.Ints2List([]int{1,2,3,4,5}),
			},
			want: structures.Ints2List([]int{5,4,3,2,1}),
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
