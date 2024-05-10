package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var name string
	// fmt.Print("Enter your name: ")
	// fmt.Scanf("%s", &name)

	// fmt.Println("Hey there,", name)

	// var surname string
	// var is_muggle bool

	// fmt.Print("Enter your name & are you a muggle: ")
	// fmt.Scanf("%s %t", &surname, &is_muggle)
	// fmt.Println(name, is_muggle)

	var a string
	var b int

	fmt.Print("Enter a string and a number: ")
	count, err := fmt.Scanf("%s %d", &a, &b)

	fmt.Println("count : ", count)
	fmt.Println("error: ", err) // This will crash a <nil> error
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	// Type of variable
	var grades int = 42
	var message string = "Hello World"
	var isCheck bool = true
	var amount float32 = 5466.54

	fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
	fmt.Printf("variable message = %v is of type %T \n", message, message)
	fmt.Printf("variable isCheck = %v is of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount = %v is of type %T \n", amount, amount)

	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("Cedric"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(46.0))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))

	fmt.Printf("Variable grades=%v is of type %v \n", grades, reflect.TypeOf(grades))
	fmt.Printf("Variable message=%v is of type %v \n", message, reflect.TypeOf(message))
}
