/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU缓存
 *
 * https://leetcode-cn.com/problems/lfu-cache/description/
 *
 * algorithms
 * Hard (34.37%)
 * Likes:    96
 * Dislikes: 0
 * Total Accepted:    3.1K
 * Total Submissions: 9K
 * Testcase Example:  '["LFUCache","put","put","get","put","get","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]'
 *
 * 设计并实现最不经常使用（LFU）缓存的数据结构。它应该支持以下操作：get 和 put。
 *
 * get(key) - 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。
 * put(key, value) -
 * 如果键不存在，请设置或插入值。当缓存达到其容量时，它应该在插入新项目之前，使最不经常使用的项目无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，最近最少使用的键将被去除。
 *
 * 进阶：
 * 你是否可以在 O(1) 时间复杂度内执行两项操作？
 *
 * 示例：
 *
 *
 * LFUCache cache = new LFUCache( 2 /* capacity (缓存容量) */ //);
//  *
//  * cache.put(1, 1);
//  * cache.put(2, 2);
//  * cache.get(1);       // 返回 1
//  * cache.put(3, 3);    // 去除 key 2
//  * cache.get(2);       // 返回 -1 (未找到key 2)
//  * cache.get(3);       // 返回 3
//  * cache.put(4, 4);    // 去除 key 1
//  * cache.get(1);       // 返回 -1 (未找到 key 1)
//  * cache.get(3);       // 返回 3
//  * cache.get(4);       // 返回 4
//  *
//  */

package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var cache LFUCache
	var s string
	ops := []string{"LFUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"}
	opcsString := "[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]"
	opcs := make([]([]int), 0)
	for _, s := range strings.Split(opcsString[2:len(opcsString)-2], "],[") {
		t := strings.Split(s[0:len(s)], ",")
		tt := make([]int, 0)
		for _, ttt := range t {
			fuck, _ := strconv.Atoi(ttt)
			tt = append(tt, fuck)
		}
		opcs = append(opcs, tt)
	}
	fmt.Println(opcs)

	for i := range ops {
		if ops[i] == "LFUCache" {
			s = fmt.Sprintf("%v: Constructor(%v)", i, opcs[i][0])
			cache = Constructor(opcs[i][0])
		} else if ops[i] == "put" {
			s = fmt.Sprintf("%v: Put(%v, %v)", i, opcs[i][0], opcs[i][1])
			cache.Put(opcs[i][0], opcs[i][1])
		} else if ops[i] == "get" {
			v := cache.Get(opcs[i][0])
			s = fmt.Sprintf("%v: Get(%v) -> %v", i, opcs[i][0], v)
		}
		if i > 0 && i < 18 {
			fmt.Print(s, ", ")
			cache.Debug()
			fmt.Println("")
		}
	}
}

func (this *LFUCache) Debug() {
	fmt.Print("capacity:", this.capacity)
	fmt.Printf(", router: [")
	for k, v := range this.router {
		fmt.Printf(" %v:(%v,%v),", k, v.Value.(*vNode).k, v.Value.(*vNode).v)
	}
	fmt.Println("]")
	for f := this.fList.Front(); f != nil; f = f.Next() {
		fmt.Printf("    %v", f.Value.(*fNode).f)
		for v := f.Value.(*fNode).vList.Front(); v != nil; v = v.Next() {
			fmt.Printf(" -> (%v,%v)", v.Value.(*vNode).k, v.Value.(*vNode).v)
		}
		fmt.Println("\n    |")
	}
}

// @see https://leetcode-cn.com/problems/lfu-cache/solution/c-ji-yu-unordered_maplist-shi-xian-o1-shi-jian-fu-/

// @lc code=start
type vNode struct {
	k        int
	v        int
	fReqItor *list.Element
}

type fNode struct {
	f     int
	vList *list.List
}

type LFUCache struct {
	capacity int
	router   map[int]*list.Element
	fList    *list.List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{capacity, make(map[int]*list.Element), list.New()}
}

func (this *LFUCache) Get(key int) int {
	ele, exist := this.router[key]
	if exist {
		this.upgradeFeq(ele)
		return ele.Value.(*vNode).v
	}
	return -1
}

func (this *LFUCache) FuckedOut() bool {
	flb := this.fList.Back()
	if flb == nil {
		return false
	}

	vl := flb.Value.(*fNode).vList
	delete(this.router, vl.Back().Value.(*vNode).k)
	if vl.Len() == 1 {
		this.fList.Remove(flb)
	} else {
		vl.Remove(vl.Back())
	}
	this.capacity++
	return true
}

func (this *LFUCache) upgradeFeq(v *list.Element) {
	vv := v.Value.(*vNode)
	fp := vv.fReqItor.Prev()
	f := vv.fReqItor
	newF := f.Value.(*fNode).f + 1

	if fp == nil {
		vl := list.New()
		vl.PushFront(vv)
		this.router[vv.k] = vl.Front()
		this.fList.PushFront(&fNode{newF, vl})
		vv.fReqItor = this.fList.Front()
	} else if fp.Value.(*fNode).f == newF {
		fp.Value.(*fNode).vList.PushFront(vv)
		this.router[vv.k] = fp.Value.(*fNode).vList.Front()
		vv.fReqItor = fp
	} else {
		vl := list.New()
		vl.PushFront(vv)
		this.router[vv.k] = vl.Front()
		this.fList.InsertBefore(&fNode{newF, vl}, f)
		vv.fReqItor = fp.Next()
	}
	if f.Value.(*fNode).vList.Len() == 1 {
		this.fList.Remove(f)
	} else {
		f.Value.(*fNode).vList.Remove(v)
	}
}

func (this *LFUCache) Put(key int, value int) {
	v, exist := this.router[key]
	if exist {
		v.Value.(*vNode).v = value
		this.upgradeFeq(v)
	} else {
		if this.capacity == 0 {
			if !this.FuckedOut() {
				return
			}
		}

		vv := &vNode{key, value, nil}
		flb := this.fList.Back()
		if flb == nil || flb.Value.(*fNode).f != 1 {
			vl := list.New()
			vl.PushFront(vv)
			this.router[key] = vl.Front()
			this.fList.PushBack(&fNode{1, vl})
			vv.fReqItor = this.fList.Back()
		} else {
			vl := flb.Value.(*fNode).vList
			vl.PushFront(vv)
			vv.fReqItor = flb
			this.router[key] = vl.Front()
		}
		this.capacity--
	}
}

// @lc code=end

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
