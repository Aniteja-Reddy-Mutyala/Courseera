// MergeSort in Go
// ----------------
// This program demonstrates the merge sort algorithm using a single slice.
// It begins with an unsorted slice of 10,000 pseudo-random integers,
// then sorts them in ascending order using merge sort.
// Author: Aniteja Reddy Mutyala


package main

import (
	"fmt"
	"math/rand"
	"time"
)

// merge combines two sorted sections of the same slice.
// It uses a temporary buffer only for merging (not for recursion).
func merge(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	// Merge elements from both halves into temp
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}

	// Copy remaining elements (if any)
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}

	// Copy back the merged result into the original slice
	copy(arr[left:right+1], temp)
}

// mergeSort recursively divides and merges the slice in place.
func mergeSort(arr []int, left, right int) {
	if left >= right {
		return // base case: single element
	}

	mid := (left + right) / 2
	mergeSort(arr, left, mid)   // sort left half
	mergeSort(arr, mid+1, right) // sort right half
	merge(arr, left, mid, right) // merge sorted halves
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate pseudo-random test data of 10,000 integers
	const size = 10000
	numbers := make([]int, size)
	for i := range numbers {
		numbers[i] = rand.Intn(10000) // random values up to 10,000
	}

	fmt.Println("Unsorted sample (first 10 elements):", numbers[:10])

	// Perform merge sort (in-place on the same slice)
	start := time.Now()
	mergeSort(numbers, 0, len(numbers)-1)
	elapsed := time.Since(start)

	fmt.Println("Sorted sample (first 10 elements):", numbers[:10])
	fmt.Println("Sorted output is ",numbers[:])
	fmt.Printf("Sorting completed in %v\n", elapsed)
}
