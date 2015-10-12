package longest_substring_without_repeating_characters_003

import (
	"fmt"
	"strconv"
)

type LongSubstring struct {
	start, end, length int
}

//https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(str string) (longestStrs []string, longestLength int) {

	longestStrs = make([]string, 1, len(str))

	// key is byte, value is last pos during iterate
	word_map := make(map[byte]int, len(str))

	var longestSubstringSlice []LongSubstring = make([]LongSubstring, 1)

	strSlice := []byte(str)

	fmt.Println(strSlice)

	startPos := 0

	endPos := 0

	for i, k := range strSlice {

		fmt.Println("For order ", i, " value k", string(k))

		if v, ok := word_map[k]; ok {

			fmt.Println("found last pos of word_map [", string(k), "] is ", strconv.Itoa(v))

			// if the duplicate element is between current startPos and endPos, then the current length of the no duplicate substring is endPos - startpos + 1
			if startPos <= v {
				
				var substringCandidate LongSubstring

				if strSlice[startPos] == strSlice[endPos] && startPos != endPos  {
					substringCandidate = LongSubstring{start: startPos, end: endPos - 1, length: endPos - startPos}
				} else {				
					substringCandidate = LongSubstring{start: startPos, end: endPos, length: endPos - startPos + 1}
				}

				fmt.Println(string(strSlice[substringCandidate.start : substringCandidate.end+1]))

				if substringCandidate.length > longestLength {
					

					longestSubstringSlice = make([]LongSubstring, 1)

					longestSubstringSlice[0] = substringCandidate
					

					longestLength = substringCandidate.length
					
					fmt.Println("Make longestSubstringSlice: ", longestSubstringSlice, " with the length ",strconv.Itoa(longestLength) )
				} else if substringCandidate.length == longestLength {
					longestSubstringSlice = append(longestSubstringSlice, substringCandidate)
					
					fmt.Println("Append ", string(strSlice[substringCandidate.start : substringCandidate.end+1]))
				}

				startPos = v + 1
			}

			if len(str)-startPos < longestLength {
				break
			}

		} else {
			fmt.Println("Not found value of word_map [", string(k), "]")
		}

		word_map[k] = i

		endPos = i

	}

	encountered := map[string]bool{}

	for _, longestSubstring := range longestSubstringSlice {

		tempStr := string(strSlice[longestSubstring.start : longestSubstring.end+1])

		if _, ok := encountered[tempStr]; ok == false {
			longestStrs = append(longestStrs, tempStr)
			encountered[tempStr] = true
		}

	}

	return longestStrs, longestLength

}
