package search_insert_position_35

import (
	"fmt"
	)

func SearchInsertTarget(nums []int, target int) (pos int) {

	pos = -1

	mid := len(nums) / 2

	start := 0

	end := len(nums) - 1
	
	i := 0
	
	loop:
	for end-start >= 1 {
		
		fmt.Println(i, "start:", start, ", end:", end, ", mid:", mid, "target:", target)

		switch {
		case nums[mid] == target:
			pos = mid
			break loop
		case target > nums[end]:
			pos = end + 1
			break loop
		case target == nums[end]:
			pos = end
			break loop
		case target <= nums[start]:
			pos = start
			break loop
	    default:
	    	fmt.Println("do nothing")
		}

		if nums[mid] > target {
			end = mid
		} else {
			start = mid
		}

		mid = (start + end) / 2
	}

	if pos == -1 {
		switch {
		case nums[mid] == target:
			pos = mid
		case mid == 0:
			pos = mid + 1
		case mid == len(nums)-1:
			pos = mid - 1
		default:
			if nums[mid] > target {
				pos = mid - 1
			} else {
				pos = mid + 1
			}

		}
	}

	return pos
}
