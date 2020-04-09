/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (37.01%)
 * Likes:    4125
 * Dislikes: 0
 * Total Accepted:    376.1K
 * Total Submissions: 1M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 *
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 *
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 * 示例：
 *
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 *
 *
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func fuck(cur int) (int, int) {
	if cur > 9 {
		return 1, cur - 10
	} else {
		return 0, cur
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var cur int
	var begin ListNode
	var p = &begin
	p1 := l1
	p2 := l2
	for p1 != nil || p2 != nil {
		if p1 == nil {
			cur = p2.Val + carry
			if cur > 9 {
				carry, cur = 1, cur-10
			} else {
				carry = 0
			}
			p.Next = &ListNode{cur, nil}
			p = p.Next
			p2 = p2.Next
		} else if p2 == nil {
			cur = p1.Val + carry
			if cur > 9 {
				carry, cur = 1, cur-10
			} else {
				carry = 0
			}
			p.Next = &ListNode{cur, nil}
			p = p.Next
			p1 = p1.Next
		} else {
			cur = p1.Val + p2.Val + carry
			if cur > 9 {
				carry, cur = 1, cur-10
			} else {
				carry = 0
			}
			p.Next = &ListNode{cur, nil}
			p = p.Next
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	if carry != 0 {
		p.Next = &ListNode{carry, nil}
	}
	return begin.Next
}

// @lc code=end
