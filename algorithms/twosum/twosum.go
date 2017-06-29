// Source : https://leetcode.com/problems/two-sum/#/description
// Author : thinkerou
// Date   : 2017-06-29

/*
 *
 * Given an array of integers, return indices of the two numbers such that they add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may not use the same element twice.
 *
 * Example:
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9
 * return [0, 1].
 *
 */

package twosum

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if t == nums[j] {
				var arrRes []int
				arrRes = append(arrRes, i)
				arrRes = append(arrRes, j)
				return arrRes
			}
		}
	}
	return nil
}
