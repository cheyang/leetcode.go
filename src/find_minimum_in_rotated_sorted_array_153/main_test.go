package find_minimum_in_rotated_sorted_array_153

import (
    "testing"
    
    . "github.com/smartystreets/goconvey/convey"
)

func Test_FindMin(t *testing.T) {
	
	Convey("Find MIN", t , func(){
			
			Convey("Find MIN in sorted array", func(){
					
					num := FindMin([]int{0,1,2,3,4,5,6,7,8,9,10})
					
					So(num, ShouldEqual, 0)
					})
			
			Convey("Find MIN in rotated sorted array1", func(){
					
					num := FindMin([]int{6,7,8,9,10,0,1,2,3,4,5})
					
					So(num, ShouldEqual, 0)
					})
			
			
			Convey("Find MIN in rotated sorted array2", func(){
					
					num := FindMin([]int{1,2,3,4,5,6,7,8,9,10,0})
					
					So(num, ShouldEqual, 0)
					})
			
			
			})
}

