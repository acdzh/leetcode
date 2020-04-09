/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (49.58%)
 * Likes:    1126
 * Dislikes: 0
 * Total Accepted:    83K
 * Total Submissions: 165K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 *
 * 示例:
 *
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 *
 */
package main

import "fmt"

func main() {
	fmt.Println(trap([]int{2, 0, 2}))
}

// @lc code=start
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func trap(height []int) int {
	l := len(height)
	if l == 0 {
		return 0
	}
	maxL := make([]int, l)
	maxR := make([]int, l)
	maxL[0] = 0
	maxR[l-1] = 0
	for i := 1; i < l; i++ {
		if height[i-1] > maxL[i-1] {
			maxL[i] = height[i-1]
		} else {
			maxL[i] = maxL[i-1]
		}
		if height[l-i] > maxR[l-i] {
			maxR[l-1-i] = height[l-i]
		} else {
			maxR[l-1-i] = maxR[l-i]
		}
	}
	sum := 0
	var rain int
	for i := 0; i < l; i++ {
		rain = min(maxL[i], maxR[i]) - height[i]
		if rain > 0 {
			sum += rain
		}
	}
	return sum
}

// @lc code=end
