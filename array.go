package main

import "fmt"

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



