func singleNonDuplicate(nums []int) int {
	leftIndex, rightIndex := 0, len(nums)-1
	resultIndex := 0

	for leftIndex <= rightIndex {
		middleIndex := leftIndex + (rightIndex-leftIndex)/2
		if moveLeft(nums, middleIndex) {
			resultIndex = middleIndex
			rightIndex = middleIndex - 1
		} else {
			leftIndex = middleIndex + 1
		}
	}

	return nums[resultIndex]
}

func moveLeft(nums []int, index int) bool {
	if index == len(nums)-1 {
		return true
	}

	if index%2 == 1 {
		return nums[index] != nums[index-1]
	}

	return nums[index] != nums[index+1]
}
