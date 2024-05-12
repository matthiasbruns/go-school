# Generating random numbers


## Generate a random integer
To generate a random integer, you can use the `rand.Intn` function from the `math/rand` package. The `rand.Intn` function returns a random integer in the range [0, n).

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    // Generate a random number between 0 and 100
    random := rand.Intn(100)
    fmt.Println(random)
}
```