package main

import "fmt"

func main() {

	slice := []int{10, 20, 30}
	fmt.Println(slice)

	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice_1 := arr[1:8]
	fmt.Println(slice_1)

	// sub slice
	sub_slice := slice_1[0:3]
	fmt.Println(sub_slice)

	// Slice using make
	fmt.Println("Slice using make")
	m_slice := make([]int, 5, 8)
	fmt.Println(m_slice)
	fmt.Println(len(m_slice))
	fmt.Println((cap(m_slice)))

	//slice from an array
	fmt.Println("slice from an array")
	a_slice := arr[1:8]
	fmt.Println(cap(arr))
	fmt.Println(cap(a_slice))
	fmt.Println(a_slice)

	// Modifying array elements
	fmt.Println("Modifying elements")
	arr_2 := [5]int{10, 20, 30, 40, 50}

	slice_2 := arr_2[:3]

	fmt.Println(arr_2)
	fmt.Println(slice_2)

	slice_2[1] = 9000
	fmt.Println("after modification")
	fmt.Println(arr_2)
	fmt.Println(slice_2)

	fmt.Println("Create new array to test appending")
	arr_4 := [4]int{10, 20, 30, 40}
	slice_4 := arr_4[1:3]

	fmt.Println(slice_4)
	fmt.Println(len(slice_4))
	fmt.Println(cap(slice_4))

	slice_4 = append(slice_4, 900, -90, 50)
	fmt.Println(slice_4)
	fmt.Println(len(slice_4))
	fmt.Println(cap(slice_4))

	// Appending to a slice
	fmt.Println("Appending to slices")
	arr_5 := [5]int{10, 20, 30, 40, 50}

	slice_5 := arr_5[:2]

	arr_6 := [5]int{5, 15, 25, 35, 45}
	slice_6 := arr_6[:2]

	new_slice := append(slice_5, slice_6...)
	fmt.Println(new_slice)

	// Deleting from a slice
	fmt.Println("Deleting from a slice")
	i := 2
	fmt.Println(new_slice)

	s_1 := arr_5[:1]
	s_2 := arr_5[i+1:]

	n_s := append(s_1, s_2...)
	fmt.Println(n_s)

	// Copying a slice
	fmt.Println("Copying a slice")
	src_slice := []int{10, 20, 30, 40, 50}
	dest_slice := make([]int, 3)

	num := copy(dest_slice, src_slice)
	fmt.Println(dest_slice)
	fmt.Println("Number of elements copied: ", num)
}
