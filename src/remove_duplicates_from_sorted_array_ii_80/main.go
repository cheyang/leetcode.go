package remove_duplicates_from_sorted_array_ii_80

import (

)


func RemoveDuplicates(array []int)(length int){
	
	if array == nil || len(array) ==0{
		return 0
	}
	
	duplicateMap := make(map[int]int, len(array))
	
	pos := 0
	
	for pos < len(array){
		
		if v, ok := duplicateMap[array[pos]]; ok{
			if v <= 2{
				length ++
			}
			
			v++
			
			duplicateMap[array[pos]] = v
			
		}else{
			duplicateMap[array[pos]] = 1
		}
		
		
		pos ++
	}
	
	return 
	
}