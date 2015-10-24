package merge_sorted_array_88

import (
    "testing"
    
    . "github.com/smartystreets/goconvey/convey"
)

func Test_MergeSortedArray(t *testing.T) {
	
	Convey("Merge Sorted Array", t, func(){
			
			Convey("Merge Sorted Array 1", func(){
					
					result := MergeSortedArray([]int{1,3,5,7,9,0,0,0,0,0},5, []int{2,4,6,8,10},5)
					
					So(result, ShouldResemble, []int{1,2,3,4,5,6,7,8,9,10})
					
					})
			
			Convey("Merge Sorted Array 2", func(){
					
					result := MergeSortedArray([]int{1,3,5,7,0,0,0,0,0,0},4, []int{2,4,6,8,9,10},6)
					
					So(result, ShouldResemble, []int{1,2,3,4,5,6,7,8,9,10})
					
					})
			
			
			})
}

