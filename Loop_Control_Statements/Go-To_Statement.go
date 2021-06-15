package main

import "fmt"

func main() {
	var x int = 0

Value:
	for x < 6 {
		if x == 3 {

			x = x + 4
			goto Value
		}
		fmt.Printf("value is: %d\n", x)
		x++
	}
}
