package longest_substring_without_repeating_characters_003

import ()

type LongSubstring struct {
	start, end, length int
}

//https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(str string) (longestStrs []string, longestLength int) {
	
	longestStrs = make(string[], 1, len(str))

	// key is byte, value is last pos during iterate
	word_map := make(map[byte]int, len(str))

	var longestSubstringSlice []LongSubstring = make([]LongSubstring, len(str))

	strSlice := []byte(str)

	startPos := 0

	endPos := 0

	for i, k := range strSlice {

		if v, ok := word_map[k]; ok {

			// if the duplicate element is between current startPos and endPos, then the current length of the no duplicate substring is endPos - startpos + 1
			if startPos <= v {

				substringCandidate := LongSubstring{start: startPos, end: endPos, length: endPos - startPos + 1}

				if substringCandidate.length > longestLength {

					longestSubstringSlice = make([]LongSubstring, len(str))

					longestSubstringSlice[0] = substringCandidate
				} else if substringCandidate.length == longestLength {
					append(longestSubstringSlice, substringCandidate)
				}

				startPos += 1
			}

			if len(str)-startPos < longestLength {
				break
			}

			word_map[k] = v

			endPos += 1

		}

	}
	
	encountered := map[string]bool{}
	
	for i, longestSubstring := range longestSubstringSlice{
		
		tempStr := string(longestStrsstrSlice[longestSubstring.start:longestSubstring.end+1])
		
		if v, ok := encountered[tempStr]; ok == false{
			append(longestStrs, tempStr)
			encountered[tempStr] = true
		}
		
	}

	return longestStrs, longestLength

}
