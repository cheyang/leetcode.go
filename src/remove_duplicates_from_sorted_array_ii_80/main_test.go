package remove_duplicates_from_sorted_array_ii_80

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func Test_RemoveDuplicates(t *testing.T) {
		Convey("remove elements", t, func() {

		length := RemoveDuplicates([]int{1,1,1,2,2,3})

		So(length, ShouldEqual, 5)
	})
}

