package main

import (
	"fmt"
	"sort"
)

//Given a sorted array, remove the duplicates in-place such that each element appear only once and return the new length.
//Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
func removeDuplicates(nums []int) int {
	l := len(nums)
	p := 0
	for p < l-1 {
		if nums[p] == nums[p+1] {
			for i := p; i < l-1; i++ {
				nums[i] = nums[i+1]
			}
			l--
		} else {
			p++
		}
	}
	fmt.Println(nums, l)
	return l
}

//Say you have an array for which the ith element is the price of a given stock on day i.
//Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times).
//However, you may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
//hints: you can buy and sell at the same day
func maxProfit(prices []int) int {
	l := len(prices)
	sum := 0
	if l > 1 {
		for i := 1; i < l; i++ {
			if prices[i] > prices[i-1] {
				sum += prices[i] - prices[i-1]
			}
		}
	}
	return sum
}

//Rotate an array of n elements to the right by k steps.
//For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].
//hints
//Original List                   : 1 2 3 4 5 6 7
//After reversing all numbers     : 7 6 5 4 3 2 1
//After reversing first k numbers : 5 6 7 4 3 2 1
//After revering last n-k numbers : 5 6 7 1 2 3 4 --> Result
func rotate(nums []int, k int) {
	index, j, n := 0, 0, len(nums)
	for i := 1; i <= n; i++ {
		index = j + i*k%n
		nums[j], nums[index] = nums[index], nums[j]
		if j == index {
			j++
		}
		fmt.Println(nums)
	}
}

//Given an array of integers, find if the array contains any duplicates.
//Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.
// hints
// use sort package or map
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	if len(nums) < 2 {
		return false
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

//Given an array of integers, every element appears twice except for one. Find that single one.
//Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
//hints
//we can XOR all bits together to find the unique number.
func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums {
		a ^= v
	}
	return a
}

//Given two arrays, write a function to compute their intersection.
//Example:
//Given nums1 = [1, 2, 2, 1], nums2 = [1, 2, 3], return [1, 2].
//Note:
//Each element in the result should appear as many times as it shows in both arrays.
//The result can be in any order.
//Follow up:
//What if the given array is already sorted? How would you optimize your algorithm?
//What if nums1's size is small compared to nums2's size? Which algorithm is better?
//What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

func intersect(nums1 []int, nums2 []int) []int {
	var tmp_nums []int
	var a []int
	var b []int
	if len(nums1) <= len(nums2) {
		a = nums1
		b = nums2
	} else {
		a = nums2
		b = nums1
	}
	sort.Ints(a)
	sort.Ints(b)
	k := 0
	for _, v := range a {
		for i := k; i < len(b); i++ {
			if v == b[i] {
				tmp_nums = append(tmp_nums, v)
				k = i + 1
				break
			}
		}
	}
	return tmp_nums
}

//Given a non-empty array of digits representing a non-negative integer, plus one to the integer.
//
//The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.
//
//You may assume the integer does not contain any leading zero, except the number 0 itself.
//
//Example 1:
//
//Input: [1,2,3]
//Output: [1,2,4]
//Explanation: The array represents the integer 123.
//Example 2:
//
//Input: [4,3,2,1]
//Output: [4,3,2,2]
//Explanation: The array represents the integer 4321.
func plusOne(digits []int) []int {
	len := len(digits)
	for i := len - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			break
		} else {
			digits[i] = 0
			if i == 0 {
				digits = append(digits, 1)
				for j := len - 1; j >= 0; j-- {
					digits[j], digits[j+1] = digits[j+1], digits[j]
				}
				break
			}
		}
	}
	return digits
}

//Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
//Example:
//
//Input: [0,1,0,3,12] -> [1,0,0,3,12] -> [1,3,0,0,12] -> [1,3,12,0,0]
//Output: [1,3,12,0,0]
//Note:
//
//You must do this in-place without making a copy of the array.
//Minimize the total number of operations.
func moveZeroes(nums []int) {
	index := 0
	flag := true
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && nums[index] == 0 {
			nums[index], nums[i] = nums[i], nums[index]
			i = index
			flag = true
		}
		if flag && nums[i] == 0 {
			index = i
			flag = false
		}
	}
}
