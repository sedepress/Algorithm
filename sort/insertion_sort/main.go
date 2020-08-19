package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}

	for i := 1; i < l; i++ {
		value := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value
	}
}

func main() {
	var nums []int
	for i := 0; i < 100000; i++ {
		nums = append(nums, rand.Int())
	}

	t1 := time.Now()
	insertionSort(nums)
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}
