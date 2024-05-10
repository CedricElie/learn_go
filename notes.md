# Data types

Go is a statically type language


Go has a concept of types that is either explicitly declared by a programmer or is infered by the compiler</li>
It is a fast, statically typed, compliled language that feels like a dynamically type, interpreted language</li>

# Integers

int - signed Integers (positive and negative)
float - numbers with decimasl, float64 is recommended, can be float32 as well


# Caracters

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

var name String = "GoDojo"
var user string = "CK"
fmt.Print(name,"\n")
fmt.Print(user)

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