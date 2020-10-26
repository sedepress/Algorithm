package main

import (
	"math/rand"
	"testing"
)

func createRandomArr(length int) []int {
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(10000)
	}
	return arr
}

func TestInsertionSort(t *testing.T) {
	arr := createRandomArr(10000)
	t.Log(arr)
	InsertionSort(arr, len(arr))
	t.Log(arr)
}