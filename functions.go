package main

import "fmt"

func returnCube(a int) int {
	cube := a * a * a
	fmt.Println("Cube of", a, "=", cube)
	return a * a * a
}

func main() {
	fmt.Print(returnCube(3))
}
