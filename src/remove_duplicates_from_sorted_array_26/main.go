//package remove_duplicates_from_sorted_array_26 

package main

import (
	"fmt"
)

func main(){
	fmt.Println(RemoveDuplicates([]int{1,1,2}))
	
	fmt.Println(RemoveDuplicates([]int{1,1,2,3,3}))
}

func RemoveDuplicates(array []int) (length int){
	
	if array == nil|| len(array) == 0 {
		return
	}
	
	pos := 0
	
	currentValue := array[0]
	
	length = 1
	
	for pos < len(array){
		
		if array[pos] != currentValue {
			length ++
			
			currentValue = array[pos]
		}
		
		
		pos ++
	}
	
	return
}
