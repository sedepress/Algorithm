package main

func selectionSort(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}

	for i := 0; i < l; i++ {
		minIndex := i
		for j := i + 1; j < l; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
