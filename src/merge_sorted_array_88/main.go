package merge_sorted_array_88
//package main

import (
		"fmt"
	)

func main(){
	
	fmt.Println(MergeSortedArray([]int{1,3,5,7,9,0,0,0,0,0},5, []int{2,4,6,8,10},5))
	
	fmt.Println(MergeSortedArray([]int{1,3,5,7,0,0,0,0,0,0},4, []int{2,4,6,8,9,10},6))
}


/**
*Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
 */
func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) []int {
	
	pos1 := m - 1
	
	pos2 := n - 1
	
	i := n+m-1
	

	for i > -1 {
		
		if pos2 < 0 {
			break
		}else if pos1 < 0{
			copy(nums1[0:], nums2[0:pos2+1])
			break
		}
		
		
		if nums1[pos1] > nums2[pos2]{
			nums1[i] = nums1[pos1]
			
			pos1 --
		}else{
			nums1[i] = nums2[pos2]
			
			pos2 --
		}
		
		i--
	}

	return nums1
}
