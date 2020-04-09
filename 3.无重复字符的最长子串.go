/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (33.58%)
 * Likes:    3392
 * Dislikes: 0
 * Total Accepted:    412.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *

 */

// 滑动窗口

package main

import (
	"fmt"
	"strings"
)

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	resLength := 0
	left := 0
	right := 0
	str := s[left:right]
	for ; right < len(s); right++ {
		if index := strings.Index(str, string(s[right])); index != -1 {
			left += index + 1
		}
		str = s[left : right+1]
		if len(str) > resLength {
			resLength = len(str)
		}
	}
	return resLength
}

// @lc code=end

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
