package remove_duplicates_from_sorted_array_26

import (
    "testing"   
    . "github.com/smartystreets/goconvey/convey"
)

func Test_RemoveDuplicates(t *testing.T) {
	Convey("remove Duplicated", t, func() {

		length := RemoveDuplicates([]int{1,1,2})

		So(length, ShouldEqual, 2)
	})
}

func Test_RemoveDuplicates1(t *testing.T) {
	Convey("remove Duplicated 1", t, func() {

		length := RemoveDuplicates([]int{1,1,2,3,3})

		So(length, ShouldEqual, 3)
	})
}