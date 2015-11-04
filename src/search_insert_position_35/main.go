package search_insert_position_35

import ()

func SearchInsertTarget(nums []int, target int) (pos int) {

	pos = -1

	mid := len(nums) / 2

	start := 0

	end := len(nums) - 1

	for end-start >= 0 {

		switch {
		case nums[mid] == target:
			pos = mid
			break
		case target > nums[end]:
			pos = end + 1
			break
		case target == nums[end]:
			pos = end
			break
		case target <= nums[start]:
			pos = start
			break
	    default:
		}

		if nums[mid] > target {
			end = mid
		} else {
			start = mid
		}

		mid = (start + end) / 2
	}

	if pos == -1 {
		switch {
		case nums[mid] == target:
			pos = mid
		case mid == 0:
			pos = mid + 1
		case mid == len(nums)-1:
			pos = mid - 1
		default:
			if nums[mid] > target {
				pos = mid - 1
			} else {
				pos = mid + 1
			}

		}
	}

	return pos
}
