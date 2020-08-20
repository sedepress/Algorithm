package main

func QuickSort(nums []int) {
	separateSort(nums, 0, len(nums) - 1)
}

func separateSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	i := partition(nums, start, end)
	separateSort(nums, start, i - 1)
	separateSort(nums, i + 1, end)
}

func partition(nums []int, start, end int) int {
	pivot := nums[end]

	var i = start
	for j := start; j < end; j++ {
		if nums[j] < pivot {
			if !(i == j) {
				nums[i], nums[j] = nums[j], nums[i]
			}
			i++
		}
	}

	nums[i], nums[end] = nums[end], nums[i]

	return i
}
