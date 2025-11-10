package main

import "fmt"

// merge combines two sorted slices into one sorted slice
func merge(left, right []int) []int {
	merged := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged[k] = left[i]
			i++
			k++
		} else {
			merged[k] = right[j]
			j++
			k++
		}
	
	}

	// Copy remaining elements
	for i < len(left) {
		merged[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		merged[k] = right[j]
		j++
		k++
	}

	return merged
}

// mergeSort recursively divides the slice and merges sorted halves
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr // base case: already sorted
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])   // divide left half
	right := mergeSort(arr[mid:])  // divide right half

	return merge(left, right) // conquer: merge sorted halves
}

func main() {
	arr := []int{10, 5, 3, 9, 12, 1, 7}
	fmt.Println("Original:", arr)

	sorted := mergeSort(arr)
	fmt.Println("Sorted:", sorted)
}
