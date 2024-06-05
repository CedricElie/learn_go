package main

import "fmt"

func main() {

	var fruits [2]string = [2]string{"apples", "oranges"}
	fmt.Println(fruits)

	// Initializing the array
	var grades [3]int = [3]int{10, 20, 30}
	fmt.Println(grades)

	marks := [3]int{10, 20, 30}
	fmt.Println(marks)

	names := [...]string{"Rachel", "Phoebe", "Monica"}
	fmt.Println(names)

	// Length of an array using the len() functiin
	var frutus [2]string = [2]string{"apples", "orange"}
	fmt.Println(len(frutus))

	var frutas [5]string = [5]string{"apples", "oranges", "grapes", "mango", "papaya"}
	fmt.Println(frutas[4])

	// changing indices
	var grades_2 [5]int = [5]int{90, 80, 70, 60, 50}
	fmt.Println(grades_2)

	grades_2[1] = 100
	fmt.Println(grades_2)

	// Looping through the array
	for i := 0; i < len(grades_2); i++ {
		fmt.Println("Element :", i, " is: ", grades_2[i])
	}

	// Loop with range
	for index, element := range grades_2 {
		fmt.Println(index, "=>", element)
	}

	// Multidimensional arrays
	arr := [3][2]int{{2, 3}, {4, 16}, {8, 64}}
	fmt.Println(arr[2][1])
}
