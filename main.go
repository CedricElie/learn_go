package main

import "fmt"

//this is a comment

/*
this is a
multi line
comment
*/

func main() {
	fmt.Println("Hello World!")

	name := "Lisa"
	fmt.Println(name)

	var greeting string = "Hello World"
	fmt.Println(greeting)

	var lab string = "GoDojo"
	var user string = "CK"

	fmt.Print("Welcome to ", lab, ", ", user)

	// Format specifiers
	var joe string = "Joe"
	var i int = 78
	fmt.Printf("Hey, %v! You have scored %d/100 in Physics", joe, i)

	// Declaring and assigning later
	var harry string
	harry = "Harry"
	fmt.Println(harry)

	city := "London"
	{
		country := "UK"
		fmt.Println(country)
		fmt.Println(city)
	}
	fmt.Println(country)
	fmt.Println(city)
}
