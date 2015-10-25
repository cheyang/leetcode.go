package find_minimum_in_rotated_sorted_array_153


func FindMin(nums []int) (min int) {

	mid := len(nums) / 2

	array := nums

	for mid > 0 && mid+1 < len(array) {

		num_1 := smaller(array[0], array[mid])

		num_2 := smaller(array[mid+1], array[len(array)-1])

		if num_1 < num_2 {
			array = array[:mid+1]
		} else {
			array = array[mid+1:]
		}

		mid = len(array) / 2
	}

	if mid == 0 {
		min = array[mid]
	} else {
		min = smaller(array[0], array[mid])
	}

	return min
}

func smaller(a, b int) (v int) {

	v = a

	if a > b {
		v = b
	}

	return

}
