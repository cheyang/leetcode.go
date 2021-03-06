package plus_one_66

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_PlusOne(t *testing.T) {

	Convey("Plus One", t, func() {
		Convey("Plus One without carry-over", func() {

			result := PlusOne([]int{1, 2, 3, 4, 5})

			expect_data := []int{1, 2, 3, 4, 6}

			So(expect_data[0], ShouldEqual, result[0])
			So(expect_data[1], ShouldEqual, result[1])
			So(expect_data[2], ShouldEqual, result[2])
			So(expect_data[3], ShouldEqual, result[3])
			So(expect_data[4], ShouldEqual, result[4])
		})

		Convey("Plus One with carry-over", func() {

			result := PlusOne([]int{9, 9, 9, 9})

			expect_data := []int{1, 0, 0, 0, 0}

			So(expect_data[0], ShouldEqual, result[0])
			So(expect_data[1], ShouldEqual, result[1])
			So(expect_data[2], ShouldEqual, result[2])
			So(expect_data[3], ShouldEqual, result[3])
			So(expect_data[4], ShouldEqual, result[4])
		})

	})

}
