package two_sum_001

import (
	"errors"
	"strconv"
)

func TwoSum(array []int, target int) ([2]int, error) {

	indexes := [2]int{-1, -1}

	found := false

	for i, v := range array {
		search_v := target - v

		if pos, e := contain(array, search_v); e == nil {
			indexes[0] = i + 1
			indexes[1] = pos + 1
			found = true
			break
		}

	}

	if found {
		return indexes, nil
	} else {
		return indexes, errors.New("Not found two sum to equal " + strconv.Itoa(target))
	}
}

func contain(array []int, value int) (int, error) {

	found := false

	pos := -1

	for i, v := range array {
		if v == value {
			pos = i
			found = true
			break
		}
	}

	if found {
		return pos, nil
	} else {
		return pos, errors.New("Failed to find " + strconv.Itoa(value) + " in array ")
	}

}
