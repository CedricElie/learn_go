package main

import "fmt"

func modify1(numbers ...int) {
	for i := range numbers {
		numbers[i] -= 5
	}
}
func main() {
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))

	// dereferencing pointers
	s := "hello"
	fmt.Printf("%T,%v \n", s, s)
	ps := &s
	*ps = "world"
	fmt.Printf("%T %v \n", s, s)

	arr := []int{10, 20, 30}
	fmt.Println(arr)
	modify1(arr...)
	fmt.Println(arr)
}
