/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 *
 * https://leetcode-cn.com/problems/edit-distance/description/
 *
 * algorithms
 * Hard (56.77%)
 * Likes:    746
 * Dislikes: 0
 * Total Accepted:    47.6K
 * Total Submissions: 80.9K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
 *
 * 你可以对一个单词进行如下三种操作：
 *
 *
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 *
 *
 * 示例 2：
 *
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 *
 *
 */
package main

import "fmt"

func main() {
	minDistance("horse", "ros")
	fmt.Println(min(0, 1, 1))
}

// @lc code=start
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func minDistance(wd1 string, wd2 string) int {
	l1, l2 := len(wd1)+1, len(wd2)+1
	dp := make([]([]int), l1)
	for i := 0; i < l1; i++ {
		dp[i] = make([]int, l2)
		dp[i][0] = i
	}
	for i := 0; i < l2; i++ {
		dp[0][i] = i
	}
	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			if wd1[i-1] == wd2[j-1] {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]-1) + 1
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	// fmt.Println("[")
	// for _, i := range dp {
	// 	fmt.Println("  ", i)
	// }
	// fmt.Println("]")

	return dp[l1-1][l2-1]
}

// @lc code=end
