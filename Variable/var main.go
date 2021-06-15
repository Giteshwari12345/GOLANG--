package main

import (
	"fmt"
)

var z = 21

func main() {
	x := 42 + 7
	y := "Hello GO"
	fmt.Println(x)
	fmt.Println(y)
	x = 50
	fmt.Println(x)
	fmt.Println(z) /* z is in scope here as well */
}

func foo() {
	fmt.Println(z) /* z is in scope anywhere throughout the package */
}
