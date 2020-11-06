/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start

package leetcode

import (
	"github.com/extraleo/alg/structures"
)

type LRUNode = structures.LRUNode

type LRUCache struct {
	keyMap map[int] LRUNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		make(map[int]structures.LRUNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
