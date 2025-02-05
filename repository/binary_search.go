package repository

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		// Periksa apakah elemen tengah adalah target
		if arr[mid] == target {
			return mid
		}

		// Jika target lebih kecil, fokus ke bagian kiri
		if target < arr[mid] {
			right = mid - 1
		} else {
			// Jika target lebih besar, fokus ke bagian kanan
			left = mid + 1
		}
	}

	// Jika tidak ditemukan, kembalikan -1
	return -1
}
