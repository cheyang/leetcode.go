package longest_substring_without_repeating_characters_003

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	Convey("longest substring 1", t , func(){
			
				strs,length := lengthOfLongestSubstring("abcabcbb")
				So(length, ShouldEqual, 2)
				So(strs, ShouldEqual, "abcd")
				
			})
	
	Convey("longest substring 2", t , func(){
			
				strs,length := lengthOfLongestSubstring("bbbbb")
				So(length, ShouldEqual, 1)
				So(strs, ShouldEqual, "bd")
				
			})
		
		
	
}


