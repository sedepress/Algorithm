package main

import (
	"math/rand"
	"testing"
)

func createRandomArr(length int) []int {
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func TestBubbleSort(t *testing.T) {
	arr := createRandomArr(10)
	t.Log(arr)
	BubbleSort(arr)
	t.Log(arr)
}