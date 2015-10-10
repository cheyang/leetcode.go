package add_two_numbers_002

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_AddTwoNumbers(t *testing.T) {
	Convey("Add two numbers", t, func() {
		ln1 := &linkedNode{
			Value: 2,
			Next: &linkedNode{
				Value: 4,
				Next: &linkedNode{
					Value: 3,
				},
			},
		}
		
		l1 := LinkedList{first: ln1}
		
		ln2 := &linkedNode{
			Value: 5,
			Next: &linkedNode{
				Value: 6,
				Next: &linkedNode{
					Value: 4,
				},
			},
		}
		
		l2 := LinkedList{first: ln2}
		
		list := AddTwoNumbers(l1, l2)
		
		node = list.first
		
		So(node.Value, ShouldEqual, 7)
		So(node.Next, ShouldNotBeNil)
		node = node.Next
		So(node.Value, ShouldEqual, 0)
		So(node.Next, ShouldNotBeNil)
		node = node.Next
		So(node.Value, ShouldEqual, 8)
		So(node.Next, ShouldBeNil)
	})
}
