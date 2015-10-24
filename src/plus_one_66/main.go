package plus_one_66

import (

)


func PlusOne(input []int) []int{
	
	carry := 1
	
	var output []int = make([]int, len(input)+1)
	
	for i := len(input) - 1 ;i > -1; i--{
		
		value := input[i] + carry
		
		if value >= 10 {
			carry = 1
			output[i+1] = value - 10
		}else{
			carry = 0
			break
		}
	}
	
	var result []int  
	
	if carry == 1{
		output[0] = 1
		
		result = output
	}else{
		result = output[1:]
	}
	
	
}
