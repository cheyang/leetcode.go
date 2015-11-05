package search_insert_position_35

import (
//	"fmt"
	)

func SearchInsertTarget(nums []int, target int) (pos int) {

	pos = -1

	
	start := 0

	end := len(nums) - 1
	
	mid := (start + end) / 2

	
	i := 0
	
	
	for end > start {
		
	//	fmt.Println(i, "start:", start, ",value:", nums[start], ", end:", end, ",value:",nums[end],", mid:", mid,",value:",nums[mid], "target:", target)
		
		if nums[mid] == target{
			return mid
		}
		
		if nums[mid] > target{
			
			end = mid -1
			
		}else{
			start = mid + 1
		}
		

		mid = (start + end) / 2
	}
	
	if nums[start] < target{
		pos = start + 1
	}else{
		pos = start
	}
	

	return pos
}
