package main

import "fmt"

func main() {
	var city string = "kolkata"
	var city_2 string = "Calcutta"

	fmt.Println(city == city_2)
	fmt.Println(city != city_2)

	var a, b string = "foo", "bar"
	fmt.Println(a + b)

	var c, d float64 = 79.02, 75.66
	fmt.Printf("%.2f \n", c-d)

	var e, f int = 12, 2
	fmt.Println(e * f)

	//Increment operators
	var i int = 1
	i++
	fmt.Println(i) //prints 2

	//Decrement operators
	var k int = 1
	k--
	fmt.Println(k) //print 0

	// Logical operators
	var x int = 10
	fmt.Println((x < 100) && (x < 200)) // true
	fmt.Println((x < 300) && (x < 0))   // false

	fmt.Println((x < 0) || (x < 200)) // true
	fmt.Println((x < 0) || (x > 200)) // false

	var j, y int = 10, 20
	fmt.Println(!(j > y))
	fmt.Println(!(true))
	fmt.Println(!(false))
}
