package longest_substring_without_repeating_characters_003

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	Convey("longest substring 1", t , func(){
			
				strs,length := lengthOfLongestSubstring("abcabcbb")
				So(length, ShouldEqual, 3)
				So(strs[0], ShouldEqual, 'abc')
				So(strs[1], ShouldEqual, 'bca')
				So(strs[2], ShouldEqual, 'cab')
				
			})
	
	Convey("longest substring 2", t , func(){
			
				strs,length := lengthOfLongestSubstring("bbbbb")
				So(length, ShouldEqual, 1)
				So(strs[0], ShouldEqual, 'b')
				
			})
		
		
	
}


