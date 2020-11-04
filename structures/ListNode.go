package structures

type ListNode struct {
	Val int
	Next *ListNode
}

func Ints2List(val []int) *ListNode{
	if len(val) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range val {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

func List2Ints(val *ListNode) []int{
	result := make([]int,0)
	for val.Next != nil{
		result = append(result, val.Val)
		val=val.Next
	}
	return result
}
