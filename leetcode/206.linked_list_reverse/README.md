## 翻转链表

206是个起始基础, 对应延伸出来的都能从206的迭代解法稍微微改动一下.

迭代:
```golang
func reverseList(head *ListNode) *ListNode {
  var prev *ListNode
  for head != nil{ // 当head移动到最后就完成了 1. 确认循环退出条件
    next := head.Next // 2.记录下 下一个 节点 
    head.Next = prev // 4. 开始交换节点
    prev = head
    head = next // 3. 开始下次循环
  }
  return prev
}
```
25 是把 K 个节点一次的翻转, 想想 206 把所有节点翻转, 是不是可以想成给定 头和尾节点, 翻转这节节点链, 翻转所有节点就是 head =  head; end = nil, 即
```golang
func reverseList(first, last *ListNode) *ListNode {
  prev := last.Next
  end := last.Next // end = nil
  for head != end{ // 当head移动到最后就完成了 1. 确认循环退出条件
    next := head.Next // 2.记录下 下一个 节点 
    head.Next = prev // 4. 开始交换节点
    prev = head
    head = next // 3. 开始下次循环
  }
  return prev
}
```


那么在 25 里面, 我们先找到k个节点的头和尾
```golang
func reverseListKGroup(head *ListNode) *ListNode {
  node := head
  for i:=0; i<k; i++{
    if node == nil{ // 如果 node == nil 就说明这一段没有 K 个节点, 就不翻转
      return head
    }
    node = node.Next
  }
  newHead := reverse(head, node)
  head.Next = reverseListKGroup(node, k)
  return newHead
}

func reverse(first, last *ListNode) *ListNode{
  prev := last
  end := last
  for first != end{
    next = first.Next
    first.Next = prev
    prev = first
    first = next
  }
}
```
那么在 24 里面, 2个一组地翻转, 就是 k = 2 的情况 