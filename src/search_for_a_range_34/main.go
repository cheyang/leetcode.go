package search_for_a_range_34

import (

)

func SearchRange(array []int, target int) (result []int){
	
	tempArray := array
	
	mid := len(tempArray)/2
	
	result = []int{-1, -1}
	
	for mid > -1 || mid + 1 < len(tempArray){
		
		part1 := tempArray[0:mid]
		
		part2 := tempArray[mid:len(tempArray)]
		
		
		if target >= part2[0] {
			
			tempArray = part2
			
		}else{
			tempArray = part1
		}
		
		
		mid = len(tempArray)/2
	}
	
	
	
}