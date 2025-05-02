package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2 // Calculate the middle index

		if arr[mid]== target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {

		swapped := false
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
	

}

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i] // The current element to be inserted

		j := i -1 // Index of the last element in the sorted portion

		// Move elements of arr[0..1] that are greater than key
		// to on position ahead of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key // Insert the key into its correct position
	}
}

// MergeSort sorts a slice of integers using the Merge Sort algorithm
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr // Base case: a slice with 0 or 1 element is already sorted
	}

	// 1. Divide: Find the middle point and divide the array into two halves
	mid := len(arr) / 2
	left := arr[:mid]  // Elements from 0 to mid-1
	right := arr[mid:] // Elements from mid to len(arr)-1

	// 2. Conquer: Recursively sort the two halves
	left = MergeSort(left)
	right = MergeSort(right)

	// 3. Combine: Merge the sorted halves
	return merge(left, right)
}

// merge merges two sorted slices into a single sorted slice
func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right)) // Pre-allocate for efficiency
	i, j := 0, 0

	// Compare elements from the left and right slices and add the smaller one to the result
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Add the remaining elements from the left slice (if any)
	for ; i < len(left); i++ {
		result = append(result, left[i])
	}

	// Add the remaining elements from the right slice (if any)
	for ; j < len(right); j++ {
		result = append(result, right[j])
	}

	return result
}


func main() {
	fmt.Println("Hello world")

	//Numbers slice must be ordered
	numbers := []int{-1, 2, 5, 7, 9, 12, 15, 18, 20, 24, 28}
	target := 15

	index := BinarySearch(numbers, target)

	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Target %d not found in the slice \n", target)
	}

	target = 10
	index = BinarySearch(numbers, target)

	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Target %d not found in the slice \n", target)
	}

	//bubble sort

	numbers = []int{5,2,9,1,5,6}
	fmt.Println("Unsorted:",numbers)
	fmt.Println(len(numbers))
	BubbleSort(numbers)
	fmt.Println("Bubble Sorted:", numbers)

	// Insertion sort
	numbers = []int{5, 2, 9, 1, 5, 6}
	fmt.Println("Unsorted:", numbers)

	InsertionSort(numbers)
	fmt.Println("Insertion Sorted:", numbers)


	// Merge Sort
	numbers = []int{5, 2, 9, 1, 5, 6, 8, 3, 7}
	fmt.Println("Merge Unsorted:", numbers)

	sortedNumbers := MergeSort(numbers)
	fmt.Println("Merge Sorted:", sortedNumbers)
}