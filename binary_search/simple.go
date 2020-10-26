package binary_search

func BSearch(arr []int, n, value int) int {
	low := 0
	high := n - 1

	for low <= high {
		mid := low + (high - low) >> 1
		if arr[mid] == value {
			return mid
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// 变体一：查找第一个值等于给定值的元素
func BSearchOne(arr []int, n, value int) int {
	var low = 0
	var high = n - 1

	for low <= high {
		var mid = low + (high - low) >> 1
		if arr[mid] < value {
			low = mid + 1
		} else if arr[mid] > value {
			high = mid - 1
		} else {
			if mid == 0 || arr[mid - 1] != value {
				return mid
			}
			high = mid - 1
		}
	}

	return - 1
}

// 变体二：查找最后一个值等于给定值的元素
func BSearchTwo(arr []int, n, value int) int {
	var low = 0
	var high = n - 1

	for low <= high {
		var mid = low + (high - low) >> 1
		if arr[mid] < value {
			low = mid + 1
		} else if arr[mid] > value {
			high = mid - 1
		} else {
			if mid == n - 1 || arr[mid + 1] != value {
				return mid
			}
			low = mid + 1
		}
	}

	return -1
}

// 变体三：查找第一个大于等于给定值的元素
func BSearchThree(arr []int, n, value int) int {
	var low = 0
	var high = n - 1

	for low <= high {
		var mid = low + (high - low) >> 1
		if arr[mid] < value {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid - 1] < value {
				return mid
			}
			high = mid - 1
		}
	}

	return - 1
}

// 变体四：查找最后一个小于等于给定值的元素
func BSearchFour(arr []int, n, value int) int {
	var low = 0
	var high = n - 1

	for low <= high {
		var mid = low + (high - low) >> 1
		if arr[mid] <= value {
			if mid == n - 1 || arr[mid + 1] > value {
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}