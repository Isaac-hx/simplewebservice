package repository

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var left []int
	var right []int

	for _, data := range arr[1:] {
		if data <= pivot {
			left = append(left, data)
		} else {
			right = append(right, data)
		}
	}

	// Recursively sort left and right, then concatenate them with pivot
	result := append(QuickSort(left), pivot)
	result = append(result, QuickSort(right)...)
	fmt.Println(left, pivot, right)
	return result
}
