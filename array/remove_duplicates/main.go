package main

import "fmt"
// 给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// 双指针法
func remove_duplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	var i = 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func main() {
	nums := []int{1, 2, 4, 4, 5}
	ret := remove_duplicates(nums)
	fmt.Println(ret)
}
