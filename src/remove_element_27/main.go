package remove_element_27

import ()

func RemoveElement(array []int, value int) int {

	

	var startPos, currentLen = 0, 0


	for startPos < len(array) {

		if array[startPos] != value {

			currentLen ++

		}

		startPos++
	}

	return currentLen
}
