package main

import (
		"fmt"
)

func sumSlice(numbers []int) int {
	sum := 0 // Initialize the sum to zero

	for _, num := range numbers {
		sum += num
	}

	return sum
}

//Function to find the max
func findMax(numbers []int)(int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("Input slice cannot be empty")
	}

	max := numbers[0]
	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
	}
	return max, nil
}

func linearSearch(slice []int, target int) int {
	for index, element := range slice {
		if element == target {
			return index
		}
	}
	return -1
}

func binarySearch(sortedSlice []int, target int) int {
	low := 0
	high := len(sortedSlice) - 1

	for low <= high {
		mid := (low + high) / 2

		if sortedSlice[mid] == target {
			return mid // Target found
		} else if sortedSlice[mid] < target {
			low = mid + 1 // Search in the right half
		} else {
			high = mid - 1 // Search in the left half
		}
	}
	return -1 // Target not found
}

func bubbleSort(slice []int) []int {
	n := len(slice)
	for i := 0; i < n-1 ; i++ {
		// Last 1 elements are already in place
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				//swap them
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}

func selectionSort(slice []int) []int {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		// Find the index of the minimum element in the unsorted part
		minIndex := 1
		for j := i + 1; j <n; j++ {
			if slice[j] < slice[minIndex] {
				minIndex = j
			}
		}
		// swap the found minimum elements with the element at index 1
		if minIndex != 1 {
			slice[i], slice[minIndex] = slice[minIndex], slice[i]
		}
	}
	return slice
}
func main() {

	//Find max
	intNumbers := []int{3,1,4,1,5}
	maxInt, errInt := findMax(intNumbers)
	total := sumSlice(intNumbers)
	if errInt != nil {
		fmt.Println("Error:", errInt)
	} else {
		fmt.Println("Max of integer slice:", maxInt)
		fmt.Println("The sum of the number is:", total)
	}

	// Linear Search
	numbers := []int{2,7,1,9,4,6,3,8,5}
	target := 4

	index := linearSearch(numbers,target)

	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target,index)
	} else {
		fmt.Printf("Target %d not found in the slice\n",target)
	}

	targetNotFound := 10
	indexNotFound := linearSearch(numbers, targetNotFound)
	if indexNotFound != -1 {
		fmt.Printf("Target %d found at index %d\n",targetNotFound,indexNotFound)
	} else {
		fmt.Printf("Target %d not found in the slice\n",targetNotFound)
	}

	// Binary search
	sortedNumbers := []int{2,3,4,5,6,7,8,9,10}
	targetFound := 7

	indexFound := binarySearch(sortedNumbers,targetFound)

	if indexFound != -1 {
		fmt.Printf("Target %d found at index %d\n",targetFound,indexFound )
	} else {
		fmt.Printf("Target %d not found\n",targetFound)
	}

	targetNF := 1
	IndexNotFound := binarySearch(sortedNumbers,targetNF)

	if IndexNotFound != -1 {
		fmt.Printf("Target %d found at index %d\n",targetNF,IndexNotFound )
	} else {
		fmt.Printf("Target %d not found\n",targetNF)	
	}

	targetAtEnd := 10
	IndexAtEnd := binarySearch(sortedNumbers,targetAtEnd)

	if IndexAtEnd != -1 {
		fmt.Printf("Target %d found at index %d\n",targetAtEnd,IndexAtEnd )
	} else {
		fmt.Printf("Target %d not found\n",targetAtEnd)	
	}

	// Bubble sort
	unsortedNumbers := []int{5,1,4,2,8}
	b_sortedNumbers := bubbleSort(unsortedNumbers)
	fmt.Println("Sorted slice:", b_sortedNumbers )

	anotherUnsorted := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	anotherSorted := bubbleSort(anotherUnsorted)
	fmt.Println("Another sorted slice:",anotherSorted)

	// Selection sort
	s_unsortednumnber := []int{6, 4, 1, 8, 3}
	s_sortedNumbers := selectionSort(s_unsortednumnber)
	fmt.Println("Sorted slice:", s_sortedNumbers)

	a_unsortednumnber := []int{9, 2, 5, 1, 7}
	a_sortedNumbers := selectionSort(a_unsortednumnber)
	fmt.Println("Another Sorted slice:", a_sortedNumbers)
}