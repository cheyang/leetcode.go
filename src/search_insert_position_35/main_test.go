package search_insert_position_35

import (
    "testing"
    
    . "github.com/smartystreets/goconvey/convey"
)

func Test_SearchInsertTarget(t *testing.T) {
	
	Convey("Search Insert Target Test", t , func(){
			
			Convey("compare [1,3,5,6], 5 → 2", func(){
					
					pos := SearchInsertTarget([]int{1,3,5,6}, 5)
					
					So(pos, ShouldEqual, 2)
					
					})
			
			Convey("compare [1,3,5,6], 3 → 2", func(){
					
					pos := SearchInsertTarget([]int{1,3,5,6}, 5)
					
					So(pos, ShouldEqual, 2)
					
					})
			
			Convey("compare [1,3,5,6], 4 → 2", func(){
					
					pos := SearchInsertTarget([]int{1,3,5,6}, 4)
					
					So(pos, ShouldEqual, 2)
					
					})
			
			Convey("compare [1,3,5,6], 2 → 1", func(){
					
					pos := SearchInsertTarget([]int{1,3,5,6}, 2)
					
					So(pos, ShouldEqual, 1)
					
					})
			
			Convey("compare [1,3,5,6], 7 → 4", func(){
					
					pos := SearchInsertTarget([]int{1,3,5,6}, 7)
					
					So(pos, ShouldEqual, 4)
					
					})
			
			Convey("compare [1,3,5,6], 0 → 0 ", func(){
					
					pos := SearchInsertTarget([]int{1,3,5,6}, 0)
					
					So(pos, ShouldEqual, 0)
					
					})
			
			
			Convey("compare [1,3,5,6], 6 → 3 ", func(){
					
					pos := SearchInsertTarget([]int{1,3,5,6}, 6)
					
					So(pos, ShouldEqual, 3)
					
					})
			
			})

}

