package main

import "fmt"

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
// 暴力破解法 时间复杂度O(n²)
//func two_sum(nums []int, target int) []int {
//	for i, v := range nums {
//		for k := i + 1; k < len(nums); k++ {
//			if target - v == nums[k] {
//				return []int{i, k}
//			}
//		}
//	}
//	return []int{}
//}

// map法
func two_sum(nums []int, target int) []int {
	result := []int{}
	m := make(map[int]int)
	for i, k := range nums {
		if value, exist := m[target-k]; exist {
			result = append(result, value)
			result = append(result, i)
		}
		m[k] = i
	}
	return result
}

func main() {
	nums := []int{5, 7, 11, 20}
	target := 12

	ret := two_sum(nums, target)
	fmt.Printf("下标1为%v,下标2为%v", ret[0], ret[1])
}
