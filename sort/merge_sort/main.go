package main

// MergeSort
// mergeSortC
// merge

func MergeSort(nums []int) {
	mergeSortC(nums, 0, len(nums) - 1)
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
	// 申请一个跟nums一样的临时数组
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

	// 判断哪个子数组中有剩余的数据
	for ; i <= mid; i++ {
		tmpArr[k] = nums[i]
		k++
	}

	// 判断哪个子数组中有剩余的数据
	for ; j <= right; j++ {
		tmpArr[k] = nums[j]
		k++
	}

	// 将tmp中的数组拷贝回nums
	copy(nums[left:right+1], tmpArr)
}