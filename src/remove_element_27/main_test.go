package remove_element_27

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func Test_RemoveElement(t *testing.T) {
	
	Convey("remove elements", t , func(){
			
			if	length := RemoveElement([]int{2, 7, 11, 15}, 9){
				
				So(length, ShouldEqual, 4)
				}
			
			})
		
	}
}


func Test_RemoveElement1(t *testing.T) {
	
	Convey("remove elements", t , func(){
			
			if	length := RemoveElement([]int{1, 2, 7, 9,  11, 15}, 9){
				
				So(length, ShouldEqual, 5)
				}
			
			})
		
	}
}


func Test_RemoveElement3(t *testing.T) {
	
	Convey("remove elements", t , func(){
			
			if	length := RemoveElement([]int{1, 9, 9, 9,  11, 9}, 9){
				
				So(length, ShouldEqual, 2)
				}
			
			})
		
	}
}