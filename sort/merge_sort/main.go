package main

func mergeSort(nums []int) {
	l := len(nums)
	mergeSortC(nums, 0, l - 1)
}

func mergeSortC(nums []int, left, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2

	mergeSortC(nums, left, mid)
	mergeSortC(nums, mid + 1, right)

	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {
	tmpArr := make([]int, right-left+1)

	i := left
	j := mid + 1
	k := 0
	for ; i <= mid && j <= right; k++ {
		if nums[i] <= nums[j] {
			tmpArr[k] = nums[i]
			i++
		} else {
			tmpArr[k] = nums[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmpArr[k] = nums[i]
		k++
	}

	for ; j <= right; j++ {
		tmpArr[k] = nums[j]
		k++
	}

	copy(nums[left:right+1], tmpArr)
}