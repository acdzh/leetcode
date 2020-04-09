/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (29.18%)
 * Likes:    1950
 * Dislikes: 0
 * Total Accepted:    223K
 * Total Submissions: 764.2K
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1：
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 *
 * 示例 2：
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
 *
 */
package main

import "fmt"

// @lc code=start
func longestPalindrome(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}
	maxL, maxR, maxLen := 0, 0, 0
	sLen := len(s)
	dp := make([][]bool, sLen)
	for i, _ := range dp {
		dp[i] = make([]bool, sLen)
	}
	for r := 1; r < sLen; r++ {
		for l := 0; l < r; l++ {
			// @comment1
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r-l > maxLen {
					maxLen = r - l
					maxL = l
					maxR = r
				}
			}
		}
	}
	// fmt.Println(dp)
	return s[maxL : maxR+1]
}

// @lc code=end

func main() {
	fmt.Println(longestPalindrome("ss"))
}

// @comment1:
// dp[l][r]	= dp[l+1][r-1] && s[l] == s[r]	, l+1 < r-1
// 			= s[l] == s[r]					, l+1 >= r-1
