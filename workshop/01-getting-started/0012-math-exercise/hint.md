# Reading inputs from Console


## Ensure types - fmt.Scanf

To read inputs from the console, you can use the `fmt.Scanf` function. The `fmt.Scanf` function reads formatted text from the standard input (console) and stores the results in the variables passed as arguments.
Bear in mind that the type of the variable passed as an argument to `fmt.Scanf` should match the type of the input.

```go
package main

import (
	"fmt"
)

func main() {
	var input float64
	fmt.Print("Enter the radius of the circle: ")
	// it's important to pass the address of the variable
	// we will dive into pointers later
	_, _ = fmt.Scanf("%f", &input)
	fmt.Println("The radius of the circle is:", input)
}
```

## Alternative - fmt.Scanln
You can use the `fmt.Scanln` function to read a single line of text from the standard input (console) and store the result in the variable passed as an argument.

```go
package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Enter the radius of the circle: ")
	// it's important to pass the address of the variable
	// we will dive into pointers later
	_, _ = fmt.Scanln(&input)
	fmt.Println("The radius of the circle is:", input)
}

```