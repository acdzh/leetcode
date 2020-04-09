/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (37.17%)
 * Likes:    2382
 * Dislikes: 0
 * Total Accepted:    167.1K
 * Total Submissions: 449.6K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 *
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 *
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * 则中位数是 2.0
 *
 *
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * 则中位数是 (2 + 3)/2 = 2.5
 *
 *
 */

package main

import "fmt"

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 == 0 {
		if l2%2 == 1 {
			return float64(nums2[l2/2])
		} else {
			return float64(nums2[l2/2-1]+nums2[l2/2]) / 2
		}
	}
	if l2 == 0 {
		if l1%2 == 1 {
			return float64(nums1[l1/2])
		} else {
			return float64(nums1[l1/2-1]+nums1[l1/2]) / 2
		}
	}
	i1, i2 := 0, 0
	l := (l1 + l2) / 2
	nums := make([]int, l+2)
	for i1 != l1 || i2 != l2 {
		if i1+i2 > l {
			break
		}
		if i1 == l1 {
			nums[i1+i2] = nums2[i2]
			i2++
		} else if i2 == l2 {
			nums[i1+i2] = nums1[i1]
			i1++
		} else {
			v1, v2 := nums1[i1], nums2[i2]
			if v1 > v2 {
				nums[i1+i2] = v2
				i2++
			} else if v1 == v2 {
				nums[i1+i2], nums[i1+i2+1] = v1, v2
				i1++
				i2++
			} else {
				nums[i1+i2] = v1
				i1++
			}
		}
	}
	if (l1+l2)%2 == 1 {
		return float64(nums[l])
	} else {
		return float64(nums[l-1]+nums[l]) / 2.0
	}
}

// @lc code=end

func main() {
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
}
