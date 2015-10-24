package remove_duplicates_from_sorted_array_ii_80

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func Test_RemoveDuplicates(t *testing.T) {
		Convey("remove elements", t, func() {
				
				Convey("remove elements with duplicate data", func(){
						length := RemoveDuplicates([]int{1,1,1,2,2,3})

						So(length, ShouldEqual, 5)
						
						})
				
				Convey("remove elements without duplicate data", func(){
						length := RemoveDuplicates([]int{1,2,3})

						So(length, ShouldEqual, 3)
						
						})

		
	})
}

