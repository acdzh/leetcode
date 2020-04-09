/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 *
 * https://leetcode-cn.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (27.06%)
 * Likes:    1065
 * Dislikes: 0
 * Total Accepted:    64.8K
 * Total Submissions: 239.6K
 * Testcase Example:  '"aa"\n"a"'
 *
 * 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
 *
 * '.' 匹配任意单个字符
 * '*' 匹配零个或多个前面的那一个元素
 *
 *
 * 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
 *
 * 说明:
 *
 *
 * s 可能为空，且只包含从 a-z 的小写字母。
 * p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
 *
 *
 * 示例 1:
 *
 * 输入:
 * s = "aa"
 * p = "a"
 * 输出: false
 * 解释: "a" 无法匹配 "aa" 整个字符串。
 *
 *
 * 示例 2:
 *
 * 输入:
 * s = "aa"
 * p = "a*"
 * 输出: true
 * 解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
 *
 *
 * 示例 3:
 *
 * 输入:
 * s = "ab"
 * p = ".*"
 * 输出: true
 * 解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
 *
 *
 * 示例 4:
 *
 * 输入:
 * s = "aab"
 * p = "c*a*b"
 * 输出: true
 * 解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
 *
 *
 * 示例 5:
 *
 * 输入:
 * s = "mississippi"
 * p = "mis*is*p*."
 * 输出: false
 *
 */
package main

import "fmt"

func main() {
	dfa := getDFA("a*a*a")
	debug(dfa)
	fmt.Println(check(dfa, "aaaaaaaaaaaaab", 0))
}

func debug(p *Node) {
	h := make(map[int]bool)
	debugCore(p, &h)
}

func debugCore(p *Node, hasVisited *map[int]bool) {
	if (*hasVisited)[p.id] {
		return
	}
	(*hasVisited)[p.id] = true
	debugNode(p)
	for _, v := range p.router {
		for _, vv := range v {
			debugCore(vv, hasVisited)
		}
	}
}

func debugNode(p *Node) {
	par := ""
	if p.parentNode == nil {
		par = "<nil>"
	} else {
		for _, pp := range p.parentNode {
			if pp != nil {
				par += fmt.Sprintf("%v (%c) | ", pp.id, pp.C)
			}
		}
	}
	fmt.Printf("%v, '%c', isEnd: %v, parent: %v\n", p.id, p.C, p.isEnd, par)
	for k, v := range p.router {
		for _, vv := range v {
			fmt.Printf("    '%c' -> %v ('%c')\n", k, vv.id, vv.C)
		}
	}
}

// @lc code=start
type Node struct {
	C          byte
	isEnd      bool
	router     map[byte]([]*Node)
	parentNode []*Node
	canBack    bool
}

func (node *Node) appendRouter(c byte, next *Node) {
	node.router[c] = append(node.router[c], next)
}

func getDFA(p string) *Node {
	pLen := len(p)
	begin := Node{'>', true, make(map[byte]([]*Node)), []*Node{nil}, false}
	last := &begin
	for i := 0; i < pLen; i++ {
		if p[i] == '.' {
			t := &Node{'.', true, make(map[byte]([]*Node)), []*Node{last}, false}
			last.appendRouter('.', t)
			last.isEnd = false
			if last.canBack {
				for _, pp := range last.parentNode {
					pp.appendRouter('.', t)
					t.parentNode = append(t.parentNode, pp)
				}
			}
			last = t
		} else if p[i] == '*' {
			last.canBack = true
			last.appendRouter(last.C, last)
			if i == pLen-1 {
				for _, p := range last.parentNode {
					p.isEnd = true
				}
			}
		} else {
			t := &Node{p[i], true, make(map[byte]([]*Node)), []*Node{last}, false}
			last.appendRouter(p[i], t)
			last.isEnd = false
			if last.canBack {
				for _, pp := range last.parentNode {
					pp.appendRouter(p[i], t)
					t.parentNode = append(t.parentNode, pp)
				}
			}
			last = t
		}
	}
	return &begin
}

func check(p *Node, s string, i int) bool {
	if i > len(s) {
		return false
	}

	if i == len(s) {
		if p.isEnd {
			return true
		} else {
			return false
		}
	}

	for k, v := range p.router {
		if k != s[i] && k != '.' {
			continue
		}
		for _, vv := range v {
			if check(vv, s, i+1) {
				return true
			}
		}
	}
	return false
}

func isMatch(s string, p string) bool {
	return check(getDFA(p), s, 0)
}

// @lc code=end
