# learn_go
My notes learning golang

# Data types

Go is a statically type language


Go has a concept of types that is either explicitly declared by a programmer or is infered by the compiler</li>
It is a fast, statically typed, compliled language that feels like a dynamically type, interpreted language</li>

# Integers

int - signed Integers (positive and negative)
float - numbers with decimasl, float64 is recommended, can be float32 as well


# Characters

string - 16 bytes of Caracters

# Booleans

bool - true, false

# Arrays, Slices and Maps
```
[1,2,3,4,9]
["foo","bar"]
[7.0,9.43,0.65]

Maps
"x" -> 30
1 -> 100
"key" - "value"
```

# Storing data in variables

Go is statically typed
var <variable_name> <data_type> =  <value>

```
ex:
var s string = "Hello world"
var i int = 100
var b bool = false
var f float64 = 77.90
```
# Printing Variables

```
fmt.Println("Hello World")

var city string = "kolkata"
fmt.Println(fmt)

var name string = "GoDojo"
var user string = "Cedric"

fmt.Println("Welcome to ", name, ", ", user)
```

\n is called the Newline character
It is used to create a new line.
Placed within string expresiions
When inserted in a string, all the characters after \n are added to a new line
```
var name String = "GoDojo"
var user string = "CK"
fmt.Print(name,"\n")
fmt.Print(user)
```
The Println statement prints and adds a newline

# Printf funtcion - format specifier
Used for inserting a custom string 
fmt.Printf("Template string %s", Object args(s))

>%v formats the value in a default format
var name string = "KodeKloud"
fmt.Printf("Nice to see you here, at %v",name)
>
>%d formats decimal integers
var grades int = 42
mt.Printf("Marks: %d",grades)
>
>%T type of the value
>
>%c characters
>
>%q quoted characters/string
>
>%s plain string
>
>%t true or false
>
>%f floating numbers
>
>%.2f floating numbers up to 2 decimal places

# Declaring variables

```
var user string
user = "Harry"
fmt.Println(user)
```

Data types is important, because it defines the variable that is going to be assigned
Short hand way, only when they are of the same data type
```
var s,t string = "foo","bar"
fmt.Println(s)
fmt.Println(t)
```
Multiple assignement
```
var (
	s string = "foo"
	i int = 5)

fmt.Println(s)
fmt.Println(i)
```

## Short Variable Declaration
When go behaves as a dynamically typed variable
```
s := "Hello World"

name := "Lisa"
name = "Peter"
fmt.Println(name)
```
# Varible Scope
From which part of a program a variable is accessible, and from where it can be referenced
```
{
// Outer blocl
	{
		// Inner block
	}
}
```

Inner blocks can access variables declared within outer blocks
Outer blocks cannot access variables declared within inner blocks

## Local and Global variables
### Local Variable
- Declared inside a function or a block
- not accessible outside the function or the block.

### Global Variables
- Declared outside of a function or a block
- Available throughout the lifetime of a program
- declared at the top of the program outside all functions or blocks
- Can be accessed from any part of the program

```
package main
import ("fmt")

var name string = "Lisa"
func main() {
	fmt.Println(name)
}
```

# Zero Values

When declaring variables in goland and not giving them a value, go gives them a default value
This value is known as default value and differs per variable type
- bool - false
- int - 0
- float64 - 0.0
- string - ""

```
package main
import "fmt"

func main() {
	var fl float64
	fmt.Printf("%.2f",fl)
}

>>> go run main.go
0.00
```

Zero values of other datatypes
- int - 0
- float64 - 0.0
- string - ""
- bool -  false
- pointers,functions, interfaces,maps - nil


# User Input - Scanf

fmt.Scanf("%<format specifier> (s)", Object_arguments)

```
package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)

	fmt.Println("Hey there,", name)
}
```

Getting multiple inputs
```
	fmt.Print("Enter your name & are you a muggle: ")
	fmt.Scanf("%s %t", &surname, &is_muggle)
	fmt.Println(name, is_muggle)
```

## Scanf return values
cout - the number of arguments that the function writes to
err - any errow thrown during the execution of the funtion

````

var a string
var b int

fmt.Print("Enter a string and a number: ")
count, err := fmt.Scanf("%s %d",&a,&b)

fmt.Println("count : ", count)
fmt.Println("error: ",err)
fmt.Println("a: ",a)
fmt.Println("b: ",b)
````

# Find the type of variable

- %T format specifier
- reflect.TypeOf function from the reflect package

```
	// Type of variable
	var grades int = 42
	var message string = "Hello World"
	var isCheck bool = true
	var amount float32 = 5466.54

	fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
	fmt.Printf("variable message = %v is of type %T \n", message, message)
	fmt.Printf("variable isCheck = %v is of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount = %v is of type %T \n", amount, amount)
```

Using reflect.TypeOf

```
package main
import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Prinf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Prinf("Type: %v \n", reflect.TypeOf("Cedric"))
	fmt.Prinf("Type: %v \n", reflect.TypeOf(46.0))
	fmt.Prinf("Type: %v \n", reflect.TypeOf(true))
}
```

Combining
```
func main() {
	var grades int = 42
	var message string = "hello world"

	fmt.Printf("Variable grades=%v is of type %v \n", grades, reflect.TypeOf(grades))
	fmt.Printf("Variable message=%v is of type %v \n", message, reflect.TypeOf(message))
}
```
# Converting between data types

Type casting :
- The process of converting one data type to another is known as Type Casting
- Data types can be converted to other data types, but this does not guarantee that the value will remain intact

````
func main() {
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n",f)

	// Convert float to integers
	var f float64 = 45.89
	var i int = int(f)
	fmt.Printf("%v\n",i)

}
````

## strconv package

>Itoa()
- converts integers to string
- returns one value - string formes with the given integer

```
import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42 
	var s string = strconv.Itoa(i) // conver int to string
	fmt.Printf("%q",s)
}
```
> Atoi()
- Converts string to integer
- returns two values - the corresponding integer, error(if any).

````
 import (
	"fmt"
	"strconv"
 )

 func main() {
	var s string = "200"
	i,err := strconv.Atoi(s)

	fmt.Printf("%v, %T",i,i)
	fmt.Printf("%v, %T",err,err)
 }
````

````
 import (
	"fmt"
	"strconv"
 )

 func main() {
	var s string = "200abc"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v, %T",i,i)
	fmt.Printf("%v, %T",err,err) // Will crash an error for string convertion 200abc to an integer
 }

````

# Constants

Variables whose variables are declared and can't be changed

syntax:
const <const_name> <data_type> = <value>

Types:
- Untyped constant, when they are not explicitly given a type at declaration. Allows flexibility
```
const age = 12
const h_name, h_age = "Hermione",12
```
- Typed Constats, where you explicitly spcify the type in the declaration. Flexibility is lost
```
const name string = "Harry Potter"
const age int = 12
```

# Operators & Operands

We five types of operators in golang
- Comparison operators
- Arithmetic operators
- Assignment operators
- bitwise operators
- Logical operators

## Comparison operators

- Compare two operands and yield a Boolean value.
- Allow values of the same data type for comparisons

== , != , <, <=, >, >=

ex:
````
package main
import "fmt"

func main() {
	var city string = "kolkata"
	var city_2 string = "Calcutta"

	fmt.Println(city == city_2) //false is printed
	fmt.Println(city != city_2) //true is printed
}
````


## Arithmetic Operators

- Used to perform common arithmetic operations, such as addition, subtraction, multiplication 

+, -, *, /, %, ++, --

ex

ex:
````
package main
import "fmt"

func main() {
	var a,b string = "foo", "bar"

	fmt.Println(a+b) // prints "foobar"
}
````


## Logical Operators

- Used to determine the logic between variables or values

&& (Logical AND), ||(Logical OR), ! (Logical NOT)

### AND ( && )
- returns true if both the statements are true
- returns false when either of the statement is false

ex:

````
package main
import "fmt"

func main() {
	var x int = 10
	fmt.Println((x < 100) && (x < 200)) // true
	fmt.Println((x < 300) && (x < 0))   // false
}
````

### OR ( || )

- returns true if one of the statement is true
- returns false when both statements are false

````
package main
import "fmt"

func main() {
	var x int = 10
	fmt.Println((x < 0) || (x < 200)) // true
	fmt.Println((x < 0) || (x > 200))   // false
}
````

### NOT ( ! )

- unary operator
- Reverses the results, returns false if the expression evaluates to true and vise versa

````
package main
import "fmt"

func main() {
	var x,y int = 10,20
	fmt.Println(!(x > y))
	fmt.Println(!(true))
	fmt.Println(!(false))
}
````


## Assignment Operators

- = , assign
- += , add and assign
- -= , subtract and assign
- *= , multiply and assign
- /= , divide and assign quotient
- %= , divide and assig modulus



ex
````
package main
import "fmt"

func main() {
	// add and assign
	// x += y means x = x + y
	var x,y int = 10,20
	x += y
	fmt.Println(x) // prints 30

	// x -= y means x = x - y
	var x,y int = 10,20
	x += y
	fmt.Println(x) //print -10

	// x *= y means x = x * y
	var x,y int = 10,20
	x *= y
	fmt.Println(x) // print 200

	// x %= y means x = x % y
	var x,y int = 210,20
	x %= y
	fmt.Println(x)
}
````


## Bitwise Operators

- bitwise AND - &
	- Takes two numbers as operands and does AND on every bit of two numbers
	12 = 00001100
	25 = 00011001
	     --------
		 00001000 - answer is 8

	````
	package main
	import "fmt"

	func main() {
		var x,y int = 12,25
		z := x & y
		fmt.Println(z) //answer is 8
	}
	````

- bitwise OR - |
	- Takes two numbers as operands and does OR on every bit of two numbers
	12 = 00001100
	25 = 00011001
	     --------
		 00011101- answer is 29
	
	````
	package main
	import "fmt"

	func main() {
		var x,y int = 12,25
		z := x | y
		fmt.Println(z) //answer is 29
	}
	````


- bitwise XOR - ^
    - takes two numbers as operands and does XOR on every bit of two numbers
	- The result of XOR is 1 if the two bits are opposite

	12 = 00001100
	25 = 00011001
	     --------
		 00010101- answer is 21
	````
	package main
	import "fmt"

	func main() {
		var x,y int = 12,25
		z := x ^ y
		fmt.Println(z) //answer is 21
	}
	````

- bitwise left shift - >>
	- Shifts all bits towards left by a certain number of specified bits

	212 = 11010100
	212 << 1
	11010100 0 = 424 
	- The bit positions that have been vacated by the left shift operator are filled with 0
	````
	package main
	import "fmt"

	func main() {
		var x int = 212
		z := x << 1
		fmt.Println(z) //answer is 424
	}
	```` 
- bitwise right shift - <<
	- Shifts all bit towards right by a certain number of specified bit
	212 = 11010100
	212 >> 2
	00 110101 = 53

	- excess bits shifted off to the right are discarded
	````
	package main
	import "fmt"

	func main() {
		var x int = 212
		z := x >> 2
		fmt.Println(z) //answer is 53
	}
	````


# If-else and else if statements

## If statement

```
if (condition) {
	// executes when condition is true
}

// if .. else
if condition {
	// executes when condition is true
} else {
	// executes when condition is false
}

// if .. else if .. else
if condition_1 {
	// execute when condition_1 is true
} else if condition_2 {
	/* execute when condition_1 is false,
	and condition_2 is true */
} else if condition_3 {
	/* execute when condition_1 and 2 are false,
	and condition_3 is true */
} else {
	// when none of the above conditions are true
}
```

ex:

```
package main
import "fmt"

func main() {
	var fruit string = "grapes"
	if fruit == "apples" {
		fmt.Println("Fruit is apple")
	} else {
		fmt.Println("Fruit is not apple")
	}

	fruit := "grapes"
	if fruit == "apple" {
		fmt.Prinln("I love apples")
	} else if fruit == "orange {
		fmt.Println("Oranges are not apples")
	} else {
		fmt.Println("no appetite")
	}
}
```

## Switch-case Statement

> syntax 

```
switch expression {

case value_1:
	// execute when exp equals to value_1
case value_2:
	// execute when exp equals to value_2
default:
	// execute when no match is found
}
```

> ex

````
package main
import "fmt"

func main() {
	var i int = 100
	switch i {
		case 10:
			fmt.Println,("i is 10")
		case 100, 200:
			fmt.Println("i is either 100 or 200")
		default:
			fmt.Println("i is either 0, 100 or 200")
	}
}
````

*fallthough* - keyword
- The *fallthough* keyword is used in switch-case to force the execution flow to fall through the successive case block

````
package main
import "fmt"

func main() {
	var i int 10
	switch i {
		case -5:
			fmt.Println("-5")
		case 10:
			fmt.Println("10")
		case 20:
			fmt.Println("20")
		default:
			fmt.Println("default")
	}

}

// Prints 
10
20 default
````

### Switch with conditions

````
package main
import "fmt"

func main() {
	var a,b int 10, 20
	switch {
		case a+b == 30:
			fmt.Println("Equal to 30")
		case a+b <= 30:
			fmt.Println("less than or equal to 30")
		default:
			fmt.Printlnt("Greater than 30")
	}
}

// Print 
equal to 30
````



## Looping with the for loop

syntax 

```
for initialization; condition; post {
	// statements
}
```

````
package main
import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Hello World")
	}
	// another for loop

	i := 1
	for i <= 5 {
		fmt.Println( i *i)
		i += 1
	}
}
````

Break and continue

- the *break* statement ends the loop immediately when it is encountered

````
package main
import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}

// prints
1
2
````


- the *continue* statement skips the current iteration of loop and continues with the next iteration
````
package main
import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

// prints
1
2
4
5
````

# Arrays, slices and maps 

## Arrays
Array is a collection of similar data elements stored at contiguous memory locations.


They are of fixed length
Elements should be of the same data type

Declaration
var <array_name> [<size_of_the_array>] <data_type>
var grades [5] int
var fruits [3] string


````
package main
import "fmt"

func main() {
	var grades [5] int
	fmt.Println(grades)
}
````

## Array initialization
var grades [3]int = [3]int{10,20,30}

grades := [3]int{10,20,30}

grades := [...]int{10,20,30}

... represents ellipses, we do not need to specify the length, the compiler calculates it for us.


## len()
The length of the array refers to the number of elements stored in the array

usage :

````
func main() {
	var fruits [2]string = [2]string{"apples","orange"}
	fmt.Println(len(fruits))
}
````

## Indexes in array, 0 <= index <= len-1

First elements: 0
Last element : len(-1)

fruits[1] => "orange"


````
var fruits [5]string = [5]string{"apples","oranges","grapes","mango","papaya"}
fmt.Println(fruits[2])

	// changing indices
	var grades_2 [5]int = [5]int{90, 80, 70, 60, 50}
	fmt.Println(grades_2)

	grades_2[1] = 100
	fmt.Println(grades_2)
````

## Looping through an array

```
// Loop with for
for i := 0; i < len(grades); i++ {
	fmt.Println(grades[i])
}

// Loop with range
for index, element := range grades {
	fmt.Println(index, "=>", element)
}
```

## Multidimensional arrays

```
func main() {
	arr := [3][2]int{{2,3},{4,16},{8,64}}
	fmt.Println(arr[2][1])
}
```


# Slice
Continuous segment of an underlying array
- variable typed (elements can be added or removed)
- more flexible

3 major compenets
- pointer	: Points to the first element of the array that is accessible throught the slice
- lenght 	: number of elements in contains, len()
- capacity  : number of elements in the un underlying array, counting from the first element in the slice, cap()

## Declaration

<slice_name> := <data_type>{<values>}
grades := []int{10,20,30}


## Create slice from array
	array[start_index : end_index]
// end_index element is not included

````
 func main() {
	arr := [10]int{10,20,30,40,50,60,70,80,90,100}
	slice_1 := arr[1:8]
	fmt.Println(slice_1)

	// sub slice
	sub_slice := slice[0:3]
	fmt.Println(sub_slice)
 }

````

Delcare throught the make function
````
slice := make([]<data_type>,length,capacity)

slice := make([]int,5,10)
````

````

//slice from an array
fmt.Println("slice from an array")
a_slice := arr[1:8]
fmt.Println(cap(arr))
fmt.Println(cap(a_slice))
fmt.Println(a_slice)

````

Modifying slice elements modifies the underlying array
````
func main() {
	arr_2 := [5]int{10,20,30,40,50}

	slice_2 := arr_2[:3]

	fmt.Println(arr_2)
	fmt.Println(slice)

	slice[1] = 10
	fmt.Println("after modification")
	fmt.Println(arr_2)
	fmt.Println(slice)
}
````

## Apprending to a slice

syntax
````
func append(s []T, vs ...T) []T
slice = append(slice, element-1, element-2)

slice = append(slice,10,20,30 )
````

We can apprend a slice to another slice
````
slice = append(slice, anotherSlice...)

// Appending to a slice
fmt.Println("Appending to slices")
arr_5 := [5]int{10, 20, 30, 40, 50}

slice_5 := arr_5[:2]

arr_6 := [5]int{5, 15, 25, 35, 45}
slice_6 := arr_6[:2]

new_slice := append(slice_5, slice_6...)
fmt.Println(new_slice)
````

## Delete from a slice

````
arr := [5]int{10,20,30,40,50}

i := 2
fmt.Println(arr)

slice_1 := arr[:i]
slice_2 := arr[i+1:]

new_slice := append(slice_1,slice_2...)
fmt.Println(new_slice)
````

## Copying from a slice
Slices must be initialized with the same data type

````
func copy(dst, src []Type) int

num := copy(dest_slice,src_slice)
````

```
src_slice := []int{10,20,30,40,50}
dest_slice := make([]int,3)

num := copy(dest_slice,src_slice)
fmt.Println(dest_slice)
fmt.Println("Number of elements copied: ", num)

arr := []int{10,20,30,40,50}
for index, value := range arr {
	fmt.Println(index, "=>", value)
}

for _, value := range arr {
	fmt.Println( value)
}
```

# Maps

Unordered collection of key/value pairs.
Implemented by hash tables
Provide efficient add, get and delete operations

Declaration
````
var <map_name> map[<key_data_type>]<value_data_type>
	var my_map map[string]int

````

zero value of a map is nil
Declaring and initializing maps
````
 <map_name> := map[<key_data_type>]<value_data_type>{<key_value-pair>}

 codes := map[string]string{"en": "English", "fr": "French"}
````


## Declare and initializing a map - make function
```
<map_name> := make(map[<key_data_type>]<value_data_type>,<initial_capacity>)
```

<initial_capacity> is an optional argument

```
codes := make(map[string]int)
fmt.Println(codes)

>> output
map[]
```

## Length of a map .len()
```
codes := map[string]string{"en": "English", "fr": "French","hi":"Hindi"}
fmt.Println(len(codes))

>> ouput
3
```


## Accessing a map

```
codes := map[string]string{"en": "English", "fr": "French","hi":"Hindi"}
fmt.Println(codes["en"])
fmt.Println(codes["ft"])
fmt.Println(codes["fr"])

>>> ouput
Engligh
French
Hindi
```


## Getting a key 

value, found := map_name[key]

```
codes := map[string]int{"en":1,"fr":2,"hi":3}
value, found := codes["en"]
fmt.Println(found,value)

value, found := codes["hh"]
fmt.Println(found,value)

>>> ouput
true 1
false 0
```

## Adding items to a map

````
codes := map[string]string{"en":"English","fr":"French","hi":"Hindi"}
codes["it"] = "Italian"

fmt.Println(codes)
````


## Update value of specific value
```
codes := map[string]string{"en":"English","fr":"French","hi":"Hindi"}
codes["en"] = "English language"
fmt.Println(codes)
```

## deleting key-value pair
delete(map,key_name)

```
codes := map[string]string{"en":"English","fr":"French","hi":"Hindi"}
fmt.Println(codes)
delete(codes,"en")
fmt.Println(codes)
```

## Looping and iterating

````
func main() {
	codes := map[string]string{"en":"English","fr":"French","hi":"Hindi"}
	for key,value := range codes {
		fmt.Println(key, "=>", value)
	}
}
````

## Truncate a map
two methods :
- Iterate over the keys and delete them one-by-one
- by deleting each key one-by-one

````
func main() {
	codes := map[string]string{"en":"English","fr":"French","hi":"Hindi"}
	for key,value := range codes {
		delete(codes, key)
	}
	fmt.Println(codes)
}
````

````
func main() {
	codes := map[string]string{"en":"English","fr":"French","hi":"Hindi"}
	// reinitialize the function
	code = make(map[string]string)
	
	fmt.Println(codes)
}
````


## Playing with maps
```
package main
import "fmt"

func main() {
  ascii_codes := make(map[string]int)
  ascii_codes["A"] = 65
  ascii_codes["F"] = 70
  ascii_codes["K"] = 75
  fmt.Println(ascii_codes)

  ascii_codes = make(map[string]int)
  ascii_codes["U"] = 85
  fmt.Println(ascii_codes)
}
```
## Playing with Arrays, slices and maps

```
package main

import "fmt"

func main() {
	arr := [5]string{"one","two","three"}
	slice := arr[:3]

	my_map := make(map[int]string)

	for i, el := range slice {
		my_map[i+1] = el
	}
	fmt.Println(my_map)
}

func main() {
	arr := [5]int{}
	my_map := make(map[string]int)
	my_map["A"] = 65
	my_map["B"] = 66
	i := 0
	
	for _, values := range my_map {
		arr[i] = values
		i += 1
	}
	fmt.Println(arr)
}

func main() {
	arr := [5]int{10,20,30,90,100}
	new_slice := append(arr[:3], arr[4:]...)
	fmt.Print(new_slice)
}

func main() {
	arr := [10]string{"a", "b", "c"}
	hashmap := make(map[string]int)
	my_slice := arr[:]
	fmt.Println(len(my_slice))
	fmt.Println(cap(my_slice))
	fmt.Println(len(hashmap))
}
```

# Functions

Functions are self contained units of code which carry out a certain job

Syntax: 

```
func <function_name>(<params>) <return typs> {
	// body of the function
}


func addNumbers(a int, b int) int {
	sun := a + b
	retun sum
}
```

Calling a function
<function_name>(<argument(s)>)

Naming conventions:
- must begin with a letter
- can have any number of additional letters and symbols
- cannot contain spaces
- case-sensitive

```
package main
import "fmt"

func printgreeting(str string)
{
	fmt.Println("Hey there,",str)
}

func main() {
	printgreeting("Joe")
}

func returnCube(a int) int {
	cube := a*a*a
	fmt.Println("Cube of",a,"=",cube)
	return a*a*a
}

func main() {
	fmt.Print(returnCube(3))
}
```

## Return types

Go makes it possible to return multiple values

``` 
func operation(a int, b int) (int,int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	sum, difference := operation(20,10)
	fmt.Println(sum,difference)
}
```

## Veriadic functions

- Function that accepts variable number of arguments
- It is possible to pass a varying number of arguments of the same type as referenced in the function signature
- to declare a variadic function, the type of the final parameter is preceded by an ellipsis "..."

syntax:

func <func_name>(param-1 type, param-2 type, para-3 ...type) <return-type>

ex:
```
func sumNumbers(numbers ...int) int

func sumNumbers(str string, numbers ...int)

package main
import "fmt"

func sumNumbers(numbers ...int) int{
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func printDetails(student string, subjects ...string) {
	fmt.Println("Hey ",student,", here are your subjects -")
	for _, sub := range subjects {
		fmt.Printf(¨%s,",sub)
	}
}

func main() {
	fmt.Println,(sumNumbers())
	fmt.Println,(sumNumbers(10,20))
	fmt.Println,(sumNumbers(10,20,30,40,50))

	printDetails("Joe","Physics","Biology")
}
```


### Blank identifier '_'

Used to ignore the values that are returned by a function.

```
func f() (int, int) {
	return 42,53
}

func main {
	k; _ := f()
	fmt.Println(v)
	 
}

>> go run main.go 
	42
```


## Recursive functions

- Recursion is a concept where a function calls itself by direct or indirect means

factorial(5) = 5*4*3*2*1

```

func factorial (n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	n := 5
	result := factorial(n)
	fmt.Println("Factorial of",n,"is :", result)
}
```

## Anonymous functions

A function that is declared without any name identifier to refer to it

ex

````
func main() {
	x := func(l int, b int) int {
		return l * b
	}
	fmt.Printf("%T \n",x)
	fmt.Println(x(20,30))
}

func main() {
	x := func(l int, b int) int {
		return l * b
	} (20,30)
	fmt.Printf("%T \n",x)
	fmt.Println(x)
}

>> go run main.go
int
600
````

## High Order functions
- function that receives a function as an argument or returns a function as an output



# Pointers

A pointer is a variable that holds memory adress of another variable
They also point to where the memory is allocated and provide ways to find or even change the value located at the memory location

## Address and Derefence Operator

- & operator : The address of a variable
- * operatir : Dereference operator. It returns the value at the address

```

package main
import "fmt"

func main() {
	i := 10
	fmt.Printf("%T %v \n",&i,&i)
	fmt.Printf("%T %v \n", *(&i), *(&i))
}

>> go run main.go
	*int 0xx00018c00
	int 10
```

## Dereferencing a pointer

	*<pointer_name>
	*<pointer_name> = <new_value>

```
package main
import "fmt"

func main() {
	s := "hello"
	fmt.Printf("%T,%v \n",s ,s)
	ps := &s
	*ps = "world"
	fmt.Printf("%T %v \n",s s)
}

```

## Pass by Value in Functions

- Function is called by directly passing the value of the variable as an argument
- The parameter is copied into another location of your memory
- So, when accessing or modyfing the variable within your function, only the copy is accessed or modified, and the original value is never modified
- All basic types are passed by value

## Pass by reference

- Go supports pointers, allowing you to pass references to values within your program
- In call by reference/pointer, the address of the variable is passed into the function call as the actual parameter
- All the operations in the function are performed on the value stored at the address of the actual parameters

````
func modify(s *string) {
	*s = "world"
}

func main() {
	a := "hello"
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)
}
````


