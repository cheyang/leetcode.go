package remove_element_27

import ()

func RemoveElement(array []int, value int) currentLen int {

	var startPos int = 0

	for startPos < len(array) {

		if array[startPos] != value {

			currentLen++

		}

		startPos++
	}

	return 
}
