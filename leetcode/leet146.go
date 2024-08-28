package leetcode

import "container/list"

type LRUCache struct {
	cap   int
	group list.List
	cache map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		group: list.List{},
		cache: map[int]*list.Element{},
	}
}

func (this *LRUCache) Get(key int) int {
	data, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.group.MoveToFront(data)
	return data.Value.(int)
}

func (this *LRUCache) Put(key int, value int) {
	data, ok := this.cache[key]
	if !ok {
		el := this.group.PushFront(value)
		this.cache[key] = el
		if len(this.cache) > this.cap {
			last := this.group.Back()
			delete(this.cache, last.Value.(int))
			this.group.Remove(last)
		}
	} else {
		data.Value = value
		this.group.MoveToFront(data)
		this.cache[key] = data
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
