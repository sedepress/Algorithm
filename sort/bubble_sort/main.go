package main

func BubbleSort(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}

	for i := 0; i < l; i++ {
		var flag = false
		for j := 0; j < l - i - 1; j++ {
			if nums[j] > nums[j + 1] {
				nums[j], nums[j + 1] = nums[j + 1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}