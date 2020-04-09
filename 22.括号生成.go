/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (73.76%)
 * Likes:    878
 * Dislikes: 0
 * Total Accepted:    101.8K
 * Total Submissions: 136.7K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例：
 *
 * 输入：n = 3
 * 输出：[
 * ⁠      "((()))",
 * ⁠      "(()())",
 * ⁠      "(())()",
 * ⁠      "()(())",
 * ⁠      "()()()"
 * ⁠    ]
 *
 *
 */
package main

func main() {
	generateParenthesis(3)
}

// @lc code=start
func generateParenthesis_dp(n int) []string {
	lTotal := make([](map[string]bool), n+1)
	lTotal[0] = map[string]bool{"": true}
	lTotal[1] = map[string]bool{"()": true}
	for i := 2; i < n+1; i++ {
		lTotal[i] = map[string]bool{}
		for p := 1; p < i; p++ {
			q := i - p // AB <- A, B
			for k1, _ := range lTotal[p] {
				for k2, _ := range lTotal[q] {
					lTotal[i][k1+k2] = true
				}
			}
		}
		// (C) <- C
		for k := range lTotal[i-1] {
			lTotal[i]["("+k+")"] = true
		}
	}
	results := []string{}
	for k := range lTotal[n] {
		results = append(results, k)
	}
	return results
}

// @lc code=start
func generateParenthesis(n int) []string {
	var dfs func(str string, left, right, n int, ret []string) []string
	dfs = func(str string, left, right, n int, ret []string) []string {
		if len(str) == 2*n {
			return append(ret, str)
		}
		if left < n {
			ret = dfs(str+"(", left+1, right, n, ret)
		}
		if right < left {
			ret = dfs(str+")", left, right+1, n, ret)
		}
		return ret
	}
	ret := make([]string, 0)
	ret = dfs("(", 1, 0, n, ret)
	return ret
}

// @lc code=end
