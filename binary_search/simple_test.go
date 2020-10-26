package binary_search

import "testing"

func TestBSearch(t *testing.T) {
	arr := []int{1, 3, 3, 5, 7}
	ret := BSearch1(arr, len(arr), 7)
	t.Log(ret)
}