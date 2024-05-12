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
	fmt.Println(!(j > y)) // output false, since condition is true
	fmt.Println(!(true))  // output false
	fmt.Println(!(false)) // output true

	day := "sunday"
	switch day {
	case "monday":
		fmt.Println("monday")
	case "tuesday":
		fmt.Println("tuesday")
	case "wednesday":
		fmt.Println("wednesday")
	case "thursday":
		fmt.Println("thursday")
	case "friday":
		fmt.Println("friday")
	case "saturday", "sunday":
		fmt.Println("weekend")
	default:
		fmt.Println("default")
	}

	day := "wednesday"
	switch day {
	case "monday":
		fmt.Println("monday")
	case "tuesday":
		fmt.Println("tuesday")
	case "wednesday":
		fmt.Println("wednesday")
		fallthrough
	case "thursday":
		fmt.Println("thursday")
		fallthrough
	case "friday":
		fmt.Println("friday")
	case "saturday", "sunday":
		fmt.Println("weekend")
	default:
		fmt.Println("default")
	}
}
