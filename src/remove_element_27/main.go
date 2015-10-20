package remove_element_27

import ()

func RemoveElement(array []int, value int) int {

	var currentLength int = len(array)

	var startPos = 0

JUMP:
	for startPos < (currentLength - 1) {

		if array[startPos] == value {

			array[startPos] = array[currentLength-1]

			currentLength--

			goto JUMP
		}

		startPos++
	}

	return currentLength
}
