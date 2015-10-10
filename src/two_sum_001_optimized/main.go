package two_sum_001_optimzed

import (
	"errors"
	"strconv"
)


// https://leetcode.com/problems/two-sum/
// O(n)
func TwoSum(array []int, target int) ([2]int, error) {

	indexes := [2]int{-1, -1}

	found := false
	
	search_map := make(map[int]int, len(array))

	for i, v := range array {
		
		search_map[v]=i + 1

	}
	
	for k, v := range search_map{
		another_k := target - k
		
		if another_v, ok := search_map[another_k]; ok {
			
			if (another_v > v){
				indexes[0] = v
				indexes[1] = another_v
			}else{
				indexes[0] = another_v 
				indexes[1] = v 
			}
			found = true
			break
		}
		
	}
	

	if found {
		return indexes, nil
	} else {
		return indexes, errors.New("Not found two sum to equal " + strconv.Itoa(target))
	}
}

func contain(array []int, value int) (int, error) {

	found := false

	pos := -1

	for i, v := range array {
		if v == value {
			pos = i
			found = true
			break
		}
	}

	if found {
		return pos, nil
	} else {
		return pos, errors.New("Failed to find " + strconv.Itoa(value) + " in array ")
	}

}
