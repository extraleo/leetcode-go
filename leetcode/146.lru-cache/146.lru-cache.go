/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start

package leetcode

// @again
// FIXME seems NPE


type ListNode struct{
	Key int
	Val int
	Prev *ListNode
	Next *ListNode
}

type LRUCache struct {
	keyMap map[int] *ListNode
	head *ListNode
	tail *ListNode
	cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		make(map[int]*ListNode, capacity),
		&ListNode{},
		&ListNode{},
		capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if this.keyMap[key] == nil{
		return -1
	}
	node:=this.keyMap[key]
	this.remove(node)
	this.appendHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if this.keyMap[key] == nil{
		newNode:=&ListNode{key, value, &ListNode{},&ListNode{}}
		if len(this.keyMap) == this.cap{
			delete(this.keyMap, this.tail.Key)
			this.remove(this.tail)
			this.appendHead(newNode)
			this.keyMap[key]=newNode
			return
		}
		this.appendHead(newNode)
		this.keyMap[key]=newNode
		return
	}
	node:=this.keyMap[key]
	node.Val = value
	this.remove(node)
	this.appendHead(node)
}

func (this *LRUCache) remove(node *ListNode){
	if this.head == this.tail && this.head == node{
		this.head, this.tail = nil, nil
		return
	}
	if this.tail == node{
		this.tail = this.tail.Prev
		this.tail.Next = nil
		node.Prev = nil
		return
	}
	if this.head == node{
		this.head = this.head.Next
		this.head.Prev = nil
		node.Next = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Prev = nil
	node.Next = nil
}

func (this *LRUCache) appendHead(node *ListNode){
	if this.head == this.tail && this.head == nil{
		this.head, this.tail = node, node
		return
	}
	node.Prev = nil
	node.Next = this.head
	this.head = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
