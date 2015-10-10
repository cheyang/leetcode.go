package two_sum_001_optimzed

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)


func Test_TwoSum(t *testing.T) {
	
	Convey("Add Two Sum", t , func(){
			
			if	idxs,e := TwoSum([]int{2, 7, 11, 15}, 9); e != nil{
				So(e, ShouldBeNil)
				So(idxs[0], ShouldEqual, 1)
				So(idx[1], ShouldEqual, 2)
				}
			
			})
		
	}
		
}