/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (26.26%)
 * Likes:    1966
 * Dislikes: 0
 * Total Accepted:    190.6K
 * Total Submissions: 724.9K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？请你找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例：
 *
 * 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	result := [][]int{}
	sort.Ints(nums)
	if nums[0] > 0 {
		return [][]int{}
	}
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			if nums[i]+nums[l]+nums[r] > 0 {
				r--
				continue
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
				continue
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for ; l < r && nums[l] == nums[l+1]; l++ {
				}
				for ; l < r && nums[r] == nums[r-1]; r-- {
				}
				l++
				r--
			}
		}
	}
	return result
}

// @lc code=end
