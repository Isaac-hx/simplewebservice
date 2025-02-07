package repository

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) < 2 { //base case
		fmt.Println(arr)
		return arr
	}
	//bagian terpenting dari quicksort
	pivot := arr[0]
	var left []int
	var right []int
	//bagian utama

	for _, data := range arr[1:] {
		if data <= pivot {
			left = append(left, data)
		} else {
			right = append(right, data)
		}
	}

	// Recursively sort left and right, then concatenate them with pivot
	result := append(QuickSort(left), pivot)     //berjalan sampai iterasi ke dua
	result = append(result, QuickSort(right)...) //berjalan sampai iterasi terakhir!
	fmt.Println(result)
	return result
}
