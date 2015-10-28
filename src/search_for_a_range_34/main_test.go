package search_for_a_range_34

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func Test_SearchRange(t *testing.T) {
		
		Convey("Search Range ", t, func(){
				
					Convey("Search Range middle", func(){
							result := SearchRange([]int{5, 7, 7, 8, 8, 10}, 8)
							So(result, ShouldResemble, []int{3,4})
							
							})
					
					Convey("Search Range 2", func(){
							result := SearchRange([]int{5, 7, 7, 8, 8, 10}, 5)
							So(result, ShouldResemble, []int{0,-1})
							
							})
					
					Convey("Search Range 3", func(){
							result := SearchRange([]int{5, 7, 7, 8, 8, 10}, 7)
							So(result, ShouldResemble, []int{1,2})
							
							})
				})
}

