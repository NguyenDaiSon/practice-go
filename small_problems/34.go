func findIndex(nums []int, target int, firstIndex bool) int {
	firstPos, lastPos := 0, len(nums)-1
	expectedPos := -1
	for firstPos <= lastPos {
		middlePos := firstPos + (lastPos-firstPos)/2
		if target == nums[middlePos] {
			expectedPos = middlePos
			if firstIndex {
				lastPos = middlePos - 1
			} else {
				firstPos = middlePos + 1
			}
		} else if target < nums[middlePos] {
			lastPos = middlePos - 1
		} else {
			firstPos = middlePos + 1
		}
	}
	return expectedPos
}

func searchRange(nums []int, target int) []int {
	firstIndex := findIndex(nums, target, true)
	lastIndex := findIndex(nums, target, false)
	return []int{firstIndex, lastIndex}
}
