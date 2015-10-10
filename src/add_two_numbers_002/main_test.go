package add_two_numbers_002

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_AddTwoNumbers(t *testing.T) {
	Convey("Add two numbers", t, func() {
		ln1 := &linkedNode{
			value: 2,
			next: &linkedNode{
				value: 4,
				next: &linkedNode{
					value: 3,
				},
			},
		}

		l1 := linkedList{first: ln1}

		ln2 := &linkedNode{
			value: 5,
			next: &linkedNode{
				value: 6,
				next: &linkedNode{
					value: 4,
				},
			},
		}

		l2 := linkedList{first: ln2}

		list := AddTwoNumbers(l1, l2)

		node := list.first

		So(node.value, ShouldEqual, 7)
		So(node.next, ShouldNotBeNil)
		node = node.next
		So(node.value, ShouldEqual, 0)
		So(node.next, ShouldNotBeNil)
		node = node.next
		So(node.value, ShouldEqual, 8)
		So(node.next, ShouldBeNil)
	})
}
