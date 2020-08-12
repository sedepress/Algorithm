package main

import "fmt"

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
func plus_one(digits []int) []int {
	for i := len(digits)- 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	return append([]int{1}, digits...)
}

func main() {
	d := []int{9}
	r := plus_one(d)
	for _, v := range r {
		fmt.Println(v)
	}
}
