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
			for i:=p; i<l-1; i++ {
				nums[i] = nums[i+1]
			}
			l--
		} else {
			p++
		}
	}
	fmt.Println(nums,l)
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
		for i:=1;i<l;i++ {
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
func rotate(nums []int, k int)  {
	index,j,n := 0,0,len(nums)
	for i:=1;i<=n;i++{
		index = j + i * k % n
		nums[j],nums[index] = nums[index], nums[j]
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
	for i:=0;i<len(nums)-1;i++ {
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
	for _,v := range nums {
		a ^= v
	}
	return a
}
