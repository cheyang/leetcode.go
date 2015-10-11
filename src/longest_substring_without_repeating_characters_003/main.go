package main

import (
	"fmt"
	)

type LongSubstring struct {
	start, end, length int
}

//https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(str string) (longestStrs []string, longestLength int) {
	
	longestStrs = make([]string, 1, len(str))

	// key is byte, value is last pos during iterate
	word_map := make(map[byte]int, len(str))

	var longestSubstringSlice []LongSubstring = make([]LongSubstring, len(str))

	strSlice := []byte(str)
	
	fmt.Println(strSlice)

	startPos := 0

	endPos := 0

	for i, k := range strSlice {

		if v, ok := word_map[k]; ok {

			// if the duplicate element is between current startPos and endPos, then the current length of the no duplicate substring is endPos - startpos + 1
			if startPos <= v {

				substringCandidate := LongSubstring{start: startPos, end: endPos, length: endPos - startPos + 1}
				
				fmt.Println(substringCandidate)

				if substringCandidate.length > longestLength {

					longestSubstringSlice = make([]LongSubstring, len(str))

					longestSubstringSlice[0] = substringCandidate
				} else if substringCandidate.length == longestLength {
					longestSubstringSlice = append(longestSubstringSlice, substringCandidate)
				}

				startPos += 1
			}

			if len(str)-startPos < longestLength {
				break
			}

			word_map[k] = v

			endPos = i

		}

	}
	
	encountered := map[string]bool{}
	
	for _, longestSubstring := range longestSubstringSlice{
		
		tempStr := string(strSlice[longestSubstring.start:longestSubstring.end+1])
		
		if _, ok := encountered[tempStr]; ok == false{
			longestStrs = append(longestStrs, tempStr)
			encountered[tempStr] = true
		}
		
	}

	return longestStrs, longestLength

}

func main(){
	fmt.Println( lengthOfLongestSubstring("abcabcbb"))
	
	fmt.Println( lengthOfLongestSubstring("bbbbb"))
}
