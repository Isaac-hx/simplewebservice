package repository

import (
	"testing"
)

var (
	testList       = [][]int{{1, 2, 3, 4, 5, 6, 7, 8}, {10, 11, 12, 13, 14, 15, 16, 17}}
	itemInput      = []int{2, 17}
	expectedOutput = []int{1, 7}
)

func TestBinarySearch(t *testing.T) {
	var result int
	for i, data := range testList {
		t.Logf("Execution function")
		result = BinarySearch(data, itemInput[i])
		if result != expectedOutput[i] {
			t.Fatalf("Task Failed output : %d expected : %d", result, expectedOutput[i])
		}
	}
}
