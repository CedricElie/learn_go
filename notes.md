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


