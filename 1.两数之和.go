/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package main

func main() {

}

// @lc code=start
func twoSum(nums []int, target int) []int {
	l := len(nums)
	w := make(map[int]int, l)

	for i := 0; i < l; i++ {
		need := target - nums[i]
		v, ok := w[need]
		if ok {
			return []int{v, i}
		}
		w[nums[i]] = i // who am i
	}
	return []int{}
}

// @lc code=end
