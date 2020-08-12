package main

import "fmt"

func search_insert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	left := 0
	right := l
	for left < right {
		mid := left + (right - left) / 2
		fmt.Printf("left=%v,right=%v,mid=%v\n", left, right, mid)
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0
	fmt.Println(search_insert(nums, target))
}
