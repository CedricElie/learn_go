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
