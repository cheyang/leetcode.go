package two_sum_001

import (
	"testing"
	"fmt"
)

func Test_TwoSum(t *testing.T) {
	
	if	idxs,e := TwoSum([]int{2, 7, 11, 15}, 9); e != nil{
		t.Errorf("twoSum failed due to not found")
	}else{
		
		fmt.Println(idxs)
		if idxs != [2]int{1,2}{
			t.Errorf("twoSum is found, not exact. ")
		}
	}
		
}
