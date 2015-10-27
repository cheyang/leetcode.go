package search_for_a_range_34

import (

)

func SearchRange(array []int, target int) (result []int){
	
	
	
	mid := len(array)/2
	
	start := 0
	
	end := len(array) - 1
	
	result = []int{-1, -1}
	
	for mid>start || mid<end {
		
		if array[mid] > target{
			
			end = mid -1
			
		}else{
			
			start = mid 
		}
		
		mid = start + (end - start)/2
	}
	
	
	switch target {
		case array[start]:
			if start - 1 >= 0 && array[start - 1] == target{
				result[0] = start -1
				result[1] = start
			} else if start + 1 < len(array) && array[start+1] == target{
				result[0] = start
				result[1] = start + 1
			}else{
				result[0] = start
			}
		case array[end]:
			if end + 1 < len(array) && array[end+1] == target {
				result[0] = end
				result[1] = end + 1
			}else if end - 1 < len(array) && array[end -1] == target{
				result[0] = end - 1
				result[1] = end 
			}else{
				result[0] = end
			}
	}
	
	return
}